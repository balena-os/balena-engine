package devnotify

import (
	"context"
	"os"

	"github.com/sirupsen/logrus"

	"github.com/docker/docker/container"
)

// Watcher is responsible for mirroring and syncing the host devfs to the
// location passed to `Prepare`.
type Watcher interface {
	Prepare(path string) error
	Start() error
	Stop() error
}

func NewWatcher(c *container.Container) Watcher {
	return &watcher{containerID: c.ID}
}

type watcher struct {
	ctx         context.Context
	cancelfn    context.CancelFunc
	containerID string
	path        string
}

// Prepare sets up a mirror of the current devfs at `path`. This does not start
// syncing any changes.
func (w *watcher) Prepare(path string) error {
	w.path = path

	logger := logrus.WithFields(logrus.Fields{
		"balenaext": "devfs-watcher",
		"container": w.containerID,
	})
	logger.Warning("hello from balena-devfs initFunc / watcher.Prepare")

	return CloneTree(logger.Logger, w.path)
}

// Start syncs changes from `/dev` to the destination that was passed to
// `Prepare`.
//
// It doesn't block and exits right away.
func (w *watcher) Start() error {
	logger := logrus.WithFields(logrus.Fields{
		"balenaext": "devfs-watcher",
		"container": w.containerID,
	})

	if _, err := os.Stat(w.path); err != nil {
		return err
	}

	ctx := context.TODO()
	w.ctx, w.cancelfn = context.WithCancel(ctx)

	logger.WithField("target-dir", w.path).Warn("Watching host devfs")
	go SyncTree(w.ctx, logger.Logger, w.path)

	return nil
}

// Stop cancels the running watcher.
//
// It doesn't take care of cleaning up the sync destination.
func (w *watcher) Stop() error {
	w.cancelfn()
	return w.ctx.Err()
}
