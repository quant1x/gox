package logger

import (
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var (
	onceApp         sync.Once
	applicationName = ""
)

func lazyLoadApplication() {
	path, _ := os.Executable()
	_, exec := filepath.Split(path)
	arr := strings.Split(exec, ".")
	applicationName = arr[0]
}

// ApplicationName 获取执行文件名
func ApplicationName() string {
	onceApp.Do(lazyLoadApplication)
	return applicationName
}
