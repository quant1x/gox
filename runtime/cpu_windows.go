//go:build windows

package licenses

import (
	"os/exec"
	"regexp"
)

func CpuPhysicalID() (PhysicalID string, err error) {
	PhysicalID = "Unknown"
	cmd := exec.Command("wmic", "cpu", "get", "ProcessorID")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return
	}
	str := string(out)
	// 匹配一个或多个空白符的正则表达式
	reg := regexp.MustCompile("\\s+")
	str = reg.ReplaceAllString(str, "")
	return str[11:], nil
}
