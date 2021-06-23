package reflection

import (
	"reflect"
)

func walk(x interface{}, fn func(input string)) {
	val := getValue(x)
	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walk(val.Index(i).Interface(), fn)
		}
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)
			walk(field.Interface(), fn)
		}
	}
}
func getValue(x interface{}) (val reflect.Value) {
	val = reflect.ValueOf(x)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	return
}
