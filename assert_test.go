package assert

import (
	"fmt"
	"testing"
)

type TestingMock struct {
	ErrorfStr []string
}

func (m *TestingMock) Errorf(p string, v ...interface{}) {
	m.ErrorfStr = append(m.ErrorfStr, fmt.Sprintf(p, v...))
}

func (TestingMock) Fail()    {}
func (TestingMock) FailNow() {}

// MyStruct is a sample struct
type MyStruct struct {
	Sub *MyStruct
}

func TestEqual_OK(t *testing.T) {
	Equal(t, "foo", "foo")
	Equal(t, true, true)
	Equal(t, MyStruct{}, MyStruct{})
}

func TestEqual_Fail(t *testing.T) {
	tm := &TestingMock{}

	Equal(tm, "foo", "bar")
	Equal(t, "! Expected `foo` to equal `bar`", tm.ErrorfStr[1])
	Equal(t, "- \"foo\" != \"bar\"", tm.ErrorfStr[2])
}

func TestNotEqual_OK(t *testing.T) {
	NotEqual(t, "foo", "bar")
	NotEqual(t, nil, false)
	NotEqual(t, MyStruct{}, MyStruct{&MyStruct{}})
	NotEqual(t, &MyStruct{}, MyStruct{})
}

func TestNotEqual_Fail(t *testing.T) {
	tm := &TestingMock{}

	NotEqual(tm, "foo", "foo")
	Equal(t, "! Expected `foo` to not equal `foo`", tm.ErrorfStr[1])
	Equal(t, "- Unexpected `foo`", tm.ErrorfStr[2])
}

func TestTrue_OK(t *testing.T) {
	True(t, true)
}

func TestTrue_Fail(t *testing.T) {
	tm := &TestingMock{}

	True(tm, false)
	Equal(t, "! Expected to be true", tm.ErrorfStr[1])
}

func TestFalse_OK(t *testing.T) {
	False(t, false)
}

func TestFalse_Fail(t *testing.T) {
	tm := &TestingMock{}

	False(tm, true)
	Equal(t, "! Expected to be false", tm.ErrorfStr[1])
}

func TestNil_OK(t *testing.T) {
	Nil(t, nil)

	var (
		nilChan      chan int
		nilFunc      func(int) int
		nilInterface interface{}
		nilMap       map[string]string
		myStruct     MyStruct
		nilSlice     []string
	)

	Nil(t, nilChan)
	Nil(t, nilFunc)
	Nil(t, nilInterface)
	Nil(t, nilMap)
	Nil(t, myStruct.Sub) // nil pointer
	Nil(t, nilSlice)
}

func TestNil_Fail(t *testing.T) {
	tm := &TestingMock{}

	Nil(tm, "hello!")
	Equal(t, "! Expected nil", tm.ErrorfStr[1])
	Equal(t, "- Unexpected `hello!`", tm.ErrorfStr[2])

}

func TestNotNil_OK(t *testing.T) {
	NotNil(t, "foo")
	NotNil(t, MyStruct{})
	NotNil(t, &MyStruct{})
}

func TestNotNil_Fail(t *testing.T) {
	tm := &TestingMock{}

	NotNil(tm, nil)
	Equal(t, "! Unexpected nil", tm.ErrorfStr[1])
}

func TestPanic_OK(t *testing.T) {
	Panic(t, "uh-oh!", func() {
		panic("uh-oh!")
	})
}

func TestPanic_Fail(t *testing.T) {
	{
		tm := &TestingMock{}

		Panic(tm, "uh-oh", func() {})
		Equal(t, "! Expected panic `uh-oh`", tm.ErrorfStr[1])
		Equal(t, "- No panic", tm.ErrorfStr[2])
	}

	{
		tm := &TestingMock{}

		Panic(tm, "uh-oh", func() {
			panic("boom")
		})
		Equal(t, "! Expected panic `uh-oh`", tm.ErrorfStr[1])
		Equal(t, "- Unexpected `boom`", tm.ErrorfStr[2])
	}
}

func TestTypeOf_OK(t *testing.T) {
	TypeOf(t, "assert.MyStruct", MyStruct{})
	TypeOf(t, "*assert.MyStruct", &MyStruct{})
}

func TestTypeOf_Fail(t *testing.T) {
	tm := &TestingMock{}

	TypeOf(tm, "assert.MyStruct", &MyStruct{})
	Equal(t, "! Expected type `assert.MyStruct` but got `*assert.MyStruct`",
		tm.ErrorfStr[1])
}

func TestExtraMessage(t *testing.T) {
	tm := &TestingMock{}

	Equal(tm, "foo", "bar", "baz", "qux")
	Equal(t, "! Expected `foo` to equal `bar`", tm.ErrorfStr[1])
	Equal(t, "- \"foo\" != \"bar\"", tm.ErrorfStr[2])
	Equal(t, "- baz", tm.ErrorfStr[3])
	Equal(t, "- qux", tm.ErrorfStr[4])
}
