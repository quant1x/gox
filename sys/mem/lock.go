package mem

import (
	"unsafe"
	_ "unsafe" // for go:linkname
)

// 返回地址和长度
func addrAndLength(data []byte) (uintptr, uintptr) {
	addr := uintptr(unsafe.Pointer(&data[0]))
	length := uintptr(len(data))
	return addr, length
}

// Lock 将进程的虚拟内存地址空间中的某段内存锁定在物理内存（RAM）中, 防止被交换到磁盘页面文件
//
//go:linkname Lock
func Lock(data []byte) error

// Unlock 解除内存区域的物理锁定, 允许操作系统再次将内存交换到磁盘
//
//go:linkname Unlock
func Unlock(data []byte) error
