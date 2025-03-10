package mem

import (
	"errors"
	"golang.org/x/sys/windows"
	"sync"
	"unsafe"
)

// mmap on Windows is a two-step process.
// First, we call CreateFileMapping to get a handle.
// Then, we call MapviewToFile to get an actual pointer into memory.
// Because we want to emulate a POSIX-style mmap, we don't want to expose
// the handle -- only the pointer. We also want to return only a byte slice,
// not a struct, so it's convenient to manipulate.

// We keep this map so that we can get back the original handle from the memory address.

type addrInfo struct {
	file     windows.Handle
	mapview  windows.Handle
	writable bool
}

var handleLock sync.Mutex
var handleMap = map[uintptr]*addrInfo{}

func mmap(len int, prot, flags, hfile uintptr, off int64) ([]byte, error) {
	flProtect := uint32(windows.PAGE_READONLY)
	dwDesiredAccess := uint32(windows.FILE_MAP_READ)
	writable := false
	switch {
	case prot&COPY != 0:
		flProtect = windows.PAGE_WRITECOPY
		dwDesiredAccess = windows.FILE_MAP_COPY
		writable = true
	case prot&RDWR != 0:
		flProtect = windows.PAGE_READWRITE
		dwDesiredAccess = windows.FILE_MAP_WRITE
		writable = true
	}
	if prot&EXEC != 0 {
		flProtect <<= 4
		dwDesiredAccess |= windows.FILE_MAP_EXECUTE
	}

	// The maximum size is the area of the file, starting from 0,
	// that we wish to allow to be mappable. It is the sum of
	// the length the user requested, plus the offset where that length
	// is starting from. This does not map the Bytes into memory.
	maxSizeHigh := uint32((off + int64(len)) >> 32)
	maxSizeLow := uint32((off + int64(len)) & 0xFFFFFFFF)
	// TODO: Do we need to set some security attributes? It might help portability.
	h, errno := windows.CreateFileMapping(windows.Handle(hfile), nil, flProtect, maxSizeHigh, maxSizeLow, nil)
	if h == 0 {
		return nil, syscallError("CreateFileMapping", errno)
	}

	// Actually map a view of the Bytes into memory. The view's size
	// is the length the user requested.
	fileOffsetHigh := uint32(off >> 32)
	fileOffsetLow := uint32(off & 0xFFFFFFFF)
	addr, errno := windows.MapViewOfFile(h, dwDesiredAccess, fileOffsetHigh, fileOffsetLow, uintptr(len))
	if addr == 0 {
		_ = windows.CloseHandle(windows.Handle(h))
		return nil, syscallError("MapViewOfFile", errno)
	}
	handleLock.Lock()
	handleMap[addr] = &addrInfo{
		file:     windows.Handle(hfile),
		mapview:  h,
		writable: writable,
	}
	handleLock.Unlock()

	pointer := (*byte)(unsafe.Pointer(addr))
	m := unsafe.Slice(pointer, len)
	_ = flags
	return m, nil
}

func mlock(data []byte) error {
	return mem_windows_lock(data)
}

func munlock(data []byte) error {
	return mem_windows_unlock(data)
}

func mflush(data []byte) error {
	addr, length := addressAndLength(data)
	errno := windows.FlushViewOfFile(addr, length)
	if errno != nil {
		return syscallError("FlushViewOfFile", errno)
	}

	handleLock.Lock()
	defer handleLock.Unlock()
	handle, ok := handleMap[addr]
	if !ok {
		// should be impossible; we would've errored above
		return errors.New("unknown base address")
	}

	if handle.writable && handle.file != windows.Handle(^uintptr(0)) {
		if err := windows.FlushFileBuffers(handle.file); err != nil {
			return syscallError("FlushFileBuffers", err)
		}
	}

	return nil
}

func munmap(data []byte) error {
	addr, _ := addressAndLength(data)
	// Lock the UnmapViewOfFile along with the handleMap deletion.
	// As soon as we unmap the view, the OS is free to give the
	// same addr to another new map. We don't want another goroutine
	// to insert and remove the same addr into handleMap while
	// we're trying to remove our old addr/handle pair.
	handleLock.Lock()
	defer handleLock.Unlock()
	err := windows.UnmapViewOfFile(addr)
	if err != nil {
		return err
	}

	handle, ok := handleMap[addr]
	if !ok {
		// should be impossible; we would've errored above
		return errors.New("unknown base address")
	}
	delete(handleMap, addr)

	e := windows.CloseHandle(windows.Handle(handle.mapview))
	return syscallError("CloseHandle", e)
}
