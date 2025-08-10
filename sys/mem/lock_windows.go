package mem

import (
	"os"
	"unsafe"   // for go:linkname
	_ "unsafe" // for go:linkname

	"golang.org/x/sys/windows"
)

// windows.VirtualLock 将进程的虚拟内存地址空间中的某段内存锁定在物理内存（RAM）中，防止被交换到磁盘页面文件
//
//go:linkname mem_windows_lock gitee.com/quant1x/gox/sys/mem.Lock
func mem_windows_lock(data []byte) error {
	addr := uintptr(unsafe.Pointer(&data[0]))
	length := uintptr(len(data))
	errno := windows.VirtualLock(addr, length)
	return os.NewSyscallError("VirtualLock", errno)
}

// windows.VirtualUnlock 解除内存区域的物理锁定，允许操作系统再次将内存交换到磁盘
//
//go:linkname mem_windows_unlock gitee.com/quant1x/gox/sys/mem.Unlock
func mem_windows_unlock(data []byte) error {
	addr := uintptr(unsafe.Pointer(&data[0]))
	length := uintptr(len(data))
	errno := windows.VirtualUnlock(addr, length)
	return os.NewSyscallError("VirtualUnlock", errno)
}
