//go:build !windows

package devices

import (
	"gitee.com/quant1x/gox/api"
	"os"
	"os/exec"
	"strings"
)

const (
	EnvPath = "PATH"
)

// CpuPhysicalID CPU序列号
func CpuPhysicalID() (PhysicalID string, err error) {
	PhysicalID = "Unknown"
	path := os.Getenv(EnvPath)
	path += ":" + "/usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin"
	err = os.Setenv(EnvPath, path)
	if err != nil {
		return
	}
	cmd := exec.Command("system_profiler", "SPHardwareDataType")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return PhysicalID, err
	}
	infos := api.Bytes2String(out)
	tmpArray := strings.Split(infos, "\n")
	for _, line := range tmpArray {
		line := strings.TrimSpace(line)
		key, value, found := strings.Cut(line, ":")
		if found {
			key = strings.TrimSpace(key)
			value = strings.TrimSpace(value)
			if len(key) == 0 || len(value) == 0 {
				continue
			}
			key = strings.ReplaceAll(key, " ", "")
			key = strings.ToLower(key)

			if strings.HasPrefix(key, "serialnumber") {
				PhysicalID = value
				break
			}
		}
	}
	return PhysicalID, nil
}
