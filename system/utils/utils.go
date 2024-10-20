package utils

import (
	"reflect"
)

func In_array(val interface{}, array interface{}) (exists bool, index int) {
	exists = false
	index = -1

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) {
				index = i
				exists = true
				return
			}
		}
	default:
		panic("unexpected reflect.Kind")
	}

	return
}

func Contains(val interface{}, array interface{}) bool {
	exists, _ := In_array(val, array)
	return exists
}
