// Copyright 2011 Evan Shaw. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file defines the common package interface and contains a little bit of
// factored out logic.

// Package mmap allows mapping files into memory. It tries to provide a simple, reasonably portable interface,
// but doesn't go out of its way to abstract away every little platform detail.
// This specifically means:
//   - forked processes may or may not inherit mappings
//   - a file's timestamp may or may not be updated by writes through mappings
//   - specifying a size larger than the file's actual size can increase the file's size
//   - If the mapped file is being modified by another process while your program's running, don't expect consistent results between platforms
package mem

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
	// If EXEC is set, the mapped memory is marked as executable.
	EXEC
)

const (
	// If the ANON flag is set, the mapped memory will not be backed by a file.
	ANON = 1 << iota
)

// OpenMapper 打开一个内存映射
func OpenMapper(len int, prot, flags, hfile uintptr, off int64) ([]byte, error) {
	return mmap(len, prot, flags, hfile, off)
}
