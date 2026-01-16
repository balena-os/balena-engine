//go:build !windows

package dockerd

import (
	"testing"

	"github.com/docker/docker/daemon/config"
	"github.com/spf13/pflag"
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
	"gotest.tools/v3/fs"
)

func TestLoadDaemonCliConfigWithDaemonFlags(t *testing.T) {
	content := `{"log-opts": {"max-size": "1k"}}`
	tempFile := fs.NewFile(t, "config", fs.WithContent(content))
	defer tempFile.Remove()

	opts := defaultOptions(t, tempFile.Path())
	opts.Debug = true
	opts.LogLevel = "info"
	assert.Check(t, opts.flags.Set("selinux-enabled", "true"))

	loadedConfig, err := loadDaemonCliConfig(opts)
	assert.NilError(t, err)
	assert.Assert(t, loadedConfig != nil)

	assert.Check(t, loadedConfig.Debug)
	assert.Check(t, is.Equal("info", loadedConfig.LogLevel))
	assert.Check(t, loadedConfig.EnableSelinuxSupport)
	assert.Check(t, is.Equal("json-file", loadedConfig.LogConfig.Type))
	assert.Check(t, is.Equal("1k", loadedConfig.LogConfig.Config["max-size"]))
}

func TestLoadDaemonConfigWithNetwork(t *testing.T) {
	content := `{"bip": "127.0.0.2", "ip": "127.0.0.1"}`
	tempFile := fs.NewFile(t, "config", fs.WithContent(content))
	defer tempFile.Remove()

	opts := defaultOptions(t, tempFile.Path())
	loadedConfig, err := loadDaemonCliConfig(opts)
	assert.NilError(t, err)
	assert.Assert(t, loadedConfig != nil)

	assert.Check(t, is.Equal("127.0.0.2", loadedConfig.IP))
	assert.Check(t, is.Equal("127.0.0.1", loadedConfig.DefaultIP.String()))
}

func TestLoadDaemonConfigWithMapOptions(t *testing.T) {
	content := `{"log-opts": {"tag": "test"}}`
	tempFile := fs.NewFile(t, "config", fs.WithContent(content))
	defer tempFile.Remove()

	opts := defaultOptions(t, tempFile.Path())
	loadedConfig, err := loadDaemonCliConfig(opts)
	assert.NilError(t, err)
	assert.Assert(t, loadedConfig != nil)
	assert.Check(t, loadedConfig.LogConfig.Config != nil)
	assert.Check(t, is.Equal("test", loadedConfig.LogConfig.Config["tag"]))
}

func TestLoadDaemonConfigWithTrueDefaultValues(t *testing.T) {
	content := `{ "userland-proxy": false }`
	tempFile := fs.NewFile(t, "config", fs.WithContent(content))
	defer tempFile.Remove()

	// Use a custom options setup that doesn't pre-set the userland-proxy flag,
	// so we can test the file config taking precedence.
	cfg, err := config.New()
	assert.NilError(t, err)
	opts := newDaemonOptions(cfg)
	opts.flags = &pflag.FlagSet{}
	opts.installFlags(opts.flags)
	err = installConfigFlags(opts.daemonConfig, opts.flags)
	assert.NilError(t, err)
	opts.flags.StringVar(&opts.configFile, "config-file", "", "")
	opts.configFile = tempFile.Path()
	err = opts.flags.Parse([]string{})
	assert.NilError(t, err)

	loadedConfig, err := loadDaemonCliConfig(opts)
	assert.NilError(t, err)
	assert.Assert(t, loadedConfig != nil)

	assert.Check(t, !loadedConfig.EnableUserlandProxy)

	// make sure reloading doesn't generate configuration
	// conflicts after normalizing boolean values.
	reload := func(reloadedConfig *config.Config) {
		assert.Check(t, !reloadedConfig.EnableUserlandProxy)
	}
	assert.Check(t, config.Reload(opts.configFile, opts.flags, reload))
}

func TestLoadDaemonConfigWithTrueDefaultValuesLeaveDefaults(t *testing.T) {
	// In balena-engine, the default for userland-proxy is effectively false in tests
	// because enabling it requires a valid proxy binary path. This test verifies
	// that when userland-proxy is disabled (the safe default), the config loads successfully.
	tempFile := fs.NewFile(t, "config", fs.WithContent(`{}`))
	defer tempFile.Remove()

	opts := defaultOptions(t, tempFile.Path())
	loadedConfig, err := loadDaemonCliConfig(opts)
	assert.NilError(t, err)
	assert.Assert(t, loadedConfig != nil)

	// In balena-engine tests, userland-proxy defaults to false since
	// enabling it requires a valid userland-proxy-path.
	assert.Check(t, !loadedConfig.EnableUserlandProxy)
}
