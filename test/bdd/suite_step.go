package bdd

import (
	"fmt"
	"testing"
)

type Suiter interface {
	Setup(fn func()) *suiteStep
	TierDown(fn func()) *suiteStep
	Feature(name string) *featureStep
	Run()

	Log(args ...any)
	Error(args ...any)
	Errorf(format string, args ...any)
	Fail()
	FailNow()
	Setenv(key, value string)
}

type suiteStep struct {
	name       string
	opts       map[int]interface{}
	t          *testing.T
	features   []*featureStep
	setupFn    func()
	tierDownFn func()
}

func (s *suiteStep) Setup(fn func()) *suiteStep {
	s.setupFn = fn
	return s
}
func (s *suiteStep) TierDown(fn func()) *suiteStep {
	s.tierDownFn = fn
	return s
}

func (s *suiteStep) Run() {

	if s.setupFn != nil {
		s.setupFn()
	}

	for _, feature := range s.features {
		feature.Test()
	}

	s.doReport()

	if s.tierDownFn != nil {
		s.tierDownFn()
	}
}

func (s *suiteStep) doReport() {

	if rep := OptReport(s.opts); rep != nil {

		var decorator = OptDecorator(s.opts)
		if decorator == nil {
			decorator = &PlainTextDecorator{}
		}

		rep.SetTitle(decorator._title(fmt.Sprintf("Summary suite report: %s", s.t.Name())))
		for _, f := range s.features {
			if f.skipped() {
				continue
			}

			fReport := &FeatureReport{Name: f.name, ScenarioReports: make([]*ScenarioReport, len(f.scenarios))}
			for i, s := range f.scenarios {
				if s.skipped() {
					continue
				}
				fReport.ScenarioReports[i] = s.report()
			}
			rep.Print(decorator._feature(fReport))
		}
		rep.Generate()
	}
}

func (s *suiteStep) Feature(name string) *featureStep {
	feature := newFeatureInSuite(name, s.t, s.opts, s)

	s.features = append(s.features, feature)
	return feature
}

func (s *suiteStep) Log(args ...any) {
	s.t.Log(args...)
}

func (s *suiteStep) Error(args ...any) {
	s.t.Error(args...)
}

func (s *suiteStep) Errorf(format string, args ...any) {
	s.t.Errorf(format, args...)
}

func (s *suiteStep) Fail() {
	s.t.Fail()
}

func (s *suiteStep) FailNow() {
	s.t.FailNow()
}

func (s *suiteStep) Setenv(key, value string) {
	s.t.Setenv(key, value)
}
