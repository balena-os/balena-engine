package delta

import (
	"net/http"

	"github.com/docker/docker/api/server/httputils"
	"github.com/docker/docker/api/types"
	"golang.org/x/net/context"
)

func (d *deltaRouter) postDeltasCreate(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error {
	if err := httputils.ParseForm(r); err != nil {
		return err
	}

	deltaSrc := r.Form.Get("src")
	deltaDest := r.Form.Get("dest")

	imgID, err := d.backend.DeltaCreate(deltaSrc, deltaDest)
	if err != nil {
		return err
	}

	return httputils.WriteJSON(w, http.StatusCreated, &types.IDResponse{
		ID: string(imgID),
	})
}
