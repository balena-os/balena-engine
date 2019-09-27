package devnotify

import (
	"github.com/docker/docker/layer"
	"github.com/docker/docker/pkg/containerfs"
)

func WrapInitFunc(initFunc layer.MountInit, watcher Watcher) layer.MountInit {
	return func(initPath containerfs.ContainerFS) error {
		err := initFunc(initPath)
		if err != nil {
			return err
		}
		return watcher.Prepare(initPath.Path())
	}
}
