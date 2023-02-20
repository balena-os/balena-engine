//go:build !windows
// +build !windows

package dockerd

import (
	"github.com/spf13/pflag"
)

func installServiceFlags(flags *pflag.FlagSet) {
}
