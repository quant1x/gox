package cache

import (
	"fmt"
	"hash/crc32"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"unsafe"
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
	size     int64        // 映射大小
	data     MemObject    // 内存映射对象
	header   *cacheHeader // 头结构指针
	checksum uint32       // 数据校验和
}

// 缓存头结构（16字节对齐）
type cacheHeader struct {
	Magic    uint32  // 文件标识
	Version  uint32  // 格式版本
	DataSize uint32  // 有效数据长度
	_        [4]byte // 填充对齐
}

const (
	headerSize  = 16         // 头结构大小
	magicNumber = 0xCAC1E5A5 // 魔数标识
	maxFileSize = 1 << 30    // 1GB最大文件限制
)

var (
	ErrInvalidFile   = fmt.Errorf("invalid cache file")
	ErrInvalidAccess = fmt.Errorf("invalid memory access")
	ErrChecksum      = fmt.Errorf("data checksum mismatch")
)

// OpenCache 创建或打开内存映射缓存
func OpenCache(name string, size int64) (*Cache, error) {
	if size < headerSize || size > maxFileSize {
		return nil, fmt.Errorf("invalid size %d", size)
	}

	// 确保目录存在
	if err := os.MkdirAll(filepath.Dir(name), 0755); err != nil {
		return nil, fmt.Errorf("create directory failed: %w", err)
	}

	// 打开文件
	f, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return nil, fmt.Errorf("open file failed: %w", err)
	}

	// 调整文件大小
	if err := f.Truncate(size); err != nil {
		f.Close()
		return nil, fmt.Errorf("truncate file failed: %w", err)
	}

	// 内存映射
	data, err := mmap(int(size), f)
	if err != nil {
		f.Close()
		return nil, fmt.Errorf("mmap failed: %w", err)
	}

	c := &Cache{
		filename: name,
		f:        f,
		size:     size,
		data:     data,
		header:   (*cacheHeader)(unsafe.Pointer(&data.Bytes()[0])),
	}

	// 初始化或验证头
	if err := c.initHeader(); err != nil {
		c.Close()
		return nil, err
	}

	return c, nil
}

// 初始化文件头
func (c *Cache) initHeader() error {
	if c.header.Magic == 0 {
		// 新文件初始化
		c.header.Magic = magicNumber
		c.header.Version = 1
		c.header.DataSize = 0
		return nil
	}

	// 验证现有文件
	if c.header.Magic != magicNumber {
		return ErrInvalidFile
	}
	return c.verifyData()
}

// Close 安全关闭缓存
func (c *Cache) Close() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.f == nil {
		return nil // 已关闭
	}

	var errs []error

	// 同步数据
	if err := c.data.Flush(); err != nil {
		errs = append(errs, fmt.Errorf("flush failed: %w", err))
	}

	// 解除映射
	if err := c.data.Unmap(); err != nil {
		errs = append(errs, fmt.Errorf("unmap failed: %w", err))
	}

	// 关闭文件
	if err := c.f.Close(); err != nil {
		errs = append(errs, fmt.Errorf("close file failed: %w", err))
	}

	c.f = nil
	if len(errs) > 0 {
		return fmt.Errorf("close errors: %v", errs)
	}
	return nil
}

// ReadData 类型安全读取（示例）
func (c *Cache) ReadData(offset uint32, dst []byte) error {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if offset+uint32(len(dst)) > uint32(c.size) {
		return ErrInvalidAccess
	}

	data := c.data.Bytes()
	copy(dst, data[headerSize+offset:])
	return nil
}

// WriteData 类型安全写入（示例）
func (c *Cache) WriteData(offset uint32, src []byte) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if offset+uint32(len(src)) > uint32(c.size)-headerSize {
		return ErrOutOfSpace
	}

	data := c.data.Bytes()
	copy(data[headerSize+offset:], src)
	c.updateChecksum()
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

	dataSize := c.header.DataSize
	if dataSize == 0 {
		return nil, io.EOF
	}

	// 检查类型对齐
	if uintptr(unsafe.Pointer(&c.data.Bytes()[headerSize]))%unsafe.Alignof(e) != 0 {
		return nil, fmt.Errorf("misaligned memory access")
	}

	sliceSize := int(dataSize) / int(eSize)
	ptr := unsafe.Pointer(&c.data.Bytes()[headerSize])
	return unsafe.Slice((*E)(ptr), sliceSize), nil
}

// 更新校验和
func (c *Cache) updateChecksum() {
	data := c.data.Bytes()[headerSize : headerSize+c.header.DataSize]
	c.checksum = crc32.ChecksumIEEE(data)
}

// 验证数据完整性
func (c *Cache) verifyData() error {
	data := c.data.Bytes()[headerSize : headerSize+c.header.DataSize]
	if crc32.ChecksumIEEE(data) != c.checksum {
		return ErrChecksum
	}
	return nil
}

// 跨平台mmap实现
func mmap(size int, f *os.File) (MemObject, error) {
	// Windows与其他系统实现差异处理
	if runtime.GOOS == "windows" {
		return windowsMmap(size, f)
	}
	return posixMmap(size, f)
}
