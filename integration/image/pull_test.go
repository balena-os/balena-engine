package image

import (
	"context"
	"fmt"
	"testing"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/internal/test/daemon"

	"gotest.tools/assert"
	"gotest.tools/skip"
)

func TestImagePullSyncDiffs(t *testing.T) {
	skip.If(t, testEnv.IsRemoteDaemon())

	var testImage = "balenalib/amd64-debian:build" // should have a few layers so it actually has an impact on pull performance

	for _, storageDriver := range []string{"aufs", "overlay2"} {
		for _, syncDiffs := range []bool{true, false} {
			t.Run(fmt.Sprintf("storageDriver=%v,syncDiffs=%v", storageDriver, syncDiffs), func(t *testing.T) {

				skip.If(t, storageDriver == "aufs", "Aufs doesn't work with dind")

				var args = []string{
					fmt.Sprintf("--storage-driver=%v", storageDriver),
					fmt.Sprintf("--storage-opt=%v.sync_diffs=%v", storageDriver, syncDiffs),
				}

				d := daemon.New(t)

				d.Start(t, args...)
				defer d.Stop(t)

				info := d.Info(t)
				assert.Equal(t, info.Driver, storageDriver)

				client, err := d.NewClient()
				assert.NilError(t, err)

				ctx := context.Background()
				_, err = client.ImagePull(ctx, testImage, types.ImagePullOptions{})
				assert.NilError(t, err)
			})
		}
	}
}
