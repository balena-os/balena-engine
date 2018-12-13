package image

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/docker/cli/cli/command"
	"github.com/docker/docker/api/types"
	apiclient "github.com/docker/docker/client"
	"github.com/docker/docker/integration-cli/daemon"
	"github.com/docker/docker/integration-cli/registry"
)

const registryURI = "127.0.0.1:5000"

// TestDeltaSimple just creates a delta and checks if it exists
func TestDeltaSimple(t *testing.T) {
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
			Tags: []string{delta},
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

// TestDeltaWithRegistry creates a delta, pushes it to the registry.
// We then remove the delta and the delta target from the host and pull delta.
// The check looks for the delta on the host
func TestDeltaWithRegistry(t *testing.T) {
	var (
		base   = "busybox:1.24"
		target = "busybox:1.29"
		delta  = fmt.Sprintf("%s/busybox:delta-1.24-1.29", registryURI)
	)

	var (
		err    error
		rc     io.ReadCloser
		ctx    = context.Background()
		client = testEnv.APIClient()
	)

	reg, encodedAuth := startRegistry(t)
	defer reg.Close()

	pullBaseAndTargetImages(t, client, base, target)

	rc, err = client.ImageDelta(ctx,
		base,
		target,
		types.ImageDeltaOptions{
			Tags: []string{delta},
		})
	if err != nil {
		t.Fatalf("Creating delta: %s", err)
	}
	io.Copy(ioutil.Discard, rc)
	rc.Close()

	rc, err = client.ImagePush(ctx, delta, types.ImagePushOptions{RegistryAuth: encodedAuth})
	if err != nil {
		t.Fatalf("Pushing %s to local registry: %s", delta, err)
	}
	io.Copy(ioutil.Discard, rc)
	rc.Close()

	_, err = client.ImageRemove(ctx, delta, types.ImageRemoveOptions{})
	if err != nil {
		t.Fatalf("Removing delta from host: %s", err)
	}

	var targetInspect types.ImageInspect
	targetInspect, _, err = client.ImageInspectWithRaw(ctx, target)
	if err != nil {
		t.Fatalf("Inspecting target image: %s", err)
	}

	_, err = client.ImageRemove(ctx, target, types.ImageRemoveOptions{})
	if err != nil {
		t.Fatalf("Removing delta target from host: %s", err)
	}

	rc, err = client.ImagePull(ctx, delta, types.ImagePullOptions{
		RegistryAuth: encodedAuth,
	})
	if err != nil {
		t.Fatalf("Pulling delta from local registry: %s", err)
	}
	parseCmdOutputForError(t, rc)
	rc.Close()

	var deltaInspect types.ImageInspect
	deltaInspect, _, err = client.ImageInspectWithRaw(ctx, delta)
	if err != nil {
		t.Fatalf("Inspecting delta: %s", err)
	}

	for _, layer := range targetInspect.RootFS.Layers {
		assert.Contains(t, deltaInspect.RootFS.Layers, layer)
	}
}

// PATH=$PATH:`pwd`/balena-engine TESTDIRS="integration/image" TESTFLAGS="-test.run Delta" hack/make.sh test-integration
func TestDeltaWithRegistryUsingSeparateDeltaStore(t *testing.T) {
	var (
		base   = "busybox:1.24"
		target = "busybox:1.29"
		delta  = fmt.Sprintf("%s/busybox:delta-1.24-1.29", registryURI)
	)

	var (
		err error
		rc  io.ReadCloser
		ctx = context.Background()
	)

	d := daemon.New(t, "", "balena-engine-daemon", daemon.Config{})
	client, err := d.NewClient()
	if err != nil {
		t.Fatalf("Starting daemon: %s", err)
	}

	var args = []string{
		"--insecure-registry=" + registryURI,
		"--debug",
	}

	d.Start(t, args...)
	defer d.Stop(t)

	reg, encodedAuth := startRegistry(t)
	defer reg.Close()

	pullBaseAndTargetImages(t, client, base, target)

	rc, err = client.ImageDelta(ctx,
		base,
		target,
		types.ImageDeltaOptions{
			Tags: []string{delta},
		})
	if err != nil {
		t.Fatalf("Creating delta: %s", err)
	}
	io.Copy(ioutil.Discard, rc)
	rc.Close()

	rc, err = client.ImagePush(ctx, delta, types.ImagePushOptions{RegistryAuth: encodedAuth})
	if err != nil {
		t.Fatalf("Pushing %s to local registry: %s", delta, err)
	}
	io.Copy(ioutil.Discard, rc)
	rc.Close()

	var targetInspect types.ImageInspect
	targetInspect, _, err = client.ImageInspectWithRaw(ctx, target)
	if err != nil {
		t.Fatalf("Inspecting target image: %s", err)
	}

	// t.Log("Stopping daemon")
	d.Stop(t)

	args = append(args, []string{
		fmt.Sprintf("--delta-data-root=%s", d.Root),
		fmt.Sprintf("--delta-storage-driver=%s", os.Getenv("DOCKER_GRAPHDRIVER")),
	}...)
	var newRootDir = fmt.Sprintf("%s/root-2", d.Folder)
	d.Root = newRootDir

	// t.Log("Starting daemon with separate delta-data-root")
	d.Start(t, args...)

	rc, err = client.ImagePull(ctx, delta, types.ImagePullOptions{
		RegistryAuth: encodedAuth,
	})
	if err != nil {
		t.Fatalf("Pulling delta from local registry: %s", err)
	}
	parseCmdOutputForError(t, rc)
	rc.Close()

	var deltaInspect types.ImageInspect
	deltaInspect, _, err = client.ImageInspectWithRaw(ctx, delta)
	if err != nil {
		t.Fatalf("Inspecting delta: %s", err)
	}

	for _, layer := range targetInspect.RootFS.Layers {
		assert.Contains(t, deltaInspect.RootFS.Layers, layer)
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

func startRegistry(t *testing.T) (reg *registry.V2, encodedRegistryAuthConfig string) {
	var err error
	reg, err = registry.NewV2(false, "htpasswd", "", registryURI)
	if err != nil {
		t.Fatalf("Starting registry: %s", err)
	}
	encodedRegistryAuthConfig, err = command.EncodeAuthToBase64(types.AuthConfig{
		Username: reg.Username(),
		Password: reg.Password(),
		Auth:     "htpasswd",
	})
	if err != nil {
		t.Fatalf("Encodign registry auth config: %s", err)
	}
	return reg, encodedRegistryAuthConfig
}

// nice to have for debugging
func listImages(t *testing.T, client apiclient.APIClient) {
	var (
		err  error
		list []types.ImageSummary
		ctx  = context.Background()
	)

	list, err = client.ImageList(ctx, types.ImageListOptions{})
	if err != nil {
		t.Fatal(err)
	}
	for _, im := range list {
		t.Logf("%v: %v", im.ID, im.RepoTags)
	}
}

// nice to have for debugging
func parseCmdOutputForError(t *testing.T, cmdOutput io.Reader) {
	sc := bufio.NewScanner(cmdOutput)
	for sc.Scan() {
		var v map[string]interface{}
		if err := json.Unmarshal(sc.Bytes(), &v); err != nil {
			t.Fatal(err)
		}
		if e, ok := v["errorDetail"]; ok {
			err, ok := e.(map[string]interface{})
			if !ok {
				t.Fail()
				break
			}
			t.Fatal(err["message"])
		}
	}
}
