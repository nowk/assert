package assert

import (
	"fmt"
	"github.com/kr/pretty"
	"reflect"
	"runtime"
	"strings"
	"time"
)

var errorPrefix = "! "

func handleMessages(t Testing, messages ...interface{}) {
	if len(messages) > 0 {
		t.Error(errorPrefix, "-", fmt.Sprint(messages...))
	}
}

// -- Assertion handlers

func assert(t Testing, success bool, f func(), callDepth int) {
	if !success {
		_, file, line, _ := runtime.Caller(callDepth + 1)
		t.Errorf("%s:%d", file, line)
		f()
		t.FailNow()
	}
}

func equal(t Testing, expected, got interface{}, callDepth int, messages ...interface{}) {
	fn := func() {
		for _, desc := range pretty.Diff(expected, got) {
			t.Error(errorPrefix, desc)
		}
		handleMessages(t, messages...)
	}
	assert(t, isEqual(expected, got), fn, callDepth+1)
}

func notEqual(t Testing, expected, got interface{}, callDepth int, messages ...interface{}) {
	fn := func() {
		t.Errorf("%s Unexpected: %#v", errorPrefix, got)
		handleMessages(t, messages...)
	}
	assert(t, !isEqual(expected, got), fn, callDepth+1)
}

func contains(t Testing, expected, got string, callDepth int, messages ...interface{}) {
	fn := func() {
		t.Errorf("%s Expected to find: %#v", errorPrefix, expected)
		t.Errorf("%s in: %#v", errorPrefix, got)
		handleMessages(t, messages...)
	}
	assert(t, strings.Contains(got, expected), fn, callDepth+1)
}

func notContains(t Testing, unexpected, got string, callDepth int, messages ...interface{}) {
	fn := func() {
		t.Errorf("%s Expected not to find: %#v", errorPrefix, unexpected)
		t.Errorf("%s in: %#v", errorPrefix, got)
		handleMessages(t, messages...)
	}
	assert(t, !strings.Contains(got, unexpected), fn, callDepth+1)
}

// -- TypeOF

func typeOf(t Testing, expected string, got interface{}, callDepth int, messages ...interface{}) {
	_type := reflect.TypeOf(got)
	v := _type.String()

	fn := func() {
		t.Errorf("%s Expected TypeOf %s, got %s", errorPrefix, expected, v)
		handleMessages(t, messages...)
	}

	assert(t, v == expected, fn, callDepth+1)
}

// -- Duration

func withinDuration(t Testing, duration time.Duration, goalTime, gotTime time.Time, callDepth int, messages ...interface{}) {
	fn := func() {
		t.Errorf("%s Expected %v to be within %v of %v", errorPrefix, gotTime, duration, goalTime)
		handleMessages(t, messages...)
	}
	actualDuration := goalTime.Sub(gotTime)
	if actualDuration < time.Duration(0) {
		actualDuration = -actualDuration
	}
	assert(t, actualDuration <= duration, fn, callDepth+1)
}

// -- Matching

func isEqual(expected, got interface{}) bool {
	if expected == nil {
		return isNil(got)
	}
	return reflect.DeepEqual(expected, got)
}

func isNil(got interface{}) bool {
	if got == nil {
		return true
	}
	value := reflect.ValueOf(got)
	switch value.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		return value.IsNil()
	}
	return false
}

// -- Public API

func Equal(t Testing, expected, got interface{}, messages ...interface{}) {
	equal(t, expected, got, 1, messages...)
}

func NotEqual(t Testing, expected, got interface{}, messages ...interface{}) {
	notEqual(t, expected, got, 1, messages...)
}

func True(t Testing, got interface{}, messages ...interface{}) {
	equal(t, true, got, 1, messages...)
}

func False(t Testing, got interface{}, messages ...interface{}) {
	equal(t, false, got, 1, messages...)
}

func Nil(t Testing, got interface{}, messages ...interface{}) {
	equal(t, nil, got, 1, messages...)
}

func NotNil(t Testing, got interface{}, messages ...interface{}) {
	notEqual(t, nil, got, 1, messages...)
}

func Contains(t Testing, expected, got string, messages ...interface{}) {
	contains(t, expected, got, 1, messages...)
}

func NotContains(t Testing, unexpected, got string, messages ...interface{}) {
	notContains(t, unexpected, got, 1, messages...)
}

func WithinDuration(t Testing, duration time.Duration, goalTime, gotTime time.Time, messages ...interface{}) {
	withinDuration(t, duration, goalTime, gotTime, 1, messages...)
}

func Panic(t Testing, err interface{}, fn func(), messages ...interface{}) {
	defer func() {
		equal(t, err, recover(), 3, messages...)
	}()
	fn()
}

func TypeOf(t Testing, expected string, got interface{}, messages ...interface{}) {
	typeOf(t, expected, got, 1, messages)
}
