package storagemigration

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"

	"github.com/docker/docker/internal/test/daemon"

	"gotest.tools/assert"
	"gotest.tools/fs"
	"gotest.tools/skip"
)

func TestAufsToOverlay2Migration(t *testing.T) {
	skip.If(t, testEnv.DaemonInfo.OSType != "linux")
	skip.If(t, testEnv.DaemonInfo.Driver != "overlay2")
	defer setupTest(t)()

	var err error

	root := fs.NewDir(t, t.Name())
	defer root.Remove()

	{
		// aufs.tar.gz contains a snapshot of /var/lib/docker after
		// building testdata/Dockerfile using dockerd which uses aufs
		// as default storage driver
		tar := exec.Command("tar", "-xzf", filepath.Join("testdata", "aufs.tar.gz"), "-C", root.Path())
		tar.Stdout = os.Stdout
		tar.Stderr = os.Stderr
		assert.NilError(t, tar.Run())
	}

	err = os.Setenv("BALENA_MIGRATE_OVERLAY", "1")
	assert.NilError(t, err)

	d := daemon.New(t)
	d.Root = root.Path()
	d.Start(t)
	defer d.Stop(t)

	ctx := context.Background()

	cl := d.NewClientT(t)

	ctr, err := cl.ContainerCreate(ctx,
		&container.Config{
			Image: "a2o-test",
		},
		nil,
		nil,
		"",
	)
	assert.NilError(t, err)

	err = cl.ContainerStart(ctx, ctr.ID, types.ContainerStartOptions{})
	assert.NilError(t, err)

	// original f1 should be removed (.wh.)
	_, err = cl.ContainerStatPath(ctx, ctr.ID, "/tmp/f1")
	assert.ErrorContains(t, err, "No such container:path")
	// original d1 should be opaque (.wh..wh.opq)
	_, err = cl.ContainerStatPath(ctx, ctr.ID, "/tmp/d1/d1f2")
	assert.NilError(t, err)
	_, err = cl.ContainerStatPath(ctx, ctr.ID, "/tmp/hlkn")
	assert.NilError(t, err)
	_, err = cl.ContainerStatPath(ctx, ctr.ID, "/tmp/slkn")
	assert.NilError(t, err)
}
