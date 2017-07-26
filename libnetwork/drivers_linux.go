package libnetwork

import (
	"github.com/docker/docker/libnetwork/drivers/bridge"
	"github.com/docker/docker/libnetwork/drivers/host"
	"github.com/docker/docker/libnetwork/drivers/null"
)

func getInitializers() []initializer {
	in := []initializer{
		{bridge.Init, "bridge"},
		{host.Init, "host"},
		{null.Init, "null"},
	}
	return in
}
