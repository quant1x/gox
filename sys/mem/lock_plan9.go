package mem

import (
	"syscall"
	_ "unsafe" // for go:linkname
)

//go:linkname mem_plan9_lock github.com/quant1x/gox/sys/mem.Lock
func mem_plan9_lock(data []byte) error {
	return syscall.EPLAN9
}

//go:linkname mem_plan9_unlock github.com/quant1x/gox/sys/mem.Lock
func mem_plan9_unlock(data []byte) error {
	return syscall.EPLAN9
}
