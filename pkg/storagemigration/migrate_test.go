package storagemigration

import (
	"os"
	"testing"

	"github.com/docker/docker/pkg/archive"

	"gotest.tools/assert"
	"gotest.tools/fs"
)

func createStorageDir(t *testing.T) (*fs.Dir, *State) {
	root := fs.NewDir(t, t.Name(),
		fs.WithDir("aufs",
			fs.WithDir("layers",
				fs.WithFile("b38c03118c1e41289cf0972f11453c9b", "",
					fs.WithMode(0755)),
				fs.WithFile("b8936bbae21948ed826207ced6fa19c5", "b38c03118c1e41289cf0972f11453c9b",
					fs.WithMode(0666)),
			),
			fs.WithDir("diff",
				fs.WithDir("b38c03118c1e41289cf0972f11453c9b", fs.WithMode(0755),
					fs.WithFile("test", "")),
				fs.WithDir("b8936bbae21948ed826207ced6fa19c5", fs.WithMode(0755),
					fs.WithFile(archive.WhiteoutPrefix+"test", "")),
			),
		),
	)
	state := &State{
		Layers: []Layer{
			{
				ID:        "b38c03118c1e41289cf0972f11453c9b",
				ParentIDs: nil,
				Meta:      nil,
			},
			{
				ID:        "b8936bbae21948ed826207ced6fa19c5",
				ParentIDs: []string{"b38c03118c1e41289cf0972f11453c9b"},
				Meta: []Meta{
					{Type: MetaWhiteout, Path: "test"},
				},
			},
		},
	}
	return root, state
}

func TestCreateState(t *testing.T) {
	root, expect := createStorageDir(t)
	defer root.Remove()

	state, err := createState(root.Join("aufs"))
	assert.NilError(t, err)
	assert.DeepEqual(t, state, expect)
}

func TestMigrate(t *testing.T) {
	root, _ := createStorageDir(t)
	defer root.Remove()

	err := Migrate(root.Path())
	assert.NilError(t, err)
}

func TestMigrateFailure(t *testing.T) {
	root, _ := createStorageDir(t)
	defer root.Remove()

	// delete diff directory to force createState to fail
	os.RemoveAll(root.Join("aufs", "diff"))

	logPath := root.Join("migrate.log")
	os.Setenv("BALENA_MIGRATE_OVERLAY_LOGFILE", logPath)

	err := Migrate(root.Path())
	assert.Assert(t, err != nil)

	_, err = os.Stat(logPath)
	assert.NilError(t, err)
}
