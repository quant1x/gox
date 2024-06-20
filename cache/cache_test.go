package cache

import (
	"fmt"
	"gitee.com/quant1x/gox/api"
	"testing"
	"unsafe"
)

const (
	filename = "data.dat"
)

type TT struct {
	A int
	B [6]byte
}

func TestToSlices(t *testing.T) {
	num := 2
	size := int64(unsafe.Sizeof(TT{}))

	fc, err := OpenCache(filename, size*int64(num))
	if err != nil {
		fmt.Println(fc, err)
		return
	}
	defer fc.Close()
	list := ToSlices[TT](fc)
	for i := 0; i < num; i++ {
		d := i + 1
		var tt TT
		tt.A = d
		str := fmt.Sprintf("%d", d)
		b := [6]byte{}
		copy(b[:], api.String2Bytes(str))
		copy(tt.B[:], b[:])
		list[i] = tt
	}
	fmt.Println(len(list))
	fmt.Printf("%+v\n", list)
}
