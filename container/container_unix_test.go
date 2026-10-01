//go:build !windows

package container

import (
	"os"
	"path/filepath"
	"syscall"
	"testing"

	"gotest.tools/v3/assert"
)

func statInode(t *testing.T, path string) uint64 {
	t.Helper()
	info, err := os.Stat(path)
	assert.NilError(t, err) // actually means that error == nil
	ino := info.Sys().(*syscall.Stat_t).Ino
	t.Logf("inode %s = %d", path, ino)
	return ino
}

func assertHardlinked(t *testing.T, srcPath, dstPath string) {
	t.Helper()
	srcIno := statInode(t, srcPath)
	dstIno := statInode(t, dstPath)
	assert.Equal(t, srcIno, dstIno, "expected hardlink: same inode for %s and %s", srcPath, dstPath)
	t.Logf("hardlink verified: %s and %s share inode %d", srcPath, dstPath, srcIno)
}

func TestLinkExistingContents(t *testing.T) {
	t.Parallel()

	tmp := t.TempDir()
	source := filepath.Join(tmp, "src")
	dest := filepath.Join(tmp, "dest")

	nestedInSource := filepath.Join(source, "nested")
	nestedFile := filepath.Join(nestedInSource, "file")
	topFile := filepath.Join(source, "top")

	t.Logf("source=%s dest=%s", source, dest)

	assert.NilError(t, os.MkdirAll(nestedInSource, 0o755))
	assert.NilError(t, os.WriteFile(nestedFile, []byte("nested"), 0o644))
	assert.NilError(t, os.WriteFile(topFile, []byte("top"), 0o644))
	assert.NilError(t, os.Mkdir(dest, 0o755))

	assert.NilError(t, linkExistingContents(source, dest))
	t.Log("linkExistingContents completed for directory tree")

	assertHardlinked(t, nestedFile, filepath.Join(dest, "nested", "file"))
	assertHardlinked(t, topFile, filepath.Join(dest, "top"))
	t.Log("passed: directory tree hardlinked without copying")
}

// this test verifies that linkExistingContents does not hardlink
// if the destination is not empty
func TestLinkExistingContentsSkipsNonEmptyDestination(t *testing.T) {
	t.Parallel()

	tmp := t.TempDir()
	source := filepath.Join(tmp, "src")
	dest := filepath.Join(tmp, "dest")
	destExistingFile := filepath.Join(dest, "existing")

	t.Logf("source=%s dest=%s (non-empty)", source, dest)

	assert.NilError(t, os.Mkdir(source, 0o755))
	assert.NilError(t, os.WriteFile(filepath.Join(source, "file"), []byte("src"), 0o644))

	assert.NilError(t, os.Mkdir(dest, 0o755))
	assert.NilError(t, os.WriteFile(destExistingFile, []byte("existing"), 0o644))

	beforeIno := statInode(t, destExistingFile)
	assert.NilError(t, linkExistingContents(source, dest))
	t.Log("linkExistingContents returned nil")

	_, err := os.Stat(filepath.Join(dest, "file"))
	assert.Assert(t, os.IsNotExist(err))
	t.Log("no file linked from source into non-empty destination")

	assert.Equal(t, beforeIno, statInode(t, destExistingFile))
	t.Log("passed: non-empty destination left unchanged")
}

func TestLinkExistingContentsSingleFile(t *testing.T) {
	t.Parallel()

	tmp := t.TempDir()
	source := filepath.Join(tmp, "src", "file")
	dest := filepath.Join(tmp, "dest")

	assert.NilError(t, os.MkdirAll(filepath.Dir(source), 0o755))
	assert.NilError(t, os.WriteFile(source, []byte("hello"), 0o644))
	assert.NilError(t, os.Mkdir(dest, 0o755))

	assert.NilError(t, linkExistingContents(source, dest))
	assertHardlinked(t, source, filepath.Join(dest, "file"))
}

func TestFindInLowerDir(t *testing.T) {
	t.Parallel()

	tmp := t.TempDir()
	lower1 := filepath.Join(tmp, "lower1", "diff")
	lower2 := filepath.Join(tmp, "lower2", "diff")

	assert.NilError(t, os.MkdirAll(filepath.Join(lower2, "boot"), 0o755))
	assert.NilError(t, os.WriteFile(filepath.Join(lower2, "boot", "vmlinuz"), []byte("kernel"), 0o644))

	metadata := map[string]string{"LowerDir": lower1 + ":" + lower2}

	path, ok := findInLowerDir(metadata, "boot")
	assert.Assert(t, ok)
	assert.Equal(t, filepath.Join(lower2, "boot"), path)

	_, ok = findInLowerDir(metadata, "missing")
	assert.Assert(t, !ok)

	// first matching layer wins (LowerDir order is uppermost first)
	assert.NilError(t, os.MkdirAll(filepath.Join(lower1, "boot"), 0o755))
	path, ok = findInLowerDir(metadata, "boot")
	assert.Assert(t, ok)
	assert.Equal(t, filepath.Join(lower1, "boot"), path)
}
