package image

import (
	"archive/tar"
	"bufio"
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/docker/docker/api/types"
	apiclient "github.com/docker/docker/client"
	"github.com/docker/docker/daemon/graphdriver/copy"
	"github.com/docker/docker/testutil/daemon"
	"github.com/docker/docker/testutil/fakecontext"
	"github.com/docker/docker/testutil/registry"
	"gotest.tools/v3/assert"
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

// deltaTestCases contains all the test cases we use to test delta sizes and
// correctness.
var deltaTestCases = []struct {
	// All images to build (including the base and target images), in the
	// desired build order. Required only when the base and target images depend
	// on still other images (which must be built before). If nil, the test code
	// will build only the base and target images.
	images []string

	// Images to delete locally before pulling (including the target image). If
	// nil, the test code will delete only the target image.
	removeBeforePull []string

	// Image to use as the delta base.
	base string

	// Image to use as the delta target.
	target string

	// The interval of delta sizes we consider acceptable. We are using an
	// interval here because the image sizes (and therefore the delta sizes)
	// vary depending on the exact union filesystem implementation used on the
	// machine running the tests. The absolute and relative differences on image
	// sizes vary a lot depending on factors that are specific to each image, so
	// we can't simply use a tolerance factor.
	wantSizeMin int64
	wantSizeMax int64
}{
	// Image 000 contains a single layer with a single empty file on it. 001
	// extends it adding a layer with another empty file.
	{
		base:        "000",
		target:      "001",
		wantSizeMin: 1044,
		wantSizeMax: 1556,
	},

	// 002 extends 000 adding a layer with a 256-byte file.
	{
		base:        "000",
		target:      "002",
		wantSizeMin: 1045,
		wantSizeMax: 1045,
	},

	// 003 extends 001, adding four layers, each with a new empty file (and one
	// of them overwrites the empty file that was in 001 with a 256-byte file).
	{
		images:      []string{"000", "001", "003"},
		base:        "001",
		target:      "003",
		wantSizeMin: 4689,
		wantSizeMax: 5201,
	},

	// Similar to the case above, but we diff between 000 and 003. (003 extends
	// 001, which extends 000.)
	{
		images:      []string{"000", "001", "003"},
		base:        "000",
		target:      "003",
		wantSizeMin: 5733,
		wantSizeMax: 6757,
	},

	// 004 and 005 equal. They are both created from scratch and contain a
	// number of files. Some layers contain a single file, others contain
	// multiple files.
	{
		base:        "004",
		target:      "005",
		wantSizeMin: 0,
		wantSizeMax: 0,
	},

	// 006 and 007 are both created from scratch, both contain the same files,
	// one per layer, but the layers are added in different orders.
	{
		base:        "006",
		target:      "007",
		wantSizeMin: 381,
		wantSizeMax: 381,
	},

	// 008 has the same files as 006, but all on a single layer.
	{
		base:        "006",
		target:      "008",
		wantSizeMin: 115,
		wantSizeMax: 115,
	},
	{
		base:        "008",
		target:      "006",
		wantSizeMin: 381,
		wantSizeMax: 381,
	},

	// This simulates a common case: a user creating a new image with a new
	// version of their app. 009 is a 3.7MB image representing one of those
	// large base images provided by vendors. 010 represents the old version of
	// the user app: it extends 009 by adding 161kB of new files to a new layer.
	// 011 represents the new version of the user app: it also extends 009 and
	// also adds a layer of 161kB -- only 1kB of which is different from those
	// in 010.
	{
		images:      []string{"009", "010", "011"},
		base:        "010",
		target:      "011",
		wantSizeMin: 1567,
		wantSizeMax: 1567,
	},

	// This simulates the case in which the large base image itself was
	// upgraded, but the user code remains the same. 009 is the old
	// vendor-provided base image; 012 is the new one: there's a 1MB file
	// changed from 009. 010 (based on 009) and 013 (based on 012) are the
	// customer app images: the only difference between them is the base image.
	{
		images:           []string{"009", "010", "012", "013"},
		base:             "010",
		target:           "012",
		removeBeforePull: []string{"012", "013"},
		wantSizeMin:      1049125,
		wantSizeMax:      1049132,
	},
}

// TestDeltaSize checks if the sizes of generated deltas have not increased. In
// other words, this test is designed to catch regressions in the delta sizes.
//
// The expected sizes (wantSize) were defined empirically so that they match
// the values we were getting by the time the case was created. This test logs
// the expected and obtained sizes, so that it is relatively easy for us to also
// check if we got any gains when working on delta improvements.
func TestDeltaSize(t *testing.T) {
	type recordedSize struct {
		desc    string
		wantMin int64
		wantMax int64
		got     int64
	}

	allSizes := []recordedSize{}

	client := testEnv.APIClient()
	ctx := context.Background()

	for _, tc := range deltaTestCases {
		delta := deltaName(tc.base, tc.target)
		t.Run(delta, func(t *testing.T) {
			// Build all required images
			if tc.images != nil {
				for _, image := range tc.images {
					defer ttrBuildImageAsserting(ctx, t, client, image)()
				}
			} else {
				defer ttrBuildImageAsserting(ctx, t, client, tc.base)()
				defer ttrBuildImageAsserting(ctx, t, client, tc.target)()
			}

			// Create the delta, check its size
			defer ttrCreateDeltaAsserting(ctx, t, client, tc.base, tc.target)()
			gotSize := ttrQueryDeltaSizeAsserting(ctx, t, client, delta)
			if gotSize > tc.wantSizeMax {
				t.Errorf("Delta too big: got %v bytes, expected at most %v",
					gotSize, tc.wantSizeMax)
			}

			allSizes = append(allSizes, recordedSize{delta, tc.wantSizeMin, tc.wantSizeMax, gotSize})
		})
	}

	// Log all obtained ratios
	format := "%-20v%-20v%-15v%v"
	t.Log("-------------------------------------------------------------")
	t.Logf(format, "Test case", "Want size", "Got size", "")
	t.Log("-------------------------------------------------------------")
	for _, r := range allSizes {
		change := ""
		if r.got < r.wantMin {
			change = "IMPROVED!"
		} else if r.got > r.wantMax {
			change = "BAD!"
		}
		wantSize := fmt.Sprintf("%v..%v", r.wantMin, r.wantMax)
		t.Logf(format, r.desc, wantSize, r.got, change)
	}
	t.Log("-------------------------------------------------------------")
}

// TestDeltaCorrectness checks if applying a delta on a base image results in an
// image with the same contents as the original target image.
func TestDeltaCorrectness(t *testing.T) {
	defer setupTemporaryTestRegistry(t)()

	client := testEnv.APIClient()
	ctx := context.Background()

	for _, tc := range deltaTestCases {
		delta := deltaName(tc.base, tc.target)
		t.Run(delta, func(t *testing.T) {
			imagesToBuild := []string{tc.base, tc.target}
			if tc.images != nil {
				imagesToBuild = tc.images
			}

			imagesToRemove := []string{tc.target}
			if tc.removeBeforePull != nil {
				imagesToRemove = tc.removeBeforePull
			}

			// Build all required images
			var imageRemovers []func()
			for _, image := range imagesToBuild {
				removeImage := ttrBuildImageAsserting(ctx, t, client, image)
				if sliceContains(imagesToRemove, image) {
					imageRemovers = append(imageRemovers, removeImage)
				} else {
					defer removeImage()
				}
			}

			// Create delta of them and push this delta.
			removeDelta := ttrCreateDeltaAsserting(ctx, t, client, tc.base, tc.target)
			ttrPushImageAsserting(ctx, t, client, delta)

			// The delta we have locally shall not be the same as the target image.
			targetHash := ttrHashImageAsserting(ctx, t, client, tc.target)
			deltaHash := ttrHashImageAsserting(ctx, t, client, delta)
			assert.Assert(t, !reflect.DeepEqual(targetHash, deltaHash))

			// Remove the delta and the target image.
			for _, removeImage := range imageRemovers {
				removeImage()
			}
			removeDelta()

			// Pull the delta. This will cause it to be applied.
			ttrPullImageAsserting(ctx, t, client, delta)

			// The (now applied) delta shall be the same the target image was.
			appliedDeltaHash := ttrHashImageAsserting(ctx, t, client, delta)
			assert.Assert(t, reflect.DeepEqual(targetHash, appliedDeltaHash))
		})
	}
}

// TestPullUsingDeltaStore checks if balenaEngine's alternative delta root
// feature is working as expected.
func TestPullUsingDeltaStore(t *testing.T) {
	// Basis and target for delta.
	const basis = "busybox:1.34"
	const target = "busybox:1.35"
	delta := deltaName(basis, target)

	// We'll push the delta to this registry.
	defer setupTemporaryTestRegistry(t)()

	// Create a delta, push it to the temporary test registry.
	func() {
		d := daemon.New(t)
		d.Start(t)
		defer d.Stop(t)

		client := d.NewClientT(t)
		ctx := context.Background()

		pullAndTagAsserting(ctx, t, client, basis, ttrImageName(basis))
		pullAndTagAsserting(ctx, t, client, target, ttrImageName(target))

		ttrCreateDeltaAsserting(ctx, t, client, basis, target)
		ttrPushImageAsserting(ctx, t, client, delta)
	}()

	// This is the path we'll eventually use as the delta data root.
	deltaDataRootDir, err := os.MkdirTemp("", "")
	assert.NilError(t, err)
	defer os.RemoveAll(deltaDataRootDir)

	// Place the basis image on `deltaDataRootDir`. We do that by pulling it from
	// the temporary registry to a temporary daemon that uses `deltaDataRootDir`
	// as its regular data root.
	func() {
		d := daemon.New(t)
		args := []string{
			fmt.Sprintf("--data-root=%s", deltaDataRootDir),
		}
		d.Start(t, args...)
		defer d.Stop(t)

		client := d.NewClientT(t)
		ctx := context.Background()

		rc, err := client.ImagePull(ctx, basis, types.ImagePullOptions{})
		_, err = readAllAndClose(rc)
		assert.NilError(t, err)

		imgs, err := client.ImageList(ctx, types.ImageListOptions{All: true})
		assert.NilError(t, err)
		assert.Equal(t, len(imgs), 1)

		assert.NilError(t, err)
	}()

	// Just to help detecting issues with the test itself, try to pull the delta
	// from the temporary test registry using a daemon that doesn't have the
	// delta data root set. It should fail, because this daemon doesn't have
	// access to the basis image.
	func() {
		d := daemon.New(t)
		d.Start(t)
		defer d.Stop(t)

		client := d.NewClientT(t)
		ctx := context.Background()

		err = ttrPullImage(ctx, client, delta)
		assert.Error(t, err, "image pull not successful")
	}()

	// Finally everything is set up for the real test. Create a daemon using
	// `deltaDataRootDir` as the delta data root, then try to pull the delta. It
	// should work.
	func() {
		d := daemon.New(t)
		args := []string{
			fmt.Sprintf("--delta-data-root=%s", deltaDataRootDir),
			fmt.Sprintf("--delta-storage-driver=%s", d.StorageDriver()),
		}
		d.Start(t, args...)
		defer d.Stop(t)

		client := d.NewClientT(t)
		ctx := context.Background()

		ttrPullImageAsserting(ctx, t, client, delta)
	}()
}

//
// Temporary Test Registry (TTR) helper functions
//
// With regards to image names: all these functions expect plain image names
// (like "002"), but under the hood they'll add the temporary test registry name
// to the images ("127.0.0.1:5000/002").
//
// In other words, this is about image tagging. Most of these functions will
// work even without the temporary test registry running. (Only the push and
// pull operations need the registry.)
//

// ttrBuildImageAsserting builds a given image and tags it. It asserts that all
// operations succeeded. Returns a function that, when called, remove the image
// built.
func ttrBuildImageAsserting(ctx context.Context, t *testing.T, client apiclient.APIClient, image string) func() {
	source := fakecontext.New(t, "")
	defer source.Close()

	err := copy.DirCopy("../../integration/testdata/delta/", source.Dir, copy.Content, false)
	assert.Assert(t, err)

	resp, err := client.ImageBuild(ctx, source.AsTarReader(t),
		types.ImageBuildOptions{
			Tags:       []string{ttrImageName(image)},
			Dockerfile: image + ".Dockerfile",
		},
	)
	assert.Assert(t, err)

	if resp.Body != nil {
		body, err := readAllAndClose(resp.Body)
		assert.Assert(t, err)
		assert.Assert(t, strings.Contains(body, "Successfully built"))
		assert.Assert(t, strings.Contains(body, "Successfully tagged"))
	}

	return func() {
		ttrRemoveImageAsserting(ctx, t, client, image)
	}
}

// ttrRemoveImageAsserting removes a given image. Asserts that the operation
// succeeded.
func ttrRemoveImageAsserting(ctx context.Context, t *testing.T, client apiclient.APIClient, image string) {
	resp, err := client.ImageRemove(ctx, ttrImageName(image), types.ImageRemoveOptions{})
	assert.Assert(t, err)
	assert.Assert(t, len(resp) > 0)
}

// ttrCreateDeltaAsserting creates a delta between base and target, and tags the
// delta as deltaName(base, target). Returns a function that, when called,
// removes the created delta.
func ttrCreateDeltaAsserting(ctx context.Context, t *testing.T, client apiclient.APIClient,
	base, target string) func() {

	delta := deltaName(base, target)

	rc, err := client.ImageDelta(ctx, ttrImageName(base), ttrImageName(target),
		types.ImageDeltaOptions{Tag: ttrImageName(delta)})
	assert.Assert(t, err)

	if rc != nil {
		body, err := readAllAndClose(rc)
		assert.Assert(t, err)
		assert.Assert(t, strings.Contains(body, "Created delta"))
		assert.Assert(t, strings.Contains(body, "Successfully tagged"))
	}

	return func() {
		ttrRemoveImageAsserting(ctx, t, client, delta)
	}
}

// ttrPushImageAsserting pushes a given image to the temporary test registry. It
// asserts that the operation succeeded.
func ttrPushImageAsserting(ctx context.Context, t *testing.T, client apiclient.APIClient, image string) {
	rc, err := client.ImagePush(ctx, ttrImageName(image), types.ImagePushOptions{RegistryAuth: "{}"})
	assert.Assert(t, err)
	if rc != nil {
		body, err := readAllAndClose(rc)
		assert.Assert(t, err)
		assert.Assert(t, strings.Contains(body, `"status":"latest: digest: `))
	}
}

// ttrPullImage pulls a given image from the temporary test registry. The image
// parameter must not include the registry name. Unlike the typical ttr*()
// function, this one returns an error instead of doing the asserts internally.
func ttrPullImage(ctx context.Context, client apiclient.APIClient, image string) error {
	rc, err := client.ImagePull(ctx, ttrImageName(image), types.ImagePullOptions{RegistryAuth: "{}"})
	if err != nil {
		return err
	}
	if rc != nil {
		body, err := readAllAndClose(rc)
		if err != nil {
			return err
		}
		if !strings.Contains(body, "Status: Downloaded newer image") {
			return errors.New("image pull not successful")
		}
	}
	return nil
}

// ttrPullImageAsserting pulls a given image from the temporary test registry.
// The image parameter must not include the registry name. Asserts that the
// operation succeeded.
func ttrPullImageAsserting(ctx context.Context, t *testing.T, client apiclient.APIClient, image string) {
	err := ttrPullImage(ctx, client, image)
	assert.NilError(t, err)
}

// ttrQueryDeltaSizeAsserting returns the size in bytes of a delta image.
// Asserts that the operation succeeded.
func ttrQueryDeltaSizeAsserting(ctx context.Context, t *testing.T, client apiclient.APIClient,
	image string) int64 {

	tarRC, err := client.ImageSave(ctx, []string{ttrImageName(image)})
	assert.Assert(t, err)
	defer tarRC.Close()
	size := deltaSizeFromTar(t, tarRC)
	return size
}

func deltaSizeFromTar(t *testing.T, r io.Reader) int64 {
	tarReader := tar.NewReader(r)
	size := int64(0)

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		assert.Assert(t, err)
		info := header.FileInfo()
		if !info.IsDir() {
			switch info.Name() {
			case "delta":
				size += info.Size()
			case "layer.tar":
				subR := io.LimitReader(tarReader, info.Size())
				size += deltaSizeFromTar(t, subR)
			}
		}
	}
	return size
}

// ttrImageName returns the image name including the test registry.
func ttrImageName(image string) string {
	return registry.DefaultURL + "/" + image
}

// ttrHashImageAsserting computes a hash based on the contents of a given image.
// Asserts that the operation succeeded.
func ttrHashImageAsserting(ctx context.Context, t *testing.T, client apiclient.APIClient, image string) []byte {
	ii, _, err := client.ImageInspectWithRaw(ctx, ttrImageName(image))
	assert.Assert(t, err)

	// We don't hash the image content per se, but the hashes of individual
	// layers, which are themselves based on the layer contents.
	hash := sha256.New()
	for _, layerHash := range ii.RootFS.Layers {
		_, err = hash.Write([]byte(layerHash))
		assert.Assert(t, err)
	}
	return hash.Sum(nil)
}

//
// Other helper functions
//

// deltaName returns the name we'll use in the tests for a delta from base to
// target.
func deltaName(base, target string) string {
	cleanBase := strings.ReplaceAll(base, ":", "_")
	cleanTarget := strings.ReplaceAll(target, ":", "_")
	return "delta-" + cleanBase + "-" + cleanTarget
}

// readAllAndClose reads everything from r, closes it, and returns whatever was
// read as a string.
func readAllAndClose(rc io.ReadCloser) (string, error) {
	data, err := io.ReadAll(rc)
	if err != nil {
		return "", err
	}
	rc.Close()
	return string(data), nil
}

// sliceContains checks if the haystack slice contains the needle string.
func sliceContains(haystack []string, needle string) bool {
	for _, e := range haystack {
		if e == needle {
			return true
		}
	}
	return false
}

// pullAndTagAsserting pulls a given image and tags it with a given tag. This
// function asserts that all operations are successful.
func pullAndTagAsserting(ctx context.Context, t *testing.T, client apiclient.APIClient, image, tag string) {
	rc, err := client.ImagePull(ctx, image, types.ImagePullOptions{})
	assert.NilError(t, err)

	_, err = readAllAndClose(rc)
	assert.NilError(t, err)

	err = client.ImageTag(ctx, image, tag)
	assert.NilError(t, err)
}
