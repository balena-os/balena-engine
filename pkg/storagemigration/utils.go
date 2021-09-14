package storagemigration

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/docker/docker/daemon/graphdriver/copy"
	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/writer"
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
	return copy.DirCopy(sourceDir, targetDir, copy.Hardlink, false)
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

	containerConfig := make(map[string]interface{})
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

func setupLogs(logpath string) (teardown func(), err error) {
	logfile, err := os.OpenFile(logpath, os.O_WRONLY|os.O_CREATE, 0444)
	if err != nil {
		return nil, err
	}

	// copy the previous hooks
	hooks := logrus.StandardLogger().Hooks
	prev := make(logrus.LevelHooks, len(hooks))
	for k, v := range hooks {
		prev[k] = v
	}

	logrus.AddHook(&writer.Hook{
		Writer: logfile,
		LogLevels: []logrus.Level{
			logrus.TraceLevel,
			logrus.DebugLevel,
			logrus.InfoLevel,
			logrus.WarnLevel,
			logrus.ErrorLevel,
			logrus.FatalLevel,
			logrus.PanicLevel,
		},
	})

	return func() {
		logrus.StandardLogger().ReplaceHooks(prev)
	}, nil
}
