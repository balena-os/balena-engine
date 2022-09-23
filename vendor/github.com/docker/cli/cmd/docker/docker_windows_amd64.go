//go:build windows && amd64
// +build windows,amd64

//go:generate goversioninfo -64=true -o=../../cli/winresources/resource.syso -icon=winresources/docker.ico -manifest=winresources/docker.exe.manifest ../../cli/winresources/versioninfo.json

package docker

import _ "github.com/docker/cli/cli/winresources"
