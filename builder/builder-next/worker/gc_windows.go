//go:build windows && !no_buildkit

package worker

func detectDefaultGCCap(root string) int64 {
	return defaultCap
}
