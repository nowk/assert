package assert

import (
	"fmt"
	. "github.com/kr/pretty"
	. "gopkg.in/nowk/assert.v2/tests"
	. "gopkg.in/nowk/go-calm.v1"
	"reflect"
	"runtime"
)

const (
	Prefix = "! "
)

type Testing interface {
	Errorf(string, ...interface{})
	Fail()
	FailNow()
}

// Fail fails a test and outputs the caller and message
// In most cases callDepth is 2. Panics are 4.
func Fail(t Testing, callDepth int, msg string, x ...interface{}) {
	_, f, ln, _ := runtime.Caller(callDepth)
	t.Errorf("%s:%d", f, ln)
	t.Errorf("%s%s", Prefix, msg)

	for _, v := range x {
		t.Errorf("- %s", v)
	}

	t.FailNow()
}

func unshift(i []interface{}, m ...interface{}) []interface{} {
	return append(m, i...)
}

func Equal(t Testing, exp, got interface{}, m ...interface{}) {
	if !IsEqual(exp, got) {
		var i []interface{}
		for _, v := range Diff(exp, got) {
			i = append(i, v)
		}
		m = unshift(m, i...)

		Fail(t, 2, fmt.Sprintf("Expected `%s` to equal `%s`", exp, got), m...)
	}
}

func NotEqual(t Testing, exp, got interface{}, m ...interface{}) {
	if IsEqual(exp, got) {
		m = unshift(m, fmt.Sprintf("Unexpected `%s`", got))
		Fail(t, 2, fmt.Sprintf("Expected `%s` to not equal `%s`", exp, got), m...)
	}
}

func True(t Testing, got interface{}, m ...interface{}) {
	if !IsEqual(true, got) {
		Fail(t, 2, "Expected to be true", m...)
	}
}

func False(t Testing, got interface{}, m ...interface{}) {
	if !IsEqual(false, got) {
		Fail(t, 2, "Expected to be false", m...)
	}
}

func Nil(t Testing, got interface{}, m ...interface{}) {
	if !IsEqual(nil, got) {
		m = unshift(m, fmt.Sprintf("Unexpected `%s`", got))
		Fail(t, 2, "Expected nil", m...)
	}
}

func NotNil(t Testing, got interface{}, m ...interface{}) {
	if IsEqual(nil, got) {
		Fail(t, 2, "Unexpected nil", m...)
	}
}

func Panic(t Testing, exp string, fn func(), m ...interface{}) {
	msg := fmt.Sprintf("Expected panic `%s`", exp)
	err := Calm(fn)
	if err == nil {
		m = unshift(m, "No panic")
		Fail(t, 2, msg, m...)
		return
	}

	got := err.Error()
	if !IsEqual(exp, got) {
		m = unshift(m, fmt.Sprintf("Unexpected `%s`", got))
		Fail(t, 4, msg, m...)
	}
}

func TypeOf(t Testing, exp string, got interface{}, m ...interface{}) {
	to := reflect.TypeOf(got)
	vs := to.String()
	if !IsEqual(exp, vs) {
		Fail(t, 2, fmt.Sprintf("Expected type `%s` but got `%s`", exp, vs), m...)
	}
}

// Assert is a general use assert to test a bool with custom message.
// If you need control over callDepth use Fail directly
func Assert(t Testing, ok bool, msg string, m ...interface{}) {
	if !ok {
		Fail(t, 2, msg, m...)
	}
}
