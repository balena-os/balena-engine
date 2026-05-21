package zfs_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	zfs "github.com/mistifyio/go-zfs/v3"
)

func TestDatasets(t *testing.T) {
	defer setupZPool(t).cleanUp()

	_, err := zfs.Datasets("")
	ok(t, err)

	ds, err := zfs.GetDataset("test")
	ok(t, err)
	equals(t, zfs.DatasetFilesystem, ds.Type)
	equals(t, "", ds.Origin)
	if runtime.GOOS != "solaris" {
		assert(t, ds.Logicalused != 0, "Logicalused is not greater than 0")
	}
}

func TestDatasetGetProperty(t *testing.T) {
	defer setupZPool(t).cleanUp()

	ds, err := zfs.GetDataset("test")
	ok(t, err)

	prop, err := ds.GetProperty("foobarbaz")
	nok(t, err)
	equals(t, "", prop)

	prop, err = ds.GetProperty("compression")
	ok(t, err)
	equals(t, "off", prop)

	// creation should be a time stamp with spaces in it
	prop, err = ds.GetProperty("creation")
	ok(t, err)
	if len(strings.Fields(prop)) != 5 {
		t.Errorf("expected a string with spaces in it, got: %v", prop)
	}
}

func TestSnapshots(t *testing.T) {
	defer setupZPool(t).cleanUp()

	snapshots, err := zfs.Snapshots("")
	ok(t, err)

	for _, snapshot := range snapshots {
		equals(t, zfs.DatasetSnapshot, snapshot.Type)
	}
}

func TestFilesystems(t *testing.T) {
	defer setupZPool(t).cleanUp()

	f, err := zfs.CreateFilesystem("test/filesystem-test", nil)
	ok(t, err)

	filesystems, err := zfs.Filesystems("")
	ok(t, err)

	for _, filesystem := range filesystems {
		equals(t, zfs.DatasetFilesystem, filesystem.Type)
	}

	ok(t, f.Destroy(zfs.DestroyDefault))
}

func TestCreateFilesystemWithProperties(t *testing.T) {
	defer setupZPool(t).cleanUp()

	props := map[string]string{
		"compression": "lz4",
	}

	f, err := zfs.CreateFilesystem("test/filesystem-test", props)
	ok(t, err)

	equals(t, "lz4", f.Compression)

	filesystems, err := zfs.Filesystems("")
	ok(t, err)

	for _, filesystem := range filesystems {
		equals(t, zfs.DatasetFilesystem, filesystem.Type)
	}

	ok(t, f.Destroy(zfs.DestroyDefault))
}

func TestVolumes(t *testing.T) {
	defer setupZPool(t).cleanUp()

	v, err := zfs.CreateVolume("test/volume-test", uint64(pow2(23)), nil)
	ok(t, err)

	// volumes are sometimes "busy" if you try to manipulate them right away
	sleep(1)

	equals(t, zfs.DatasetVolume, v.Type)
	volumes, err := zfs.Volumes("")
	ok(t, err)

	for _, volume := range volumes {
		equals(t, zfs.DatasetVolume, volume.Type)
	}

	ok(t, v.Destroy(zfs.DestroyDefault))
}

func TestSnapshot(t *testing.T) {
	defer setupZPool(t).cleanUp()

	f, err := zfs.CreateFilesystem("test/snapshot-test", nil)
	ok(t, err)

	filesystems, err := zfs.Filesystems("")
	ok(t, err)

	for _, filesystem := range filesystems {
		equals(t, zfs.DatasetFilesystem, filesystem.Type)
	}

	s, err := f.Snapshot("test", false)
	ok(t, err)

	equals(t, zfs.DatasetSnapshot, s.Type)

	equals(t, "test/snapshot-test@test", s.Name)

	ok(t, s.Destroy(zfs.DestroyDefault))

	ok(t, f.Destroy(zfs.DestroyDefault))
}

func TestClone(t *testing.T) {
	defer setupZPool(t).cleanUp()

	f, err := zfs.CreateFilesystem("test/snapshot-test", nil)
	ok(t, err)

	filesystems, err := zfs.Filesystems("")
	ok(t, err)

	for _, filesystem := range filesystems {
		equals(t, zfs.DatasetFilesystem, filesystem.Type)
	}

	s, err := f.Snapshot("test", false)
	ok(t, err)

	equals(t, zfs.DatasetSnapshot, s.Type)
	equals(t, "test/snapshot-test@test", s.Name)

	c, err := s.Clone("test/clone-test", nil)
	ok(t, err)

	equals(t, zfs.DatasetFilesystem, c.Type)

	ok(t, c.Destroy(zfs.DestroyDefault))

	ok(t, s.Destroy(zfs.DestroyDefault))

	ok(t, f.Destroy(zfs.DestroyDefault))
}

func TestSendSnapshot(t *testing.T) {
	defer setupZPool(t).cleanUp()

	f, err := zfs.CreateFilesystem("test/snapshot-test", nil)
	ok(t, err)

	filesystems, err := zfs.Filesystems("")
	ok(t, err)

	for _, filesystem := range filesystems {
		equals(t, zfs.DatasetFilesystem, filesystem.Type)
	}

	s, err := f.Snapshot("test", false)
	ok(t, err)

	file, _ := ioutil.TempFile("/tmp/", "zfs-")
	defer file.Close()
	err = file.Truncate(pow2(30))
	ok(t, err)
	defer os.Remove(file.Name())

	err = s.SendSnapshot(file)
	ok(t, err)

	ok(t, s.Destroy(zfs.DestroyDefault))

	ok(t, f.Destroy(zfs.DestroyDefault))
}

func TestChildren(t *testing.T) {
	defer setupZPool(t).cleanUp()

	f, err := zfs.CreateFilesystem("test/snapshot-test", nil)
	ok(t, err)

	s, err := f.Snapshot("test", false)
	ok(t, err)

	equals(t, zfs.DatasetSnapshot, s.Type)
	equals(t, "test/snapshot-test@test", s.Name)

	children, err := f.Children(0)
	ok(t, err)

	equals(t, 1, len(children))
	equals(t, "test/snapshot-test@test", children[0].Name)

	ok(t, s.Destroy(zfs.DestroyDefault))
	ok(t, f.Destroy(zfs.DestroyDefault))
}

func TestListZpool(t *testing.T) {
	defer setupZPool(t).cleanUp()

	pools, err := zfs.ListZpools()
	ok(t, err)
	for _, pool := range pools {
		if pool.Name == "test" {
			equals(t, "test", pool.Name)
			return
		}
	}
	t.Fatal("Failed to find test pool")
}

func TestRollback(t *testing.T) {
	defer setupZPool(t).cleanUp()

	f, err := zfs.CreateFilesystem("test/snapshot-test", nil)
	ok(t, err)

	filesystems, err := zfs.Filesystems("")
	ok(t, err)

	for _, filesystem := range filesystems {
		equals(t, zfs.DatasetFilesystem, filesystem.Type)
	}

	s1, err := f.Snapshot("test", false)
	ok(t, err)

	_, err = f.Snapshot("test2", false)
	ok(t, err)

	s3, err := f.Snapshot("test3", false)
	ok(t, err)

	err = s3.Rollback(false)
	ok(t, err)

	err = s1.Rollback(false)
	assert(t, err != nil, "should error when rolling back beyond most recent snapshot without destroyMoreRecent = true")

	err = s1.Rollback(true)
	ok(t, err)

	ok(t, s1.Destroy(zfs.DestroyDefault))

	ok(t, f.Destroy(zfs.DestroyDefault))
}

func TestDiff(t *testing.T) {
	defer setupZPool(t).cleanUp()

	fs, err := zfs.CreateFilesystem("test/origin", nil)
	ok(t, err)

	linkedFile, err := os.Create(filepath.Join(fs.Mountpoint, "linked"))
	ok(t, err)

	movedFile, err := os.Create(filepath.Join(fs.Mountpoint, "file"))
	ok(t, err)

	snapshot, err := fs.Snapshot("snapshot", false)
	ok(t, err)

	unicodeFile, err := os.Create(filepath.Join(fs.Mountpoint, "i ❤ unicode"))
	ok(t, err)

	err = os.Rename(movedFile.Name(), movedFile.Name()+"-new")
	ok(t, err)

	err = os.Link(linkedFile.Name(), linkedFile.Name()+"_hard")
	ok(t, err)

	inodeChanges, err := fs.Diff(snapshot.Name)
	ok(t, err)
	equals(t, 4, len(inodeChanges))

	unicodePath := "/test/origin/i\x040\x1c2\x135\x144\x040unicode"
	wants := map[string]*zfs.InodeChange{
		"/test/origin/linked": {
			Type:                 zfs.File,
			Change:               zfs.Modified,
			ReferenceCountChange: 1,
		},
		"/test/origin/file": {
			Type:    zfs.File,
			Change:  zfs.Renamed,
			NewPath: "/test/origin/file-new",
		},
		"/test/origin/i ❤ unicode": {
			Path:   "❤❤ unicode ❤❤",
			Type:   zfs.File,
			Change: zfs.Created,
		},
		unicodePath: {
			Path:   "❤❤ unicode ❤❤",
			Type:   zfs.File,
			Change: zfs.Created,
		},
		"/test/origin/": {
			Type:   zfs.Directory,
			Change: zfs.Modified,
		},
	}
	for _, change := range inodeChanges {
		want := wants[change.Path]
		want.Path = change.Path
		delete(wants, change.Path)

		equals(t, want, change)
	}

	equals(t, 1, len(wants))
	for _, want := range wants {
		equals(t, "❤❤ unicode ❤❤", want.Path)
	}

	ok(t, movedFile.Close())
	ok(t, unicodeFile.Close())
	ok(t, linkedFile.Close())
	ok(t, snapshot.Destroy(zfs.DestroyForceUmount))
	ok(t, fs.Destroy(zfs.DestroyForceUmount))
}
