package config // import "github.com/docker/docker/daemon/config"

import (
	"fmt"

	"github.com/docker/docker/api/types/swarm"
)

// ParseGenericResources parses and validates the specified string as a list of GenericResource
func ParseGenericResources(value []string) ([]swarm.GenericResource, error) {
	return nil, fmt.Errorf("Unsupported feature")
}
