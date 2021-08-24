package storagemigration

import (
	"fmt"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

// FailCleanup should be run after a failed migration.
// It will remove any files left over from the migration process
// and migrate containers back to aufs.
//
func failCleanup(root string) error {
	logrus.WithField("storage_root", root).Warning("recovering from failed aufs to overlay migration")

	var err error
	err = SwitchAllContainersStorageDriver(root, "aufs")
	if err != nil {
		return fmt.Errorf("Error migrating containers to aufs: %v", err)
	}

	err = removeDirIfExists(tempTargetRoot(root))
	if err != nil {
		return err
	}

	err = removeDirIfExists(overlayRoot(root))
	if err != nil {
		return err
	}

	overlayImageDir := filepath.Join(root, "image", "overlay2")
	err = removeDirIfExists(overlayImageDir)
	if err != nil {
		return err
	}

	return nil
}
