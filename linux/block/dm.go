package block

import (
	"os"
	"path/filepath"
	"strings"
)

// IsDeviceDm returns true if the given device name starts with "dm-", indicating
// that it is a device mapper (dm) device.
func IsDeviceDm(name string) bool {
	return strings.HasPrefix(name, "dm-")
}

// CanonicalizeDmName reads the canonical name of a device mapper (dm) device
// from /sys/block/<name>/dm/name and returns it as a string.
// If the device is not a dm device, or if the name cannot be read, an error is returned.
func CanonicalizeDmName(name string) (string, error) {
	path := filepath.Join("/sys/block", name, "dm/name")
	bname, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(bname), nil
}
