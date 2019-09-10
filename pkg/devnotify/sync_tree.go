package devnotify

import (
	"context"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"github.com/fsnotify/fsnotify"
	libcontainer_configs "github.com/opencontainers/runc/libcontainer/configs"

	"github.com/docker/docker/pkg/filenotify"
)

// SyncTree will watch for changes under /dev/... and
// apply them to the cloned tree under `dest`.
//
// This blocks and handles filesystem events until it is
// manually aborted using the passed context.
//
func SyncTree(ctx context.Context, dest string) error {
	logger := logrus.WithFields(logrus.Fields{
		"is":   "SyncTree",
		"dest": dest,
	})

	logger.Debug("Setting up watcher")
	watcher, err := filenotify.New()
	if err != nil {
		return err
	}

	for _, d := range devices.t {
		err := watcher.Add(d.Path)
		if err != nil {
			return err
		}
	}

	for {
		select {
		// this is how we stop the file watcher
		case <-ctx.Done():
			logger.Warning("Syncing stopped")
			watcher.Close()
			return nil

		case err := <-watcher.Errors():
			logger.WithError(err).Debug("Watcher error")
			return err

		// this is where the magic happens...
		case ev := <-watcher.Events():
			logger.WithField("file", ev.Name).Debug("recv event")

			d, ok := getDeviceFromEvent(ev)
			if !ok {
				continue
			}

			// TODO: do we need rate-limiting for when
			// something get's really busy?

			logger = logger.WithFields(logrus.Fields{
				"event": ev.String(),
				"dev":   d.Path,
			})

			err := syncDeviceChange(logger, dest, ev, d)
			if err != nil {
				logger.WithError(err).Error("Error syncing device")
			}
		}
	}
}

func getDeviceFromEvent(ev fsnotify.Event) (*libcontainer_configs.Device, bool) {
	// update the device table
	// TODO(urgent) need to think of a better way to do this we can't rerun
	// the GetDevices func every time we receive a create or delete event
	devices.Sync()
	return devices.Lookup(ev.Name)
}

func syncDeviceChange(logger logrus.FieldLogger,
	dest string,
	ev fsnotify.Event,
	d *libcontainer_configs.Device) error {

	// logger.WithFields(logrus.Fields{
	// 	"type":   d.Type,
	// 	"mkdev":  d.Mkdev(),
	// 	"cgroup": d.CgroupString,
	// 	"perm":   d.Permissions,
	// 	"uid":    d.Uid,
	// 	"gid":    d.Gid,
	// }).Debugf(">>> %v", filepath.Join(dest, d.Path))

	switch ev.Op {
	case fsnotify.Create:
		err := createDevice(dest, d)
		if err != nil {
			return err
		}
		logger.Info("Created new device node")

	case fsnotify.Remove:
		err := os.Remove(filepath.Join(dest, d.Path))
		if err != nil {
			// if os.IsNotExist(err) {
			// 	logger.WithError(err).Warning("Unable to delete device node")
			// 	return nil
			// }
			return err
		}
		logger.Info("Deleted device node")

	case fsnotify.Rename:
		logger.Warning("TODO: Rename")
		logger.Warningf("%#+v", ev)
		// err := os.Rename()
		// if err != nil {
		// 	return err
		// }

	}

	return nil
}
