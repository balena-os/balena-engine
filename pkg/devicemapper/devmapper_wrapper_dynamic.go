//go:build linux && cgo && !static_build && !no_devmapper
// +build linux,cgo,!static_build,!no_devmapper

package devicemapper // import "github.com/docker/docker/pkg/devicemapper"

// #cgo pkg-config: devmapper
import "C"
