package image

import (
	"context"
	"fmt"
	"testing"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/versions"
	"github.com/docker/docker/errdefs"
	"github.com/docker/docker/testutil/daemon"

	"gotest.tools/v3/assert"
	"gotest.tools/v3/skip"
)

func TestImagePullPlatformInvalid(t *testing.T) {
	skip.If(t, versions.LessThan(testEnv.DaemonAPIVersion(), "1.40"), "experimental in older versions")
	defer setupTest(t)()
	client := testEnv.APIClient()
	ctx := context.Background()

	_, err := client.ImagePull(ctx, "docker.io/library/hello-world:latest", types.ImagePullOptions{Platform: "foobar"})
	assert.Assert(t, err != nil)
	assert.ErrorContains(t, err, "unknown operating system or architecture")
	assert.Assert(t, errdefs.IsInvalidParameter(err))
}

func TestImagePullNoSyncDiffsOverlay2(t *testing.T) {
	skip.If(t, testEnv.IsRemoteDaemon)

	testPullWithSyncDiffs(t, false)
}

func testPullWithSyncDiffs(t *testing.T, syncDiffs bool) {
	// should have a few layers so it actually has an impact on pull performance
	testImage := "balenalib/amd64-debian:build"

	var args = []string{
		"--storage-driver=overlay2",
		fmt.Sprintf("--storage-opt=overlay2.sync_diffs=%t", syncDiffs),
	}

	d := daemon.New(t)
	d.Start(t, args...)
	defer d.Stop(t)

	client := d.NewClientT(t)
	_, err := client.ImagePull(context.Background(), testImage, types.ImagePullOptions{})
	assert.NilError(t, err)
}
