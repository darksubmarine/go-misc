package bdd

import (
	"fmt"
	"github.com/stretchr/testify/assert"
)

type whenStep struct {
	name   string
	fn     func()
	parent *scenarioStep
}

func (w *whenStep) String() string {
	return fmt.Sprintf("When: %s", w.name)
}

func (w *whenStep) Then(name string, fn func(assertion *Assertions)) *thenStep {
	then := &thenStep{name: name, fn: fn, parent: w.parent, assertion: &Assertions{assert: assert.New(w.parent.parent.t), sErr: []*ContextErr{}}}
	w.parent.thens = append(w.parent.thens, then)
	return then
}
