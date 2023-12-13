package api

import (
	"slices"
)

type canUnique interface {
	~string | ~int | ~int32 | ~int64 | ~float32 | ~float64
}

// Unique 切片去重, 按照升序排序
func Unique[E canUnique](s []E) []E {
	slices.Sort(s)
	return slices.Compact(s)
}
