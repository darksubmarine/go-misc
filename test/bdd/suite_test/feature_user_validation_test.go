package suite_test

import "github.com/darksubmarine/versatile/test/bdd"

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

func featureUserValidation(suite bdd.Suiter) {
	var token string
	var profile *UserProfile

	suite.Feature("User validation").
		Scenario("Request profile").
		Given("A valid user token", func() {
			token = VALID_TOKEN
			suite.Log("Some log from suite")
		}).
		When("The user request his profile info", func() {
			profile = fetchUserInfo(token)
		}).
		Then("the profile got is valid", func(assertion *bdd.Assertions) {
			assertion.Nil(profile, "user profile is nil")
			assertion.Equal(USER_INFO.Email, profile.Email)
			//assertion.NotEqual(USER_INFO.Plan, profile.Plan)
		})
}
