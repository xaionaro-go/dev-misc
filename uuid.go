// +build linux

package misc

import (
	"os/exec"

	"github.com/satori/go.uuid"
)

func GetUUIDOfDevice(devicePath string) (uuid.UUID, error) {
	result, err := exec.Command("blkid", "-s", "UUID", "-o", "value", devicePath).Output()
	if err != nil {
		return uuid.UUID{}, err
	}
	return uuid.FromString(string(result))
}

