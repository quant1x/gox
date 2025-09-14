//go:build !windows
// +build !windows

package mem

import (
	"unsafe"
	_ "unsafe" // for go:linkname

	"golang.org/x/sys/unix"
)

//go:linkname mem_unix_lock gitee.com/quant1x/gox/sys/mem.Lock
func mem_unix_Lock(data []byte) error {
	addr := uintptr(unsafe.Pointer(&data[0]))
	length := uintptr(len(data))
	return unix.Mlock(addr, length)
}

//go:linkname mem_unix_unlock gitee.com/quant1x/gox/sys/mem.Unlock
func mem_unix_Unlock(data []byte) error {
	addr := uintptr(unsafe.Pointer(&data[0]))
	length := uintptr(len(data))
	return unix.Munlock(addr, length)
}
