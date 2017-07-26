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
	case "docker":
		docker.Main()
	case "dockerd":
		dockerd.Main()
	case "docker-containerd":
		containerd.Main()
	case "docker-containerd-shim":
		containerdShim.Main()
	case "docker-containerd-ctr":
		ctr.Main()
	case "docker-runc":
		runc.Main()
	case "docker-proxy":
		proxy.Main()

	default:
		fmt.Println("Unknown command")
	}
}
