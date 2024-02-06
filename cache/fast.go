package cache

import (
	"fmt"
	"os"
	"path/filepath"
	"syscall"
	"unsafe"
)

type FastCache struct {
	filename string
	f        *os.File
	size     int64
	//Data     []byte
	Data MMap
}

func OpenCache(name string, size int64) (*FastCache, error) {
	dir := filepath.Dir(name)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return nil, err
	}
	filename := name
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if nil != err {
		return nil, err
	}
	err = f.Truncate(size)
	if nil != err {
		return nil, err
	}
	//data, err := syscall.Mmap(int(f.Fd()), 0, int(size), syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)
	data, err := mmap(int(size), RDWR, 0, f.Fd(), 0)
	if nil != err {
		return nil, err
	}

	return &FastCache{
		filename: filename,
		f:        f,
		size:     size,
		Data:     data,
	}, nil
}

func (fc *FastCache) Close() {
	if fc == nil {
		return
	}
	err := fc.Data.Flush()
	_ = syscall.Munmap(fc.Data)
	if fc.f != nil {
		_ = fc.f.Close()
	}
	_ = err
}

func ToSlices2[E any](fc *FastCache) []E {
	ptr := &fc.Data[0]
	fmt.Printf("fc.Data addr=%p\n", fc.Data)
	//unsafe.Slice(ptr, 2)
	p := (*[1 << 32]E)(unsafe.Pointer(ptr))[:]
	fmt.Printf("p addr=%p\n", p)
	return p
}

func ToSlices[E any](fc *FastCache) []E {
	var tmpValue E
	size := int(unsafe.Sizeof(tmpValue))
	num := int(fc.size) / size
	ptr := &fc.Data[0]
	//fmt.Printf("fc.Data addr=%p\n", fc.Data)
	//unsafe.Slice(ptr, 2)
	p := (*[1 << 32]E)(unsafe.Pointer(ptr))[:]
	//fmt.Printf("p addr=%p\n", p)
	//p = slices.Grow(p, num)
	return p[:0:num]
}
