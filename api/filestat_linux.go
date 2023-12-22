//go:build linux

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
	// linux环境下代码如下
	stat, _ := finfo.Sys().(*syscall.Stat_t)
	return &FileStat{
		CreationTime:   time.Unix(stat.Ctim.Sec, stat.Ctim.Nsec),
		LastAccessTime: time.Unix(stat.Atim.Sec, stat.Atim.Nsec),
		LastWriteTime:  time.Unix(stat.Mtim.Sec, stat.Mtim.Nsec),
	}, nil
}
