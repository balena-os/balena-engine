package storagemigration

import (
	"os"
	"path/filepath"
	"testing"

	"gotest.tools/assert"
	"gotest.tools/fs"
)

func TestCreateLayerLink(t *testing.T) {
	testCases := []string{
		"9d5d02966ef6413f9b185cb017d183bf",
		"51d2de199c894a038aa795e6a0614bb1",
		"2db82ff9e33b4ce4b23c05c5086d949b",
	}

	root := fs.NewDir(t, t.Name())
	defer root.Remove()

	for _, layerID := range testCases {
		// prepare
		os.Mkdir(root.Join(layerID), os.ModeDir|0700)

		ref, err := CreateLayerLink(root.Path(), layerID)
		assert.NilError(t, err)
		assert.Assert(t, fs.Equal(root.Join("l"), fs.Expected(t,
			fs.WithSymlink(ref, filepath.Join("..", layerID, "diff")))))

		// cleanup
		os.RemoveAll(root.Join("l"))
	}
}

func TestAppendLower(t *testing.T) {
	testCases := []struct {
		layerID string
		parents []string
		lower   string
	}{
		{
			layerID: "9d5d02966ef6413f9b185cb017d183bf",
			parents: []string{},
			lower:   "",
		},
		{
			layerID: "3e0790f98ab04048b34f0d88a5766230",
			parents: []string{
				"9d5d02966ef6413f9b185cb017d183bf", // [0]
				"8981b529cffa4c17a79de62be23ebd0a",
			},
			lower: "l/9d5d029:l/8981b52",
		},
		{
			layerID: "05d4fd0c51fb4c93b0ad06af567ea91d",
			parents: []string{
				"3e0790f98ab04048b34f0d88a5766230", // [1]
				"a608235d89e94708a60abc497e0aa0f8",
				"8981b529cffa4c17a79de62be23ebd0a",
				"9d5d02966ef6413f9b185cb017d183bf", // [0]
			},
			lower: "l/3e0790f:l/a608235:l/8981b52:l/9d5d029",
		},
	}

	for _, tc := range testCases {
		var lowerString string
		for _, parentID := range tc.parents {
			parentRef := parentID[:7] // spoof here
			lowerString = AppendLower(lowerString, parentRef)
		}
		assert.Equal(t, lowerString, tc.lower)
	}
}
