package api

import (
	"github.com/mymmsc/gox/errors"
	"reflect"
	"strconv"
)

var (
	mapTag map[reflect.Type]map[int]reflect.StructField = nil
)

func init() {
	//fmt.Println("reflect_array init...")
	mapTag = make(map[reflect.Type]map[int]reflect.StructField)
	//fmt.Println("reflect_array init...OK")
}

func initTag(t reflect.Type) map[int]reflect.StructField {
	ma, mok := mapTag[t]
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
						mapTag[t] = ma
					}
					ma[index] = field
				}
			}
		}
	}
	return ma
}

func Convert(data []string, v interface{}) error {
	val := reflect.ValueOf(v)
	//t := reflect.TypeOf(v)
	//fieldNum := val.NumField()
	//_ = fieldNum
	obj := reflect.ValueOf(v)
	t := val.Type()
	if val.Kind() == reflect.Ptr {
		t = t.Elem()
		obj = obj.Elem()
	}
	ma := initTag(t)
	if ma == nil {
		return errors.New("can not Convert")
	}
	dl := len(data)
	for i := 0; i < dl; i++ {
		field, ok := ma[i]
		if ok {
			dv := data[i]
			ov := obj.FieldByName(field.Name)
			if ov.CanSet() {
				var value interface{}
				switch ov.Interface().(type) {
				case string:
					value = string(dv)
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
