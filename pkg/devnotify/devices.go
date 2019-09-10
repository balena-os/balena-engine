package devnotify

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"syscall"

	libcontainer_configs "github.com/opencontainers/runc/libcontainer/configs"
	libcontainer_devices "github.com/opencontainers/runc/libcontainer/devices"
	"github.com/sirupsen/logrus"
)

const slashDev = "/dev"

var devices *DeviceTable

func init() {
	devices = NewDeviceTable()
	devices.Sync()
}

// DeviceTable allows package wide access to the list of devices under
// /dev, as filtered by
// github.com/opencontainers/runc/libcontainer/devices.GetDevices
//
type DeviceTable struct {
	t      map[string]*libcontainer_configs.Device
	mtx    *sync.Mutex
	logger logrus.FieldLogger
}

func NewDeviceTable() *DeviceTable {
	return &DeviceTable{
		t:      make(map[string]*libcontainer_configs.Device),
		mtx:    new(sync.Mutex),
		logger: logrus.WithField("is", "DeviceTable"),
	}
}

func (dt *DeviceTable) Lookup(path string) (d *libcontainer_configs.Device, ok bool) {
	d, ok = dt.t[path]
	return d, ok
}

func (dt *DeviceTable) Sync() {
	dt.logger.Debug("Syncing devices")
	devs, err := libcontainer_devices.GetDevices(slashDev)
	if err != nil {
		dt.logger.WithError(err).Fatal("Failed to get devices")
		return
	}
	dt.mtx.Lock()
	defer dt.mtx.Unlock()
	// reset first to catch disappearing devices
	// dt.logger.Debug("Resetting devices")
	dt.t = make(map[string]*libcontainer_configs.Device)
	for _, dev := range devs {
		// dt.logger.Debugf("Adding %v", dev.Path)
		dt.t[dev.Path] = dev
	}
}

// helper function to create a device node
//
// taken from createDeviceNode + mknodDevice:
// https://github.com/opencontainers/runc/blob/v0.1.1/libcontainer/rootfs_linux.go#L439
//
//
func createDevice(root string, node *libcontainer_configs.Device) error {
	dest := filepath.Join(root, node.Path)
	if err := os.MkdirAll(filepath.Dir(dest), 0755); err != nil {
		return err
	}
	if err := mknodDevice(dest, node); err != nil {
		if os.IsExist(err) {
			return nil
		}
		return err
	}
	return nil
}
func mknodDevice(dest string, node *libcontainer_configs.Device) error {
	fileMode := node.FileMode
	switch node.Type {
	case 'c':
		fileMode |= syscall.S_IFCHR
	case 'b':
		fileMode |= syscall.S_IFBLK
	default:
		return fmt.Errorf("%c is not a valid device type for device %s", node.Type, node.Path)
	}
	if err := syscall.Mknod(dest, uint32(fileMode), node.Mkdev()); err != nil {
		return err
	}
	return syscall.Chown(dest, int(node.Uid), int(node.Gid))
}
