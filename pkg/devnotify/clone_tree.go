package devnotify

import (
	"github.com/sirupsen/logrus"
)

// CloneTree will create a duplicate of /dev under `dest`.
func CloneTree(logger logrus.FieldLogger, dest string) error {
	logger.Infof("Cloning %v device(s)", len(devices.t))

	for _, d := range devices.t {
		if err := createDevice(dest, d); err != nil {
			return err
		}
		logger.Debugf("Cloned %v", d.Path)
	}

	return nil
}
