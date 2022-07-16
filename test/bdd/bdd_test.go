package bdd_test

import (
	"fmt"
	"github.com/darksubmarine/versatile/report"
	"github.com/darksubmarine/versatile/test/bdd"
	"testing"
)

const VALID_TOKEN = "qwerty-123456"

type UserProfile struct {
	Name  string
	Email string
	Plan  string
}

var USER_INFO = &UserProfile{Name: "some-name", Email: "some@email.com", Plan: "platinum"}

func fetchUserInfo(tkn string) *UserProfile {
	if tkn == VALID_TOKEN {
		return USER_INFO
	}

	return nil // will produce a panic! on tests (runtime error: invalid memory address or nil pointer dereference)
}

func TestBDD(t *testing.T) {
	var token string
	var profile *UserProfile

	opts := []bdd.Option{
		bdd.WithDecorator(bdd.NewPlainTextDecorator()),
		bdd.WithReport(report.NewStdOut()),
	}

	bdd.Feature("User validation", t, opts...).
		Scenario("Request profile").
		Given("A valid user token", func() {
			token = VALID_TOKEN
		}).
		When("The user request his profile info", func() {
			profile = fetchUserInfo(token)
		}).
		Then("the profile got is valid", func(assertion *bdd.Assertions) {
			assertion.NotNil(profile, "user profile is nil")
			assertion.Equal(USER_INFO.Email, profile.Email)
			assertion.NotEqual(USER_INFO.Plan, profile.Plan)
		}).
		Test()
}

func TestFeature_(t *testing.T) {

	opts := []bdd.Option{
		bdd.WithDecorator(bdd.NewPlainTextDecorator()),
		bdd.WithReport(report.NewStdOut()),
	}

	featureA := bdd.Feature("some feature", t, opts...).
		Before(func() {
			fmt.Println("*** before feature")
		})

	featureA.Scenario("scenarioA").
		Before(func() {
			fmt.Println("+++ before scenario A")
		}).
		Given("some given step",
			func() {

			}).
		And("some AND step", func() {

		}).
		When("when something happens", func() {

		}).
		Then("check some values", func(assertion *bdd.Assertions) {
			assertion.True(false)
		})

	featureA.Scenario("scenarioB").
		Before(func() {
			fmt.Println("--- before scenario B")
		}).
		Given("some given step", func() {}).
		And("some AND step", func() {}).
		When("when something happens", func() {}).
		Then("check some values", func(assertion *bdd.Assertions) {})

	featureA.Test()
}

func TestFeatureWithScenariosNoBefores(t *testing.T) {

	var shoppingCart map[string]bool
	var mustHave []string

	opts := []bdd.Option{
		bdd.WithDecorator(bdd.NewPlainTextDecorator()),
		bdd.WithReport(report.NewStdOut()),
	}

	featureShoppingCartCheck := bdd.Feature("Shopping Cart check", t, opts...)

	featureShoppingCartCheck.
		Scenario("All done").
		Given("I am out shopping", func() {
			shoppingCart = map[string]bool{}
		}).
		And("I have eggs", func() {
			shoppingCart["eggs"] = true
		}).
		And("I have milk", func() {
			shoppingCart["milk"] = true
			//panic("something")
		}).
		And("I have butter", func() {
			shoppingCart["butter"] = true
		}).When("I check my list", func() {
		mustHave = []string{"eggs", "milk", "butter", "none"}
	}).
		Then("I don't need anything", func(assertion *bdd.Assertions) {
			for _, item := range mustHave {
				assertion.True(shoppingCart[item])
			}
		})

	featureShoppingCartCheck.
		Scenario("Missing juice").
		Given("I am out shopping", func() {
			shoppingCart = map[string]bool{}
		}).
		And("I have eggs, milk and butter", func() {
			shoppingCart["eggs"] = true
			shoppingCart["milk"] = true
			shoppingCart["butter"] = true
		}).
		When("I check my list", func() {
			mustHave = []string{"eggs", "milk", "butter"}
		}).
		Then("I forgot the juice", func(assertion *bdd.Assertions) {
			for _, item := range mustHave {

				assertion.NotEqual(shoppingCart[item], "juice")
			}
		})

	featureShoppingCartCheck.Test()
}
