package runtime

import (
	"os"
	"path/filepath"
)

// ApplicationName 获取执行文件名
func ApplicationName() string {
	path, _ := os.Executable()
	_, exec := filepath.Split(path)
	return exec
}
