package config // import "github.com/docker/docker/daemon/config"

import (
	"fmt"

	"github.com/docker/docker/api/types/swarm"
)

// ParseGenericResources parses and validates the specified string as a list of GenericResource
// balena: Generic resources are not supported, but we allow empty values to not break config validation
func ParseGenericResources(value []string) ([]swarm.GenericResource, error) {
	if len(value) > 0 {
		return nil, fmt.Errorf("Unsupported feature: generic resources")
	}
	return nil, nil
}
