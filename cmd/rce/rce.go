package main

import (
	"fmt"
	"github.com/docker/containerd/containerd"
	containerdShim "github.com/docker/containerd/containerd-shim"
	"github.com/docker/containerd/ctr"
	"github.com/docker/docker/cmd/docker"
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
