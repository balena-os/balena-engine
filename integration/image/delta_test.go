package image

import (
	"context"
	"io"
	"io/ioutil"
	"testing"

	"github.com/docker/docker/api/types"
	apiclient "github.com/docker/docker/client"
)

// TestDelta just creates a delta and checks if it exists
func TestDelta(t *testing.T) {
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
	if err != nil {
		t.Fatalf("Creating delta: %s", err)
	}
	io.Copy(ioutil.Discard, rc)
	rc.Close()

	_, _, err = client.ImageInspectWithRaw(ctx, delta)
	if err != nil {
		t.Fatalf("Inspecting delta: %s", err)
	}
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
	if err != nil {
		t.Fatalf("Pulling delta base: %s", err)
	}
	io.Copy(ioutil.Discard, rc)
	rc.Close()

	rc, err = client.ImagePull(ctx,
		target,
		types.ImagePullOptions{})
	if err != nil {
		t.Fatalf("Pulling delta target: %s", err)
	}
	io.Copy(ioutil.Discard, rc)
	rc.Close()
}
