package network // import "github.com/docker/docker/api/server/router/network"

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/backend"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/network"
)

// Backend is all the methods that need to be implemented
// to provide network specific functionality.
type Backend interface {
	GetNetworks(filters.Args, backend.NetworkListConfig) ([]types.NetworkResource, error)
	CreateNetwork(nc types.NetworkCreateRequest) (*types.NetworkCreateResponse, error)
	ConnectContainerToNetwork(containerName, networkName string, endpointConfig *network.EndpointSettings) error
	DisconnectContainerFromNetwork(containerName string, networkName string, force bool) error
	DeleteNetwork(networkID string) error
	NetworksPrune(ctx context.Context, pruneFilters filters.Args) (*types.NetworksPruneReport, error)
}
