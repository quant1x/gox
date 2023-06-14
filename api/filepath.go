package api

import (
	"gitee.com/quant1x/gox/util/homedir"
	"os"
	"path/filepath"
)

const (
	// CACHE_DIR_MODE 目录权限
	CACHE_DIR_MODE os.FileMode = 0755
	// CACHE_FILE_MODE 文件权限
	CACHE_FILE_MODE os.FileMode = 0644

	// DEBUG 调试开关
	DEBUG = false
	// CACHE_REPLACE 文件替换模式, 会用到os.TRUNC
	CACHE_REPLACE = os.O_CREATE | os.O_RDWR | os.O_TRUNC
	// CACHE_UPDATE 更新
	CACHE_UPDATE = os.O_CREATE | os.O_WRONLY
)

// CheckFilepath
//
//	检查filename 文件路径, 如果不存在就创建
func CheckFilepath(filename string, notExistToCreate ...bool) error {
	filename, _ = homedir.Expand(filename)
	path := filepath.Dir(filename)
	dir, err := os.Lstat(path)
	if err == nil {
		// 已存在
		return nil
	}
	if os.IsExist(err) {
		// 已存在
		return nil
	}
	__create := false
	if len(notExistToCreate) > 0 {
		__create = notExistToCreate[0]
	}
	if !__create {
		return os.ErrNotExist
	}
	// 不存在, 创建
	err = os.MkdirAll(path, CACHE_DIR_MODE)
	if err != nil {
		return err
	}
	dir, err = os.Stat(path)
	if err != nil {
		return err
	}
	if dir.IsDir() {
		return nil
	}
	return os.ErrNotExist
}

// FileExist 路径是否存在
func FileExist(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}

// FileIsValid 检查文件是否有效
func FileIsValid(path string) bool {
	finfo, err := os.Lstat(path)
	found := !os.IsNotExist(err)
	if !found {
		return false
	}
	if finfo.Size() > 2 {
		return true
	}
	return false
}
