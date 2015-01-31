package goutil

import (
	"errors"
	"testing"
)

func TestPanicOnError(t *testing.T) {
	defer func() {
		s := recover()
		if s == nil {
			t.Fatal("PanicOnError did not panic")
		}
	}()

	err := errors.New("serious error")
	PanicOnError(err)
}
