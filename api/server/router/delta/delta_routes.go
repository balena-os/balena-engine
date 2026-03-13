package delta // import "github.com/docker/docker/api/server/router/delta"

import (
	"context"
	"net/http"

	"github.com/docker/docker/api/server/httputils"
	"github.com/docker/docker/api/types"
)

func (d *deltaRouter) postDeltasCreate(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error {
	if err := httputils.ParseForm(r); err != nil {
		return err
	}

	options := types.ImageDeltaOptions{
		SrcName:  r.Form.Get("src"),
		DestName: r.Form.Get("dest"),
		Tag:      r.Form.Get("t"),
	}

	w.Header().Set("Content-Type", "application/json")

	return d.backend.DeltaCreate(options.SrcName, options.DestName, options, w)
}
