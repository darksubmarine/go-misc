package failure

import (
	"errors"
	"testing"
)

func TestExtendedError_Error(t *testing.T) {
	errs := errors.New("some error")
	err := NewExtendedError(errs).
		WithString("string-key", "some meta value").
		WithInt("int-key", 123).
		WithFloat("float-key", 456.34).
		WithBool("bool-key", true)

	if !err.Is(errs) {
		t.FailNow()
	}

	t.Log(err)
}
