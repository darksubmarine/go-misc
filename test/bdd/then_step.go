package bdd

import "fmt"

type thenStep struct {
	name      string
	fn        func(assertion *Assertions)
	parent    *scenarioStep
	assertion *Assertions
}

func (th *thenStep) HasError() bool {
	return len(th.assertion.sErr) > 0
}

func (th *thenStep) Assert() bool {
	th.assertion.Reset()
	th.fn(th.assertion)
	return th.HasError()
}

func (th *thenStep) String() string {
	return fmt.Sprintf("Then: %s", th.name)
}

func (th *thenStep) Report() []string {
	return th.assertion.Report()
}

func (th *thenStep) Test() {
	th.parent.parent.Test()
}
