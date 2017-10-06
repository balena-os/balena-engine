package main

import (
	"fmt"
	"github.com/containerd/containerd/cmd/containerd"
	containerdShim "github.com/containerd/containerd/cmd/containerd-shim"
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
	switch filepath.Base(os.Args[0]) {
	case "balena":
		docker.Main()
	case "balenad":
		dockerd.Main()
	case "balena-containerd":
		containerd.Main()
	case "balena-containerd-shim":
		containerdShim.Main()
	case "balena-containerd-ctr":
		ctr.Main()
	case "balena-runc":
		runc.Main()
	case "balena-proxy":
		proxy.Main()

	default:
		fmt.Println("Unknown command")
	}
}
