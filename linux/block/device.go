package block

import "path/filepath"

// DevicePathFromName returns the path to a block device given its name.
// If the name starts with "dm-", it is assumed to be a device mapper (dm) device
// and the canonical name is used to construct the path.
// Otherwise, the name is used directly.
// If an error occurs during the construction of the path, it is returned.
func DevicePathFromName(name string) (string, error) {

	if IsDeviceDm(name) {
		cname, err := CanonicalizeDmName(name)
		if err != nil {
			return "", err
		}
		return filepath.Join("/dev/mapper", cname), nil
	}
	return filepath.Join("/dev", name), nil
}
