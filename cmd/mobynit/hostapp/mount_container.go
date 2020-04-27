package hostapp

import (
	"log"
	"path/filepath"

	_ "github.com/docker/docker/daemon/graphdriver/aufs"
	_ "github.com/docker/docker/daemon/graphdriver/overlay2"
	"github.com/docker/docker/layer"
	"github.com/docker/docker/pkg/idtools"
	"golang.org/x/sys/unix"
)

func MountContainer(layer_root, containerID, graphDriver string) string {
	ls, err := layer.NewStoreFromOptions(layer.StoreOptions{
		Root:                      layer_root,
		MetadataStorePathTemplate: filepath.Join(layer_root, "image", "%s", "layerdb"),
		IDMapping:                 &idtools.IdentityMapping{},
		GraphDriver:               graphDriver,
		OS:                        "linux",
	})
	if err != nil {
		log.Fatal("error loading layer store:", err)
	}

	rwlayer, err := ls.GetRWLayer(containerID)
	if err != nil {
		log.Fatal("error getting container layer:", err)
	}

	newRoot, err := rwlayer.Mount("")
	if err != nil {
		log.Fatal("error mounting container fs:", err)
	}
	newRootPath := newRoot.Path()

	if err := unix.Mount("", newRootPath, "", unix.MS_REMOUNT, ""); err != nil {
		log.Fatal("error remounting container as read/write:", err)
	}

	return newRootPath
}
