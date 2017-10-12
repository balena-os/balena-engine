// +build !daemon

package docker

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
)

func newDaemonCommand() *cobra.Command {
	return &cobra.Command{
		Use:                "daemon",
		Hidden:             true,
		Args:               cobra.ArbitraryArgs,
		DisableFlagParsing: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runDaemon()
		},
	}
}

func runDaemon() error {
	return fmt.Errorf(
		"`balena daemon` is not supported on %s. Please run `balenad` directly",
		strings.Title(runtime.GOOS))
}
