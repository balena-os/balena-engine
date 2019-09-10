package devnotify

import (
	"github.com/sirupsen/logrus"
)

// CloneTree will create a duplicate of /dev under `dest`.
func CloneTree(dest string) error {
	logger := logrus.WithFields(logrus.Fields{
		"is":   "CloneTree",
		"dest": dest,
	})
	logger.Infof("Cloning %v device(s)", len(devices.t))

	for _, d := range devices.t {
		if err := createDevice(dest, d); err != nil {
			return err
		}
		logger.Debugf("Cloned %v", d.Path)
	}

	return nil
}
