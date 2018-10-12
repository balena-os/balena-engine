// +build go1.8,!windows,amd64,!static_build

package plugin

func loadPlugins(path string) error {
	// no plugin support to avoid binary size costs in balena
	return nil
}
