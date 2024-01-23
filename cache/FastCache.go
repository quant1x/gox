package cache

import (
	"os"
	"path/filepath"
	"syscall"
)

type FastCache struct {
	filename string
	f        *os.File
	size     int64
	Data     []byte
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
	data, err := syscall.Mmap(int(f.Fd()), 0, int(size), syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)
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
	if fc != nil && fc.f != nil {
		_ = fc.f.Close()
		_ = syscall.Munmap(fc.Data)
	}
}
