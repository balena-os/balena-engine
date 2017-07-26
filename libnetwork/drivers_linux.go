package libnetwork

import (
	"fmt"

	"github.com/docker/docker/libnetwork/driverapi"
	"github.com/docker/docker/libnetwork/drivers/bridge"
	"github.com/docker/docker/libnetwork/drivers/host"
	"github.com/docker/docker/libnetwork/drivers/null"
)

func registerNetworkDrivers(r driverapi.Registerer, driverConfig func(string) map[string]interface{}) error {
	noConfig := func(fn func(driverapi.Registerer) error) func(driverapi.Registerer, map[string]interface{}) error {
		return func(r driverapi.Registerer, _ map[string]interface{}) error { return fn(r) }
	}

	for _, nr := range []struct {
		ntype    string
		register func(driverapi.Registerer, map[string]interface{}) error
	}{
		{ntype: bridge.NetworkType, register: bridge.Register},
		{ntype: host.NetworkType, register: noConfig(host.Register)},
		{ntype: null.NetworkType, register: noConfig(null.Register)},
	} {
		if err := nr.register(r, driverConfig(nr.ntype)); err != nil {
			return fmt.Errorf("failed to register %q driver: %w", nr.ntype, err)
		}
	}

	return nil
}
