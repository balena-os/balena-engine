package storagemigration

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"golang.org/x/sys/unix"

	"github.com/docker/docker/daemon/graphdriver/overlay2"
)

// CreateLayerLink creates a link file in the layer root dir and a corresponding file in the l directory
// The returned layerRef is the content of the created link file
func CreateLayerLink(root, layerID string) (layerRef string, err error) {
	layerLinkFile := filepath.Join(root, layerID, "link")
	ok, err := exists(layerLinkFile, false)
	if err != nil {
		return "", fmt.Errorf("Error checking for %s: %v", layerLinkFile, err)
	}
	if ok {
		// Return early if it already exists.
		// Happens when we process layer that
		// previously appeared as a parent layer.
		ref, err := ioutil.ReadFile(layerLinkFile)
		if err != nil {
			return "", fmt.Errorf("Error reading %s: %v", layerLinkFile, err)
		}
		return string(ref), nil
	}
	layerRefDir := filepath.Join(root, "l")
	ok, err = exists(layerRefDir, true)
	if err != nil {
		return "", fmt.Errorf("Error checking for %s: %v", layerRefDir, err)
	}
	if !ok {
		// create layer ref dir
		// to avoid having to do this outside of this function
		err := os.MkdirAll(layerRefDir, 0700)
		if err != nil {
			return "", fmt.Errorf("Error creating directory %s: %v", layerRefDir, err)
		}
	}
	// idLength
	// daemon/graphdriver/overlay2/overlay#L87
	layerRef = overlay2.GenerateID(overlay2.IDLength)
	err = ioutil.WriteFile(layerLinkFile, []byte(layerRef), 0644)
	if err != nil {
		return "", fmt.Errorf("Error writing to %s: %v", layerLinkFile, err)
	}
	layerDiffDir := filepath.Join("..", layerID, "diff")
	layerLinkRef := filepath.Join(layerRefDir, layerRef)
	err = os.Symlink(layerDiffDir, layerLinkRef)
	if err != nil {
		return "", fmt.Errorf("Error creating symlink %s -> %s: %v", layerDiffDir, layerLinkRef, err)
	}
	return layerRef, nil
}

// AppendLower adds parentID to the list of lower directories written to /:layer_id/lower
func AppendLower(lower, parentID string) string {
	if lower == "" {
		return "l/" + parentID
	}
	return lower + ":l/" + parentID
}

// SetOpaque marks the directory to appera empty
// by setting the xattr "trusted.overlay.opaque" to "y"
func SetOpaque(path string) error {
	return unix.Setxattr(path, "trusted.overlay.opaque", []byte("y"), 0)
}

// SetWhiteout marks the file as deleted
// by creating a character device with 0/0 device number
func SetWhiteout(path string) error {
	return unix.Mknod(path, unix.S_IFCHR, int(unix.Mkdev(0, 0)))
}
