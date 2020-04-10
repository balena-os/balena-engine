// +build linux

package chrootarchive // import "github.com/docker/docker/pkg/chrootarchive"

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"

	"github.com/pkg/errors"
	"golang.org/x/sys/unix"
)

func set_ionice(dst string) error {
	ioniceStr, ok := os.LookupEnv("MOBY_UNTAR_IONICE")
	if !ok {
		// noop
		return nil
	}

	// find block device for unpack destination
	//
	var stat unix.Stat_t
	if err := unix.Stat(dst, &stat); err != nil {
		// failed to statfs
		return errors.Wrapf(err, "stat %s", dst)
	}
	blkdevPath, err := getBlkdev(fmt.Sprintf("%d:%d",
		unix.Major(stat.Dev),
		unix.Minor(stat.Dev)))
	if err != nil {
		//
		return err
	}
	parentDev, err := ioutil.ReadFile(path.Join(string(blkdevPath), "../dev"))
	if !os.IsNotExist(err) {
		blkdevPath, err = getBlkdev(string(bytes.TrimSpace(parentDev)))
		if err != nil {
			//
			return err
		}
	}
	if err != nil {
		return err
	}
	blkdev := path.Base(blkdevPath)

	var bfqSched = []byte("bfq")
	schedulers, err := ioutil.ReadFile(path.Join("/sys/block/", blkdev, "/queue/scheduler"))
	if err != nil {
		return err
	}
	if !bytes.Contains(schedulers, bfqSched) {
		// noop
		return nil
	}
	var prevSched []byte
	for _, sched := range bytes.Split(schedulers, []byte(" ")) {
		if bytes.HasPrefix(sched, []byte("[")) && bytes.HasSuffix(sched, []byte("]")) {
			prevSched = sched
		}
	}
	if bytes.Compare(prevSched, bfqSched) == 0 {
		// no need to switch scheduler
		goto setioprio
	}
	if err := ioutil.WriteFile(path.Join("/sys/block/", blkdev, "/queue/iosched"), bfqSched, 0644); err != nil {
		// something went wrong?
		return err
	}
	defer func() {
		// switch back to previous scheduler
		ioutil.WriteFile(path.Join("/sys/block/", blkdev, "/queue/iosched"), prevSched, 0644)
	}()

setioprio:
	fmt.Printf("IONICE: destination: %s\n", dst)
	fmt.Printf("IONICE: MOBY_UNTAR_IONICE: %s\n", ioniceStr)

	if ioniceStr == "idle" {
		return unix.IoprioSet(unix.IOPRIO_WHO_PROCESS, 0, unix.IOPRIO_CLASS_IDLE, 0)
	}

	val, err := strconv.ParseInt(ioniceStr, 10, 8)
	if err != nil {
		return err
	}
	return unix.IoprioSet(unix.IOPRIO_WHO_PROCESS, 0, unix.IOPRIO_CLASS_BE, int(val))
}

func getBlkdev(devids string) (string, error) {
	sysPath := path.Join("/sys/dev/block/", devids)
	var blkdevPath = make([]byte, 1024)
	n, err := unix.Readlink(sysPath, blkdevPath)
	if err != nil || n <= 0 {
		return "", errors.Wrapf(err, "readlink %s", sysPath)
	}
	return path.Join(sysPath, "..", string(blkdevPath[:n])), nil
}
