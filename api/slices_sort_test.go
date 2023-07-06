package api

import (
	"fmt"
	"testing"
)

func TestSliceUnique(t *testing.T) {
	s := []int{1, 2, 3, 4, 1, 1, 2, 2, 5, 1, 1, 1, 1, 1, 5, 5, 5, 3, 3, 3}
	s = SliceUnique(s, func(a, b int) int {
		return a - b
	})
	fmt.Println(s)
}
