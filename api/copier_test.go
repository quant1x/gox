package api

import (
	"fmt"
	"testing"
)

type TT struct {
	A int
	B int
}

func TestCopy(t *testing.T) {
	t1 := TT{
		A: 1,
		B: 2,
	}
	var t2 TT
	_ = Copy(&t2, &t1)
	fmt.Println(t2)
}
