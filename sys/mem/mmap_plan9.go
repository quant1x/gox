// Copyright 2020 Evan Shaw. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mem

import "syscall"

func mmap(len int, inprot, inflags, fd uintptr, off int64) ([]byte, error) {
	return nil, syscall.EPLAN9
}

func (m MemObject) flush() error {
	return syscall.EPLAN9
}

func (m MemObject) lock() error {
	return syscall.EPLAN9
}

func (m MemObject) unlock() error {
	return syscall.EPLAN9
}

func (m MemObject) unmap() error {
	return syscall.EPLAN9
}
