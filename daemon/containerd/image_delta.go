package containerd

import (
	"errors"
	"io"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/errdefs"
)

// DeltaCreate creates a delta of the specified src and dest images.
func (i *ImageService) DeltaCreate(deltaSrc, deltaDest string, options types.ImageDeltaOptions, outStream io.Writer) error {
	return errdefs.NotImplemented(errors.New("not implemented"))
}
