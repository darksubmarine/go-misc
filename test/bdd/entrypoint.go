package bdd

import "testing"

// Suite creates a test suite that supports a list of Features
func Suite(name string, t *testing.T, opts ...Option) Suiter {
	mOpts := sliceOptionToMapOption(opts)
	return &suiteStep{name: name, t: t, opts: mOpts}
}

// Feature creates a test feature that contains a list of scenarios
func Feature(name string, t *testing.T, opts ...Option) Featurer {
	mOpts := sliceOptionToMapOption(opts)
	return newFeature(name, t, mOpts)
}

// Scenario creates a test Scenario to describe a test with Given, When and Then
func Scenario(name string, t *testing.T, opts ...Option) Scenery {
	feature := Feature("default feature", t, opts...)
	return feature.Scenario(name)
}

// Given creates a Given step
func Given(name string, t *testing.T, fn func(), opts ...Option) Givener {
	scenario := Scenario("default scenario", t, opts...)
	return scenario.Given(name, fn)
}

//--- Non exported functions

func sliceOptionToMapOption(opts []Option) map[int]interface{} {
	mOpts := make(map[int]interface{}, len(opts))
	if len(opts) > 0 {
		for _, opt := range opts {
			mOpts[opt.key] = opt.val
		}
	}
	return mOpts
}

func newFeature(name string, t *testing.T, mOpts map[int]interface{}) *featureStep {
	return &featureStep{name: name, t: t, opts: mOpts, scenarios: []*scenarioStep{}}
}

func newFeatureInSuite(name string, t *testing.T, mOpts map[int]interface{}, parent *suiteStep) *featureStep {
	feature := newFeature(name, t, mOpts)
	feature.parent = parent
	return feature
}
