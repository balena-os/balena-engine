package daemon

import (
	"fmt"

	"github.com/docker/docker/container"
	"github.com/docker/docker/pkg/devnotify"
)

func (daemon *Daemon) createDevfsWatcher(c *container.Container) (devnotify.Watcher, error) {
	if daemon.devfsWatchers == nil {
		daemon.devfsWatchers = make(map[string]devnotify.Watcher)
	}

	if _, ok := daemon.devfsWatchers[c.ID]; ok {
		return nil, fmt.Errorf("devfs watcher for %v already registerd", c.ID)
	}

	w := devnotify.NewWatcher(c)
	daemon.devfsWatchers[c.ID] = w
	return w, nil
}
