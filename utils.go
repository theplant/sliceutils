package sliceutils

import (
	"reflect"
)

func Unwrap(arr []interface{}, toValue interface{}) {
	v := reflect.ValueOf(toValue)
	if v.Kind() != reflect.Ptr {
		panic("toValue must be pointer")
	}

	ve := v.Elem()
	if v.Elem().Kind() != reflect.Slice {
		panic("toValue must be array")
	}

	len := len(arr)
	newSlice := reflect.MakeSlice(v.Elem().Type(), len, len)
	for i, ele := range arr {
		newSlice.Index(i).Set(reflect.ValueOf(ele))
	}
	ve.Set(newSlice)
}

func Wrap(val interface{}) (r []interface{}) {
	v := reflect.ValueOf(val)
	len := v.Len()
	r = make([]interface{}, len, len)
	for i := 0; i < len; i++ {
		r[i] = v.Index(i).Interface()
	}
	return
}
