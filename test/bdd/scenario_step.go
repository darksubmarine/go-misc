package bdd

import (
	"fmt"
	"runtime/debug"
	"testing"
)

type Scenery interface {
	Given(name string, fn func()) Givener
	Before(fn func()) ScenarioSteper
}

type ScenarioSteper interface {
	Given(name string, fn func()) Givener
}

type scenarioStep struct {
	name         string
	before       func()
	lastRunFails bool
	executed     bool
	testName     string
	panicStack   string
	panicMsg     string
	parent       *featureStep
	givens       []*givenStep
	whens        []*whenStep
	thens        []*thenStep
}

func (s *scenarioStep) init() {
	s.panicStack = ""
	s.lastRunFails = false
}

func (s *scenarioStep) test(t *testing.T) {
	s.init()
	t.Run(s.name, func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				s.panicStack = stack_(debug.Stack())
				s.panicMsg = fmt.Sprintf("%v", r)
				recover_(r, s.panicStack, t)
			}
		}()

		s.executed = true
		s.testName = t.Name()

		if s.before != nil {
			s.before()
		}

		for _, g := range s.givens {
			g.fn()
		}

		for _, w := range s.whens {
			w.fn()
		}

		for _, th := range s.thens {
			if th.Assert() {
				t.Fail()
				s.lastRunFails = true        // scenario fails
				s.parent.lastRunFails = true // feature fails
			}
		}

		s.reportTLog(t)

	})
}

func (s *scenarioStep) panicReport() *ScenarioReport {
	return &ScenarioReport{
		TestName:   s.testName,
		PanicMsg:   s.panicMsg,
		Trace:      s.panicStack,
		TestStatus: "PANIC",
	}
}

func (s *scenarioStep) report() *ScenarioReport {

	if s.panicStack != "" {
		return s.panicReport()
	}

	return s.testReport()
}

func (s *scenarioStep) skipped() bool {
	return !s.executed
}

func (s *scenarioStep) testReport() *ScenarioReport {

	tplSt := &ScenarioReport{
		Feature:  s.parent.name,
		Scenario: s.name,
		Steps:    []string{},
		Errors:   []string{},
	}

	tplSt.TestName = s.testName

	if s.lastRunFails {
		tplSt.TestStatus = "FAIL"
	} else {
		tplSt.TestStatus = "PASS"
	}

	for _, g := range s.givens {
		tplSt.Steps = append(tplSt.Steps, g.String())
	}

	for _, w := range s.whens {
		tplSt.Steps = append(tplSt.Steps, w.String())
	}

	for _, th := range s.thens {
		if th.HasError() {
			tplSt.Steps = append(tplSt.Steps, fmt.Sprintf("%s %s", th.String(), testFailMarker))
			for _, line := range th.Report() {
				tplSt.Errors = append(tplSt.Errors, line)
			}
		} else {
			tplSt.Steps = append(tplSt.Steps, fmt.Sprintf("%s %s", th.String(), testPassMarker))
		}
	}

	return tplSt
}

func (s *scenarioStep) reportTLog(t *testing.T) {
	decorator := &PlainTextDecorator{}
	t.Log(decorator._scenario(s.report()))
}

func (s *scenarioStep) Before(fn func()) ScenarioSteper {
	s.before = fn
	return s
}

func (s *scenarioStep) Given(name string, fn func()) Givener {
	return s.given(name, "Given", fn)
}

func (s *scenarioStep) given(name string, type_ string, fn func()) Givener {
	given := &givenStep{name: name, type_: type_, fn: fn, parent: s}
	s.givens = append(s.givens, given)
	return given
}
