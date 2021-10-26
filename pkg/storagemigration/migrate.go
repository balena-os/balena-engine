package storagemigration // import "github.com/docker/docker/pkg/storagemigration"

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"golang.org/x/sys/unix"

	"github.com/docker/docker/pkg/archive"
)

// Migrate migrates the state of the storage from aufs -> overlay2
func Migrate(root string) (err error) {
	// rollback partial migration
	defer func() {
		if err != nil {
			logrus.WithField("storage_root", root).WithError(err).Error("failed aufs to overlay2 migration")
			if cleanupErr := failCleanup(root); cleanupErr != nil {
				err = errors.Wrapf(err, "error cleaning up: %v", cleanupErr)
				return
			}
			// clear the error; we don't want to propagate it to the daemon,
			// otherwise it won't be able to start at all
			err = nil
		}
	}()

	if logpath, ok := os.LookupEnv("BALENA_MIGRATE_OVERLAY_LOGFILE"); ok {
		// setup a logrus hook to duplicate logs at <logpath>
		var teardownLogs func()
		teardownLogs, err = setupLogs(logpath)
		if err != nil {
			return err
		}
		// remove the logrus hook
		defer func() {
			teardownLogs()
			if err == nil {
				// cleanup the logs file if there was no errors during Migrate
				os.Remove(logpath)
			}
		}()
	}

	logrus.WithField("storage_root", root).Debug("starting aufs to overlay2 migration")

	// make sure we actually have an aufs tree to migrate from
	aufsRootExists, err := exists(aufsRoot(root), true)
	if err != nil {
		return err
	}
	// if we don't there is nothing to do
	if !aufsRootExists {
		logrus.Infof("Storage migration skipped: %s", err)
		return nil
	}

	// make sure there isn't an overlay2 tree already
	overlayRootExists, err := exists(overlayRoot(root), true)
	if err != nil {
		return err
	}
	if overlayRootExists {
		// if both roots exist, assume migration succeeded during a previous run
		// and commit (dropping aufs data)
		logrus.Infof("Storage migration completed, cleaning up aufs storage")
		if err := commit(root); err != nil {
			return errors.Wrap(err, "failed to commit storage migration")
		}
		return nil
	}

	logrus.Info("Storage migration from aufs to overlay2 starting")
	startT := time.Now()
	defer func() {
		logrus.Infof("Storage migration finished, took %s", time.Now().Sub(startT))
	}()

	// Scan aufs layer data and build structure holding all the relevant information
	// needed to replicate on overlayfs.
	// We need to pay special attention to the whiteout metadata files used by aufs to
	// mark deleted files and empty directories.
	state, err := createState(aufsRoot(root))
	if err != nil {
		return err
	}

	logrus.Infof("transforming %d layers(s) to overlay2", len(state.Layers))

	// Build up the overlayfs layer data from the state structure.
	// This ignores the special files at first and just replicates the data
	// (using hardlinks to save space).
	// In a second step we use the state.Meta data to delete aufs whiteout files
	// and create the special files / set file attributes used by overlayfs.
	if err := transformStateToOverlay(root, state); err != nil {
		return err
	}

	// Finalize the migration:
	// - duplicate aufs images to $storageRoot/image/overlay2
	// - move temp dir holding overlay layer data to $storageRoot/overlay
	// - edit container config to use overlay storage driver

	var (
		aufsImageDir    = filepath.Join(root, "image", "aufs")
		overlayImageDir = filepath.Join(root, "image", "overlay2")
	)
	if ok, _ := exists(aufsImageDir, true); ok {
		logrus.Debug("moving aufs images to overlay2")
		err = replicate(aufsImageDir, overlayImageDir)
		if err != nil {
			return fmt.Errorf("Error moving images from aufs to overlay: %v", err)
		}
	}

	logrus.Debug("moving layer data from temporary location to overlay2 root")
	err = os.Rename(tempTargetRoot(root), overlayRoot(root))
	if err != nil {
		return fmt.Errorf("Error moving from temporary root: %v", err)
	}

	err = SwitchAllContainersStorageDriver(root, "overlay2")
	if err != nil {
		return fmt.Errorf("Error migrating containers to overlay2: %v", err)
	}

	return nil
}

func createState(aufsDir string) (*State, error) {
	var state State

	diffDir := filepath.Join(aufsDir, "diff")

	// get all layers
	layerIDs, err := LoadIDs(diffDir)
	if err != nil {
		return nil, fmt.Errorf("Error loading layer ids: %v", err)
	}
	logrus.Debugf("processing %d aufs layers(s)", len(layerIDs))

	for _, layerID := range layerIDs {
		layer := Layer{ID: layerID}

		// get parent layers
		parentIDs, err := GetParentIDs(aufsDir, layerID)
		if err != nil {
			return nil, fmt.Errorf("Error loading parent IDs for %s: %v", layerID, err)
		}
		layer.ParentIDs = parentIDs

		layerDir := filepath.Join(diffDir, layerID)
		err = filepath.Walk(layerDir, func(path string, fi os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			absPath, err := filepath.Rel(layerDir, path)
			if err != nil {
				return err
			}

			if IsWhiteout(fi.Name()) {
				if IsWhiteoutMeta(fi.Name()) {
					if IsOpaqueParentDir(fi.Name()) {
						layer.Meta = append(layer.Meta, Meta{
							Path: filepath.Dir(absPath),
							Type: MetaOpaque,
						})
						return nil
					}

					// other whiteout metadata
					layer.Meta = append(layer.Meta, Meta{
						Path: absPath,
						Type: MetaOther,
					})
					return nil
				}

				// simple whiteout file
				layer.Meta = append(layer.Meta, Meta{
					Path: filepath.Join(filepath.Dir(absPath), StripWhiteoutPrefix(fi.Name())),
					Type: MetaWhiteout,
				})
				return nil
			}

			return nil
		})
		if err != nil {
			return nil, fmt.Errorf("Error walking filetree for %s: %v", layerID, err)
		}

		state.Layers = append(state.Layers, layer)
	}
	return &state, nil
}

func transformStateToOverlay(root string, state *State) error {
	// move to overlay filetree
	for _, layer := range state.Layers {

		layerDir := filepath.Join(tempTargetRoot(root), layer.ID)

		// create /:layer_id dir
		err := os.MkdirAll(layerDir, os.ModeDir|0700)
		if err != nil {
			return fmt.Errorf("Error creating layer directory for %s: %v", layer.ID, err)
		}

		// create /:layer_id/link file and /l/:layer_ref file
		_, err = CreateLayerLink(tempTargetRoot(root), layer.ID)
		if err != nil {
			return fmt.Errorf("Error creating layer link dir for %s: %v", layer.ID, err)
		}

		// create /:layer_id/lower
		var lower string
		for _, parentID := range layer.ParentIDs {

			parentLayerDir := filepath.Join(tempTargetRoot(root), parentID)
			ok, err := exists(parentLayerDir, true)
			if err != nil {
				return fmt.Errorf("Error checking for parent layer dir for %s: %v", layer.ID, err)
			}
			if !ok {
				// parent layer hasn't been processed separately yet.
				err := os.MkdirAll(parentLayerDir, os.ModeDir|0700)
				if err != nil {
					return fmt.Errorf("Error creating layer directory for parent layer %s: %v", parentID, err)
				}
			}
			parentRef, err := CreateLayerLink(tempTargetRoot(root), parentID)
			if err != nil {
				return fmt.Errorf("Error creating layer link dir for parent layer %s: %v", parentID, err)
			}
			lower = AppendLower(lower, parentRef)
		}
		// if this layer had parents
		if lower != "" {
			lowerFile := filepath.Join(layerDir, "lower")
			err := ioutil.WriteFile(lowerFile, []byte(lower), 0644)
			if err != nil {
				return fmt.Errorf("Error creating lower file for %s: %v", layer.ID, err)
			}
			layerWorkDir := filepath.Join(layerDir, "work")
			err = os.MkdirAll(layerWorkDir, os.ModeDir|0700)
			if err != nil {
				return fmt.Errorf("Error creating work dir for %s: %v", layer.ID, err)
			}
		}

		logrus.WithField("layer_id", layer.ID).Debug("hardlinking aufs data to overlay")
		var (
			overlayLayerDir = filepath.Join(layerDir, "diff")
			aufsLayerDir    = filepath.Join(aufsRoot(root), "diff", layer.ID)
		)
		err = replicate(aufsLayerDir, overlayLayerDir)
		if err != nil {
			return fmt.Errorf("Error moving layer data to overlay2: %v", err)
		}

		// migrate metadata files
		logrus.WithField("layer_id", layer.ID).Debugf("processing %d metadata file(s)", len(layer.Meta))
		for _, meta := range layer.Meta {
			metaPath := filepath.Join(overlayLayerDir, meta.Path)

			switch meta.Type {
			case MetaOpaque:
				// set the opque xattr
				err := SetOpaque(metaPath)
				if err != nil {
					return fmt.Errorf("Error marking %s as opque: %v", metaPath, err)
				}
				// remove aufs metadata file
				aufsMetaFile := filepath.Join(metaPath, archive.WhiteoutOpaqueDir)
				err = os.Remove(aufsMetaFile)
				if err != nil {
					return fmt.Errorf("Error removing opque meta file: %v", err)
				}

			case MetaWhiteout:
				// create the 0x0 char device
				err := SetWhiteout(metaPath)
				if err != nil {
					return fmt.Errorf("Error marking %s as whiteout: %v", metaPath, err)
				}
				metaDir, metaFile := filepath.Split(metaPath)
				aufsMetaFile := filepath.Join(metaDir, archive.WhiteoutPrefix+metaFile)

				// chown the new char device with the old uid/gid
				uid, gid, err := getUIDAndGID(aufsMetaFile)
				if err != nil {
					return fmt.Errorf("Error getting UID and GID: %v", err)
				}
				err = unix.Chown(metaPath, uid, gid)
				if err != nil {
					return fmt.Errorf("Error chowning character device: %v", err)
				}

				err = os.Remove(aufsMetaFile)
				if err != nil {
					return fmt.Errorf("Error removing aufs whiteout file: %v", err)
				}

			case MetaOther:
				logrus.WithField("meta_type", "whiteoutmeta").Debugf("removing %s from overlay", metaPath)
				err = os.Remove(metaPath)
				if err != nil {
					return fmt.Errorf("Error removing useless aufs meta file at: %v", err)
				}
			}
		}
	}
	return nil
}

// SwitchAllContainersStorageDriver iterates over all containers and configures
// them to use `newStorageDriver`.
func SwitchAllContainersStorageDriver(root, newStorageDriver string) error {
	containerDir := filepath.Join(root, "containers")
	if ok, _ := exists(containerDir, true); !ok {
		return nil
	}

	containerIDs, err := LoadIDs(containerDir)
	if err != nil {
		return fmt.Errorf("Error listing containers: %v", err)
	}
	logrus.Debugf("migrating %v container(s) to %s", len(containerIDs), newStorageDriver)
	for _, containerID := range containerIDs {
		err := switchContainerStorageDriver(root, containerID, newStorageDriver)
		if err != nil {
			return fmt.Errorf("Error rewriting container config for %s: %v", containerID, err)
		}
		logrus.WithField("container_id", containerID).Debugf("reconfigured storage-driver to %s", newStorageDriver)
	}
	return nil
}
