//go:build windows

package api

import (
	"os"
	"syscall"
)

// GetFileStat 获取文件状态(创建,修改和访问时间)
func GetFileStat(name string) (*FileStat, error) {
	finfo, err := os.Lstat(name)
	if err != nil {
		return nil, err
	}
	// windows下代码如下
	fileAttr, ok := finfo.Sys().(*syscall.Win32FileAttributeData)
	if !ok || stat == nil {
		return nil, ErrInvaildFileStat
	}
	return &FileStat{
		CreationTime:   NanosecondToTime(fileAttr.CreationTime.Nanoseconds()),
		LastAccessTime: NanosecondToTime(fileAttr.LastAccessTime.Nanoseconds()),
		LastWriteTime:  NanosecondToTime(fileAttr.LastWriteTime.Nanoseconds()),
	}, nil
}
