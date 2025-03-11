// cache.go
package cache

import (
	"fmt"
	"hash/crc32"
	"os"
	"path/filepath"
	"sync"
	"unsafe"
)

const (
	dirMode  = 0755
	fileMode = 0644
)

// MemObject 封装内存映射操作接口
type MemObject interface {
	Flush() error
	Unmap() error
	Bytes() []byte
}

// Cache 使用内存映射的跨进程安全缓存
type Cache struct {
	mu       sync.RWMutex // 读写锁
	filename string       // 文件路径
	f        *os.File     // 文件对象
	userSize int64        // 用户指定的数据区容量（不含header）
	data     MemObject    // 内存映射对象
	header   *cacheHeader // 头结构指针
}

// 缓存头结构（16字节对齐）
type cacheHeader struct {
	headerSize  uint32 // 头信息长度, 包括headerSize字段
	magic       uint32 // 魔法数
	version     uint32 // 版本
	dataSize    uint32 // 数据长度
	checksum    uint32 // 数据校验和
	elementSize uint32 // 元素尺寸
	arrayLen    uint32 // 数组有效长度
	arrayCap    uint32 // 数组容量
	//_           [4]byte // 填充对齐
}

const (
	version     = 1
	headerSize  = 32
	magicNumber = 0xCAC1E5A5
	maxFileSize = 1 << 30
)

var (
	ErrInvalidFile   = fmt.Errorf("invalid cache file")
	ErrInvalidAccess = fmt.Errorf("invalid memory access")
	ErrChecksum      = fmt.Errorf("data checksum mismatch")
	ErrOutOfSpace    = fmt.Errorf("out of space")
)

// OpenCache 创建或打开内存映射缓存
func OpenCache(name string, userSize int64) (*Cache, error) {
	totalSize := headerSize + userSize
	if totalSize > maxFileSize || userSize < 0 {
		return nil, fmt.Errorf("invalid size %d (max allowed %d)",
			userSize, maxFileSize-headerSize)
	}

	if err := os.MkdirAll(filepath.Dir(name), dirMode); err != nil {
		return nil, fmt.Errorf("create directory failed: %w", err)
	}

	f, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE, fileMode)
	if err != nil {
		return nil, fmt.Errorf("open file failed: %w", err)
	}

	if err := f.Truncate(totalSize); err != nil {
		f.Close()
		return nil, fmt.Errorf("truncate failed: %w", err)
	}

	data, err := mmap(int(totalSize), f)
	if err != nil {
		f.Close()
		return nil, fmt.Errorf("mmap failed: %w", err)
	}

	c := &Cache{
		filename: name,
		f:        f,
		userSize: userSize,
		data:     data,
		header:   (*cacheHeader)(unsafe.Pointer(&data.Bytes()[0])),
	}

	if err := c.initHeader(); err != nil {
		c.Close()
		return nil, err
	}

	return c, nil
}

// 初始化文件头
func (c *Cache) initHeader() error {
	if c.header.magic == 0 {
		c.header.magic = magicNumber
		c.header.version = version
		c.header.dataSize = 0
		return nil
	}

	if c.header.magic != magicNumber {
		return ErrInvalidFile
	}
	return c.verifyData()
}

// WriteData 类型安全写入
func (c *Cache) WriteData(offset uint32, src []byte) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	end := offset + uint32(len(src))
	if end > uint32(c.userSize) {
		return ErrOutOfSpace
	}

	if end > c.header.dataSize {
		c.header.dataSize = end
	}

	data := c.data.Bytes()
	copy(data[headerSize+offset:], src)
	c.updateChecksum()
	return nil
}

// 更新校验和
func (c *Cache) updateChecksum() {
	data := c.data.Bytes()[headerSize : headerSize+c.header.dataSize]
	c.header.checksum = crc32.ChecksumIEEE(data)
}

// 验证数据完整性
func (c *Cache) verifyData() error {
	data := c.data.Bytes()[headerSize : headerSize+c.header.dataSize]
	if crc32.ChecksumIEEE(data) != c.header.checksum {
		return ErrChecksum
	}
	return nil
}

// Close 安全关闭
func (c *Cache) Close() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.f == nil {
		return nil
	}

	var errs []error
	if err := c.data.Flush(); err != nil {
		errs = append(errs, err)
	}
	if err := c.data.Unmap(); err != nil {
		errs = append(errs, err)
	}
	if err := c.f.Close(); err != nil {
		errs = append(errs, err)
	}

	c.f = nil
	if len(errs) > 0 {
		return fmt.Errorf("close errors: %v", errs)
	}
	return nil
}

func (c *Cache) Add(delta int) error {
	if delta < 0 {
		return fmt.Errorf("index out of range [%d]", delta)
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.header.arrayCap < c.header.arrayLen+uint32(delta) {
		return ErrOutOfSpace
	}
	c.header.arrayLen += uint32(delta)
	return nil
}

// ToSlice 安全转换为类型切片
func ToSlice[E any](c *Cache) ([]E, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	var e E
	eSize := unsafe.Sizeof(e)
	if eSize == 0 {
		return nil, fmt.Errorf("zero-sized type")
	}

	dataStart := uintptr(unsafe.Pointer(&c.data.Bytes()[headerSize]))
	if dataStart%unsafe.Alignof(e) != 0 {
		return nil, fmt.Errorf("memory address %x not aligned for %T (alignment %d)",
			dataStart, e, unsafe.Alignof(e))
	}

	usedElements := int(c.header.dataSize) / int(eSize)
	usedElements = int(c.userSize) / int(eSize)
	c.header.arrayCap = uint32(usedElements)
	ptr := unsafe.Pointer(dataStart)
	return unsafe.Slice((*E)(ptr), usedElements), nil
}
