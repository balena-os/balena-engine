//go:build linux && !seccomp

package seccomp

import (
	"errors"

	"github.com/opencontainers/runtime-spec/specs-go"
)

// GetDefaultProfile returns no profile when the engine is built without
// the seccomp build tag. The embedded runc is then compiled without
// libseccomp and rejects any seccomp configuration, so callers such as the
// BuildKit executor must run containers unconfined, matching the daemon.
func GetDefaultProfile(_ *specs.Spec) (*specs.LinuxSeccomp, error) {
	return nil, nil
}

// LoadProfile returns an error because a custom profile cannot be applied
// without seccomp support.
func LoadProfile(_ string, _ *specs.Spec) (*specs.LinuxSeccomp, error) {
	return nil, errors.New("seccomp support is not compiled into this engine")
}
