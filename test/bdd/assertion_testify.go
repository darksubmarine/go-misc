package bdd

import "fmt"

// Testify assertions wrapper: github.com/stretchr/testify/assert

// False asserts that the specified value is false.
//
//    a.False(myBool)
func (a *Assertions) False(value bool) {
	a._e(a.assert.False(value), "False", "expected FALSE but got TRUE")
}

func (a *Assertions) Fail(failureMessage string) {
	a._e(a.assert.Fail(failureMessage), "Fail", fmt.Sprintf("%v", failureMessage))
}

func (a *Assertions) FailNow(failureMessage string) {
	a._e(a.assert.FailNow(failureMessage), "FailNow", fmt.Sprintf("%v", failureMessage))
}

func (a *Assertions) Equal(expected, actual interface{}) {
	a._e(a.assert.Equal(expected, actual), "Equal", fmt.Sprintf("expected %v but got %v", expected, actual))
}

func (a *Assertions) NotEqual(expected, actual interface{}) {
	a._e(a.assert.NotEqual(expected, actual), "NotEqual", fmt.Sprintf("items should not be equal: expected %v but got %v", expected, actual))
}

func (a *Assertions) True(actual bool) {
	a._e(a.assert.True(actual), "True", "expected TRUE but got FALSE")
}

func (a *Assertions) Nil(object interface{}) {
	a._e(a.assert.Nil(object), "Nil", fmt.Sprintf("expected NIL but got %T", object))
}

func (a *Assertions) NotNil(object interface{}, msg string) {
	a._e(a.assert.NotNil(object), "NotNil", fmt.Sprintf("Actual is a NIL pointer: %s", msg))
}

func (a *Assertions) Greater(e1 interface{}, e2 interface{}) {
	a._e(a.assert.Greater(e1, e2), "Greater", fmt.Sprintf("%v must be greater than %v", e1, e2))
}
