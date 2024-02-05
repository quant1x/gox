package api

import "reflect"

func GetConcreteContainerInnerType(in reflect.Type) (inInnerWasPointer bool, inInnerType reflect.Type) {
	inInnerType = in.Elem()
	inInnerWasPointer = false
	if inInnerType.Kind() == reflect.Ptr {
		inInnerWasPointer = true
		inInnerType = inInnerType.Elem()
	}
	return inInnerWasPointer, inInnerType
}

func GetConcreteReflectValueAndType(in any) (reflect.Value, reflect.Type) {
	value := reflect.ValueOf(in)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	return value, value.Type()
}

var errorInterface = reflect.TypeOf((*error)(nil)).Elem()

func IsErrorType(outType reflect.Type) bool {
	if outType.Kind() != reflect.Interface {
		return false
	}

	return outType.Implements(errorInterface)
}
