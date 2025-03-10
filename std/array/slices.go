package array

import "unsafe"

func ZeroCopyResize(s []int, newCap int) []int {
	if newCap <= cap(s) {
		return s[:newCap]
	}

	// 获取原始切片底层数组指针
	ptr := unsafe.Pointer(unsafe.SliceData(s))

	// 创建新切片指向同一内存区域
	newSlice := unsafe.Slice((*int)(ptr), newCap)

	// 拷贝原数据到新切片（仅拷贝有效数据）
	copy(newSlice, s)
	return newSlice
}
