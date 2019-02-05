// +build linux freebsd openbsd darwin

package client // import "github.com/docker/docker/client"

// DefaultDockerHost defines os specific default if DOCKER_HOST is unset
const DefaultDockerHost = "unix:///var/run/balena-engine.sock"

const defaultProto = "unix"
const defaultAddr = "/var/run/balena-engine.sock"
