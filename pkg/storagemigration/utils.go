package storagemigration

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/sys/unix"
)

// exists checks if a file  (or if isDir is set to "true" a directory) exists
func exists(path string, isDir bool) (bool, error) {
	fi, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	if fi.IsDir() != isDir {
		return false, nil
	}
	return true, nil
}

// getUIDAndGID retrieves user and group id for path
func getUIDAndGID(path string) (uid, gid int, err error) {
	var fi unix.Stat_t
	err = unix.Stat(path, &fi)
	if err != nil {
		return 0, 0, err
	}
	return int(fi.Uid), int(fi.Gid), nil
}

func removeDirIfExists(path string) error {
	ok, err := exists(path, true)
	if err != nil {
		return err
	}
	if ok {
		err = os.RemoveAll(path)
		if err != nil {
			return err
		}
	}
	return nil
}

// replicate hardlinks all files from sourceDir to targetDir, reusing the same
// file structure
func replicate(sourceDir, targetDir string) error {
	return filepath.Walk(sourceDir, func(path string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		var (
			targetPath = strings.Replace(path, sourceDir, targetDir, 1)
		)
		if fi.IsDir() {
			err = os.MkdirAll(targetPath, os.ModeDir|0755)
			if err != nil {
				return err
			}
		} else {
			err = os.Link(path, targetPath)
			if err != nil {
				return err
			}
		}

		return nil
	})
}

// switchContainerStorageDriver rewrites the container config to use a new storage driver,
// this is the only change needed to make it work after the migration
func switchContainerStorageDriver(root, containerID, newStorageDriver string) error {
	containerConfigPath := filepath.Join(root, "containers", containerID, "config.v2.json")
	f, err := os.OpenFile(containerConfigPath, os.O_RDWR, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()

	var containerConfig = make(map[string]interface{})
	err = json.NewDecoder(f).Decode(&containerConfig)
	if err != nil {
		return err
	}
	containerConfig["Driver"] = newStorageDriver

	err = f.Truncate(0)
	if err != nil {
		return err
	}
	_, err = f.Seek(0, 0)
	if err != nil {
		return err
	}
	err = json.NewEncoder(f).Encode(&containerConfig)
	if err != nil {
		return err
	}
	err = f.Sync()
	if err != nil {
		return err
	}
	return nil
}
