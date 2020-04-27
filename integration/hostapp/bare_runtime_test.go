package hostapp

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/internal/test/request"

	"gotest.tools/assert"
	"gotest.tools/skip"
)

func TestBareRuntime(t *testing.T) {
	skip.If(t, testEnv.DaemonInfo.OSType != "linux")
	skip.If(t, testEnv.IsRemoteDaemon, "cannot start daemon on remote test run")
	defer setupTest(t)()

	client := request.NewAPIClient(t)
	ctx := context.Background()

	c, err := client.ContainerCreate(ctx,
		&container.Config{Image: "busybox:latest"},
		&container.HostConfig{Runtime: "bare"},
		&network.NetworkingConfig{},
		"",
	)
	assert.NilError(t, err)

	j, err := client.ContainerInspect(ctx, c.ID)
	assert.NilError(t, err)

	containerDir, ok := j.GraphDriver.Data["MergedDir"]
	assert.Check(t, ok)

	_, err = os.Stat(filepath.Join(containerDir, ".dockerenv"))
	assert.Check(t, os.IsNotExist(err))
}
