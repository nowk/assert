package assert

type Testing interface {
	Error(...interface{})
	Errorf(string, ...interface{})
	Fail()
	FailNow()
}
