package v1

import (
	"gitee.com/quant1x/gox/sys/mem"
	"os"
	"path/filepath"
	"unsafe"
)

// Cache 使用跨平台的mmap构建的快速缓存
type Cache struct {
	filename string     // 文件名
	f        *os.File   // 文件对象
	size     int64      // 尺寸
	data     mem.Object // 内存映射对象
}

// OpenCache 打开本地缓存
func OpenCache(name string, size int64) (*Cache, error) {
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
	//data , err :=mem.FileMap(f, mem.RDWR, 0)
	data, err := mem.OpenMapper(f, mem.RDWR, 0)
	if nil != err {
		return nil, err
	}

	return &Cache{
		filename: filename,
		f:        f,
		size:     size,
		data:     data,
	}, nil
}

// Close 关闭缓存
func (c *Cache) Close() {
	if c == nil {
		return
	}
	err := c.data.Flush()
	err = c.data.Unmap()
	if c.f != nil {
		_ = c.f.Close()
	}
	_ = err
}

// Flush 同步数据到磁盘
func (c *Cache) Flush() error {
	if c == nil {
		return os.ErrInvalid
	}
	return c.data.Flush()
}

// ToSlices 内存转切片
func ToSlices[E any](cache *Cache) []E {
	return v2ToSlices[E](cache)
}

// ToSlices 内存转切片
func v1ToSlices[E any](cache *Cache) []E {
	var tmpValue E
	size := int(unsafe.Sizeof(tmpValue))
	num := int(cache.size) / size
	ptr := &cache.data[0]
	s := (*[1 << 32]E)(unsafe.Pointer(ptr))[:num:num]
	return s
}

// ToSlices 内存转切片
func v2ToSlices[E any](cache *Cache) []E {
	var tmpValue E
	size := int(unsafe.Sizeof(tmpValue))
	num := int(cache.size) / size
	ptr := unsafe.SliceData(cache.data)
	pointer := (*E)(unsafe.Pointer(ptr))
	s := unsafe.Slice(pointer, num)
	return s
}
