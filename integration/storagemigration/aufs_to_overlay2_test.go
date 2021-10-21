package storagemigration

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"

	"github.com/docker/docker/testutil/daemon"

	"golang.org/x/sys/unix"
	"gotest.tools/v3/assert"
	"gotest.tools/v3/fs"
	"gotest.tools/v3/skip"
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

		// FIXME(robertgzr): come up with something better than hardcoding this here
		sockpath := root.Join("aufs/diff/34999927091de7ae41baa85f0be576c4b70ff81bf30d6850bf84c16d5d8cb9f5/run/udev/control")
		assert.NilError(t, os.MkdirAll(filepath.Dir(sockpath), 0666))
		// create a socket
		assert.NilError(t, unix.Mknod(sockpath, 0755|unix.S_IFSOCK, 0))
	}

	err = os.Setenv("BALENA_MIGRATE_OVERLAY", "1")
	assert.NilError(t, err)

	d := daemon.New(t)
	d.Root = root.Path()
	d.Start(t)
	defer d.Stop(t)

	ctx := context.Background()

	cl := d.NewClientT(t)

	info, err := cl.Info(ctx)
	assert.NilError(t, err)
	assert.Equal(t, info.Driver, "overlay2")

	images, err := cl.ImageList(ctx, types.ImageListOptions{})
	assert.NilError(t, err)
	assert.Equal(t, len(images), 2)

	containers, err := cl.ContainerList(ctx, types.ContainerListOptions{All: true})
	assert.NilError(t, err)
	assert.Equal(t, len(containers), 0)

	ctr, err := cl.ContainerCreate(ctx,
		&container.Config{
			Image: "a2o-test",
		},
		nil,
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
