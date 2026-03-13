package delta // import "github.com/docker/docker/api/server/router/delta"

import (
	"io"

	"github.com/docker/docker/api/types"
)

// Backend is the methods that need to be implemented to provide
// delta specific functionality.
type Backend interface {
	DeltaCreate(deltaSrc, deltaDest string, options types.ImageDeltaOptions, outStream io.Writer) error
}
