//go:build linux

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
	// linux环境下代码如下
	stat, ok := finfo.Sys().(*syscall.Stat_t)
	if !ok || stat == nil {
		return nil, ErrInvalidFileStat
	}
	return &FileStat{
		CreationTime:   timespecToTime(stat.Ctim),
		LastAccessTime: timespecToTime(stat.Atim),
		LastWriteTime:  timespecToTime(stat.Mtim),
	}, nil
}
