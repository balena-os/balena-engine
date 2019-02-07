package container // import "github.com/docker/docker/integration/container"

import (
	"context"
	"strconv"
	"testing"
	"time"

	"github.com/docker/docker/integration/internal/container"
	"github.com/docker/docker/internal/test/request"
	"gotest.tools/assert"
	"gotest.tools/poll"
	"gotest.tools/skip"
)

// TestStopContainerWithTimeout checks that ContainerStop with
// a timeout works as documented, i.e. in case of negative timeout
// waiting is not limited (issue #35311).
func TestStopContainerWithTimeout(t *testing.T) {
	skip.If(t, testEnv.OSType == "windows")
	defer setupTest(t)()
	client := request.NewAPIClient(t)
	ctx := context.Background()

	testCmd := container.WithCmd("sh", "-c", "sleep 2 && exit 42")
	testData := []struct {
		doc              string
		timeout          int
		expectedExitCode int
	}{
		// In case container is forcefully killed, 137 is returned,
		// otherwise the exit code from the above script
		{
			"zero timeout: expect forceful container kill",
			1, 0x40010004,
		},
		{
			"too small timeout: expect forceful container kill",
			2, 0x40010004,
		},
		{
			"big enough timeout: expect graceful container stop",
			120, 42,
		},
		{
			"unlimited timeout: expect graceful container stop",
			-1, 42,
		},
	}

	for _, d := range testData {
		d := d
		t.Run(strconv.Itoa(d.timeout), func(t *testing.T) {
			t.Parallel()
			id := container.Run(t, ctx, client, testCmd)

			timeout := time.Duration(d.timeout) * time.Second
			err := client.ContainerStop(ctx, id, &timeout)
			assert.NilError(t, err)

			poll.WaitOn(t, container.IsStopped(ctx, client, id),
				poll.WithDelay(100*time.Millisecond))

			inspect, err := client.ContainerInspect(ctx, id)
			assert.NilError(t, err)
			assert.Equal(t, inspect.State.ExitCode, d.expectedExitCode)
		})
	}
}
