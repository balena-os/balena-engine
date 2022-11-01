//go:build linux
// +build linux

package daemon

import (
	"testing"

	containertypes "github.com/docker/docker/api/types/container"
	"github.com/docker/docker/container"
	"github.com/docker/docker/daemon/config"
	"github.com/docker/docker/daemon/exec"
	"github.com/opencontainers/runc/libcontainer/apparmor"
	specs "github.com/opencontainers/runtime-spec/specs-go"
	"gotest.tools/v3/assert"
)

func TestExecSetPlatformOpt(t *testing.T) {
	if !apparmor.IsEnabled() {
		t.Skip("requires AppArmor to be enabled")
	}
	d := &Daemon{configStore: &config.Config{}}
	c := &container.Container{
		AppArmorProfile: "my-custom-profile",
		HostConfig:      &containertypes.HostConfig{Privileged: false},
	}
	ec := &exec.Config{}
	p := &specs.Process{}

	err := d.execSetPlatformOpt(c, ec, p)
	assert.NilError(t, err)
	assert.Equal(t, "my-custom-profile", p.ApparmorProfile)
}

// TestExecSetPlatformOptPrivileged verifies that `docker exec --privileged`
// does not disable AppArmor profiles. Exec currently inherits the `Privileged`
// configuration of the container. See https://github.com/moby/moby/pull/31773#discussion_r105586900
//
// This behavior may change in future, but test for the behavior to prevent it
// from being changed accidentally.
//
// balenaEngine: this test was failing after we upgraded several components
// while updating containerd to 1.6.6. We changed it so it resembles the
// following test case in the more recent Moby codebase:
// https://github.com/moby/moby/blob/572ca799db4b67b7be35904e487f0cc51c3f9f06/daemon/exec_linux_test.go#L37-L39
func TestExecSetPlatformOptPrivileged(t *testing.T) {
	if !apparmor.IsEnabled() {
		t.Skip("requires AppArmor to be enabled")
	}
	d := &Daemon{configStore: &config.Config{}}
	c := &container.Container{
		AppArmorProfile: "",
		HostConfig:      &containertypes.HostConfig{Privileged: true},
	}
	ec := &exec.Config{Privileged: false}
	p := &specs.Process{}

	err := d.execSetPlatformOpt(c, ec, p)
	assert.NilError(t, err)
	assert.Equal(t, unconfinedAppArmorProfile, p.ApparmorProfile)
}
