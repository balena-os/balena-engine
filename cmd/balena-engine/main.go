package main

import (
	"fmt"
	"os"

	"github.com/containerd/containerd/cmd/containerd"
	containerdShim "github.com/containerd/containerd/cmd/containerd-shim"
	"github.com/containerd/containerd/cmd/ctr"
	"github.com/docker/cli/cmd/docker"
	"github.com/docker/docker/cmd/dockerd"
	"github.com/docker/docker/pkg/reexec"
	"github.com/docker/libnetwork/cmd/proxy"
	"github.com/opencontainers/runc"

	filepath "path/filepath"
)

func init() {
	reexec.Register("balena-engine", docker.Main)
	reexec.Register("balena", docker.Main)

	reexec.Register("balena-engine-daemon", dockerd.Main)
	reexec.Register("balenad", dockerd.Main)

	reexec.Register("containerd", containerd.Main)
	reexec.Register("balena-engine-containerd", containerd.Main)
	reexec.Register("balena-containerd", containerd.Main)

	reexec.Register("containerd-shim", containerdShim.Main)
	reexec.Register("balena-engine-containerd-shim", containerdShim.Main)
	reexec.Register("balena-containerd-shim", containerdShim.Main)

	reexec.Register("ctr", ctr.Main)
	reexec.Register("balena-engine-containerd-ctr", ctr.Main)
	reexec.Register("balena-containerd-ctr", ctr.Main)

	reexec.Register("balena-proxy", proxy.Main)
	reexec.Register("balena-engine-proxy", proxy.Main)

	reexec.Register("runc", runc.Main)
	reexec.Register("balena-engine-runc", runc.Main)
	reexec.Register("balena-runc", runc.Main)
}

func main() {
	os.Args[0] = filepath.Base(os.Args[0])
	if reexec.Init() {
		return
	}

	fmt.Fprintf(os.Stderr, "reexec failed: unkown command: %v\n", os.Args[0])
	os.Exit(1)
}
