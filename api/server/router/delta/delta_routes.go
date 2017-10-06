package delta

import (
	"net/http"

	"github.com/docker/docker/api/server/httputils"
	"github.com/docker/docker/pkg/ioutils"
	"github.com/docker/docker/pkg/streamformatter"
	"golang.org/x/net/context"
)

func (d *deltaRouter) postDeltasCreate(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error {
	if err := httputils.ParseForm(r); err != nil {
		return err
	}

	deltaSrc := r.Form.Get("src")
	deltaDest := r.Form.Get("dest")

	output := ioutils.NewWriteFlusher(w)
	defer output.Close()

	w.Header().Set("Content-Type", "application/json")

	if err := d.backend.DeltaCreate(deltaSrc, deltaDest, output); err != nil {
		if !output.Flushed() {
			return err
		}
		output.Write(streamformatter.FormatError(err))
	}
	return nil
}
