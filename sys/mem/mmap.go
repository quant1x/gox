package mem

import (
	"errors"
	"os"
	"unsafe"
)

func syscallError(funcname string, e error) error {
	return os.NewSyscallError(funcname, e)
}

const (
	// RDONLY maps the memory read-only.
	// Attempts to write to the MemObject object will result in undefined behavior.
	RDONLY = 0
	// RDWR maps the memory as read-write. Writes to the MemObject object will update the
	// underlying file.
	RDWR = 1 << iota
	// COPY maps the memory as copy-on-write. Writes to the MemObject object will affect
	// memory, but the underlying file will remain unchanged.
	COPY
	// EXEC If EXEC is set, the mapped memory is marked as executable.
	EXEC
)

const (
	// ANON If the ANON flag is set, the mapped memory will not be backed by a file.
	ANON = 1 << iota
)

type sliceHeader struct {
	Data uintptr
	Len  int
	Cap  int
}

func addressAndLength(data []byte) (uintptr, uintptr) {
	//addr := (uintptr)(unsafe.Pointer(&Bytes))
	//length := uintptr(len(Bytes))
	header := (*sliceHeader)(unsafe.Pointer(&data))
	addr := header.Data
	length := uintptr(header.Len)
	return addr, length
}

// mmap_region maps part of a file into memory.
// The offset parameter must be a multiple of the system's page size.
// If length < 0, the entire file will be mapped.
// If ANON is set in flags, f is ignored.
func mmap_region(f *os.File, length int, prot, flags int, offset int64) ([]byte, error) {
	if offset%int64(os.Getpagesize()) != 0 {
		return nil, errors.New("offset parameter must be a multiple of the system's page size")
	}

	var fd uintptr
	if flags&ANON == 0 {
		fd = uintptr(f.Fd())
		if length < 0 {
			fi, err := f.Stat()
			if err != nil {
				return nil, err
			}
			length = int(fi.Size())
		}
	} else {
		if length <= 0 {
			return nil, errors.New("anonymous mapping requires non-zero length")
		}
		fd = ^uintptr(0)
	}
	return mmap(length, uintptr(prot), uintptr(flags), fd, offset)
}

// FileMap maps an entire file into memory.
// If ANON is set in flags, f is ignored.
func FileMap(f *os.File, prot, flags int) ([]byte, error) {
	return mmap_region(f, -1, prot, flags, 0)
}

// Object represents a file mapped into memory.
type Object []byte

// OpenMapper maps an entire file into memory.
//
//	If ANON is set in flags, f is ignored.
func OpenMapper(f *os.File, prot, flags int) (Object, error) {
	return mmap_region(f, -1, prot, flags, 0)
}

func OpenMmap(size int, f *os.File) (Object, error) {
	return mmap_region(f, -1, RDWR, size, 0)
}

func (m *Object) header() *sliceHeader {
	return (*sliceHeader)(unsafe.Pointer(m))
}

func (m *Object) addrLen() (uintptr, uintptr) {
	header := m.header()
	return header.Data, uintptr(header.Len)
}
func (m *Object) Bytes() []byte {
	buf := ([]byte)(*m)
	return buf
}

// Lock keeps the mapped region in physical memory, ensuring that it will not be
// swapped out.
func (m *Object) Lock() error {
	return mlock(m.Bytes())
}

// Unlock reverses the effect of Lock, allowing the mapped region to potentially
// be swapped out.
// If m is already unlocked, aan error will result.
func (m *Object) Unlock() error {
	return munlock(m.Bytes())
}

// Flush synchronizes the mapping's contents to the file's contents on disk.
func (m *Object) Flush() error {
	return mflush(m.Bytes())
}

// Unmap deletes the memory mapped region, flushes any remaining changes, and sets
// m to nil.
// Trying to read or write any remaining references to m after Unmap is called will
// result in undefined behavior.
// Unmap should only be called on the slice value that was originally returned from
// a call to FileMap. Calling Unmap on a derived slice may cause errors.
func (m *Object) Unmap() error {
	err := munmap(m.Bytes())
	*m = nil
	return err
}
