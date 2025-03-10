package mem

import (
	"syscall"
	_ "unsafe" // for go:linkname
)

//go:linkname mem_plan9_lock gitee.com/quant1x/gox/sys/mem.Lock
func mem_plan9_lock(data []byte) error {
	return syscall.EPLAN9
}

//go:linkname mem_plan9_unlock gitee.com/quant1x/gox/sys/mem.Lock
func mem_plan9_unlock(data []byte) error {
	return syscall.EPLAN9
}
