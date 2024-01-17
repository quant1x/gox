package api

import "os"

// Touch 创建一个空文件
func Touch(filename string) error {
	_ = CheckFilepath(filename, true)
	return os.WriteFile(filename, nil, CACHE_FILE_MODE)
}
