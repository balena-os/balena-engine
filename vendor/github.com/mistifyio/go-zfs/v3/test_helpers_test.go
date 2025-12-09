package zfs_test

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
	"time"

	"github.com/mistifyio/go-zfs/v3"
)

func sleep(delay int) {
	time.Sleep(time.Duration(delay) * time.Second)
}

func pow2(x int) int64 {
	return int64(math.Pow(2, float64(x)))
}

// https://github.com/benbjohnson/testing
// assert fails the test if the condition is false.
func _assert(t *testing.T, condition bool, msg string, v ...interface{}) {
	t.Helper()

	if !condition {
		_, file, line, _ := runtime.Caller(2)
		t.Logf("\033[31m%s:%d: "+msg+"\033[39m\n\n", append([]interface{}{filepath.Base(file), line}, v...)...)
		t.FailNow()
	}
}

func assert(t *testing.T, condition bool, msg string, v ...interface{}) {
	t.Helper()
	_assert(t, condition, msg, v...)
}

// ok fails the test if an err is not nil.
func ok(t *testing.T, err error) {
	t.Helper()
	_assert(t, err == nil, "unexpected error: %v", err)
}

// nok fails the test if an err is nil.
func nok(t *testing.T, err error) {
	t.Helper()
	_assert(t, err != nil, "expected error, got nil")
}

// equals fails the test if exp is not equal to act.
func equals(t *testing.T, exp, act interface{}) {
	t.Helper()
	_assert(t, reflect.DeepEqual(exp, act), "exp: %#v\n\ngot: %#v", exp, act)
}

type cleanUpFunc func()

func (f cleanUpFunc) cleanUp() {
	f()
}

// do something like Restorer in github.com/packethost/pkg/internal/testenv/clearer.go
func setupZPool(t *testing.T) cleanUpFunc {
	t.Helper()

	d, err := ioutil.TempDir("/tmp/", "zfs-test-*")
	ok(t, err)

	var skipRemoveAll bool
	defer func() {
		if !skipRemoveAll {
			t.Logf("cleaning up")
			os.RemoveAll(d)
		}
	}()

	tempfiles := make([]string, 3)
	for i := range tempfiles {
		f, err := ioutil.TempFile(d, fmt.Sprintf("loop%d", i))
		ok(t, err)

		ok(t, f.Truncate(pow2(30)))

		f.Close()
		tempfiles[i] = f.Name()
	}

	pool, err := zfs.CreateZpool("test", nil, tempfiles...)
	ok(t, err)

	skipRemoveAll = true
	return func() {
		ok(t, pool.Destroy())
		os.RemoveAll(d)
	}
}
