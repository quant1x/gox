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
	//addr := (uintptr)(unsafe.Pointer(&data))
	//length := uintptr(len(data))
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
