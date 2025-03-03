package api

import (
	"reflect"
	"sync/atomic"
	"unsafe"
)

// ChanIsClosed 判断channel是否关闭
func ChanIsClosed(ch any) bool {
	return v2ChanIsClosed(ch)
}

// 官方版本: 虽然可以检测, 但是对正常的业务逻辑会丢数据
func v1ChanIsClosed(ch chan int) (closed bool) {
	select {
	case _, ok := <-ch:
		if !ok {
			closed = true // channel已关闭
		}
	default:
		closed = false // channel未关闭
	}
	return
}

func v2ChanIsClosed(ch any) bool {
	if reflect.TypeOf(ch).Kind() != reflect.Chan {
		panic("非channel类型")
	}

	// 获取channel的hchan指针
	// this function will return true if chan.closed > 0
	// see hchan on https://github.com/golang/go/blob/master/src/runtime/chan.go
	// type hchan struct {
	// qcount   uint           // total data in the queue
	// dataqsiz uint           // size of the circular queue
	// buf      unsafe.Pointer // points to an array of dataqsiz elements
	// elemsize uint16
	// closed   uint32
	// **
	//c := reflect.ValueOf(ch)
	hchanPtr := reflect.ValueOf(ch).UnsafePointer() // Go 1.18+

	// 计算closed字段偏移量（需根据Go版本调整）
	// 假设hchan结构：qcount(uint), dataqsiz(uint), buf(unsafe.Pointer), elemsize(uint16), closed(uint32)
	closedOffset := unsafe.Sizeof(uint(0))*2 + unsafe.Sizeof(unsafe.Pointer(nil)) + unsafe.Sizeof(uint16(0))
	closedPtr := (*uint32)(unsafe.Add(hchanPtr, closedOffset))

	return atomic.LoadUint32(closedPtr) > 0
}
