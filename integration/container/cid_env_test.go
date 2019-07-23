package container // import "github.com/docker/docker/integration/container"

import (
	"context"
	"fmt"
	"testing"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	testcontainer "github.com/docker/docker/integration/internal/container"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestContainerIDEnvOK(t *testing.T) {
	defer setupTest(t)()

	cidEnv := "CONTAINER_ID"
	ctx := context.Background()
	apiclient := testEnv.APIClient()

	config := &testcontainer.TestContainerConfig{
		Config: &container.Config{
			Image: "busybox",
			Cmd:   []string{"top"},
		},
		HostConfig: &container.HostConfig{
			ContainerIDEnv: cidEnv,
		},
		NetworkingConfig: &network.NetworkingConfig{},
	}
	resp, err := apiclient.ContainerCreate(ctx, config.Config, config.HostConfig, config.NetworkingConfig, nil, config.Name)
	assert.NilError(t, err)

	c, err := apiclient.ContainerInspect(ctx, resp.ID)
	assert.NilError(t, err)
	expected := fmt.Sprintf("%s=%s", cidEnv, resp.ID)
	assert.Check(t, is.Contains(c.Config.Env, expected))
}

func TestContainerIDEnvVariableExists(t *testing.T) {
	defer setupTest(t)()

	cidEnv := "PATH"
	ctx := context.Background()
	apiclient := testEnv.APIClient()

	config := &testcontainer.TestContainerConfig{
		Config: &container.Config{
			Image: "busybox",
			Cmd:   []string{"top"},
		},
		HostConfig: &container.HostConfig{
			ContainerIDEnv: cidEnv,
		},
		NetworkingConfig: &network.NetworkingConfig{},
	}
	_, err := apiclient.ContainerCreate(ctx, config.Config, config.HostConfig, config.NetworkingConfig, nil, config.Name)
	assert.ErrorContains(t, err, fmt.Sprintf("environment variable %s already defined", cidEnv))
}
