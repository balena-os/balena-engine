package image

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/docker/cli/cli/command"
	"github.com/docker/docker/api/types"
	apiclient "github.com/docker/docker/client"
	"github.com/docker/docker/integration-cli/daemon"
	"github.com/docker/docker/integration-cli/registry"
)

const registryURI = "127.0.0.1:5000"

func TestDelta(t *testing.T) {
	var (
		base           = "busybox:1.24"
		target         = "busybox:1.29"
		delta          = fmt.Sprintf("%s/busybox:delta-1.24-1.29", registryURI)
		expectedImages = []string{
			"sha256:47bcc53f74dc94b1920f0b34f6036096526296767650f223433fe65c35f149eb", // busybox:1.24
			"sha256:59788edf1f3e78cd0ebe6ce1446e9d10788225db3dedcfd1a59f764bad2b2690", // busybox:1.29
		}
	)

	var (
		err    error
		rc     io.ReadCloser
		ctx    = context.Background()
		client = testEnv.APIClient()
	)

	pullDeltaImages(t, client, base, target)

	t.Log("Creating delta")
	rc, err = client.ImageDelta(ctx,
		base,
		target)
	if err != nil {
		t.Fatal(err)
	}
	deltaID := parseDeltaImageID(t, rc)
	expectedImages = append(expectedImages, deltaID)
	rc.Close()

	err = client.ImageTag(ctx, deltaID, delta)
	if err != nil {
		t.Fatal(err)
	}

	assert.True(t, hasImages(t, client, expectedImages))
}

func TestDeltaWithRegistry(t *testing.T) {
	var (
		base           = "busybox:1.24"
		target         = "busybox:1.29"
		delta          = fmt.Sprintf("%s/busybox:delta-1.24-1.29", registryURI)
		expectedImages = []string{
			"sha256:47bcc53f74dc94b1920f0b34f6036096526296767650f223433fe65c35f149eb", // busybox:1.24
		}
	)

	var (
		err    error
		rc     io.ReadCloser
		ctx    = context.Background()
		client = testEnv.APIClient()
	)

	reg, err := registry.NewV2(false, "htpasswd", "", registryURI)
	if err != nil {
		t.Fatal(err)
	}
	defer reg.Close()

	pullDeltaImages(t, client, base, target)

	t.Log("Creating delta")
	rc, err = client.ImageDelta(ctx,
		base,
		target)
	if err != nil {
		t.Fatal(err)
	}
	deltaID := parseDeltaImageID(t, rc)
	expectedImages = append(expectedImages, deltaID)
	rc.Close()

	err = client.ImageTag(ctx, deltaID, delta)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Pushing delta to local registry")
	encodedAuth, err := command.EncodeAuthToBase64(types.AuthConfig{
		Username: reg.Username(),
		Password: reg.Password(),
		Auth:     "htpasswd",
	})
	if err != nil {
		t.Fatal(err)
	}
	rc, err = client.ImagePush(ctx, delta, types.ImagePushOptions{
		RegistryAuth: encodedAuth,
	})
	if err != nil {
		t.Fatal(err)
	}
	io.Copy(ioutil.Discard, rc)
	rc.Close()

	t.Log("Removing delta target from host")
	_, err = client.ImageRemove(ctx, target, types.ImageRemoveOptions{})
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Pulling delta from local registry")
	rc, err = client.ImagePull(ctx, delta, types.ImagePullOptions{
		RegistryAuth: encodedAuth,
	})
	if err != nil {
		t.Fatal(err)
	}
	parseCmdOutputForError(t, rc)
	rc.Close()

	assert.True(t, hasImages(t, client, expectedImages))
}

// PATH=$PATH:`pwd`/balena-engine TESTDIRS="integration/image" TESTFLAGS="-test.run Delta" hack/make.sh test-integration
func TestDeltaWithRegistryUsingSeparateDeltaStore(t *testing.T) {
	var (
		base           = "busybox:1.24"
		target         = "busybox:1.29"
		delta          = fmt.Sprintf("%s/busybox:delta-1.24-1.29", registryURI)
		expectedImages = []string{}
	)

	var (
		err error
		rc  io.ReadCloser
		ctx = context.Background()
	)

	d := daemon.New(t, "", "balena-engine-daemon", daemon.Config{})
	client, err := d.NewClient()
	if err != nil {
		t.Fatal(err)
	}

	var args = []string{
		"--insecure-registry=" + registryURI,
		"--debug",
	}

	d.Start(t, args...)
	defer d.Stop(t)

	reg, err := registry.NewV2(false, "htpasswd", "", registryURI)
	if err != nil {
		t.Fatal(err)
	}
	defer reg.Close()

	pullDeltaImages(t, client, base, target)

	t.Log("Creating delta")
	rc, err = client.ImageDelta(ctx,
		base,
		target)
	if err != nil {
		t.Fatal(err)
	}
	deltaID := parseDeltaImageID(t, rc)
	rc.Close()

	expectedImages = append(expectedImages, deltaID)

	err = client.ImageTag(ctx, deltaID, delta)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Pushing delta to local registry")
	encodedAuth, err := command.EncodeAuthToBase64(types.AuthConfig{
		Username: reg.Username(),
		Password: reg.Password(),
		Auth:     "htpasswd",
	})
	if err != nil {
		t.Fatal(err)
	}
	rc, err = client.ImagePush(ctx, delta, types.ImagePushOptions{
		RegistryAuth: encodedAuth,
	})
	if err != nil {
		t.Fatal(err)
	}
	io.Copy(ioutil.Discard, rc)
	rc.Close()

	t.Log("Stopping daemon")
	d.Stop(t)

	args = append(args, []string{
		fmt.Sprintf("--delta-data-root=%s", d.Root),
		fmt.Sprintf("--delta-storage-driver=%s", os.Getenv("DOCKER_GRAPHDRIVER")),
	}...)
	var newRootDir = fmt.Sprintf("%s/root-2", d.Folder)
	d.Root = newRootDir

	t.Log("Starting daemon with separate delta-data-root")
	d.Start(t, args...)

	t.Log("Pulling delta from local registry")
	rc, err = client.ImagePull(ctx, delta, types.ImagePullOptions{
		RegistryAuth: encodedAuth,
	})
	if err != nil {
		t.Fatal(err)
	}
	parseCmdOutputForError(t, rc)
	rc.Close()

	assert.True(t, hasImages(t, client, expectedImages))
}

func pullDeltaImages(t *testing.T, client apiclient.APIClient, base, target string) {
	var (
		err error
		rc  io.ReadCloser
		ctx = context.Background()
	)

	t.Log("Pulling delta base")
	rc, err = client.ImagePull(ctx,
		base,
		types.ImagePullOptions{})
	if err != nil {
		t.Fatal(err)
	}
	io.Copy(ioutil.Discard, rc)
	rc.Close()

	t.Log("Pulling delta target")
	rc, err = client.ImagePull(ctx,
		target,
		types.ImagePullOptions{})
	if err != nil {
		t.Fatal(err)
	}
	io.Copy(ioutil.Discard, rc)
	rc.Close()
}

var digestRe = regexp.MustCompile("sha256:([a-fA-F0-9]{64})$")

func parseDeltaImageID(t *testing.T, deltaCmdOutput io.Reader) (digest string) {
	sc := bufio.NewScanner(deltaCmdOutput)
	for sc.Scan() {
		var v map[string]interface{}
		if err := json.Unmarshal(sc.Bytes(), &v); err != nil {
			t.Fatal(err)
		}
		if _, ok := v["errorDetail"]; ok {
			t.Fatal(errors.New(v["errorDetail"].(string)))
		}
		_, ok := v["progressDetail"]
		if ok {
			continue
		}
		s, ok := v["status"]
		if !ok {
			continue
		}
		status, ok := s.(string)
		if !ok {
			continue
		}
		r := digestRe.FindStringSubmatch(status)
		if len(r) < 1 {
			continue
		}
		digest = r[0]
	}
	if err := sc.Err(); err != nil {
		t.Fatal(err)
	}
	if digest == "" {
		t.Fatal("Unable to parse delta image id from progress output")
	}
	return digest
}

func hasImages(t *testing.T, client apiclient.APIClient, imageIDs []string) (ok bool) {
	var (
		err  error
		list []types.ImageSummary
		ctx  = context.Background()
	)

	list, err = client.ImageList(ctx, types.ImageListOptions{})
	if err != nil {
		return false
	}
	var foundImageIDs []string
	for _, im := range list {
		foundImageIDs = append(foundImageIDs, im.ID)
	}

	for _, id := range imageIDs {
		if !assert.Contains(t, foundImageIDs, id) {
			return false
		}
	}
	return true
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
