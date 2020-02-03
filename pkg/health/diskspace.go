package health

import (
	"context"
	"syscall"

	"github.com/pkg/errors"
)

// diskSpace checks for sufficient available disk space on the partition holding
// layer store etc.
type diskSpace struct {}

func (c diskSpace) Description() string {
	return "Available space on layer store partition"
}

func (c diskSpace) Check(ctx context.Context) error {
	var stat syscall.Statfs_t
	var statfsC = make(chan struct{})
	go func() {
		syscall.Statfs("/var/lib/balena-engine/", &stat)
		close(statfsC)
	}()
	// to enable cancelling of this check we run the blocking operation
	// in a coroutine and wait on either it's completion (via statfsC)
	// or the timeout from the passed context
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-statfsC:
	}

	total := (stat.Blocks * uint64(stat.Bsize))
	avail := (stat.Bavail * uint64(stat.Bsize))
	
	if ((100.0 / total)*avail) > 90.0 {
		return errors.New("layer store partition >90%% full")
	}
	return nil
}
