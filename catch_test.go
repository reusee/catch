package catch

import (
	"errors"
	"testing"
)

func TestCatch(t *testing.T) {
	var err error
	check := PkgChecker("test")
	func() {
		defer Catch(&err)
		e := errors.New("foobar")
		check(e, "FOO")
	}()
	if err == nil || err.Error() != "test: FOO\nfoobar" {
		t.Fail()
	}
}

func TestCatch2(t *testing.T) {
	err := func() (err error) {
		defer Catch(&err)
		return
	}()
	if err != nil {
		t.Fail()
	}
}

func TestCatch3(t *testing.T) {
	err := func() (err error) {
		defer Catch(&err)
		return errors.New("Err")
	}()
	if err == nil || err.Error() != "Err" {
		t.Fail()
	}
}

func TestNoError(t *testing.T) {
	var err error
	check := PkgChecker("test")
	func() {
		defer Catch(&err)
		check(nil, "test")
	}()
	if err != nil {
		t.Fail()
	}
}

func BenchmarkCatchError(b *testing.B) {
	var err error
	e := errors.New("foo")
	check := PkgChecker("bench")
	for i := 0; i < b.N; i++ {
		func() {
			defer Catch(&err)
			check(e, "bench")
		}()
	}
}

func BenchmarkNoError(b *testing.B) {
	var err error
	check := PkgChecker("bench")
	for i := 0; i < b.N; i++ {
		func() {
			defer Catch(&err)
			check(nil, "bench")
		}()
	}
}
