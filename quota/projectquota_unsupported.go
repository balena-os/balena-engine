//go:build (linux && exclude_disk_quota) || (linux && !cgo) || !linux
// +build linux,exclude_disk_quota linux,!cgo !linux

package quota // import "github.com/docker/docker/quota"

import "testing"

func NewControl(basePath string) (*Control, error) {
	return nil, ErrQuotaNotSupported
}

// SetQuota - assign a unique project id to directory and set the quota limits
// for that project id
func (q *Control) SetQuota(targetPath string, quota Quota) error {
	return ErrQuotaNotSupported
}

// GetQuota - get the quota limits of a directory that was configured with SetQuota
func (q *Control) GetQuota(targetPath string, quota *Quota) error {
	return ErrQuotaNotSupported
}

// testhelpers

func CanTestQuota() (string, bool) {
	return ErrQuotaNotSupported.Error(), false
}

func PrepareQuotaTestImage(t *testing.T) (string, error) {
	return "", ErrQuotaNotSupported
}

func WrapMountTest(imageFileName string, enableQuota bool, testFunc func(t *testing.T, mountPoint, backingFsDev, testDir string)) func(*testing.T) {
	return func(t *testing.T) {
		t.Skip(ErrQuotaNotSupported.Error())
	}
}
