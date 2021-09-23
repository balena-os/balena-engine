package main

import (
	"io/ioutil"
	"math"
	"strconv"
	"time"
	"unsafe"

	"github.com/sirupsen/logrus"
	"golang.org/x/sys/unix"
)

// Scheduling policies.
const (
	SCHED_NORMAL        = 0
	SCHED_FIFO          = 1
	SCHED_RR            = 2
	SCHED_BATCH         = 3
	SCHED_RESET_ON_FORK = 0x40000000 // Meant to be ORed with the others
)

// setScheduler sets the scheduling policy and priority of all threads in the
// current process.
func setScheduler(policy, prio int) {
	// The Go runtime can (and does) create new threads after the program
	// startup. We therefore keep looping for some minutes in a goroutine,
	// setting the priorities of any new threads. From experimentation, we see
	// no new threads created after a handful of minutes, so we don't need an
	// infinite loop.
	go func() {
		type sched_param struct {
			sched_priority int
		}
		s := &sched_param{int(prio)}
		p := unsafe.Pointer(s)
		doneTIDs := map[int]bool{}

		for i := 1; i <= 8; i++ {
			tids, err := getAllTasks()
			if err != nil {
				logrus.Errorf("Error getting task list: %v\n", err)
			}
			for _, tid := range tids {
				if _, done := doneTIDs[tid]; done {
					continue
				}
				_, _, errno := unix.Syscall(unix.SYS_SCHED_SETSCHEDULER, uintptr(tid), uintptr(policy), uintptr(p))
				if errno != 0 {
					logrus.Errorf("Syscall SYS_SCHED_SETSCHEDULER failed for tid=%v policy=%v prio=%v: %v\n",
						tid, policy, prio, errno)
				}
				doneTIDs[tid] = true
				logrus.Debugf("Priority of tid=%v set to policy=%v prio=%v", tid, policy, prio)
			}

			// From experimentation, we see that most threads are created in the
			// first seconds since program startup, so we increase the sleep
			// time exponentially to reduce the load we add to the system.
			sleepTime := time.Duration(math.Pow(2, float64(i))) * time.Second
			time.Sleep(sleepTime)
		}
	}()
}

// getAllTasks returns a list with the IDs of all tasks (threads) of the
// currently running process.
func getAllTasks() ([]int, error) {
	files, err := ioutil.ReadDir("/proc/self/task/")
	if err != nil {
		return nil, err
	}

	tids := []int{}
	for _, file := range files {
		tid, err := strconv.Atoi(file.Name())
		if err != nil {
			return tids, err
		}
		tids = append(tids, tid)
	}

	return tids, nil
}
