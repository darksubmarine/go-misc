package bdd

import (
	"github.com/stretchr/testify/assert"
)

type Assertions struct {
	assert *assert.Assertions
	sErr   []*ContextErr
}

func (a *Assertions) Reset() {
	a.sErr = make([]*ContextErr, 0)
}

func (a *Assertions) Report() []string {
	toRet := make([]string, len(a.sErr))
	for i, e := range a.sErr {
		toRet[i] = e.Error()
	}
	return toRet
}

func (a *Assertions) _e(v bool, fnName string, message string) {
	if !v {
		a.sErr = append(a.sErr, newContextErrFromAssertion(fnName, message))
	}
}
