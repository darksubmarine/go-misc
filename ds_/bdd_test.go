package ds__test

import (
	"github.com/darksubmarine/versatile/ds_"
	"github.com/darksubmarine/versatile/report"
	"github.com/darksubmarine/versatile/test/bdd"
	"testing"
)

func TestBDD(t *testing.T) {

	opts := []bdd.Option{
		bdd.WithDecorator(bdd.NewPlainTextDecorator()),
		bdd.WithReport(report.NewStdOut()),
	}

	suite := ds_.Suite("some test suite", t, opts...)
	suite.Feature("some Feature A").Scenario("ScenarioA").Given("some Given", func() {
		t.Log("given>>>>")
	}).When("some when", func() {
		t.Log("when<<<<<")
	}).Then("Some then", func(assertion *bdd.Assertions) {
		assertion.True(true)
	})

	suite.Run()
}
