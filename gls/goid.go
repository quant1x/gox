package gls

import (
	reflect "gitee.com/quant1x/gox/util/reflect2"
	"unsafe"
)

var (
	goidOffset uintptr = 128 // offset for go1.4
)

func init() {
	gType := reflect.TypeByName("runtime.g").(reflect.StructType)
	if gType == nil {
		panic("failed to get runtime.g type")
	}
	goidField := gType.FieldByName("goid")
	goidOffset = goidField.Offset()
}

// GoID returns the goroutine id of current goroutine
//
//go:nocheckptr
func GoID() int64 {
	g := getg()
	// TODO: fatal error: checkptr: pointer arithmetic result points to invalid allocation
	p_goid := (*int64)(unsafe.Pointer(g + goidOffset))
	return *p_goid
}

func getg() uintptr
