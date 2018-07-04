//+build !windows

package chrootarchive

import (
	"fmt"
	"os"

	"github.com/containerd/cgroups"
	specs "github.com/opencontainers/runtime-spec/specs-go"
)

const MEMORY_LIMIT = 20 * 1024 * 1024 // 20MB

func constrainMemory() {
	// constrain the unpacking process using a memory cgroup controller to avoid thrashing the page cache
	memoryLimit := int64(MEMORY_LIMIT)
	control, err := cgroups.New(cgroups.V1, cgroups.StaticPath("/balena/apply-layer"), &specs.LinuxResources{
		Memory: &specs.LinuxMemory{
			Limit: &memoryLimit,
		},
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not memory constrain chrootarchive: %v", err)
	}
	if err := control.Add(cgroups.Process{Pid: os.Getpid()}); err != nil {
		fmt.Fprintf(os.Stderr, "could not memory constrain chrootarchive: %v", err)
	}
}
