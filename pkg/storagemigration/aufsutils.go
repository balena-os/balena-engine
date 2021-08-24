// TODO: do we need to handle .wh..wh.plnk layer hardlinks?
package storagemigration

import (
	"bufio"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/docker/docker/pkg/archive"
)

func IsWhiteout(filename string) bool {
	return strings.HasPrefix(filename, archive.WhiteoutPrefix)
}

func IsWhiteoutMeta(filename string) bool {
	return strings.HasPrefix(filename, archive.WhiteoutMetaPrefix)
}

func IsOpaqueParentDir(filename string) bool {
	return filename == archive.WhiteoutOpaqueDir
}

func StripWhiteoutPrefix(filename string) string {
	out := filename
	for IsWhiteout(out) && !IsWhiteoutMeta(out) {
		out = strings.TrimPrefix(out, archive.WhiteoutPrefix)
	}
	return out
}

// Return all the directories
//
// from daemon/graphdriver/aufs/dirs.go
func LoadIDs(root string) ([]string, error) {
	dirs, err := ioutil.ReadDir(root)
	if err != nil {
		return nil, err
	}
	var out []string
	for _, d := range dirs {
		if d.IsDir() {
			out = append(out, d.Name())
		}
	}
	return out, nil
}

// Read the layers file for the current id and return all the
// layers represented by new lines in the file
//
// If there are no lines in the file then the id has no parent
// and an empty slice is returned.
//
// from daemon/graphdriver/aufs/dirs.go
func GetParentIDs(root, id string) ([]string, error) {
	f, err := os.Open(path.Join(root, "layers", id))
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var out []string
	s := bufio.NewScanner(f)

	for s.Scan() {
		if t := s.Text(); t != "" {
			out = append(out, s.Text())
		}
	}
	return out, s.Err()
}
