package balena

import (
	"os"
	"os/exec"
	"path"

	"github.com/containerd/containerd/plugin"
)

// DisableShimPlugins returns true if the plugin represented by r should be
// disabled.
//
// While migrating to the shim runtime v2, I shed many tears to understand this,
// so I'll try to explain it in detail for the future generations and the future
// me. I hope I got this right.
//
// containerd-shim-runc-v2 seems to be strongly based on a plugin system that is
// basically the same containerd itself used before. So, at startup, both will
// load all registered plugins. How are these plugins registered? The plugins
// register themselves by calling `plugin.Register()` (from
// `github.com/containerd/containerd/plugin`) in an `init()` function. So far so
// good.
//
// Things start to get tricky when we realize that containerd and
// containerd-shim-runc-v2 need a different set of plugins. In standard Moby
// this is not a problem, because each of these binaries will import a different
// set of packages, and therefore will only run the correct subset of `init()`s,
// and thus will register only the plugins they need.
//
// The problem is that in balenaEngine we have one single binary that
// amalgamates containerd, containerd-shim-runc-v2, and more. So, that single
// big binary will run every `init()` from every package imported by everyone.
// As a result, containerd-shim-runc-v2 will try to load all plugins in the
// world -- including some that will cause it to crash.
//
// (Sidebar: Why do some plugins cause containerd-shim-runc-v2 to crash? AFAIU,
// it's essentially because the shim plugin system is not as advanced as the one
// in containerd. For example, containerd-shim-runc-v2 will not pass
// "configurations" to the plugins, so plugins that expect to get configs will
// crash. For details, see the `run()` function in
// `vendor/github.com/containerd/containerd/runtime/v2/shim/shim.go` and search
// for "TODO". There is some code that would prepare the configs so they could
// be later accessed by the plugins, but it is commented out since forever. This
// commented out code seems to come from
// `vendor/github.com/containerd/containerd/services/server/server.go`, which I
// believe is the equivalent containerd code.)
//
// And this finally brings us to this function. We use it to filter out the
// containerd-shim-runc-v2 plugins that we don't want to load. Or rather, to
// filter out all plugins except the few ones that really needed.
//
// The End.
func DisableShimPlugins(r *plugin.Registration) bool {
	// These are the plugins we want to load.
	want := (r.Type == plugin.TTRPCPlugin && r.ID == "task") ||
		(r.Type == plugin.EventPlugin && r.ID == "publisher") ||
		(r.Type == plugin.InternalPlugin && r.ID == "shutdown")

	// Filter out anything we don't want.
	return !want
}

// Executable returns the absolute path to the executable binary or symlink used
// to start the current process. This is similar to `os.Executable()`, but it's
// guaranteed to not follow symlinks.
//
// The problem with following symlinks is that this breaks the amalgamated
// balenaEngine binary. For example, if running from
// `/usr/bin/balena-engine-containerd` (a symlink), `os.Executable()` would
// return `/usr/bin/balena-engine` (the linked binary).
//
// This is not a general replacement for `os.Executable()`, though. In
// particular, this implementation assumes the executable or symlink is present
// somewhere in the $PATH.
func Executable() (string, error) {
	exeName := path.Base(os.Args[0])
	return exec.LookPath(exeName)
}
