package daemon // import "github.com/docker/docker/daemon"

import (
	"archive/tar"
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"os"
	"runtime"
	"strings"
	"time"

"github.com/containerd/log"
	"github.com/containerd/platforms"
	"github.com/distribution/reference"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/backend"
	containertypes "github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/events"
	networktypes "github.com/docker/docker/api/types/network"
	"github.com/docker/docker/container"
	"github.com/docker/docker/daemon/config"
	"github.com/docker/docker/daemon/images"
	"github.com/docker/docker/errdefs"
	"github.com/docker/docker/image"
	"github.com/docker/docker/internal/multierror"
	"github.com/docker/docker/layer"
	"github.com/docker/docker/pkg/idtools"
	"github.com/docker/docker/pkg/ioutils"
	"github.com/docker/docker/pkg/progress"
	"github.com/docker/docker/pkg/streamformatter"
	"github.com/docker/docker/pkg/stringid"
	"github.com/docker/docker/runconfig"
	librsync "github.com/balena-os/librsync-go"
	units "github.com/docker/go-units"
	"github.com/opencontainers/go-digest"
	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
	"github.com/opencontainers/selinux/go-selinux"
	"github.com/pkg/errors"
	archvariant "github.com/tonistiigi/go-archvariant"
)

type createOpts struct {
	params                  backend.ContainerCreateConfig
	managed                 bool
	ignoreImagesArgsEscaped bool
}

// CreateManagedContainer creates a container that is managed by a Service
func (daemon *Daemon) CreateManagedContainer(ctx context.Context, params backend.ContainerCreateConfig) (containertypes.CreateResponse, error) {
	return daemon.containerCreate(ctx, daemon.config(), createOpts{
		params:  params,
		managed: true,
	})
}

// ContainerCreate creates a regular container
func (daemon *Daemon) ContainerCreate(ctx context.Context, params backend.ContainerCreateConfig) (containertypes.CreateResponse, error) {
	return daemon.containerCreate(ctx, daemon.config(), createOpts{
		params: params,
	})
}

// ContainerCreateIgnoreImagesArgsEscaped creates a regular container. This is called from the builder RUN case
// and ensures that we do not take the images ArgsEscaped
func (daemon *Daemon) ContainerCreateIgnoreImagesArgsEscaped(ctx context.Context, params backend.ContainerCreateConfig) (containertypes.CreateResponse, error) {
	return daemon.containerCreate(ctx, daemon.config(), createOpts{
		params:                  params,
		ignoreImagesArgsEscaped: true,
	})
}

func (daemon *Daemon) containerCreate(ctx context.Context, daemonCfg *configStore, opts createOpts) (containertypes.CreateResponse, error) {
	start := time.Now()
	if opts.params.Config == nil {
		return containertypes.CreateResponse{}, errdefs.InvalidParameter(runconfig.ErrEmptyConfig)
	}
	// TODO(thaJeztah): remove logentries check and migration code in release v26.0.0.
	if opts.params.HostConfig != nil && opts.params.HostConfig.LogConfig.Type == "logentries" {
		return containertypes.CreateResponse{}, errdefs.InvalidParameter(fmt.Errorf("the logentries logging driver has been deprecated and removed"))
	}

	// Normalize some defaults. Doing this "ad-hoc" here for now, as there's
	// only one field to migrate, but we should consider having a better
	// location for this (and decide where in the flow would be most appropriate).
	//
	// TODO(thaJeztah): we should have a more visible, more canonical location for this.
	if opts.params.HostConfig != nil && opts.params.HostConfig.RestartPolicy.Name == "" {
		// Set the default restart-policy ("none") if no restart-policy was set.
		opts.params.HostConfig.RestartPolicy.Name = containertypes.RestartPolicyDisabled
	}

	warnings, err := daemon.verifyContainerSettings(daemonCfg, opts.params.HostConfig, opts.params.Config, false)
	if err != nil {
		return containertypes.CreateResponse{Warnings: warnings}, errdefs.InvalidParameter(err)
	}

	if opts.params.Platform == nil && opts.params.Config.Image != "" {
		img, err := daemon.imageService.GetImage(ctx, opts.params.Config.Image, backend.GetImageOpts{Platform: opts.params.Platform})
		if err != nil {
			return containertypes.CreateResponse{}, err
		}
		if img != nil {
			p := maximumSpec()
			imgPlat := ocispec.Platform{
				OS:           img.OS,
				Architecture: img.Architecture,
				Variant:      img.Variant,
			}

			if !images.OnlyPlatformWithFallback(p).Match(imgPlat) {
				warnings = append(warnings, fmt.Sprintf("The requested image's platform (%s) does not match the detected host platform (%s) and no specific platform was requested", platforms.Format(imgPlat), platforms.Format(p)))
			}
		}
	}

	err = daemon.validateNetworkingConfig(opts.params.NetworkingConfig)
	if err != nil {
		return containertypes.CreateResponse{Warnings: warnings}, errdefs.InvalidParameter(err)
	}

	if opts.params.HostConfig == nil {
		opts.params.HostConfig = &containertypes.HostConfig{}
	}
	err = daemon.adaptContainerSettings(&daemonCfg.Config, opts.params.HostConfig)
	if err != nil {
		return containertypes.CreateResponse{Warnings: warnings}, errdefs.InvalidParameter(err)
	}

	ctr, err := daemon.create(ctx, &daemonCfg.Config, opts)
	if err != nil {
		return containertypes.CreateResponse{Warnings: warnings}, err
	}
	containerActions.WithValues("create").UpdateSince(start)

	if warnings == nil {
		warnings = make([]string, 0) // Create an empty slice to avoid https://github.com/moby/moby/issues/38222
	}

	return containertypes.CreateResponse{ID: ctr.ID, Warnings: warnings}, nil
}

// Create creates a new container from the given configuration with a given name.
func (daemon *Daemon) create(ctx context.Context, daemonCfg *config.Config, opts createOpts) (retC *container.Container, retErr error) {
	var (
		ctr         *container.Container
		img         *image.Image
		imgManifest *ocispec.Descriptor
		imgID       image.ID
		err         error
		os          = runtime.GOOS
	)

	if opts.params.Config.Image != "" {
		img, err = daemon.imageService.GetImage(ctx, opts.params.Config.Image, backend.GetImageOpts{Platform: opts.params.Platform})
		if err != nil {
			return nil, err
		}
		// when using the containerd store, we need to get the actual
		// image manifest so we can store it and later deterministically
		// resolve the specific image the container is running
		if daemon.UsesSnapshotter() {
			imgManifest, err = daemon.imageService.GetImageManifest(ctx, opts.params.Config.Image, backend.GetImageOpts{Platform: opts.params.Platform})
			if err != nil {
				log.G(ctx).WithError(err).Error("failed to find image manifest")
				return nil, err
			}
		}
		os = img.OperatingSystem()
		imgID = img.ID()
	} else if isWindows {
		os = "linux" // 'scratch' case.
	}

	// On WCOW, if are not being invoked by the builder to create this container (where
	// ignoreImagesArgEscaped will be true) - if the image already has its arguments escaped,
	// ensure that this is replicated across to the created container to avoid double-escaping
	// of the arguments/command line when the runtime attempts to run the container.
	if os == "windows" && !opts.ignoreImagesArgsEscaped && img != nil && img.RunConfig().ArgsEscaped {
		opts.params.Config.ArgsEscaped = true
	}

	if err := daemon.mergeAndVerifyConfig(opts.params.Config, img); err != nil {
		return nil, errdefs.InvalidParameter(err)
	}

	if err := daemon.mergeAndVerifyLogConfig(&opts.params.HostConfig.LogConfig); err != nil {
		return nil, errdefs.InvalidParameter(err)
	}

	if ctr, err = daemon.newContainer(opts.params.Name, os, opts.params.Config, opts.params.HostConfig, imgID, opts.managed); err != nil {
		return nil, err
	}
	defer func() {
		if retErr != nil {
			err = daemon.cleanupContainer(ctr, backend.ContainerRmConfig{
				ForceRemove:  true,
				RemoveVolume: true,
			})
			if err != nil {
				log.G(ctx).WithError(err).Error("failed to cleanup container on create error")
			}
		}
	}()

	if err := daemon.setSecurityOptions(daemonCfg, ctr, opts.params.HostConfig); err != nil {
		return nil, err
	}

	ctr.HostConfig.StorageOpt = opts.params.HostConfig.StorageOpt
	ctr.ImageManifest = imgManifest

	initFunc := setupInitLayer(daemon.idMapping)
	if opts.params.HostConfig.Runtime == "bare" {
		initFunc = nil
	}

	if daemon.UsesSnapshotter() {
		if err := daemon.imageService.PrepareSnapshot(ctx, ctr.ID, opts.params.Config.Image, opts.params.Platform, initFunc); err != nil {
			return nil, err
		}
	} else {
		// Set RWLayer for container after mount labels have been set
		rwLayer, err := daemon.imageService.CreateLayer(ctr, initFunc)
		if err != nil {
			return nil, errdefs.System(err)
		}
		ctr.RWLayer = rwLayer
	}

	current := idtools.CurrentIdentity()
	if err := idtools.MkdirAndChown(ctr.Root, 0o710, idtools.Identity{UID: current.UID, GID: daemon.IdentityMapping().RootPair().GID}); err != nil {
		return nil, err
	}
	if err := idtools.MkdirAndChown(ctr.CheckpointDir(), 0o700, current); err != nil {
		return nil, err
	}

	if err := daemon.setHostConfig(ctr, opts.params.HostConfig, opts.params.DefaultReadOnlyNonRecursive); err != nil {
		return nil, err
	}

	if err := daemon.createContainerOSSpecificSettings(ctx, ctr, opts.params.Config, opts.params.HostConfig); err != nil {
		return nil, err
	}

	var endpointsConfigs map[string]*networktypes.EndpointSettings
	if opts.params.NetworkingConfig != nil {
		endpointsConfigs = opts.params.NetworkingConfig.EndpointsConfig
	}
	// Make sure NetworkMode has an acceptable value. We do this to ensure
	// backwards API compatibility.
	if ctr.HostConfig != nil && ctr.HostConfig.NetworkMode == "" {
		ctr.HostConfig.NetworkMode = networktypes.NetworkDefault
	}

	daemon.updateContainerNetworkSettings(ctr, endpointsConfigs)
	if err := daemon.Register(ctr); err != nil {
		return nil, err
	}
	stateCtr.set(ctr.ID, "stopped")
	daemon.LogContainerEvent(ctr, events.ActionCreate)
	return ctr, nil
}

func toHostConfigSelinuxLabels(labels []string) []string {
	for i, l := range labels {
		labels[i] = "label=" + l
	}
	return labels
}

func (daemon *Daemon) generateSecurityOpt(hostConfig *containertypes.HostConfig) ([]string, error) {
	for _, opt := range hostConfig.SecurityOpt {
		con := strings.Split(opt, "=")
		if con[0] == "label" {
			// Caller overrode SecurityOpts
			return nil, nil
		}
	}
	ipcMode := hostConfig.IpcMode
	pidMode := hostConfig.PidMode
	privileged := hostConfig.Privileged
	if ipcMode.IsHost() || pidMode.IsHost() || privileged {
		return toHostConfigSelinuxLabels(selinux.DisableSecOpt()), nil
	}

	var ipcLabel []string
	var pidLabel []string
	ipcContainer := ipcMode.Container()
	pidContainer := pidMode.Container()
	if ipcContainer != "" {
		c, err := daemon.GetContainer(ipcContainer)
		if err != nil {
			return nil, err
		}
		ipcLabel, err = selinux.DupSecOpt(c.ProcessLabel)
		if err != nil {
			return nil, err
		}
		if pidContainer == "" {
			return toHostConfigSelinuxLabels(ipcLabel), err
		}
	}
	if pidContainer != "" {
		c, err := daemon.GetContainer(pidContainer)
		if err != nil {
			return nil, err
		}

		pidLabel, err = selinux.DupSecOpt(c.ProcessLabel)
		if err != nil {
			return nil, err
		}
		if ipcContainer == "" {
			return toHostConfigSelinuxLabels(pidLabel), err
		}
	}

	if pidLabel != nil && ipcLabel != nil {
		for i := 0; i < len(pidLabel); i++ {
			if pidLabel[i] != ipcLabel[i] {
				return nil, fmt.Errorf("--ipc and --pid containers SELinux labels aren't the same")
			}
		}
		return toHostConfigSelinuxLabels(pidLabel), nil
	}
	return nil, nil
}

func (daemon *Daemon) mergeAndVerifyConfig(config *containertypes.Config, img *image.Image) error {
	if img != nil && img.Config != nil {
		if err := merge(config, img.Config); err != nil {
			return err
		}
	}
	// Reset the Entrypoint if it is [""]
	if len(config.Entrypoint) == 1 && config.Entrypoint[0] == "" {
		config.Entrypoint = nil
	}
	if len(config.Entrypoint) == 0 && len(config.Cmd) == 0 {
		return fmt.Errorf("no command specified")
	}
	return nil
}

// validateNetworkingConfig checks whether a container's NetworkingConfig is valid.
func (daemon *Daemon) validateNetworkingConfig(nwConfig *networktypes.NetworkingConfig) error {
	if nwConfig == nil {
		return nil
	}

	var errs []error
	for k, v := range nwConfig.EndpointsConfig {
		if v == nil {
			errs = append(errs, fmt.Errorf("invalid config for network %s: EndpointsConfig is nil", k))
			continue
		}

		// The referenced network k might not exist when the container is created, so just ignore the error in that case.
		nw, _ := daemon.FindNetwork(k)
		if err := validateEndpointSettings(nw, k, v); err != nil {
			errs = append(errs, fmt.Errorf("invalid config for network %s: %w", k, err))
		}
	}

	if len(errs) > 0 {
		return errdefs.InvalidParameter(multierror.Join(errs...))
	}

	return nil
}

// maximumSpec returns the distribution platform with maximum compatibility for the current node.
func maximumSpec() ocispec.Platform {
	p := platforms.DefaultSpec()
	if p.Architecture == "amd64" {
		p.Variant = archvariant.AMD64Variant()
	}
	return p
}

// DeltaCreate creates a delta of the specified src and dest images
// This is called directly from the Engine API
func (daemon *Daemon) DeltaCreate(deltaSrc, deltaDest string, outStream io.Writer) error {
	progressOutput := streamformatter.NewJSONProgressOutput(outStream, false)

	srcImg, err := daemon.GetImage(deltaSrc)
	if err != nil {
		return errors.Wrapf(err, "no such image: %s", deltaSrc)
	}

	dstImg, err := daemon.GetImage(deltaDest)
	if err != nil {
		return errors.Wrapf(err, "no such image: %s", deltaDest)
	}

	is := daemon.stores[dstImg.OperatingSystem()].imageStore
	ls := daemon.stores[dstImg.OperatingSystem()].layerStore

	srcData, err := is.GetTarSeekStream(srcImg.ID())
	if err != nil {
		return err
	}
	defer srcData.Close()

	srcDataLen, err := ioutils.SeekerSize(srcData)
	if err != nil {
		return err
	}

	progressReader := progress.NewProgressReader(srcData, progressOutput, srcDataLen, deltaSrc, "Fingerprinting")
	defer progressReader.Close()

	srcSig, err := librsync.Signature(bufio.NewReaderSize(progressReader, 65536), ioutil.Discard, 512, 32, librsync.BLAKE2_SIG_MAGIC)
	if err != nil {
		return err
	}

	progress.Update(progressOutput, deltaSrc, "Fingerprint complete")

	deltaRootFS := image.NewRootFS()

	for _, diffID := range dstImg.RootFS.DiffIDs {
		progress.Update(progressOutput, stringid.TruncateID(diffID.String()), "Waiting")
	}

	statTotalSize := int64(0)
	statDeltaSize := int64(0)

	for i, diffID := range dstImg.RootFS.DiffIDs {
		var (
			layerData io.Reader
			platform layer.OS
		)
		commonLayer := false

		// We're only interested in layers that are different. Push empty
		// layers for common layers
		if i < len(srcImg.RootFS.DiffIDs) && srcImg.RootFS.DiffIDs[i] == diffID {
			commonLayer = true
			layerData, _ = layer.EmptyLayer.TarStream()
			platform = layer.EmptyLayer.OS()
		} else {
			dstRootFS := *dstImg.RootFS
			dstRootFS.DiffIDs = dstRootFS.DiffIDs[:i+1]

			l, err := ls.Get(dstRootFS.ChainID())
			if err != nil {
				return err
			}
			defer layer.ReleaseAndLog(ls, l)

			platform = l.OS()

			input, err := l.TarStream()
			if err != nil {
				return err
			}
			defer input.Close()

			inputSize, err := l.DiffSize()
			if err != nil {
				return err
			}

			statTotalSize += inputSize

			progressReader := progress.NewProgressReader(input, progressOutput, inputSize, stringid.TruncateID(diffID.String()), "Computing delta")
			defer progressReader.Close()

			pR, pW := io.Pipe()

			layerData = pR

			tmpDelta, err := ioutil.TempFile("", "docker-delta-")
			if err != nil {
				return err
			}
			defer os.Remove(tmpDelta.Name())

			go func() {
				w := bufio.NewWriter(tmpDelta)
				err := librsync.Delta(srcSig, bufio.NewReader(progressReader), w)
				if err != nil {
					pW.CloseWithError(err)
					return
				}
				w.Flush()

				info, err := tmpDelta.Stat()
				if err != nil {
					pW.CloseWithError(err)
					return
				}

				tw := tar.NewWriter(pW)

				hdr := &tar.Header{
					Name: "delta",
					Mode: 0600,
					Size: info.Size(),
				}

				if err := tw.WriteHeader(hdr); err != nil {
					pW.CloseWithError(err)
					return
				}

				if _, err := tmpDelta.Seek(0, io.SeekStart); err != nil {
					pW.CloseWithError(err)
					return
				}

				if _, err := io.Copy(tw, tmpDelta); err != nil {
					pW.CloseWithError(err)
					return
				}

				if err := tw.Close(); err != nil {
					pW.CloseWithError(err)
					return
				}

				pW.Close()
			}()
		}

		newLayer, err := ls.Register(layerData, deltaRootFS.ChainID(), platform)
		if err != nil {
			return err
		}
		defer layer.ReleaseAndLog(ls, newLayer)

		if commonLayer {
			progress.Update(progressOutput, stringid.TruncateID(diffID.String()), "Skipping common layer")
		} else {
			deltaSize, err := newLayer.DiffSize()
			if err != nil {
				return err
			}
			statDeltaSize += deltaSize
			progress.Update(progressOutput, stringid.TruncateID(diffID.String()), "Delta complete")
		}

		deltaRootFS.Append(newLayer.DiffID())
	}

	config := image.Image{
		RootFS: deltaRootFS,
		V1Image: image.V1Image{
			Created: time.Now().UTC(),
			Config: &containertypes.Config{
				Labels: map[string]string{
					"io.resin.delta.base": srcImg.ID().String(),
					"io.resin.delta.config": string(dstImg.RawJSON()),
				},
			},
		},
	}

	rawConfig, err := json.Marshal(config)
	if err != nil {
		return err
	}

	id, err := is.Create(rawConfig)
	if err != nil {
		return err
	}

	humanTotal := units.HumanSize(float64(statTotalSize))
	humanDelta := units.HumanSize(float64(statDeltaSize))
	deltaRatio := float64(statTotalSize) / float64(statDeltaSize)
	if statTotalSize == 0 {
		deltaRatio = 1
	}

	outStream.Write(streamformatter.FormatStatus("", "Normal size: %s, Delta size: %s, %.2fx improvement", humanTotal, humanDelta, deltaRatio))
	outStream.Write(streamformatter.FormatStatus("", "Created delta: %s", id.String()))
	return nil
}
