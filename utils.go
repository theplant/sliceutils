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

	newSlice := reflect.MakeSlice(v.Elem().Type(), 0, 0)
	for _, ele := range arr {
		newSlice = reflect.Append(newSlice, reflect.ValueOf(ele))
	}

	ve.Set(newSlice)
}

func Wrap(val interface{}) (r []interface{}) {
	v := reflect.ValueOf(val)
	for i := 0; i< v.Len() ; i++ {
		r = append(r, v.Index(i).Interface())
	}
	return
}
