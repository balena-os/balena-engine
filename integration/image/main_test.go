package image // import "github.com/docker/docker/integration/image"

import (
	"context"
	"os"
	"testing"

	"github.com/docker/docker/testutil"
	"github.com/docker/docker/testutil/environment"
	"github.com/docker/docker/testutil/registry"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
)

var (
	testEnv     *environment.Execution
	baseContext context.Context
)

func TestMain(m *testing.M) {
	shutdown := testutil.ConfigureTracing()
	ctx, span := otel.Tracer("").Start(context.Background(), "integration/daemon/TestMain")
	baseContext = ctx

	var err error
	testEnv, err = environment.New(ctx)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
		span.End()
		shutdown(ctx)
		panic(err)
	}
	err = environment.EnsureFrozenImagesLinux(ctx, testEnv)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
		span.End()
		shutdown(ctx)
		panic(err)
	}

	testEnv.Print()
	os.Exit(m.Run())
}

func setupTest(t *testing.T) context.Context {
	ctx := testutil.StartSpan(baseContext, t)
	environment.ProtectAll(ctx, t, testEnv)
	t.Cleanup(func() { testEnv.Clean(ctx, t) })
	return ctx
}

// setupTemporaryTestRegistry creates a temporary image registry to be used
// during testing. Returns a function that must be called to tear down this
// registry.
func setupTemporaryTestRegistry(t *testing.T) func() {
	reg := registry.NewV2(t)
	reg.WaitReady(t)
	return reg.Close
}
