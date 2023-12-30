package api

import (
	"errors"
	"reflect"
	"strconv"
	"sync"
)

var (
	ErrNotConvert = errors.New("can not Convert")
)

var (
	__tagArrayMutex sync.RWMutex
	// 结构体 tag array的反射的字段缓存
	__mapTagArray = map[reflect.Type]map[int]reflect.StructField{}
)

func initTag(t reflect.Type) map[int]reflect.StructField {
	__tagArrayMutex.RLock()
	ma, mok := __mapTagArray[t]
	__tagArrayMutex.RUnlock()
	if mok {
		return ma
	}
	ma = nil
	fieldNum := t.NumField()
	for i := 0; i < fieldNum; i++ {
		field := t.Field(i)
		tag := field.Tag
		if len(tag) > 0 {
			tv, ok := tag.Lookup("array")
			if ok {
				index, err := strconv.Atoi(tv)
				if err == nil {
					if ma == nil {
						ma = make(map[int]reflect.StructField)
						__mapTagArray[t] = ma
					}
					__tagArrayMutex.Lock()
					ma[index] = field
					__tagArrayMutex.Unlock()

				}
			}
		}
	}
	return ma
}

// Convert 将字符串数组按照下标的序号反射给一个结构体
func Convert[T any](data []string, v *T) error {
	obj := reflect.ValueOf(v)
	t := obj.Type()
	if obj.Kind() == reflect.Ptr {
		t = t.Elem()
		obj = obj.Elem()
	}
	ma := initTag(t)
	if ma == nil {
		return ErrNotConvert
	}
	//fieldNum := t.NumField()
	fieldNum := len(data)
	for i := 0; i < fieldNum; i++ {
		field, ok := ma[i]
		if ok {
			dv := data[i]
			ov := obj.FieldByName(field.Name)
			if ov.CanSet() {
				var value interface{}
				switch ov.Interface().(type) {
				case string:
					value = dv
				case int8:
					t := ParseInt(dv)
					value = int8(t)
				case int16:
					t := ParseInt(dv)
					value = int16(t)
				case int32:
					t := ParseInt(dv)
					value = int32(t)
				case int64:
					t := ParseInt(dv)
					value = int64(t)
				case uint8:
					t := ParseUint(dv)
					value = uint8(t)
				case uint16:
					t := ParseUint(dv)
					value = uint16(t)
				case uint32:
					t := ParseUint(dv)
					value = uint32(t)
				case uint64:
					t := ParseUint(dv)
					value = t
				case float32:
					t := ParseFloat(dv)
					value = float32(t)
				case float64:
					t := ParseFloat(dv)
					value = t
				case bool:
					t, _ := strconv.ParseBool(dv)
					value = t
				default:
					value = dv
				}
				ov.Set(reflect.ValueOf(value))
			}
		}
	}
	return nil
}
