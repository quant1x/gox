package api

import (
	"reflect"
)

func init() {
	//fmt.Println("array init...")
	//fmt.Println("array init...OK")
}

// Reverse 反转切片
func Reverse(slice interface{}) interface{} {
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice /*&& v.Kind() != reflect.Array*/ {
		return slice
	}
	//elemType := reflect.TypeOf(slice).Elem()
	//et := elemType.Kind()
	//_ = et
	count := v.Len()
	count1 := v.Cap()
	_ = count1
	for i, j := 0, count -1; i < j; i, j = i+1, j-1 {
		//arr[i], arr[j] = arr[j], arr[i]
		//if v.Index(i).CanAddr() && v.Index(j).CanAddr() {
		//	continue
		//}
		tmpI := v.Index(i).Interface()
		tmpJ := v.Index(j).Interface()
		v.Index(i).Set(reflect.ValueOf(tmpJ))
		v.Index(j).Set(reflect.ValueOf(tmpI))
	}

	return slice
}
