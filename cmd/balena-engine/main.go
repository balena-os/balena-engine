package main

import (
	"fmt"

	"github.com/containerd/containerd/cmd/containerd"
	containerdShimRuncV2 "github.com/containerd/containerd/cmd/containerd-shim-runc-v2"
	"github.com/containerd/containerd/cmd/ctr"
	"github.com/docker/cli/cmd/docker"
	"github.com/docker/docker/cmd/dockerd"
	"github.com/docker/docker/pkg/reexec"
	"github.com/docker/libnetwork/cmd/proxy"
	"github.com/opencontainers/runc"

	"os"
	filepath "path/filepath"
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
	case "balena-containerd-shim-runc-v2":
		containerdShimRuncV2.Main()
	case "balena-containerd-ctr", "balena-engine-containerd-ctr":
		ctr.Main()
	case "balena-runc", "balena-engine-runc":
		runc.Main()
	case "balena-proxy", "balena-engine-proxy":
		proxy.Main()
	default:
		fmt.Fprintf(os.Stderr, "error: unknown command: %v\n", command)
		os.Exit(1)
	}
}
