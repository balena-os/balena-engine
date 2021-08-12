package storagemigration

import (
	"path/filepath"

	"github.com/sirupsen/logrus"
)

// Commit finalises the migration by deleting aufs storage root and images.
func commit(root string) error {
	logrus.WithField("storage_root", root).Debug("committing changes")

	// remove aufs layer data
	err := removeDirIfExists(aufsRoot(root))
	if err != nil {
		return err
	}

	// remove images
	aufsImageDir := filepath.Join(root, "image", "aufs")
	err = removeDirIfExists(aufsImageDir)
	if err != nil {
		return err
	}

	return nil
}
