// +build !windows,!linux

package chrootarchive // import "github.com/docker/docker/pkg/chrootarchive"

func set_ionice(dst string) {
	// noop
}
