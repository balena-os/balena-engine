package container

import (
	"context"
	"testing"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	ctr "github.com/docker/docker/integration/internal/container"
	"github.com/docker/docker/internal/test/request"
	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
	"gotest.tools/poll"
	"gotest.tools/skip"
)

func TestContainerWithCPUConstraints(t *testing.T) {
	skip.If(t, testEnv.DaemonInfo.OSType == "windows")

	defer setupTest(t)()
	name := "create-cpu-constraint"
	ctx := context.Background()
	client := request.NewAPIClient(t)

	config := container.Config{
		Image: "busybox",
		Cmd:   []string{"sh", "-c", "x=a; while true; do x=$x; done"},
	}
	hc := container.HostConfig{
		Resources: container.Resources{
			// equivalent to --cpus=".5"
			CPUPeriod: 100000,
			CPUQuota:  50000,
		},
	}

	c, err := client.ContainerCreate(ctx,
		&config,
		&hc,
		&network.NetworkingConfig{},
		name,
	)
	assert.NilError(t, err)

	err = client.ContainerStart(ctx, c.ID, types.ContainerStartOptions{})
	assert.NilError(t, err)

	poll.WaitOn(t, ctr.IsInState(ctx, client, c.ID, "running"), poll.WithDelay(100*time.Millisecond))

	timeout := 1 * time.Second
	err = client.ContainerStop(ctx, c.ID, &timeout)
	assert.NilError(t, err)

	poll.WaitOn(t, ctr.IsStopped(ctx, client, c.ID), poll.WithDelay(100*time.Millisecond))

	inspect, err := client.ContainerInspect(ctx, c.ID)
	assert.NilError(t, err)
	assert.Check(t, is.Equal(hc.CPUPeriod, inspect.HostConfig.CPUPeriod))
	assert.Check(t, is.Equal(hc.CPUQuota, inspect.HostConfig.CPUQuota))
}

func TestContainerWithMemoryConstraints(t *testing.T) {
	skip.If(t, testEnv.DaemonInfo.OSType == "windows" || !testEnv.DaemonInfo.MemoryLimit || !testEnv.DaemonInfo.SwapLimit)

	defer setupTest(t)()
	name := "create-memory-constraint"
	ctx := context.Background()
	client := request.NewAPIClient(t)

	config := container.Config{
		Image: "busybox",
		Cmd:   []string{"sh", "-c", "x=a; while true; do x=$x$x$x$x; done"},
	}
	hc := container.HostConfig{
		Resources: container.Resources{
			Memory: 32 * 1024 * 1024,
		},
	}

	c, err := client.ContainerCreate(ctx,
		&config,
		&hc,
		&network.NetworkingConfig{},
		name,
	)
	assert.NilError(t, err)

	err = client.ContainerStart(ctx, c.ID, types.ContainerStartOptions{})
	assert.NilError(t, err)

	poll.WaitOn(t, ctr.IsInState(ctx, client, c.ID, "running"), poll.WithDelay(100*time.Millisecond))

	timeout := 1 * time.Second
	err = client.ContainerStop(ctx, c.ID, &timeout)
	assert.NilError(t, err)

	poll.WaitOn(t, ctr.IsStopped(ctx, client, c.ID), poll.WithDelay(100*time.Millisecond))

	inspect, err := client.ContainerInspect(ctx, c.ID)
	assert.NilError(t, err)
	assert.Check(t, is.Equal(hc.Memory, inspect.HostConfig.Memory))
}
