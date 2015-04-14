package catch

import (
	"errors"
	"testing"
)

func TestCatch(t *testing.T) {
	err := func() (err error) {
		defer Catch(&err)
		Check(func() error {
			return errors.New("Err")
		}(), Id)
		return
	}()
	if err == nil || err.Error() != "Err" {
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

func BenchmarkCatch(b *testing.B) {
	var err error
	e := errors.New("foo")
	for i := 0; i < b.N; i++ {
		func() {
			defer Catch(&err)
			Check(e, Id)
		}()
	}
}
