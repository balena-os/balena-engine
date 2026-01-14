//go:build !windows

package dockerd

import (
	"io"

	"github.com/containerd/log"
)

func runDaemon(opts *daemonOptions) error {
	daemonCli := NewDaemonCli()
	return daemonCli.start(opts)
}

func initLogging(_, stderr io.Writer) {
	log.L.Logger.SetOutput(stderr)
}
