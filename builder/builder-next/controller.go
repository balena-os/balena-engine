package buildkit

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/containerd/containerd/content/local"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/builder/builder-next/adapters/containerimage"
	"github.com/docker/docker/builder/builder-next/adapters/snapshot"
	containerimageexp "github.com/docker/docker/builder/builder-next/exporter"
	"github.com/docker/docker/builder/builder-next/imagerefchecker"
	mobyworker "github.com/docker/docker/builder/builder-next/worker"
	"github.com/docker/docker/daemon/config"
	"github.com/docker/docker/daemon/graphdriver"
	units "github.com/docker/go-units"
	"github.com/moby/buildkit/cache"
	"github.com/moby/buildkit/cache/metadata"
	registryremotecache "github.com/moby/buildkit/cache/remotecache/registry"
	"github.com/moby/buildkit/client"
	"github.com/moby/buildkit/control"
	"github.com/moby/buildkit/exporter"
	"github.com/moby/buildkit/frontend"
	dockerfile "github.com/moby/buildkit/frontend/dockerfile/builder"
	"github.com/moby/buildkit/frontend/gateway"
	"github.com/moby/buildkit/frontend/gateway/forwarder"
	"github.com/moby/buildkit/snapshot/blobmapping"
	"github.com/moby/buildkit/solver/bboltcachestorage"
	"github.com/moby/buildkit/worker"
	"github.com/pkg/errors"
)

func newController(rt http.RoundTripper, opt Opt) (*control.Controller, error) {
	if err := os.MkdirAll(opt.Root, 0700); err != nil {
		return nil, err
	}

	dist := opt.Dist
	root := opt.Root

	var driver graphdriver.Driver
	if ls, ok := dist.LayerStore.(interface {
		Driver() graphdriver.Driver
	}); ok {
		driver = ls.Driver()
	} else {
		return nil, errors.Errorf("could not access graphdriver")
	}

	sbase, err := snapshot.NewSnapshotter(snapshot.Opt{
		GraphDriver: driver,
		LayerStore:  dist.LayerStore,
		Root:        root,
	})
	if err != nil {
		return nil, err
	}

	store, err := local.NewStore(filepath.Join(root, "content"))
	if err != nil {
		return nil, err
	}
	store = &contentStoreNoLabels{store}

	md, err := metadata.NewStore(filepath.Join(root, "metadata.db"))
	if err != nil {
		return nil, err
	}

	snapshotter := blobmapping.NewSnapshotter(blobmapping.Opt{
		Content:       store,
		Snapshotter:   sbase,
		MetadataStore: md,
	})

	layerGetter, ok := sbase.(imagerefchecker.LayerGetter)
	if !ok {
		return nil, errors.Errorf("snapshotter does not implement layergetter")
	}

	refChecker := imagerefchecker.New(imagerefchecker.Opt{
		ImageStore:  dist.ImageStore,
		LayerGetter: layerGetter,
	})

	cm, err := cache.NewManager(cache.ManagerOpt{
		Snapshotter:     snapshotter,
		MetadataStore:   md,
		PruneRefChecker: refChecker,
	})
	if err != nil {
		return nil, err
	}

	src, err := containerimage.NewSource(containerimage.SourceOpt{
		SessionManager:  opt.SessionManager,
		CacheAccessor:   cm,
		ContentStore:    store,
		DownloadManager: dist.DownloadManager,
		MetadataStore:   dist.V2MetadataService,
		ImageStore:      dist.ImageStore,
		ReferenceStore:  dist.ReferenceStore,
		ResolverOpt:     opt.ResolverOpt,
	})
	if err != nil {
		return nil, err
	}

	exec, err := newExecutor(root, opt.DefaultCgroupParent, opt.NetworkController)
	if err != nil {
		return nil, err
	}

	differ, ok := sbase.(containerimageexp.Differ)
	if !ok {
		return nil, errors.Errorf("snapshotter doesn't support differ")
	}

	exp, err := containerimageexp.New(containerimageexp.Opt{
		ImageStore:     dist.ImageStore,
		ReferenceStore: dist.ReferenceStore,
		Differ:         differ,
	})
	if err != nil {
		return nil, err
	}

	cacheStorage, err := bboltcachestorage.NewStore(filepath.Join(opt.Root, "cache.db"))
	if err != nil {
		return nil, err
	}

	gcPolicy, err := getGCPolicy(opt.BuilderConfig, root)
	if err != nil {
		return nil, errors.Wrap(err, "could not get builder GC policy")
	}

	wopt := mobyworker.Opt{
		ID:                "moby",
		SessionManager:    opt.SessionManager,
		MetadataStore:     md,
		ContentStore:      store,
		CacheManager:      cm,
		GCPolicy:          gcPolicy,
		Snapshotter:       snapshotter,
		Executor:          exec,
		ImageSource:       src,
		DownloadManager:   dist.DownloadManager,
		V2MetadataService: dist.V2MetadataService,
		Exporters: map[string]exporter.Exporter{
			"moby": exp,
		},
		Transport: rt,
	}

	wc := &worker.Controller{}
	w, err := mobyworker.NewWorker(wopt)
	if err != nil {
		return nil, err
	}
	wc.Add(w)

	frontends := map[string]frontend.Frontend{
		"dockerfile.v0": forwarder.NewGatewayForwarder(wc, dockerfile.Build),
		"gateway.v0":    gateway.NewGatewayFrontend(wc),
	}

	return control.NewController(control.Opt{
		SessionManager:           opt.SessionManager,
		WorkerController:         wc,
		Frontends:                frontends,
		CacheKeyStorage:          cacheStorage,
		ResolveCacheImporterFunc: registryremotecache.ResolveCacheImporterFunc(opt.SessionManager, opt.ResolverOpt),
		// TODO: set ResolveCacheExporterFunc for exporting cache
	})
}

func getGCPolicy(conf config.BuilderConfig, root string) ([]client.PruneInfo, error) {
	var gcPolicy []client.PruneInfo
	if conf.GC.Enabled {
		var (
			defaultKeepStorage int64
			err                error
		)

		if conf.GC.DefaultKeepStorage != "" {
			defaultKeepStorage, err = units.RAMInBytes(conf.GC.DefaultKeepStorage)
			if err != nil {
				return nil, errors.Wrapf(err, "could not parse '%s' as Builder.GC.DefaultKeepStorage config", conf.GC.DefaultKeepStorage)
			}
		}

		if conf.GC.Policy == nil {
			gcPolicy = mobyworker.DefaultGCPolicy(root, defaultKeepStorage)
		} else {
			gcPolicy = make([]client.PruneInfo, len(conf.GC.Policy))
			for i, p := range conf.GC.Policy {
				b, err := units.RAMInBytes(p.KeepStorage)
				if err != nil {
					return nil, err
				}
				if b == 0 {
					b = defaultKeepStorage
				}
				gcPolicy[i], err = toBuildkitPruneInfo(types.BuildCachePruneOptions{
					All:         p.All,
					KeepStorage: b,
					Filters:     p.Filter,
				})
				if err != nil {
					return nil, err
				}
			}
		}
	}
	return gcPolicy, nil
}
