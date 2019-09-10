package daemon

import (
	"context"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/balena-os/ctrdev"
	"github.com/docker/docker/container"
	"github.com/opencontainers/runtime-spec/specs-go"
	"github.com/sirupsen/logrus"
)

const devfsWatcherRoot = "/var/run/balena-engine/devfs-watch"

type devfsWatcher struct {
	ctx      context.Context
	cancelfn context.CancelFunc
	path     string
}

func (watcher *devfsWatcher) stop() error {
	watcher.cancelfn()
	return watcher.ctx.Err()
}

func makeWatcherDir(c *container.Container) string {
	return filepath.Join(devfsWatcherRoot, c.ID)
}

func (daemon *Daemon) attachDevfsWatcher(c *container.Container, s *specs.Spec) error {
	logger := logrus.WithFields(logrus.Fields{
		"balenaext": "devfs-watch",
		"container": c.ID,
	})

	if daemon.devfsWatchers == nil {
		daemon.devfsWatchers = make(map[string]*devfsWatcher)
	}

	w := &devfsWatcher{
		path: makeWatcherDir(c),
	}

	ctx, cancel := context.WithCancel(context.TODO())

	logger.WithField("target-dir", w.path).Warn("Watching host devfs")

	if err := ctrdev.CloneTree(w.path); err != nil {
		return fmt.Errorf("Failed to clone host devfs: %v", err)
	}
	go ctrdev.SyncTree(ctx, w.path)

	w.ctx = ctx
	w.cancelfn = cancel

	// manipulate the mount map to use our devfs
	wpathDev := filepath.Join(w.path, "dev")
	files, err := ioutil.ReadDir(wpathDev)
	if err != nil {
		return err
	}
	for _, f := range files {
		s.Mounts = append(s.Mounts, specs.Mount{
			Destination: filepath.Join("/dev", f.Name()),
			Source:      filepath.Join(wpathDev, f.Name()),
			Type:        "bind",
			Options:     []string{},
		})
	}
	// inspect mount map
	logger.Warn(">>>! inspecting mount map")
	for _, m := range s.Mounts {
		logger.Warnf(">>> %#+v", m)
	}
	return nil
}
