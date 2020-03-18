package quota // import "github.com/docker/docker/quota"

import "sync"

// Quota limit params - currently we only control blocks hard limit
type Quota struct {
	Size uint64
}

// Control - Context to be used by storage driver (e.g. overlay)
// who wants to apply project quotas to container dirs
type Control struct {
	//nolint:unused // only used when !exclude_disk_quota; see projectquota.go
	backingFsBlockDev string
	sync.RWMutex      // protect nextProjectID and quotas map
	//nolint:unused // only used when !exclude_disk_quota; see projectquota.go
	quotas map[string]uint32
}
