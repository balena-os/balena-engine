//go:build linux && !exclude_graphdriver_btrfs
// +build linux,!exclude_graphdriver_btrfs

package btrfs // import "github.com/docker/docker/daemon/graphdriver/btrfs"

import (
	"testing"
)

func TestLibVersion(t *testing.T) {
	if btrfsLibVersion() <= 0 {
		t.Error("expected output from btrfs lib version > 0")
	}
}
