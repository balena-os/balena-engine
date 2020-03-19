//go:build no_buildkit

package buildkit

import (
	"context"

	"github.com/containerd/log"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/backend"
	"github.com/docker/docker/builder"
	"google.golang.org/grpc"
)

type Opt struct {
	SessionManager      interface{}
	Root                interface{}
	EngineID            interface{}
	Dist                interface{}
	ImageTagger         interface{}
	NetworkController   interface{}
	DefaultCgroupParent interface{}
	RegistryHosts       interface{}
	BuilderConfig       interface{}
	Rootless            interface{}
	IdentityMapping     interface{}
	DNSConfig           interface{}
	ApparmorProfile     interface{}
	UseSnapshotter      interface{}
	Snapshotter         interface{}
	ContainerdAddress   interface{}
	ContainerdNamespace interface{}
}

type Builder struct{}

func New(ctx context.Context, opt Opt) (*Builder, error) {
	log.G(ctx).Debug("buildkit isn't supported")
	return &Builder{}, nil
}

func (b *Builder) Close() error {
	return nil
}

func (b *Builder) RegisterGRPC(s *grpc.Server) {
}

func (b *Builder) Build(ctx context.Context, opt backend.BuildConfig) (*builder.Result, error) {
	log.G(ctx).Warn("buildkit isn't supported")
	panic("buildkit isn't supported")
}

func (b *Builder) Prune(ctx context.Context, opts types.BuildCachePruneOptions) (int64, []string, error) {
	return 0, []string{}, nil
}

func (b *Builder) Cancel(ctx context.Context, id string) error {
	return nil
}

func (b *Builder) DiskUsage(ctx context.Context) ([]*types.BuildCache, error) {
	return make([]*types.BuildCache, 0), nil
}
