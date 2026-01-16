package daemon // import "github.com/docker/docker/daemon"

import (
	"context"

	"github.com/containerd/log"
	swarmtypes "github.com/docker/docker/api/types/swarm"
	"github.com/moby/swarmkit/v2/agent/exec"
)

// SetContainerDependencyStore sets the dependency store for the container
func (daemon *Daemon) SetContainerDependencyStore(name string, store exec.DependencyGetter) error {
	c, err := daemon.GetContainer(name)
	if err != nil {
		return err
	}

	c.DependencyStore = store
	return nil
}

// SetContainerSecretReferences sets the container secret references needed
func (daemon *Daemon) SetContainerSecretReferences(name string, refs []*swarmtypes.SecretReference) error {
	if !secretsSupported() && len(refs) > 0 {
		log.G(context.TODO()).Warn("secrets are not supported on this platform")
		return nil
	}

	c, err := daemon.GetContainer(name)
	if err != nil {
		return err
	}

	c.SecretReferences = refs

	return nil
}
