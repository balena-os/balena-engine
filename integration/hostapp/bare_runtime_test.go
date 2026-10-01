package hostapp

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"testing"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/docker/docker/testutil/request"

	"gotest.tools/v3/assert"
	"gotest.tools/v3/skip"
)

func TestBareRuntime(t *testing.T) {
	skip.If(t, testEnv.DaemonInfo.OSType != "linux")
	skip.If(t, testEnv.IsRemoteDaemon, "cannot start daemon on remote test run")
	ctx := setupTest(t)

	client := request.NewAPIClient(t)

	c, err := client.ContainerCreate(ctx,
		&container.Config{Image: "busybox:latest"},
		&container.HostConfig{Runtime: "bare"},
		&network.NetworkingConfig{},
		nil,
		"",
	)
	assert.NilError(t, err)

	j, err := client.ContainerInspect(ctx, c.ID)
	assert.NilError(t, err)

	containerDir, ok := j.GraphDriver.Data["MergedDir"]
	assert.Check(t, ok)

	_, err = os.Stat(filepath.Join(containerDir, ".dockerenv"))
	assert.Check(t, os.IsNotExist(err))
}

// findInLowerDir mirrors container.findInLowerDir: first image layer diff containing relPath.
func findInLowerDir(lowerDir, relPath string) (string, bool) {
	for _, layer := range strings.Split(lowerDir, ":") {
		candidate := filepath.Join(layer, relPath)
		if _, err := os.Stat(candidate); err == nil {
			return candidate, true
		}
	}
	return "", false
}

func createBareContainerWithVolume(ctx context.Context, t *testing.T, cli client.APIClient, mount string) string {
	t.Helper()
	c, err := cli.ContainerCreate(ctx,
		&container.Config{Image: "busybox:latest"},
		&container.HostConfig{
			Runtime: "bare",
			Binds:   []string{mount},
		},
		&network.NetworkingConfig{},
		nil,
		"",
	)
	assert.NilError(t, err)
	t.Cleanup(func() {
		_ = cli.ContainerRemove(ctx, c.ID, container.RemoveOptions{Force: true, RemoveVolumes: true})
	})
	return c.ID
}

func volumeMountSource(t *testing.T, inspect types.ContainerJSON, destination string) string {
	t.Helper()
	for _, m := range inspect.Mounts {
		if m.Destination == destination {
			return m.Source
		}
	}
	t.Fatalf("expected volume mount at %s", destination)
	return ""
}

func assertHardlinked(t *testing.T, srcPath, dstPath string) {
	t.Helper()

	srcStat, err := os.Stat(srcPath)
	assert.NilError(t, err)
	dstStat, err := os.Stat(dstPath)
	assert.NilError(t, err)

	srcIno := srcStat.Sys().(*syscall.Stat_t).Ino
	dstIno := dstStat.Sys().(*syscall.Stat_t).Ino
	t.Logf("layer inode:  %d (%s)", srcIno, srcPath)
	t.Logf("volume inode: %d (%s)", dstIno, dstPath)
	assert.Equal(t, srcIno, dstIno, "expected hardlink: same inode for %s and %s", srcPath, dstPath)
	t.Log("volume populated via hardlink")
}

func TestBareRuntimeVolumePopulateHardlinks(t *testing.T) {
	skip.If(t, testEnv.DaemonInfo.OSType != "linux")
	skip.If(t, testEnv.IsRemoteDaemon, "cannot start daemon on remote test run")
	ctx := setupTest(t)
	cli := request.NewAPIClient(t)

	const mount = "/bin"
	const fileName = "busybox"

	id := createBareContainerWithVolume(ctx, t, cli, mount)

	j, err := cli.ContainerInspect(ctx, id)
	assert.NilError(t, err)
	skip.If(t, j.GraphDriver.Name != "overlay2", "bare volume populate resolves paths from overlay2 LowerDir")

	lowerDir, ok := j.GraphDriver.Data["LowerDir"]
	assert.Assert(t, ok)

	relDir := strings.TrimPrefix(filepath.Clean(mount), string(filepath.Separator))
	sourceDir, ok := findInLowerDir(lowerDir, relDir)
	assert.Assert(t, ok, "expected %q in LowerDir", relDir)

	srcPath := filepath.Join(sourceDir, fileName)
	dstPath := filepath.Join(volumeMountSource(t, j, mount), fileName)

	assertHardlinked(t, srcPath, dstPath)
}
