package image // import "github.com/docker/docker/integration/image"

import (
	"archive/tar"
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"os"
	"runtime"
	"strings"
	"testing"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/testutil"
	"github.com/docker/docker/testutil/daemon"
	"gotest.tools/v3/assert"
	"gotest.tools/v3/skip"
)

// Ensure we don't regress on CVE-2017-14992.
func TestImportExtremelyLargeImageWorks(t *testing.T) {
	skip.If(t, testEnv.IsRemoteDaemon, "cannot run daemon when remote daemon")
	skip.If(t, runtime.GOARCH == "arm64", "effective test will be time out")
	skip.If(t, testEnv.OSType == "windows", "TODO enable on windows")
	t.Parallel()

	// Spin up a new daemon, so that we can run this test in parallel (it's a slow test)
	d := daemon.New(t)
	d.Start(t)
	defer d.Stop(t)

	client := d.NewClientT(t)

	// Construct an empty tar archive with about 8GB of junk padding at the
	// end. This should not cause any crashes (the padding should be mostly
	// ignored).
	var tarBuffer bytes.Buffer

	tw := tar.NewWriter(&tarBuffer)
	err := tw.Close()
	assert.NilError(t, err)
	imageRdr := io.MultiReader(&tarBuffer, io.LimitReader(testutil.DevZero, 8*1024*1024*1024))
	reference := strings.ToLower(t.Name()) + ":v42"

	_, err = client.ImageImport(context.Background(),
		types.ImageImportSource{Source: imageRdr, SourceName: "-"},
		reference,
		types.ImageImportOptions{})
	assert.NilError(t, err)
}

// Test if we properly apply balena deltas when doing an `image load`.
func TestDeltaOnImageLoad(t *testing.T) {
	const (
		base   = "busybox:1.24"
		target = "busybox:1.29"
		delta  = "busybox:delta-1.24-1.29"
	)

	ctx := context.Background()
	client := testEnv.APIClient()

	// Pull base and target images, get the ID of the target image
	pullBaseAndTargetImages(t, client, base, target)
	targetImageInfo, _, err := client.ImageInspectWithRaw(ctx, target)
	targetImageID := targetImageInfo.ID

	// Generate the delta between them
	rc, err := client.ImageDelta(ctx,
		base,
		target,
		types.ImageDeltaOptions{
			Tag: delta,
		})
	assert.NilError(t, err)
	io.Copy(ioutil.Discard, rc)
	err = rc.Close()
	assert.NilError(t, err)

	// Export the delta image to a tar file
	rc, err = client.ImageSave(ctx, []string{delta})
	assert.NilError(t, err)
	defer rc.Close()
	deltaImageFile, err := os.CreateTemp("", "deltaImageFile-*.tar")
	assert.NilError(t, err)
	defer func() {
		deltaImageFile.Close()
		os.Remove(deltaImageFile.Name())
	}()
	io.Copy(deltaImageFile, rc)

	// Remove the target and delta images, make sure they are gone
	_, err = client.ImageRemove(ctx, target, types.ImageRemoveOptions{})
	assert.NilError(t, err)
	_, err = client.ImageRemove(ctx, delta, types.ImageRemoveOptions{})
	assert.NilError(t, err)
	_, _, err = client.ImageInspectWithRaw(ctx, target)
	assert.Error(t, err, "Error: No such image: "+target)
	_, _, err = client.ImageInspectWithRaw(ctx, delta)
	assert.Error(t, err, "Error: No such image: "+delta)

	// Load the exported delta image from the tar file
	loadedDeltaImageFile, err := os.Open(deltaImageFile.Name())
	assert.NilError(t, err)
	_, err = client.ImageLoad(ctx, loadedDeltaImageFile, true)
	assert.NilError(t, err)

	// Check if the ID of the delta-imported image matches the target image's
	restoredImageInfo, _, err := client.ImageInspectWithRaw(ctx, delta)
	assert.Equal(t, targetImageID, restoredImageInfo.ID)
}

// TODO(LMB): Could add other test cases, especially for the unhappy paths. Like
// checking we get a reasonable error if the base image doesn't exist. Or if
// nothing breaks if we tryu to load an image that is already present.
