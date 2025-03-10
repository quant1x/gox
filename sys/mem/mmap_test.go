package mem

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

var (
	name = "data.dat"
)

func Test_mmap(t *testing.T) {
	size := 1024
	dir := filepath.Dir(name)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		fmt.Println(err)
		return
	}
	filename := name
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if nil != err {
		fmt.Println(err)
		return
	}
	err = f.Truncate(int64(size))
	if nil != err {
		fmt.Println(err)
		return
	}
	//data , err :=mem.FileMap(f, mem.RDWR, 0)
	//data, err := mem.OpenMapper(int(size), mem.RDWR, 0, f.Fd(), 0)
	data, err := mmap(size, RDWR, 0, f.Fd(), 0)
	if nil != err {
		fmt.Println(err)
		return
	}
	fmt.Println(data)
	data[1] = 49
	err = mlock(data)
	if nil != err {
		fmt.Println(err)
		return
	}
	err = munlock(data)
	if nil != err {
		fmt.Println(err)
		return
	}
	err = mflush(data)
	if nil != err {
		fmt.Println(err)
		return
	}
	err = munmap(data)
	if nil != err {
		fmt.Println(err)
		return
	}
	fmt.Println("ok")
	err = f.Close()
	if nil != err {
		fmt.Println(err)
		return
	}
}
