//go:build darwin

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
	// mac环境下代码如下
	stat, ok := finfo.Sys().(*syscall.Stat_t)
	if !ok || stat == nil {
		return nil, ErrInvalidFileStat
	}
	return &FileStat{
		CreationTime:   timespecToTime(stat.Ctimespec),
		LastAccessTime: timespecToTime(stat.Atimespec),
		LastWriteTime:  timespecToTime(stat.Mtimespec),
	}, nil
}
