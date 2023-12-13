package api

import (
	"fmt"
	"testing"
)

func TestUnique(t *testing.T) {
	s := []int{1, 2, 3, 4, 1, 1, 2, 2, 5, 1, 1, 1, 1, 1, 5, 5, 5, 3, 3, 3}
	s = Unique(s)
	fmt.Println(s)
}
