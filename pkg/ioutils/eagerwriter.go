package ioutils

import (
	"io"
	"os"

	"golang.org/x/sys/unix"
)

const STEP = 2 * 1024 * 1024 // 2MB

const (
	SYNC_FILE_RANGE_WAIT_BEFORE = 1
	SYNC_FILE_RANGE_WRITE       = 2
	SYNC_FILE_RANGE_WAIT_AFTER  = 4
)

// EagerFileWriter will schedule immediate writeback of dirty pages. The writes
// will only block if the device's write queue is full, which provides
// throttling without incuring the latency cost of fsync.
//
// The purpose of this is twofold. Firstly, it instructs the kernel to
// immediatelly start flushing data to disk which will reduce the delay of a
// subsequent sync/fsync call since there will be less data to sync.
//
// Secondly and more importantly, it makes sure that there are always dirty
// pages under writeback. This is important because of what seems to be a bug
// in linux[1] where the cgroup v1 implementation only throttles dirty page
// creation if there are pages under writeback.
//
// [1] https://marc.info/?l=linux-mm&m=153069062419988&w=2
type eagerFileWriter struct {
	f       *os.File
	written int64
	synced  int64
}

func (e *eagerFileWriter) Write(b []byte) (int, error) {
	n, err := e.f.Write(b)
	e.written += int64(n)
	if e.written-e.synced > STEP {
		unix.SyncFileRange(int(e.f.Fd()), e.synced, STEP, SYNC_FILE_RANGE_WRITE)
		e.synced += STEP
	}
	return n, err
}

func (e *eagerFileWriter) Close() error {
	unix.SyncFileRange(int(e.f.Fd()), 0, 0, SYNC_FILE_RANGE_WRITE)
	return e.f.Close()
}

func NewEagerFileWriter(f *os.File) io.WriteCloser {
	return &eagerFileWriter{f: f}
}
