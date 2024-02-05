package api

import (
	"fmt"
	"testing"
)

func TestGetConcreteReflectValueAndType(t *testing.T) {
	type TT struct {
		A int
		B string
	}

	var list []TT

	a, b := GetConcreteReflectValueAndType(&list)
	fmt.Println(a, b)
	isPointer, bt := GetConcreteContainerInnerType(b)
	fmt.Println(isPointer, bt)
}
