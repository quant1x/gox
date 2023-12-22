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
	winFileAttr, _ := finfo.Sys().(*syscall.Win32FileAttributeData)
	return &FileStat{
		CreationTime:   NanosecondToTime(winFileAttr.CreationTime.Nanoseconds()),
		LastAccessTime: NanosecondToTime(winFileAttr.LastAccessTime.Nanoseconds()),
		LastWriteTime:  NanosecondToTime(winFileAttr.LastWriteTime.Nanoseconds()),
	}, nil
}
