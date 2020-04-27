package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/docker/docker/pkg/mount"
	"golang.org/x/sys/unix"
	"github.com/docker/docker/cmd/mobynit/hostapp"
)

const (
	LAYER_ROOT = "/balena"
	PIVOT_PATH = "/mnt/sysroot/active"
)

func prepareForPivot(containerID, graphDriver string) string {
	if err := os.MkdirAll("/dev/shm", os.ModePerm); err != nil {
		log.Fatal("creating /dev/shm failed:", err)
	}

	if err := unix.Mount("shm", "/dev/shm", "tmpfs", 0, ""); err != nil {
		log.Fatal("error mounting /dev/shm:", err)
	}
	defer unix.Unmount("/dev/shm", unix.MNT_DETACH)

	newRootPath := hostapp.MountContainer(filepath.Join("", LAYER_ROOT), containerID, graphDriver)

	defer unix.Mount("", newRootPath, "", unix.MS_REMOUNT|unix.MS_RDONLY, "")

	if err := os.MkdirAll(filepath.Join(newRootPath, PIVOT_PATH), os.ModePerm); err != nil {
		log.Fatal("creating /mnt/sysroot failed:", err)
	}

	return newRootPath
}

func getStorageDriverAndContainerID(sysroot string) (string, string) {
	rawGraphDriver, err := ioutil.ReadFile(filepath.Join(sysroot, "/current/boot/storage-driver"))
	if err != nil {
		log.Fatal("could not get storage driver:", err)
	}
	graphDriver := strings.TrimSpace(string(rawGraphDriver))

	current, err := os.Readlink(filepath.Join(sysroot, "/current"))
	if err != nil {
		log.Fatal("could not get container ID:", err)
	}
	containerID := filepath.Base(current)

	return graphDriver, containerID
}

func main() {
	sysrootPtr := flag.String("sysroot", "", "root of partition e.g. /mnt/sysroot/inactive. Mount destination is returned in stdout")
	flag.Parse()

	// Any mounts done by initrd will be transfered in the new root
	mounts, err := mount.GetMounts(nil)
	if err != nil {
		log.Fatal("could not get mounts:", err)
	}

	var graphDriver, containerID string

	// If a custom sysroot is passed, use it instead of LAYER_ROOT
	if *sysrootPtr != "" {
		graphDriver, containerID = getStorageDriverAndContainerID(*sysrootPtr)
		newRootPath := hostapp.MountContainer(filepath.Join(*sysrootPtr, LAYER_ROOT), containerID, graphDriver)
		fmt.Print(newRootPath)
	} else {
		graphDriver, containerID = getStorageDriverAndContainerID("")

		if err := unix.Mount("", "/", "", unix.MS_REMOUNT, ""); err != nil {
			log.Fatal("error remounting root as read/write:", err)
		}

		newRoot := prepareForPivot(containerID, graphDriver)

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
}
