package image // import "github.com/docker/docker/integration/image"

import (
	"archive/tar"
	"bytes"
	"context"
	"io"
	"os"
	"runtime"
	"strings"
	"testing"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/testutil"
	"github.com/docker/docker/testutil/daemon"
	"gotest.tools/v3/assert"
	"gotest.tools/v3/assert/cmp"
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
func TestDeltaOnImageLoadHappyPath(t *testing.T) {
	d := daemon.New(t)
	d.Start(t)
	defer d.Stop(t)
	ctx := context.Background()
	client := d.NewClientT(t)

	// Load the basis image.
	loadDeltaOnLoadTestTarAsserting(t, ctx, client, "busybox-1.24")

	// Load (and apply) the delta.
	respBody, err := loadDeltaOnLoadTestTar(t, ctx, client, "busybox-delta")

	// The load operation needs to succeed, and the response body must
	// acknowledge that the image was loaded.
	assert.NilError(t, err)
	assert.Assert(t, cmp.Contains(respBody, `"stream":`))
	assert.Assert(t, cmp.Contains(respBody, `"Loaded image: busybox:delta`))

	// Check if the ID of the delta-imported image matches the target image's.
	restoredImageInfo, _, err := client.ImageInspectWithRaw(ctx, "busybox:delta")
	assert.NilError(t, err)
	assert.Equal(t, "sha256:758ec7f3a1ee85f8f08399b55641bfb13e8c1109287ddc5e22b68c3d653152ee", restoredImageInfo.ID)
}

// Test if we get the expected error when using `image load` for a balena delta
// whose basis image is missing locally.
func TestDeltaOnImageLoadBasisNotFound(t *testing.T) {
	d := daemon.New(t)
	d.Start(t)
	defer d.Stop(t)
	ctx := context.Background()
	client := d.NewClientT(t)

	// Load (and apply) the delta; notice we didn't load the basis image before.
	respBody, err := loadDeltaOnLoadTestTar(t, ctx, client, "busybox-delta")

	// The loading itself should complete without errors, but the response body
	// shall indicate that the base image wasn't found.
	assert.NilError(t, err)
	assert.Assert(t, cmp.Contains(respBody, `"errorDetail":`))
	assert.Assert(t, cmp.Contains(respBody, "failed to get digest sha256:47bcc53f74dc94b1920f0b34f6036096526296767650f223433fe65c35f149eb"))
	assert.Assert(t, cmp.Contains(respBody, "loading delta base image"))
}

// Test the case in which the target image already exists locally (and therefore
// the delta image doesn't need to be loaded).
func TestDeltaOnImageLoadTargetAlreadyExists(t *testing.T) {
	d := daemon.New(t)
	d.Start(t)
	defer d.Stop(t)
	ctx := context.Background()
	client := d.NewClientT(t)

	// Load the basis and the target image.
	loadDeltaOnLoadTestTarAsserting(t, ctx, client, "busybox-1.24")
	loadDeltaOnLoadTestTarAsserting(t, ctx, client, "busybox-1.29")

	// Load (and apply) the delta.
	respBody, err := loadDeltaOnLoadTestTar(t, ctx, client, "busybox-delta")

	// The loading shall succeed, but we also expect a message telling that the
	// target image already exists locally.
	assert.NilError(t, err)
	assert.Assert(t, cmp.Contains(respBody, `"stream":`))
	assert.Assert(t, cmp.Contains(respBody, "Target image already exists locally, no need to load it"))
}

// Test the case in which we try to load a file that is not a valid tar file.
// Not a "delta on load" test, strictly speaking, but exercises code we have
// touched.
func TestDeltaOnImageLoadInvalidTar(t *testing.T) {
	ctx := context.Background()
	client := testEnv.APIClient()

	brokenTarFileName := ""

	// Create a file that is not really a tar.
	func() {
		brokenTarFile, err := os.CreateTemp("", "brokenTarFile-*.tar")
		assert.NilError(t, err)
		defer func() {
			brokenTarFile.Close()
		}()
		brokenTarFile.Write([]byte("this is definitely not a valid tar file!"))
		brokenTarFileName = brokenTarFile.Name()
	}()

	// Load the broken tar file. The load operation itself shall succeed, but we
	// expect an error reported in the response body.
	loadedBrokenTarFile, err := os.Open(brokenTarFileName)
	assert.NilError(t, err)
	defer loadedBrokenTarFile.Close()
	resp, err := client.ImageLoad(ctx, loadedBrokenTarFile, true)
	assert.NilError(t, err)
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	assert.NilError(t, err)
	assert.Assert(t, cmp.Contains(string(respBody), `"errorDetail":`))
}

// Test the case in which we try to load a file that is a valid tar file, but
// contains inside it a delta that ends abruptly.
func TestDeltaOnImageLoadTruncatedDelta(t *testing.T) {
	d := daemon.New(t)
	d.Start(t)
	defer d.Stop(t)
	ctx := context.Background()
	client := d.NewClientT(t)

	loadDeltaOnLoadTestTarAsserting(t, ctx, client, "busybox-1.24")

	// Load the tar file with a truncated delta. The load operation itself shall
	// succeed, but we expect a meaningful error reported in the response body.
	respBody, err := loadDeltaOnLoadTestTar(t, ctx, client, "busybox-delta-truncated")
	assert.NilError(t, err)
	assert.Assert(t, cmp.Contains(string(respBody), `"errorDetail":`))
	assert.Assert(t, cmp.Contains(string(respBody), `unexpected EOF`))
}

// Test the case in which we try to load a file that is a valid tar file, but
// contains a broken delta inside it (specifically, the delta was manually
// edited so that it has an incorrect "magic number" at its beginning).
func TestDeltaOnImageLoadInvalidDelta(t *testing.T) {
	d := daemon.New(t)
	d.Start(t)
	defer d.Stop(t)
	ctx := context.Background()
	client := d.NewClientT(t)

	loadDeltaOnLoadTestTarAsserting(t, ctx, client, "busybox-1.24")

	// Load the tar file with a broken delta. The load operation itself shall
	// succeed, but we expect a meaningful error reported in the response body.
	respBody, err := loadDeltaOnLoadTestTar(t, ctx, client, "busybox-delta-bad-magic")
	assert.NilError(t, err)
	assert.Assert(t, cmp.Contains(string(respBody), `"errorDetail":`))
	assert.Assert(t, cmp.Contains(string(respBody), `Got magic number baaaaaad rather than expected value 72730236`))
}

//
// balena "delta on load" helpers
//

// loadDeltaOnLoadTestTar loads an image from one of the test tar files included
// in the delta-on-load test data. The name argument is the name of the desired
// tar file, without path nor extension. Returns the response body of the image
// load operation, and an error.
func loadDeltaOnLoadTestTar(t *testing.T, ctx context.Context, client client.APIClient, name string) (string, error) {
	tarFile, err := os.Open("../../integration/testdata/delta-on-load/" + name + ".tar.gz")
	assert.NilError(t, err)
	defer tarFile.Close()

	resp, err := client.ImageLoad(ctx, tarFile, true)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(respBody), nil
}

// loadDeltaOnLoadTestTarAsserting is similar to loadDeltaOnLoadTestTar, but it
// does all asserts internally to make sure the tar file was successfully
// loaded.
func loadDeltaOnLoadTestTarAsserting(t *testing.T, ctx context.Context, client client.APIClient, name string) {
	respBody, err := loadDeltaOnLoadTestTar(t, ctx, client, name)
	assert.NilError(t, err)
	assert.Assert(t, cmp.Contains(string(respBody), `"stream":`))
	assert.Assert(t, cmp.Contains(string(respBody), `"Loaded image:`))
}
