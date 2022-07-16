package ds__test

import (
	"github.com/darksubmarine/versatile/ds_"
	"io/fs"
	"os"
	"testing"
)

func TestExtendedError_Error(t *testing.T) {
	invalidFileName := "non-existing"
	var extErr ds_.Error

	_, err := os.Open(invalidFileName)
	if err == nil {
		t.Errorf("file %s must not exists!", invalidFileName)
		t.FailNow()
	}

	extErr = ds_.NewError(fs.ErrNotExist).WithString("filename", invalidFileName)

	if !extErr.Is(fs.ErrNotExist) {
		t.Errorf("the error must be '%s' but got '%s'", fs.ErrNotExist, extErr)
	}

	if extErr.Is(fs.ErrClosed) {
		t.Errorf("the error must be '%s' but got '%s'", fs.ErrClosed, extErr)
	}

}
