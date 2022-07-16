package bdd

import "fmt"

type Givener interface {
	fmt.Stringer
	And(name string, fn func()) Givener
	But(name string, fn func()) Givener
	When(name string, fn func()) *whenStep
}

type givenStep struct {
	name         string
	type_        string
	fn           func()
	lastRunFails bool
	parent       *scenarioStep
}

func (g *givenStep) And(name string, fn func()) Givener {
	return g.parent.given(name, "And", fn)
}

func (g *givenStep) But(name string, fn func()) Givener {
	return g.parent.given(name, "But", fn)
}

func (g *givenStep) String() string {
	return fmt.Sprintf("%s: %s", g.type_, g.name)
}

func (g *givenStep) When(name string, fn func()) *whenStep {
	when := &whenStep{name: name, fn: fn, parent: g.parent}
	g.parent.whens = append(g.parent.whens, when)
	return when
}
