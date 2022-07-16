package bdd

import (
	"fmt"
	"runtime/debug"
	"testing"
)

type Featurer interface {
	Test()
	Before(fn func()) Featurer
	Scenario(name string) Scenery
}

type featureStep struct {
	name         string
	before       func()
	t            *testing.T
	opts         map[int]interface{}
	scenarios    []*scenarioStep
	lastRunFails bool
	executed     bool
	parent       *suiteStep
}

func (f *featureStep) Test() {

	f.lastRunFails = false // this one is set TRUE from the failed scenario
	f.t.Run(f.name, func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				recover_(r, stack_(debug.Stack()), t)
			}
		}()

		f.executed = true

		if f.before != nil {
			f.before()
		}

		for _, s := range f.scenarios {
			s.test(t)
		}

	})

	f.doReport()
}

func (f *featureStep) skipped() bool {
	return !f.executed
}

func (f *featureStep) doReport() {

	if f.parent != nil {
		return
	}

	var decorator = OptDecorator(f.opts)
	if decorator == nil {
		decorator = &PlainTextDecorator{}
	}

	if reporter := OptReport(f.opts); reporter != nil {
		reporter.SetTitle(decorator._title(fmt.Sprintf("Summary feature: %s", f.t.Name())))
		for _, s := range f.scenarios {
			if s.skipped() {
				continue
			}

			reporter.Print(decorator._scenario(s.report()))
		}
		reporter.Generate()
	}
}

func (f *featureStep) Before(fn func()) Featurer {
	f.before = fn
	return f
}

func (f *featureStep) Scenario(name string) Scenery {
	scenario := &scenarioStep{name: name, parent: f, givens: []*givenStep{}}
	f.scenarios = append(f.scenarios, scenario)
	return scenario
}
