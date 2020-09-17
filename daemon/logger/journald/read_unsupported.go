// +build !linux !cgo !journald

package journald // import "github.com/docker/docker/daemon/logger/journald"

func (s *journald) Close() error {
	return nil
}
