package image

import (
	"bufio"
	"context"
	"crypto/sha256"
	"io"
	"io/ioutil"
	"reflect"
	"strings"
	"testing"

	"github.com/docker/docker/api/types"
	apiclient "github.com/docker/docker/client"
	"github.com/docker/docker/daemon/graphdriver/copy"
	"github.com/docker/docker/internal/test/fakecontext"
	"github.com/docker/docker/internal/test/registry"
	"gotest.tools/assert"
)

// TestDeltaCreate creates a delta and checks if it exists
func TestDeltaCreate(t *testing.T) {
	var (
		base   = "busybox:1.24"
		target = "busybox:1.29"
		delta  = "busybox:delta-1.24-1.29"
	)

	var (
		err    error
		rc     io.ReadCloser
		ctx    = context.Background()
		client = testEnv.APIClient()
	)

	pullBaseAndTargetImages(t, client, base, target)

	rc, err = client.ImageDelta(ctx,
		base,
		target,
		types.ImageDeltaOptions{
			Tag: delta,
		})
	assert.NilError(t, err)
	io.Copy(ioutil.Discard, rc)
	rc.Close()

	inspectDelta, _, err := client.ImageInspectWithRaw(ctx, delta)
	assert.NilError(t, err)

	inspectBase, _, err := client.ImageInspectWithRaw(ctx, base)
	assert.NilError(t, err)
	assert.Assert(t, inspectDelta.Config.Labels["io.resin.delta.base"] == inspectBase.ID)
}

// TestDeltaCreateDestinationLock triggers a delta generation job, waits for it
// to start and removes the destination image.
//
// Until 1e83de47e32ac0caf3b4e02094aeb76da28b90b3 this would remove unprocessed
// destination layers and break delta generation.
func TestDeltaCreateDestinationLock(t *testing.T) {
	var (
		base   = "debian:8"
		target = "debian:10"
		delta  = "debian:delta"
	)

	var (
		err    error
		rc     io.ReadCloser
		ctx    = context.Background()
		client = testEnv.APIClient()
	)

	pullBaseAndTargetImages(t, client, base, target)

	rc, err = client.ImageDelta(ctx,
		base,
		target,
		types.ImageDeltaOptions{
			Tag: delta,
		})
	assert.NilError(t, err)
	defer rc.Close()

	var (
		waitFingerprinting = make(chan struct{})
		waitDelta          = make(chan struct{})
	)
	go func() {
		defer close(waitFingerprinting)
		defer close(waitDelta)
		for br := bufio.NewReader(rc); ; {
			line, err := br.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					break
				}
				t.Fatal(err)
			}
			if strings.Contains(line, "Fingerprint complete") {
				waitFingerprinting <- struct{}{}
			}
			if strings.Contains(line, "layer does not exist") {
				t.Fail()
			}
		}
	}()

	inspectTarget, _, err := client.ImageInspectWithRaw(ctx, target)

	<-waitFingerprinting
	deleted, err := client.ImageRemove(ctx, target, types.ImageRemoveOptions{})
	assert.NilError(t, err)
	for _, item := range deleted {
		for i := 0; i < len(inspectTarget.RootFS.Layers); i++ {
			assert.Assert(t, item.Deleted != inspectTarget.RootFS.Layers[i], "deleted target image layer")
		}
	}

	<-waitDelta
	inspectDelta, _, err := client.ImageInspectWithRaw(ctx, delta)
	assert.NilError(t, err)

	inspectBase, _, err := client.ImageInspectWithRaw(ctx, base)
	assert.NilError(t, err)
	assert.Assert(t, inspectDelta.Config.Labels["io.resin.delta.base"] == inspectBase.ID)
}

// TestDeltaSizes checks if the sizes of generated deltas are within the
// expected margins. This test is designed to catch regressions in the delta
// sizes.
//
// The "expected margins" (wantRatio) were defined empirically so that they are
// close to the values we were getting by the time the case was created. This
// test logs the expected and obtained ratios, so that it is relatively easy for
// us to also check if we got any substantial improvements when working on delta
// improvements.
func TestDeltaSize(t *testing.T) {
	type recordedRatio struct {
		desc string
		want float64
		got  float64
	}

	allRatios := []recordedRatio{}

	testCases := []struct {
		base      string
		target    string
		wantRatio float64 // (targetSize / deltaSize) must be at least this much
	}{
		{
			base:      "image-001",
			target:    "image-002",
			wantRatio: 5.0,
		},
		{
			base:      "image-001",
			target:    "image-003",
			wantRatio: 2.5,
		},
		{
			base:      "image-004",
			target:    "image-005",
			wantRatio: 180.0,
		},
		{
			base:      "image-004",
			target:    "image-006",
			wantRatio: 4.0,
		},
	}

	client := testEnv.APIClient()
	ctx := context.Background()

	for _, tc := range testCases {
		desc := tc.base + "-" + tc.target
		base := fullImageName(tc.base)
		target := fullImageName(tc.target)

		t.Run(desc, func(t *testing.T) {
			defer buildImage(ctx, t, client, base)()
			defer buildImage(ctx, t, client, target)()

			delta := "delta-" + desc
			defer createDelta(ctx, t, client, base, target, delta)()

			gotRatio := queryDeltaRatio(ctx, t, client, target, delta)
			if gotRatio < tc.wantRatio {
				t.Errorf("Delta ratio too small: got %.2f, expected at least %.2f",
					gotRatio, tc.wantRatio)
			}

			allRatios = append(allRatios, recordedRatio{desc, tc.wantRatio, gotRatio})
		})
	}

	// Log all obtained ratios
	t.Log("-------------------------------------------------------------")
	t.Logf("%-24s%-14s%-14s", "Test case", "Want ratio", "Got ratio")
	t.Log("-------------------------------------------------------------")
	for _, r := range allRatios {
		t.Logf("%-24s%-14.2f%-14.2f", r.desc, r.want, r.got)
	}
	t.Log("-------------------------------------------------------------")
}

// TestDeltaCorrectness checks if applying a delta on a base image results in an
// image with the same contents as the original target image.
func TestDeltaCorrectness(t *testing.T) {
	defer setupTestRegistry(t)()

	testCases := []struct {
		base   string
		target string
	}{
		{
			base:   "image-001",
			target: "image-002",
		},
		{
			base:   "image-001",
			target: "image-003",
		},
		{
			base:   "image-004",
			target: "image-005",
		},
		{
			base:   "image-004",
			target: "image-006",
		},
	}

	client := testEnv.APIClient()
	ctx := context.Background()

	for _, tc := range testCases {
		desc := tc.base + "-" + tc.target
		t.Run(desc, func(t *testing.T) {
			base := fullImageName(tc.base)
			target := fullImageName(tc.target)
			delta := fullImageName("delta-" + tc.base + "-" + tc.target)

			// Build two images, create delta of them and push this delta.
			defer buildImage(ctx, t, client, base)()
			removeTarget := buildImage(ctx, t, client, target)
			removeDelta := createDelta(ctx, t, client, base, target, delta)
			pushImageToTestRegistry(ctx, t, client, delta)

			// The delta we have locally shall not be the same as the target image.
			targetHash := imageContentHash(ctx, t, client, target)
			deltaHash := imageContentHash(ctx, t, client, delta)
			assert.Assert(t, !reflect.DeepEqual(targetHash, deltaHash))

			// Remove the delta and the target image.
			removeTarget()
			removeDelta()

			// Pull the delta. This will cause it to be applied.
			pullImageFromTestRegistry(ctx, t, client, delta)

			// The (now applied) delta shall be the same the target image was.
			appliedDeltaHash := imageContentHash(ctx, t, client, delta)
			assert.Assert(t, reflect.DeepEqual(targetHash, appliedDeltaHash))
		})
	}
}

// buildImage builds and tags a given image and returns a function that can be
// used to remove this image.
func buildImage(ctx context.Context, t *testing.T, client apiclient.APIClient, image string) func() {
	// Make sure we support both "my-image" and "registry/my-image".
	imageSegs := strings.Split(image, "/")
	assert.Assert(t, len(imageSegs) >= 1 && len(imageSegs) <= 2)
	imageDir := imageSegs[len(imageSegs)-1]

	source := fakecontext.New(t, "")
	defer source.Close()

	copy.DirCopy("../testdata/delta/"+imageDir, source.Dir, copy.Content, true)
	resp, err := client.ImageBuild(ctx, source.AsTarReader(t),
		types.ImageBuildOptions{Tags: []string{image}},
	)
	assert.Assert(t, err)

	if resp.Body != nil {
		body, err := readAllAndClose(resp.Body)
		assert.Assert(t, err)
		assert.Assert(t, strings.Contains(body, "Successfully built"))
		assert.Assert(t, strings.Contains(body, "Successfully tagged"))
	}

	return func() {
		removeImage(ctx, t, client, image)
	}
}

// removeImage removes a given image.
func removeImage(ctx context.Context, t *testing.T, client apiclient.APIClient, image string) {
	resp, err := client.ImageRemove(ctx, image, types.ImageRemoveOptions{})
	assert.Assert(t, err)
	assert.Assert(t, len(resp) > 0)
}

// createDelta creates a delta between base and target, tagging is as delta. It
// returns a function that can be used to remove this delta.
func createDelta(ctx context.Context, t *testing.T, client apiclient.APIClient,
	base, target, delta string) func() {

	rc, err := client.ImageDelta(ctx, base, target, types.ImageDeltaOptions{Tag: delta})
	assert.Assert(t, err)

	if rc != nil {
		body, err := readAllAndClose(rc)
		assert.Assert(t, err)
		assert.Assert(t, strings.Contains(body, "Created delta"))
		assert.Assert(t, strings.Contains(body, "Successfully tagged"))
	}

	return func() {
		removeImage(ctx, t, client, delta)
	}
}

// pushImageToTestRegistry pushes a given image to the temporary test registry.
func pushImageToTestRegistry(ctx context.Context, t *testing.T, client apiclient.APIClient, image string) {
	rc, err := client.ImagePush(ctx, image, types.ImagePushOptions{RegistryAuth: "{}"})
	assert.Assert(t, err)
	if rc != nil {
		body, err := readAllAndClose(rc)
		assert.Assert(t, err)
		assert.Assert(t, strings.Contains(body, "Pushed"))
	}
}

// pullImageFromTestRegistry pushes a given image to the temporary test registry.
func pullImageFromTestRegistry(ctx context.Context, t *testing.T, client apiclient.APIClient, image string) {
	rc, err := client.ImagePull(ctx, image, types.ImagePullOptions{RegistryAuth: "{}"})
	assert.Assert(t, err)
	if rc != nil {
		body, err := readAllAndClose(rc)
		assert.Assert(t, err)
		assert.Assert(t, strings.Contains(body, "Status: Downloaded newer image"))
	}
}

// imageHash computes a hash based on the contents of a given image. This is
// done indirectly, relying on the layer IDs which are already
// "content-addressable".
func imageContentHash(ctx context.Context, t *testing.T, client apiclient.APIClient, image string) []byte {
	ii, _, err := client.ImageInspectWithRaw(ctx, image)
	assert.Assert(t, err)

	hash := sha256.New()
	for _, layerHash := range ii.RootFS.Layers {
		_, err = hash.Write([]byte(layerHash))
		assert.Assert(t, err)
	}
	return hash.Sum(nil)
}

// queryImageSize returns the size in bytes of image.
func queryImageSize(ctx context.Context, t *testing.T, client apiclient.APIClient,
	image string) int64 {

	ii, _, err := client.ImageInspectWithRaw(ctx, image)
	assert.Assert(t, err)
	return ii.Size
}

// queryDeltaRatio queries image sizes and returns how many times target is
// larger than delta.
func queryDeltaRatio(ctx context.Context, t *testing.T, client apiclient.APIClient,
	target, delta string) float64 {

	targetSize := queryImageSize(ctx, t, client, target)
	deltaSize := queryImageSize(ctx, t, client, delta)
	deltaRatio := float64(targetSize) / float64(deltaSize)
	if targetSize == 0 {
		deltaRatio = 1.0
	}
	return deltaRatio
}

func pullBaseAndTargetImages(t *testing.T, client apiclient.APIClient, base, target string) {
	var (
		err error
		rc  io.ReadCloser
		ctx = context.Background()
	)

	rc, err = client.ImagePull(ctx,
		base,
		types.ImagePullOptions{})
	assert.NilError(t, err)
	io.Copy(ioutil.Discard, rc)
	rc.Close()

	rc, err = client.ImagePull(ctx,
		target,
		types.ImagePullOptions{})
	assert.NilError(t, err)
	io.Copy(ioutil.Discard, rc)
	rc.Close()
}

// fullImageName returns the image name including the test registry.
func fullImageName(image string) string {
	return registry.DefaultURL + "/" + image
}

// readAllAndClose reads everything from r, closes it, and returns whatever was
// read as a string.
func readAllAndClose(rc io.ReadCloser) (string, error) {
	// TODO: Simplify this code once we adopt Go 1.16 or later. This version of
	// Go brought io.ReadAll(), of which the code below is a simple variation.
	b := make([]byte, 0, 512)
	for {
		if len(b) == cap(b) {
			// Add more capacity (let append pick how much).
			b = append(b, 0)[:len(b)]
		}
		n, err := rc.Read(b[len(b):cap(b)])
		b = b[:len(b)+n]
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			// This call to Close() and the conversion to string are the only
			// changes from io.ReadAll().
			closeErr := rc.Close()
			if err == nil {
				err = closeErr
			}
			return string(b), err
		}
	}
}
