# Assert

[![Build Status](https://travis-ci.org/nowk/assert.svg?branch=master)](https://travis-ci.org/nowk/assert)
[![GoDoc](https://godoc.org/github.com/nowk/assert?status.svg)](http://godoc.org/github.com/nowk/assert)

Simple test assertions for Go. This is a fork of a fork of a fork of [bmizerany/assert][original]
with improved support for things like nil pointers, etc.

[original]: http://github.com/bmizerany/assert

---

###### *This is a fully refactored version and no longer tracks it's birthed forks.*
* Removes some additional assertions that were from the previous forks. 
* Exposes some of the internal API to allow for extending within your project.
* New failure messages

Installation
------------

    $ go get gopkg.in/nowk/assert.v2

API
---

Equality

    func Equal(t Testing, exp interface{}, got interface{}, m ...interface{})
---

    func NotEqual(t Testing, exp interface{}, got interface{}, m ...interface{})

Truth

    func True(t Testing, got interface{}, m ...interface{})
---

    func False(t Testing, got interface{}, m ...interface{})

Nil

    func Nil(t Testing, got interface{}, m ...interface{})
---

    func NotNil(t Testing, got interface{}, m ...interface{})

TypeOf

    func TypeOf(t Testing, exp string, got interface{}, m ...interface{})

Panic

    func Panic(t Testing, exp string, fn func(), m ...interface{})

---

`Testing` is an interface to allow for use across multiple testing libraries. 

    type Testing interface {
      Errorf(string, ...interface{})
      Fail()
      FailNow()
    }

Currently tested are the core `testing` and `gocheck`.

Example
-------

```go
package main

import "gopkg.in/nowk/assert.v2"

type CoolStruct struct{}

func TestThings(t *testing.T) {
  myString := "cool"

  assert.Equal(t, "cool", myString, "myString should be equal")
  assert.NotEqual(t, "nope", myString)

  var myStruct CoolStruct
  assert.Nil(t, myStruct)
}
```

See [assert_test.go][assert_test] for more usage examples.

Output
------

You can add extra information to test failures by passing in any number of extra
arguments:

```go
assert.Equal(t, "foo", "bar", "when did foo ever equal bar?")
```

```console
% go test
--- FAIL: TestImportantFeature (0.00 seconds)
	assert.go:18: /Users/tyson/go/src/github.com/foo/bar/main_test.go:31
	assert.go:38: ! Expected `foo` to equal `bar`
	assert.go:38: - \"foo\" != \"bar\"
	assert.go:40: - when did foo ever equal bar?
FAIL
exit status 1
FAIL	github.com/foo/bar	0.017s
```

[assert_test]: https://github.com/nowk/assert/blob/master/assert_test.go

License
-------

Copyright Blake Mizerany and Keith Rarick. Licensed under the [MIT
license](http://opensource.org/licenses/MIT). Additional modifications by
Stovepipe Studios.

