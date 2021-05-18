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
	if err != nil {
		t.Fatalf("Creating delta: %s", err)
	}
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

	<-waitFingerprinting
	_, err = client.ImageRemove(ctx, target, types.ImageRemoveOptions{})
	if err != nil {
		t.Fatalf("Removing target image failed: %s", err)
	}

	<-waitDelta
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
