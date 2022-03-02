package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/containerd/containerd/cmd/containerd"
	containerdShim "github.com/containerd/containerd/cmd/containerd-shim"
	"github.com/containerd/containerd/cmd/ctr"
	containerdRuntimev2 "github.com/containerd/containerd/runtime/v2/runc/v2"
	containerdRuntimeShim "github.com/containerd/containerd/runtime/v2/shim"
	"github.com/docker/cli/cmd/docker"
	"github.com/docker/docker/cmd/dockerd"
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
	case "balena-containerd", "balena-engine-containerd", "containerd":
		containerd.Main()
	case "balena-containerd-shim", "balena-engine-containerd-shim":
		containerdShim.Main()
	case "containerd-shim-runc-v2":
		containerdRuntimeShim.Run("io.containerd.runc.v2", containerdRuntimev2.New)
	case "balena-containerd-ctr", "balena-engine-containerd-ctr":
		ctr.Main()
	case "balena-runc", "balena-engine-runc", "runc":
		runc.Main()
	case "balena-proxy", "balena-engine-proxy":
		proxy.Main()
	default:
		fmt.Fprintf(os.Stderr, "error: unkown command: %v\n", command)
		os.Exit(1)
	}
}
