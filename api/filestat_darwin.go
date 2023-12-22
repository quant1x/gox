//go:build darwin

package api

import (
	"os"
	"syscall"
	"time"
)

// GetFileStat 获取文件状态(创建,修改和访问时间)
func GetFileStat(name string) (*FileStat, error) {
	finfo, err := os.Lstat(name)
	if err != nil {
		return nil, err
	}
	// mac环境下代码如下
	stat, _ := finfo.Sys().(*syscall.Stat_t)
	return &FileStat{
		CreationTime:   time.Unix(stat.Ctimespec.Sec, stat.Ctimespec.Nsec),
		LastAccessTime: time.Unix(stat.Atimespec.Sec, stat.Atimespec.Nsec),
		LastWriteTime:  time.Unix(stat.Mtimespec.Sec, stat.Mtimespec.Nsec),
	}, nil
}
