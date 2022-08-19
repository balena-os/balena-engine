//go:build !windows
// +build !windows

package dockerd

import (
	"io"

	"github.com/sirupsen/logrus"
)

func runDaemon(opts *daemonOptions) error {
	daemonCli := NewDaemonCli()
	return daemonCli.start(opts)
}

func initLogging(_, stderr io.Writer) {
	logrus.SetOutput(stderr)
}
