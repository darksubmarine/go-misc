package bdd

import (
	"errors"
	"fmt"
)

func newContextErrFromAssertion(fn string, err string) *ContextErr {
	return &ContextErr{assertFn: fn, err: err}
}

type ContextErr struct {
	assertFn string
	err      string
}

func (e *ContextErr) Err() error {
	return errors.New(e.Error())
}

func (e *ContextErr) Error() string {
	if e.err == "" {
		return fmt.Sprintf("assertion=%s", e.assertFn)
	}
	return fmt.Sprintf("assertion=%s %s", e.assertFn, e.err)
}
