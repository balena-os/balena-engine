// +build linux,cgo,journald,static_build

package journald // import "github.com/docker/docker/daemon/logger/journald"

// #cgo LDFLAGS: -lsystemd
import "C"
