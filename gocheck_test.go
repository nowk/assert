package assert

import "testing"
import "gopkg.in/check.v1"

func TestSuite(t *testing.T) {
	check.TestingT(t)
}

type tSuite struct{}

var _ = check.Suite(&tSuite{})

func (s *tSuite) TestGocheck(c *check.C) {
	Equal(c, "foo", "foo")
	Equal(c, true, true)

	myStructA := MyStruct{}
	myStructB := MyStruct{}
	Equal(c, myStructA, myStructB)

	// Equal(c, "foo", "bar", "this should blow up")
}
