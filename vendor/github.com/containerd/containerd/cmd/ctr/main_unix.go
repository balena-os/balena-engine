// +build !windows

package ctr

import "github.com/containerd/containerd/cmd/ctr/commands/shim"

func init() {
	extraCmds = append(extraCmds, shim.Command)
}
