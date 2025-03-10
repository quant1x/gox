//go:build darwin || dragonfly || freebsd || linux || openbsd || solaris || netbsd
// +build darwin dragonfly freebsd linux openbsd solaris netbsd

package mem

import (
	"golang.org/x/sys/unix"
)

// 内存映射
// 参数说明：
// fd:    文件描述符
// offset: 文件偏移量（必须页对齐）
// length: 映射长度
// prot:  保护标志 (PROT_READ, PROT_WRITE 等)
// flags: 映射标志 (MAP_SHARED, MAP_PRIVATE 等)

func mmap(len int, inprot, inflags, fd uintptr, off int64) ([]byte, error) {
	flags := unix.MAP_SHARED
	prot := unix.PROT_READ
	switch {
	case inprot&COPY != 0:
		prot |= unix.PROT_WRITE
		flags = unix.MAP_PRIVATE
	case inprot&RDWR != 0:
		prot |= unix.PROT_WRITE
	}
	if inprot&EXEC != 0 {
		prot |= unix.PROT_EXEC
	}
	if inflags&ANON != 0 {
		flags |= unix.MAP_ANON
	}

	b, err := unix.Mmap(int(fd), off, len, prot, flags)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func mflush(data []byte) error {
	return unix.Msync(data, unix.MS_SYNC)
}

func mlock(data []byte) error {
	return unix.Mlock(data)
}

func munlock(data []byte) error {
	return unix.Munlock(data)
}

func munmap(data []byte) error {
	return unix.Munmap(data)
}
