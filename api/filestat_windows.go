//go:build windows

package api

import (
	"os"
	"syscall"
)

// GetFileStat 获取文件状态(创建,修改和访问时间)
func GetFileStat(name string) FileStat {
	finfo, _ := os.Lstat(name)
	// linux环境下代码如下
	//linuxFileAttr := finfo.Sys().(*syscall.Stat_t)
	//fmt.Println("文件创建时间", SecondToTime(linuxFileAttr.Ctim.Sec))
	//fmt.Println("最后访问时间", SecondToTime(linuxFileAttr.Atim.Sec))
	//fmt.Println("最后修改时间", SecondToTime(linuxFileAttr.Mtim.Sec))

	// windows下代码如下
	winFileAttr := finfo.Sys().(*syscall.Win32FileAttributeData)
	return FileStat{
		CreationTime:   NanosecondToTime(winFileAttr.CreationTime.Nanoseconds()),
		LastAccessTime: NanosecondToTime(winFileAttr.LastAccessTime.Nanoseconds()),
		LastWriteTime:  NanosecondToTime(winFileAttr.LastWriteTime.Nanoseconds()),
	}
}
