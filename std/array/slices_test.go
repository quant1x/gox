package array

import (
	"fmt"
	"testing"
)

func TestZeroCopyResize(t *testing.T) {
	a := []int{1, 2}
	b := ZeroCopyResize(a, 4)
	fmt.Println(a, b)
}
