package main

import (
	"fmt"
	"os"
	filepath "path/filepath"

	containerdShim "github.com/containerd/containerd/cmd/containerd-shim"
	"github.com/containerd/containerd/cmd/containerd"
	"github.com/containerd/containerd/cmd/ctr"
	"github.com/docker/cli/cmd/docker"
	"github.com/docker/docker/cmd/dockerd"
	"github.com/docker/docker/pkg/health"
	"github.com/docker/docker/pkg/reexec"
	"github.com/docker/libnetwork/cmd/proxy"
	"github.com/opencontainers/runc"
)

func main() {
	if reexec.Init() {
		return
	}

	command := filepath.Base(os.Args[0])

	switch command {
	case "balena", "balena-engine":
		docker.Main()
	case "balenad", "balena-engine-daemon":
		dockerd.Main()
	case "balena-containerd", "balena-engine-containerd":
		containerd.Main()
	case "balena-containerd-shim", "balena-engine-containerd-shim":
		containerdShim.Main()
	case "balena-containerd-ctr", "balena-engine-containerd-ctr":
		ctr.Main()
	case "balena-runc", "balena-engine-runc":
		runc.Main()
	case "balena-proxy", "balena-engine-proxy":
		proxy.Main()
	case "balena-engine-healthcheck":
		if err := health.RunHealthChecks(); err != nil {
			os.Exit(1)
		}
		os.Exit(0)
	default:
		fmt.Fprintf(os.Stderr, "error: unkown command: %v\n", command)
		os.Exit(1)
	}
}
