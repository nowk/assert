package assert

import "testing"
import "gopkg.in/check.v1"

func TestSuite(t *testing.T) {
	check.TestingT(t)
}

type tSuite struct{}

var _ = check.Suite(&tSuite{})

func (s *tSuite) TestGocheck_OK(c *check.C) {
	Equal(c, "foo", "foo")
	Equal(c, true, true)
	Equal(c, MyStruct{}, MyStruct{})
}

func (s *tSuite) TestGocheck_Fail(c *check.C) {
	tm := &TestingMock{}

	Equal(tm, "foo", "bar")
	Equal(c, "! Expected `foo` to equal `bar`", tm.ErrorfStr[1])
	Equal(c, "- \"foo\" != \"bar\"", tm.ErrorfStr[2])
}
