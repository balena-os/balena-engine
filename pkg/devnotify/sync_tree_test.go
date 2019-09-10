package devnotify

import (
	"context"
	"io/ioutil"
	"os"
	"path"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
)

func init()  {
	logrus.SetLevel(logrus.DebugLevel)
}

func TestSyncTree_Smoke(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	// create a directory tree to test on
	newdir, err := ioutil.TempDir("", "ctrdev_treewatch_test")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		os.RemoveAll(newdir)
	}()

	// worker thread that creates events
	go func() {
		newfile, err := os.Create(path.Join(newdir, "test"))
		if err != nil {
			t.Log(err)
			return
		}
		t.Logf("created new file: %v", newfile.Name())
		time.Sleep(1*time.Second)

		content := time.Now().String()
		_, err = newfile.WriteString(content)
		if err != nil {
			t.Log(err)
			return
		}
		t.Logf("wrote to new file: %v", content)
		time.Sleep(1*time.Second)

		err = os.Remove(newfile.Name())
		if err != nil {
			t.Log(err)
			return
		}
		t.Logf("deleted new file: %v", newfile.Name())
		time.Sleep(1*time.Second)

		cancel()
	}()

	if err := SyncTree(ctx, newdir);  err != nil {
		t.Fatalf("Error watching file tree: %v", err)
	}
}
