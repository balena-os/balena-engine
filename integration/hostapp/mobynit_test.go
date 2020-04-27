package hostapp

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/cmd/mobynit/hostapp"
	"github.com/docker/docker/testutil/daemon"

	"gotest.tools/v3/assert"
	"gotest.tools/v3/skip"
)

func TestMobynitMountContainer(t *testing.T) {
	skip.If(t, testEnv.DaemonInfo.OSType != "linux")
	skip.If(t, testEnv.IsRemoteDaemon, "cannot start daemon on remote test run")
	defer setupTest(t)()

	d := daemon.New(t)
	d.StartWithBusybox(t)
	defer d.Stop(t)

	client := d.NewClientT(t)

	c, err := client.ContainerCreate(context.Background(),
		&container.Config{Image: "busybox:latest"},
		&container.HostConfig{Runtime: "bare"},
		&network.NetworkingConfig{},
		nil,
		"",
	)
	assert.NilError(t, err)

	var (
		layerRoot     = d.RootDir()
		storageDriver = d.StorageDriver()
		containerID   = c.ID
	)
	newRootPath := hostapp.MountContainer(layerRoot, containerID, storageDriver)

	// give the daemon's layer store cleanup goroutine some time to run
	// TODO probably use gotest.tools/poll
	time.Sleep(1 * time.Second)

	// only checks if the target mount point exists
	fi, err := os.Stat(newRootPath)
	assert.NilError(t, err)
	assert.Check(t, fi.IsDir())
}
