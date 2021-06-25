package image

import (
	"bufio"
	"context"
	"io"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/docker/docker/api/types"
	apiclient "github.com/docker/docker/client"
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
