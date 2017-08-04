package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"syscall"

	_ "github.com/docker/docker/daemon/graphdriver/aufs"
	_ "github.com/docker/docker/daemon/graphdriver/overlay2"
	"github.com/docker/docker/layer"
	"github.com/docker/docker/pkg/idtools"
	"github.com/docker/docker/pkg/mount"
	"golang.org/x/sys/unix"
)

const (
	LAYER_ROOT = "/docker"
	PIVOT_PATH = "/mnt/sysroot"
)

var graphDriver string

func init() {
	flag.StringVar(&graphDriver, "storage-driver", "aufs", "Storage driver to use")
	flag.StringVar(&graphDriver, "s", "aufs", "Storage driver to use")
}

func mountContainer(containerID string) string {
	if err := unix.Mount("", "/", "", unix.MS_REMOUNT, ""); err != nil {
		log.Fatal("error remounting root as read/write:", err)
	}
	defer unix.Mount("", "/", "", unix.MS_REMOUNT | unix.MS_RDONLY, "")

	if err := os.MkdirAll("/dev/shm", os.ModePerm); err != nil {
		log.Fatal("creating /dev/shm failed:", err)
	}

	if err := unix.Mount("shm", "/dev/shm", "tmpfs", 0, ""); err != nil {
		log.Fatal("error mounting /dev/shm:", err)
	}
	defer unix.Unmount("/dev/shm", unix.MNT_DETACH)

	ls, err := layer.NewStoreFromOptions(layer.StoreOptions{
		Root:                      LAYER_ROOT,
		MetadataStorePathTemplate: filepath.Join(LAYER_ROOT, "image", "%s", "layerdb"),
		IDMapping:                 &idtools.IdentityMapping{},
		GraphDriver:               graphDriver,
		Platform:                  "linux",
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

	if err := unix.Mount("", newRoot, "", unix.MS_REMOUNT, ""); err != nil {
		log.Fatal("error remounting container as read/write:", err)
	}
	defer unix.Mount("", newRoot, "", unix.MS_REMOUNT | unix.MS_RDONLY, "")

	if err := os.MkdirAll(filepath.Join(newRoot, PIVOT_PATH), os.ModePerm); err != nil {
		log.Fatal("creating /mnt/sysroot failed:", err)
	}

	return newRoot
}

func main() {
	flag.Parse()

	// Any mounts done by initrd will be transfered in the new root
	mounts, err := mount.GetMounts(nil)
	if err != nil {
		log.Fatal("could not get container ID:", err)
	}
	containerID := filepath.Base(current)

	newRoot := mountContainer(containerID)

	for _, mount := range mounts {
		if mount.Mountpoint == "/" {
			continue
		}
		if err := unix.Mount(mount.Mountpoint, filepath.Join(newRoot, mount.Mountpoint), "", unix.MS_MOVE, ""); err != nil {
			log.Println("could not move mountpoint:", mount.Mountpoint, err)
		}
	}

	if err := syscall.PivotRoot(newRoot, filepath.Join(newRoot, PIVOT_PATH)); err != nil {
		log.Fatal("error while pivoting root:", err)
	}

	if err := unix.Chdir("/"); err != nil {
		log.Fatal(err)
	}

	if err := syscall.Exec("/sbin/init", os.Args, os.Environ()); err != nil {
		log.Fatal("error executing init:", err)
	}
}
