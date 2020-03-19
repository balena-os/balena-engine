// +build no_buildkit

package buildkit

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/backend"
	"github.com/docker/docker/builder"
	"github.com/sirupsen/logrus"
)

type Opt struct {
	SessionManager      interface{}
	Root                interface{}
	Dist                interface{}
	NetworkController   interface{}
	DefaultCgroupParent interface{}
	ResolverOpt         interface{}
	BuilderConfig       interface{}
	Rootless            interface{}
	IdentityMapping     interface{}
	DNSConfig           interface{}
}

type Builder struct{}

func New(opt Opt) (*Builder, error) {
	logrus.Debug("buildkit isn't supported")
	return &Builder{}, nil
}

func (b *Builder) RegisterGRPC(s interface{}) {
	logrus.Debug("buildkit isn't supported: noop")
	return
}

func (b *Builder) Build(ctx context.Context, opt backend.BuildConfig) (*builder.Result, error) {
	logrus.Warning("buildkit isn't supported")
	panic("buildkit isn't supported")
}

func (b *Builder) Prune(ctx context.Context, opts types.BuildCachePruneOptions) (int64, []string, error) {
	logrus.Warning("buildkit isn't supported: noop")
	return 0, []string{}, nil
}

func (b *Builder) Cancel(ctx context.Context, id string) error {
	logrus.Debug("buildkit isn't supported: noop")
	return nil
}

func (b *Builder) DiskUsage(ctx context.Context) ([]*types.BuildCache, error) {
	logrus.Debug("buildkit isn't supported: noop")
	var bc = make([]*types.BuildCache, 0)
	return bc, nil
}
