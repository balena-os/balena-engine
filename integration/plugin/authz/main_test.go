//go:build !windows

package authz // import "github.com/docker/docker/integration/plugin/authz"

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/docker/docker/testutil"
	"github.com/docker/docker/testutil/daemon"
	"github.com/docker/docker/testutil/environment"
	"gotest.tools/v3/skip"
)

var (
	testEnv     *environment.Execution
	d           *daemon.Daemon
	server      *httptest.Server
	baseContext context.Context
)

func TestMain(m *testing.M) {
	// Plugins are not supported
}

func setupTest(t *testing.T) context.Context {
	skip.If(t, testEnv.IsRemoteDaemon, "cannot run daemon when remote daemon")
	skip.If(t, testEnv.DaemonInfo.OSType == "windows")
	skip.If(t, testEnv.IsRootless, "rootless mode has different view of localhost")

	ctx := testutil.StartSpan(baseContext, t)
	environment.ProtectAll(ctx, t, testEnv)

	d = daemon.New(t, daemon.WithExperimental())

	t.Cleanup(func() {
		if d != nil {
			d.Stop(t)
		}
		testEnv.Clean(ctx, t)
	})
	return ctx
}
