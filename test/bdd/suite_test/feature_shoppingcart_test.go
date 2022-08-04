package suite_test

import "github.com/darksubmarine/versatile/test/bdd"

func featureShoppingCart(suite bdd.Suiter) {
	var shoppingCart map[string]bool
	var mustHave []string

	feature := suite.Feature("Shopping Cart check")

	feature.Scenario("All done").
		Given("I am out shopping",
			func() {
				shoppingCart = map[string]bool{}
			}).
		And("I have eggs",
			func() {
				shoppingCart["eggs"] = true
			}).
		And("I have milk",
			func() {
				shoppingCart["milk"] = true
				//panic("something")
			}).
		And("I have butter",
			func() {
				shoppingCart["butter"] = true
			}).
		When("I check my list",
			func() {
				mustHave = []string{"eggs", "milk", "butter", "none"}
			}).
		Then("I don't need anything",
			func(assertion *bdd.Assertions) {
				for _, item := range mustHave {
					assertion.True(shoppingCart[item])
				}
				assertion.Equal(shoppingCart, mustHave)
			})

	feature.Scenario("Missing juice").
		Given("I am out shopping",
			func() {
				shoppingCart = map[string]bool{}
			}).
		And("I have eggs, milk and butter",
			func() {
				shoppingCart["eggs"] = true
				shoppingCart["milk"] = true
				shoppingCart["butter"] = true
			}).
		When("I check my list",
			func() {
				mustHave = []string{"eggs", "milk", "butter", "juice"}
			}).
		Then("I forgot the juice",
			func(assertion *bdd.Assertions) {
				for _, item := range mustHave {
					assertion.True(shoppingCart[item], "expected item", item)
				}
			})
}
