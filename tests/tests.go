package tests

import (
	"reflect"
)

func IsEqual(exp, got interface{}) bool {
	if exp == nil {
		return IsNil(got)
	}

	return reflect.DeepEqual(exp, got)
}

func IsNil(got interface{}) bool {
	if got == nil {
		return true
	}

	switch vo := reflect.ValueOf(got); vo.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr,
		reflect.Slice:

		return vo.IsNil()
	}

	return false
}
