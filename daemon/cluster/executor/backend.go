package executor // import "github.com/docker/docker/daemon/cluster/executor"

import (
	"context"
	"io"

	"github.com/distribution/reference"
	"github.com/docker/distribution"
	"github.com/docker/docker/api/types/backend"
	"github.com/docker/docker/api/types/registry"
	"github.com/docker/docker/image"
	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
)

// ImageBackend is used by an executor to perform image operations
type ImageBackend interface {
	PullImage(ctx context.Context, ref reference.Named, platform *ocispec.Platform, metaHeaders map[string][]string, authConfig *registry.AuthConfig, outStream io.Writer) error
	GetRepositories(context.Context, reference.Named, *registry.AuthConfig) ([]distribution.Repository, error)
	GetImage(ctx context.Context, refOrID string, options backend.GetImageOpts) (*image.Image, error)
}
