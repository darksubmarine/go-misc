package ds_

import (
	"github.com/darksubmarine/versatile/test/bdd"
	"testing"
)

// Suite creates a test suite that supports a list of Features
func Suite(name string, t *testing.T, opts ...bdd.Option) bdd.Suiter {
	return bdd.Suite(name, t, opts...)
}

// Feature creates a test feature that contains a list of scenarios
func Feature(name string, t *testing.T, opts ...bdd.Option) bdd.Featurer {
	return bdd.Feature(name, t, opts...)
}

// Scenario creates a test Scenario to describe a test with Given, When and Then
func Scenario(name string, t *testing.T, opts ...bdd.Option) bdd.Scenery {
	return bdd.Scenario(name, t, opts...)
}

// Given creates a Given step
func Given(name string, t *testing.T, fn func(), opts ...bdd.Option) bdd.Givener {
	return bdd.Given(name, t, fn, opts...)
}
