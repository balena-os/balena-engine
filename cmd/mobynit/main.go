package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"syscall"

	_ "github.com/docker/docker/daemon/graphdriver/overlay2"
	"github.com/docker/docker/layer"
	"github.com/docker/docker/pkg/idtools"
	"golang.org/x/sys/unix"
)

const (
	LAYER_ROOT = "/docker"
	PIVOT_PATH = "/mnt/sysroot/active"
)

func mountContainer(containerID string) string {
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
		IDMapping:                 idtools.IdentityMapping{},
		GraphDriver:               graphDriver,
	})
	if err != nil {
		log.Fatal("error loading layer store:", err)
	}

	rwlayer, err := ls.GetRWLayer(containerID)
	if err != nil {
		log.Fatal("error getting container layer:", err)
	}

	newRootPath, err := rwlayer.Mount("")
	if err != nil {
		log.Fatal("error mounting container fs:", err)
	}

	if err := unix.Mount("", newRootPath, "", unix.MS_REMOUNT, ""); err != nil {
		log.Fatal("error remounting container as read/write:", err)
	}
	defer unix.Mount("", newRootPath, "", unix.MS_REMOUNT|unix.MS_RDONLY, "")

	if err := os.MkdirAll(filepath.Join(newRootPath, PIVOT_PATH), os.ModePerm); err != nil {
		log.Fatal("creating /mnt/sysroot failed:", err)
	}

	return newRootPath
}

func main() {
	flag.Parse()

	current, err := os.Readlink("/current")
	if err != nil {
		log.Fatal("could not get container ID:", err)
	}
	containerID := filepath.Base(current)

	if err := unix.Mount("", "/", "", unix.MS_REMOUNT, ""); err != nil {
		log.Fatal("error remounting root as read/write:", err)
	}

	newRoot := mountContainer(containerID)

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
