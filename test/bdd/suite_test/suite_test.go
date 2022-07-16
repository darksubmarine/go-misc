package suite_test

import (
	"github.com/darksubmarine/versatile/report"
	"github.com/darksubmarine/versatile/test/bdd"
	"testing"
)

/*
To run only one specific test

Suite: "testing suite" (is the Go root test)
	go test ./test/bdd/suite_test/... -run TestBDD_Suite -v

Feature: "Shopping Cart check"
	go test ./test/bdd/suite_test/... -run TestBDD_Suite/Shopping_Cart_check -v

Scenario: "Missing juice"
	go test ./test/bdd/suite_test/... -run TestBDD_Suite/Shopping_Cart_check/Missing_juice -v
*/

func TestBDD_Suite(t *testing.T) {
	mySuite := createTestingSuite(t)

	featureUserValidation(mySuite)
	featureShoppingCart(mySuite)

	mySuite.Run()
}

func createTestingSuite(t *testing.T) bdd.Suiter {

	opts := []bdd.Option{
		// Markdown
		//bdd.WithReport(report.NewFile(fmt.Sprintf("%s.md", t.Name()))), bdd.WithDecorator(bdd.NewMarkdownDecorator()),

		// HTML
		//bdd.WithReport(report.NewHTMLFile(t.Name())), bdd.WithDecorator(bdd.NewHTMLDecorator()),

		// Stdout
		bdd.WithReport(report.NewStdOut()), bdd.WithDecorator(bdd.NewPlainTextDecorator()),
	}

	return bdd.Suite("testing suite", t, opts...).
		Setup(
			func() {
				t.Log("Suite Setup")
			}).
		TierDown(
			func() {
				t.Log("Suite TIERDOWN!")
			})
}
