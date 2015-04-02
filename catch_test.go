package catch

import (
	"errors"
	"testing"
)

func TestCatch(t *testing.T) {
	err := func() (err error) {
		defer Catch(&err)
		Check("foo", func() error {
			return errors.New("Err")
		}())
		return
	}()
	if err == nil || err.Error() != "foo error - Err" {
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
