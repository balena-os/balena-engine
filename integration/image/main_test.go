package image // import "github.com/docker/docker/integration/image"

import (
	"fmt"
	"os"
	"testing"

	"github.com/docker/docker/testutil/environment"
	"github.com/docker/docker/testutil/registry"
)

var testEnv *environment.Execution

func TestMain(m *testing.M) {
	var err error
	testEnv, err = environment.New()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = environment.EnsureFrozenImagesLinux(testEnv)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	testEnv.Print()
	os.Exit(m.Run())
}

func setupTest(t *testing.T) func() {
	environment.ProtectAll(t, testEnv)
	return func() { testEnv.Clean(t) }
}

// setupTemporaryTestRegistry creates a temporary image registry to be used
// during testing. Returns a function that must be called to tear down this
// registry.
func setupTemporaryTestRegistry(t *testing.T) func() {
	reg := registry.NewV2(t)
	reg.WaitReady(t)
	return reg.Close
}
