package api

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

type T1 struct {
	f1     string "f one"
	f2     string
	f3     string `f three`
	f4, f5 int64  `f four and five`
}

type T2 struct {
	F1 string `array:"0"`
	F2 string `array:"2"`
}

func Test_main(t *testing.T) {
	t0 := T1{f1: "1", f2: "2", f3: "3", f4: 4, f5: 5}
	t1 := reflect.TypeOf(t0)
	f1, _ := t1.FieldByName("f1")
	fmt.Println(f1.Tag) // f one
	f4, _ := t1.FieldByName("f4")
	fmt.Println(f4.Tag) // f four and five
	f5, _ := t1.FieldByName("f5")
	fmt.Println(f5.Tag) // f four and five

	ts, err := json.Marshal(t0)
	if err != nil {
		t.Error(err)
	} else {
		var t2 T1
		err = json.Unmarshal(ts, &t2)
		if err != nil {
			t.Error(err)
		}
	}

	fmt.Println(ts)
}

func TestConvert(t *testing.T) {
	str := "a0,a1,a2,a3,a4,a5"
	sa := strings.Split(str, ",")
	var t2 T2
	_ = Convert(sa, &t2)
	fmt.Println(t2)
}
