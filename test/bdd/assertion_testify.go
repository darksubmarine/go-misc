/*
* CODE GENERATED AUTOMATICALLY WITH github.com/stretchr/testify/_codegen
* THIS FILE MUST NOT BE EDITED BY HAND
 */

package bdd

import (
	assert "github.com/stretchr/testify/assert"
	http "net/http"
	url "net/url"
	time "time"
)

// Condition uses a Comparison to assert a complex condition.
func (a *Assertions) Condition(comp assert.Comparison, msgAndArgs ...interface{}) {
	a._e(a.assert.Condition(comp, msgAndArgs...), "Condition", _errorString(msgAndArgs))
}

// Contains asserts that the specified string, list(array, slice...) or map contains the
// specified substring or element.
//
//    a.Contains("Hello World", "World")
//    a.Contains(["Hello", "World"], "World")
//    a.Contains({"Hello": "World"}, "Hello")
func (a *Assertions) Contains(s interface{}, contains interface{}, msgAndArgs ...interface{}) {
	a._e(a.assert.Contains(s, contains, msgAndArgs...), "Contains", _errorString(msgAndArgs))
}

// DirExists checks whether a directory exists in the given path. It also fails
// if the path is a file rather a directory or there is an error checking whether it exists.
func (a *Assertions) DirExists(path string, msgAndArgs ...interface{}) {
	a._e(a.assert.DirExists(path, msgAndArgs...), "DirExists", _errorString(msgAndArgs))
}

// ElementsMatch asserts that the specified listA(array, slice...) is equal to specified
// listB(array, slice...) ignoring the order of the elements. If there are duplicate elements,
// the number of appearances of each of them in both lists should match.
//
// a.ElementsMatch([1, 3, 2, 3], [1, 3, 3, 2])
func (a *Assertions) ElementsMatch(listA interface{}, listB interface{}, msgAndArgs ...interface{}) {
	a._e(a.assert.ElementsMatch(listA, listB, msgAndArgs...), "ElementsMatch", _errorString(msgAndArgs))
}

// Empty asserts that the specified object is empty.  I.e. nil, "", false, 0 or either
// a slice or a channel with len == 0.
//
//  a.Empty(obj)
func (a *Assertions) Empty(object interface{}, msgAndArgs ...interface{}) {
	a._e(a.assert.Empty(object, msgAndArgs...), "Empty", _errorString(msgAndArgs))
}

// Equal asserts that two objects are equal.
//
//    a.Equal(123, 123)
//
// Pointer variable equality is determined based on the equality of the
// referenced values (as opposed to the memory addresses). Function equality
// cannot be determined and will always fail.
func (a *Assertions) Equal(expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	a._e(a.assert.Equal(expected, actual, msgAndArgs...), "Equal", _errorString(msgAndArgs))
}

// EqualError asserts that a function returned an error (i.e. not `nil`)
// and that it is equal to the provided error.
//
//   actualObj, err := SomeFunction()
//   a.EqualError(err,  expectedErrorString)
func (a *Assertions) EqualError(theError error, errString string, msgAndArgs ...interface{}) {
	a._e(a.assert.EqualError(theError, errString, msgAndArgs...), "EqualError", _errorString(msgAndArgs))
}

// EqualValues asserts that two objects are equal or convertable to the same types
// and equal.
//
//    a.EqualValues(uint32(123), int32(123))
func (a *Assertions) EqualValues(expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	a._e(a.assert.EqualValues(expected, actual, msgAndArgs...), "EqualValues", _errorString(msgAndArgs))
}

// Error asserts that a function returned an error (i.e. not `nil`).
//
//   actualObj, err := SomeFunction()
//   if a.Error(err) {
// 	   assert.Equal(t, expectedError, err)
//   }
func (a *Assertions) Error(err error, msgAndArgs ...interface{}) {
	a._e(a.assert.Error(err, msgAndArgs...), "Error", _errorString(msgAndArgs))
}

// ErrorAs asserts that at least one of the errors in err's chain matches target, and if so, sets target to that error value.
// This is a wrapper for errors.As.
func (a *Assertions) ErrorAs(err error, target interface{}, msgAndArgs ...interface{}) {
	a._e(a.assert.ErrorAs(err, target, msgAndArgs...), "ErrorAs", _errorString(msgAndArgs))
}

// ErrorContains asserts that a function returned an error (i.e. not `nil`)
// and that the error contains the specified substring.
//
//   actualObj, err := SomeFunction()
//   a.ErrorContains(err,  expectedErrorSubString)
func (a *Assertions) ErrorContains(theError error, contains string, msgAndArgs ...interface{}) {
	a._e(a.assert.ErrorContains(theError, contains, msgAndArgs...), "ErrorContains", _errorString(msgAndArgs))
}

// ErrorIs asserts that at least one of the errors in err's chain matches target.
// This is a wrapper for errors.Is.
func (a *Assertions) ErrorIs(err error, target error, msgAndArgs ...interface{}) {
	a._e(a.assert.ErrorIs(err, target, msgAndArgs...), "ErrorIs", _errorString(msgAndArgs))
}

// Eventually asserts that given condition will be met in waitFor time,
// periodically checking target function each tick.
//
//    a.Eventually(func() bool { return true; }, time.Second, 10*time.Millisecond)
func (a *Assertions) Eventually(condition func() bool, waitFor time.Duration, tick time.Duration, msgAndArgs ...interface{}) {
	a._e(a.assert.Eventually(condition, waitFor, tick, msgAndArgs...), "Eventually", _errorString(msgAndArgs))
}

// Exactly asserts that two objects are equal in value and type.
//
//    a.Exactly(int32(123), int64(123))
func (a *Assertions) Exactly(expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	a._e(a.assert.Exactly(expected, actual, msgAndArgs...), "Exactly", _errorString(msgAndArgs))
}

// Fail reports a failure through
func (a *Assertions) Fail(failureMessage string, msgAndArgs ...interface{}) {
	a._e(a.assert.Fail(failureMessage, msgAndArgs...), "Fail", _errorString(msgAndArgs))
}

// FailNow fails test
func (a *Assertions) FailNow(failureMessage string, msgAndArgs ...interface{}) {
	a._e(a.assert.FailNow(failureMessage, msgAndArgs...), "FailNow", _errorString(msgAndArgs))
}

// False asserts that the specified value is false.
//
//    a.False(myBool)
func (a *Assertions) False(value bool, msgAndArgs ...interface{}) {
	a._e(a.assert.False(value, msgAndArgs...), "False", _errorString(msgAndArgs))
}

// FileExists checks whether a file exists in the given path. It also fails if
// the path points to a directory or there is an error when trying to check the file.
func (a *Assertions) FileExists(path string, msgAndArgs ...interface{}) {
	a._e(a.assert.FileExists(path, msgAndArgs...), "FileExists", _errorString(msgAndArgs))
}

// Greater asserts that the first element is greater than the second
//
//    a.Greater(2, 1)
//    a.Greater(float64(2), float64(1))
//    a.Greater("b", "a")
func (a *Assertions) Greater(e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) {
	a._e(a.assert.Greater(e1, e2, msgAndArgs...), "Greater", _errorString(msgAndArgs))
}

// GreaterOrEqual asserts that the first element is greater than or equal to the second
//
//    a.GreaterOrEqual(2, 1)
//    a.GreaterOrEqual(2, 2)
//    a.GreaterOrEqual("b", "a")
//    a.GreaterOrEqual("b", "b")
func (a *Assertions) GreaterOrEqual(e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) {
	a._e(a.assert.GreaterOrEqual(e1, e2, msgAndArgs...), "GreaterOrEqual", _errorString(msgAndArgs))
}

// HTTPBodyContains asserts that a specified handler returns a
// body that contains a string.
//
//  a.HTTPBodyContains(myHandler, "GET", "www.google.com", nil, "I'm Feeling Lucky")
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) HTTPBodyContains(handler http.HandlerFunc, method string, url string, values url.Values, str interface{}, msgAndArgs ...interface{}) {
	a._e(a.assert.HTTPBodyContains(handler, method, url, values, str, msgAndArgs...), "HTTPBodyContains", _errorString(msgAndArgs))
}

// HTTPBodyNotContains asserts that a specified handler returns a
// body that does not contain a string.
//
//  a.HTTPBodyNotContains(myHandler, "GET", "www.google.com", nil, "I'm Feeling Lucky")
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) HTTPBodyNotContains(handler http.HandlerFunc, method string, url string, values url.Values, str interface{}, msgAndArgs ...interface{}) {
	a._e(a.assert.HTTPBodyNotContains(handler, method, url, values, str, msgAndArgs...), "HTTPBodyNotContains", _errorString(msgAndArgs))
}

// HTTPError asserts that a specified handler returns an error status code.
//
//  a.HTTPError(myHandler, "POST", "/a/b/c", url.Values{"a": []string{"b", "c"}}
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) HTTPError(handler http.HandlerFunc, method string, url string, values url.Values, msgAndArgs ...interface{}) {
	a._e(a.assert.HTTPError(handler, method, url, values, msgAndArgs...), "HTTPError", _errorString(msgAndArgs))
}

// HTTPRedirect asserts that a specified handler returns a redirect status code.
//
//  a.HTTPRedirect(myHandler, "GET", "/a/b/c", url.Values{"a": []string{"b", "c"}}
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) HTTPRedirect(handler http.HandlerFunc, method string, url string, values url.Values, msgAndArgs ...interface{}) {
	a._e(a.assert.HTTPRedirect(handler, method, url, values, msgAndArgs...), "HTTPRedirect", _errorString(msgAndArgs))
}

// HTTPStatusCode asserts that a specified handler returns a specified status code.
//
//  a.HTTPStatusCode(myHandler, "GET", "/notImplemented", nil, 501)
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) HTTPStatusCode(handler http.HandlerFunc, method string, url string, values url.Values, statuscode int, msgAndArgs ...interface{}) {
	a._e(a.assert.HTTPStatusCode(handler, method, url, values, statuscode, msgAndArgs...), "HTTPStatusCode", _errorString(msgAndArgs))
}

// HTTPSuccess asserts that a specified handler returns a success status code.
//
//  a.HTTPSuccess(myHandler, "POST", "http://www.google.com", nil)
//
// Returns whether the assertion was successful (true) or not (false).
func (a *Assertions) HTTPSuccess(handler http.HandlerFunc, method string, url string, values url.Values, msgAndArgs ...interface{}) {
	a._e(a.assert.HTTPSuccess(handler, method, url, values, msgAndArgs...), "HTTPSuccess", _errorString(msgAndArgs))
}

// Implements asserts that an object is implemented by the specified interface.
//
//    a.Implements((*MyInterface)(nil), new(MyObject))
func (a *Assertions) Implements(interfaceObject interface{}, object interface{}, msgAndArgs ...interface{}) {
	a._e(a.assert.Implements(interfaceObject, object, msgAndArgs...), "Implements", _errorString(msgAndArgs))
}

// InDelta asserts that the two numerals are within delta of each other.
//
// 	 a.InDelta(math.Pi, 22/7.0, 0.01)
func (a *Assertions) InDelta(expected interface{}, actual interface{}, delta float64, msgAndArgs ...interface{}) {
	a._e(a.assert.InDelta(expected, actual, delta, msgAndArgs...), "InDelta", _errorString(msgAndArgs))
}

// InDeltaMapValues is the same as InDelta, but it compares all values between two maps. Both maps must have exactly the same keys.
func (a *Assertions) InDeltaMapValues(expected interface{}, actual interface{}, delta float64, msgAndArgs ...interface{}) {
	a._e(a.assert.InDeltaMapValues(expected, actual, delta, msgAndArgs...), "InDeltaMapValues", _errorString(msgAndArgs))
}

// InDeltaSlice is the same as InDelta, except it compares two slices.
func (a *Assertions) InDeltaSlice(expected interface{}, actual interface{}, delta float64, msgAndArgs ...interface{}) {
	a._e(a.assert.InDeltaSlice(expected, actual, delta, msgAndArgs...), "InDeltaSlice", _errorString(msgAndArgs))
}

// InEpsilon asserts that expected and actual have a relative error less than epsilon
func (a *Assertions) InEpsilon(expected interface{}, actual interface{}, epsilon float64, msgAndArgs ...interface{}) {
	a._e(a.assert.InEpsilon(expected, actual, epsilon, msgAndArgs...), "InEpsilon", _errorString(msgAndArgs))
}

// InEpsilonSlice is the same as InEpsilon, except it compares each value from two slices.
func (a *Assertions) InEpsilonSlice(expected interface{}, actual interface{}, epsilon float64, msgAndArgs ...interface{}) {
	a._e(a.assert.InEpsilonSlice(expected, actual, epsilon, msgAndArgs...), "InEpsilonSlice", _errorString(msgAndArgs))
}

// IsDecreasing asserts that the collection is decreasing
//
//    a.IsDecreasing([]int{2, 1, 0})
//    a.IsDecreasing([]float{2, 1})
//    a.IsDecreasing([]string{"b", "a"})
func (a *Assertions) IsDecreasing(object interface{}, msgAndArgs ...interface{}) {
	a._e(a.assert.IsDecreasing(object, msgAndArgs...), "IsDecreasing", _errorString(msgAndArgs))
}

// IsIncreasing asserts that the collection is increasing
//
//    a.IsIncreasing([]int{1, 2, 3})
//    a.IsIncreasing([]float{1, 2})
//    a.IsIncreasing([]string{"a", "b"})
func (a *Assertions) IsIncreasing(object interface{}, msgAndArgs ...interface{}) {
	a._e(a.assert.IsIncreasing(object, msgAndArgs...), "IsIncreasing", _errorString(msgAndArgs))
}

// IsNonDecreasing asserts that the collection is not decreasing
//
//    a.IsNonDecreasing([]int{1, 1, 2})
//    a.IsNonDecreasing([]float{1, 2})
//    a.IsNonDecreasing([]string{"a", "b"})
func (a *Assertions) IsNonDecreasing(object interface{}, msgAndArgs ...interface{}) {
	a._e(a.assert.IsNonDecreasing(object, msgAndArgs...), "IsNonDecreasing", _errorString(msgAndArgs))
}

// IsNonIncreasing asserts that the collection is not increasing
//
//    a.IsNonIncreasing([]int{2, 1, 1})
//    a.IsNonIncreasing([]float{2, 1})
//    a.IsNonIncreasing([]string{"b", "a"})
func (a *Assertions) IsNonIncreasing(object interface{}, msgAndArgs ...interface{}) {
	a._e(a.assert.IsNonIncreasing(object, msgAndArgs...), "IsNonIncreasing", _errorString(msgAndArgs))
}

// IsType asserts that the specified objects are of the same type.
func (a *Assertions) IsType(expectedType interface{}, object interface{}, msgAndArgs ...interface{}) {
	a._e(a.assert.IsType(expectedType, object, msgAndArgs...), "IsType", _errorString(msgAndArgs))
}

// JSONEq asserts that two JSON strings are equivalent.
//
//  a.JSONEq(`{"hello": "world", "foo": "bar"}`, `{"foo": "bar", "hello": "world"}`)
func (a *Assertions) JSONEq(expected string, actual string, msgAndArgs ...interface{}) {
	a._e(a.assert.JSONEq(expected, actual, msgAndArgs...), "JSONEq", _errorString(msgAndArgs))
}

// Len asserts that the specified object has specific length.
// Len also fails if the object has a type that len() not accept.
//
//    a.Len(mySlice, 3)
func (a *Assertions) Len(object interface{}, length int, msgAndArgs ...interface{}) {
	a._e(a.assert.Len(object, length, msgAndArgs...), "Len", _errorString(msgAndArgs))
}

// Less asserts that the first element is less than the second
//
//    a.Less(1, 2)
//    a.Less(float64(1), float64(2))
//    a.Less("a", "b")
func (a *Assertions) Less(e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) {
	a._e(a.assert.Less(e1, e2, msgAndArgs...), "Less", _errorString(msgAndArgs))
}

// LessOrEqual asserts that the first element is less than or equal to the second
//
//    a.LessOrEqual(1, 2)
//    a.LessOrEqual(2, 2)
//    a.LessOrEqual("a", "b")
//    a.LessOrEqual("b", "b")
func (a *Assertions) LessOrEqual(e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) {
	a._e(a.assert.LessOrEqual(e1, e2, msgAndArgs...), "LessOrEqual", _errorString(msgAndArgs))
}

// Negative asserts that the specified element is negative
//
//    a.Negative(-1)
//    a.Negative(-1.23)
func (a *Assertions) Negative(e interface{}, msgAndArgs ...interface{}) {
	a._e(a.assert.Negative(e, msgAndArgs...), "Negative", _errorString(msgAndArgs))
}

// Never asserts that the given condition doesn't satisfy in waitFor time,
// periodically checking the target function each tick.
//
//    a.Never(func() bool { return false; }, time.Second, 10*time.Millisecond)
func (a *Assertions) Never(condition func() bool, waitFor time.Duration, tick time.Duration, msgAndArgs ...interface{}) {
	a._e(a.assert.Never(condition, waitFor, tick, msgAndArgs...), "Never", _errorString(msgAndArgs))
}

// Nil asserts that the specified object is nil.
//
//    a.Nil(err)
func (a *Assertions) Nil(object interface{}, msgAndArgs ...interface{}) {
	a._e(a.assert.Nil(object, msgAndArgs...), "Nil", _errorString(msgAndArgs))
}

// NoDirExists checks whether a directory does not exist in the given path.
// It fails if the path points to an existing _directory_ only.
func (a *Assertions) NoDirExists(path string, msgAndArgs ...interface{}) {
	a._e(a.assert.NoDirExists(path, msgAndArgs...), "NoDirExists", _errorString(msgAndArgs))
}

// NoError asserts that a function returned no error (i.e. `nil`).
//
//   actualObj, err := SomeFunction()
//   if a.NoError(err) {
// 	   assert.Equal(t, expectedObj, actualObj)
//   }
func (a *Assertions) NoError(err error, msgAndArgs ...interface{}) {
	a._e(a.assert.NoError(err, msgAndArgs...), "NoError", _errorString(msgAndArgs))
}

// NoFileExists checks whether a file does not exist in a given path. It fails
// if the path points to an existing _file_ only.
func (a *Assertions) NoFileExists(path string, msgAndArgs ...interface{}) {
	a._e(a.assert.NoFileExists(path, msgAndArgs...), "NoFileExists", _errorString(msgAndArgs))
}

// NotContains asserts that the specified string, list(array, slice...) or map does NOT contain the
// specified substring or element.
//
//    a.NotContains("Hello World", "Earth")
//    a.NotContains(["Hello", "World"], "Earth")
//    a.NotContains({"Hello": "World"}, "Earth")
func (a *Assertions) NotContains(s interface{}, contains interface{}, msgAndArgs ...interface{}) {
	a._e(a.assert.NotContains(s, contains, msgAndArgs...), "NotContains", _errorString(msgAndArgs))
}

// NotEmpty asserts that the specified object is NOT empty.  I.e. not nil, "", false, 0 or either
// a slice or a channel with len == 0.
//
//  if a.NotEmpty(obj) {
//    assert.Equal(t, "two", obj[1])
//  }
func (a *Assertions) NotEmpty(object interface{}, msgAndArgs ...interface{}) {
	a._e(a.assert.NotEmpty(object, msgAndArgs...), "NotEmpty", _errorString(msgAndArgs))
}

// NotEqual asserts that the specified values are NOT equal.
//
//    a.NotEqual(obj1, obj2)
//
// Pointer variable equality is determined based on the equality of the
// referenced values (as opposed to the memory addresses).
func (a *Assertions) NotEqual(expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	a._e(a.assert.NotEqual(expected, actual, msgAndArgs...), "NotEqual", _errorString(msgAndArgs))
}

// NotEqualValues asserts that two objects are not equal even when converted to the same type
//
//    a.NotEqualValues(obj1, obj2)
func (a *Assertions) NotEqualValues(expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	a._e(a.assert.NotEqualValues(expected, actual, msgAndArgs...), "NotEqualValues", _errorString(msgAndArgs))
}

// NotErrorIs asserts that at none of the errors in err's chain matches target.
// This is a wrapper for errors.Is.
func (a *Assertions) NotErrorIs(err error, target error, msgAndArgs ...interface{}) {
	a._e(a.assert.NotErrorIs(err, target, msgAndArgs...), "NotErrorIs", _errorString(msgAndArgs))
}

// NotNil asserts that the specified object is not nil.
//
//    a.NotNil(err)
func (a *Assertions) NotNil(object interface{}, msgAndArgs ...interface{}) {
	a._e(a.assert.NotNil(object, msgAndArgs...), "NotNil", _errorString(msgAndArgs))
}

// NotPanics asserts that the code inside the specified PanicTestFunc does NOT panic.
//
//   a.NotPanics(func(){ RemainCalm() })
func (a *Assertions) NotPanics(f assert.PanicTestFunc, msgAndArgs ...interface{}) {
	a._e(a.assert.NotPanics(f, msgAndArgs...), "NotPanics", _errorString(msgAndArgs))
}

// NotRegexp asserts that a specified regexp does not match a string.
//
//  a.NotRegexp(regexp.MustCompile("starts"), "it's starting")
//  a.NotRegexp("^start", "it's not starting")
func (a *Assertions) NotRegexp(rx interface{}, str interface{}, msgAndArgs ...interface{}) {
	a._e(a.assert.NotRegexp(rx, str, msgAndArgs...), "NotRegexp", _errorString(msgAndArgs))
}

// NotSame asserts that two pointers do not reference the same object.
//
//    a.NotSame(ptr1, ptr2)
//
// Both arguments must be pointer variables. Pointer variable sameness is
// determined based on the equality of both type and value.
func (a *Assertions) NotSame(expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	a._e(a.assert.NotSame(expected, actual, msgAndArgs...), "NotSame", _errorString(msgAndArgs))
}

// NotSubset asserts that the specified list(array, slice...) contains not all
// elements given in the specified subset(array, slice...).
//
//    a.NotSubset([1, 3, 4], [1, 2], "But [1, 3, 4] does not contain [1, 2]")
func (a *Assertions) NotSubset(list interface{}, subset interface{}, msgAndArgs ...interface{}) {
	a._e(a.assert.NotSubset(list, subset, msgAndArgs...), "NotSubset", _errorString(msgAndArgs))
}

// NotZero asserts that i is not the zero value for its type.
func (a *Assertions) NotZero(i interface{}, msgAndArgs ...interface{}) {
	a._e(a.assert.NotZero(i, msgAndArgs...), "NotZero", _errorString(msgAndArgs))
}

// Panics asserts that the code inside the specified PanicTestFunc panics.
//
//   a.Panics(func(){ GoCrazy() })
func (a *Assertions) Panics(f assert.PanicTestFunc, msgAndArgs ...interface{}) {
	a._e(a.assert.Panics(f, msgAndArgs...), "Panics", _errorString(msgAndArgs))
}

// PanicsWithError asserts that the code inside the specified PanicTestFunc
// panics, and that the recovered panic value is an error that satisfies the
// EqualError comparison.
//
//   a.PanicsWithError("crazy error", func(){ GoCrazy() })
func (a *Assertions) PanicsWithError(errString string, f assert.PanicTestFunc, msgAndArgs ...interface{}) {
	a._e(a.assert.PanicsWithError(errString, f, msgAndArgs...), "PanicsWithError", _errorString(msgAndArgs))
}

// PanicsWithValue asserts that the code inside the specified PanicTestFunc panics, and that
// the recovered panic value equals the expected panic value.
//
//   a.PanicsWithValue("crazy error", func(){ GoCrazy() })
func (a *Assertions) PanicsWithValue(expected interface{}, f assert.PanicTestFunc, msgAndArgs ...interface{}) {
	a._e(a.assert.PanicsWithValue(expected, f, msgAndArgs...), "PanicsWithValue", _errorString(msgAndArgs))
}

// Positive asserts that the specified element is positive
//
//    a.Positive(1)
//    a.Positive(1.23)
func (a *Assertions) Positive(e interface{}, msgAndArgs ...interface{}) {
	a._e(a.assert.Positive(e, msgAndArgs...), "Positive", _errorString(msgAndArgs))
}

// Regexp asserts that a specified regexp matches a string.
//
//  a.Regexp(regexp.MustCompile("start"), "it's starting")
//  a.Regexp("start...$", "it's not starting")
func (a *Assertions) Regexp(rx interface{}, str interface{}, msgAndArgs ...interface{}) {
	a._e(a.assert.Regexp(rx, str, msgAndArgs...), "Regexp", _errorString(msgAndArgs))
}

// Same asserts that two pointers reference the same object.
//
//    a.Same(ptr1, ptr2)
//
// Both arguments must be pointer variables. Pointer variable sameness is
// determined based on the equality of both type and value.
func (a *Assertions) Same(expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	a._e(a.assert.Same(expected, actual, msgAndArgs...), "Same", _errorString(msgAndArgs))
}

// Subset asserts that the specified list(array, slice...) contains all
// elements given in the specified subset(array, slice...).
//
//    a.Subset([1, 2, 3], [1, 2], "But [1, 2, 3] does contain [1, 2]")
func (a *Assertions) Subset(list interface{}, subset interface{}, msgAndArgs ...interface{}) {
	a._e(a.assert.Subset(list, subset, msgAndArgs...), "Subset", _errorString(msgAndArgs))
}

// True asserts that the specified value is true.
//
//    a.True(myBool)
func (a *Assertions) True(value bool, msgAndArgs ...interface{}) {
	a._e(a.assert.True(value, msgAndArgs...), "True", _errorString(msgAndArgs))
}

// WithinDuration asserts that the two times are within duration delta of each other.
//
//   a.WithinDuration(time.Now(), time.Now(), 10*time.Second)
func (a *Assertions) WithinDuration(expected time.Time, actual time.Time, delta time.Duration, msgAndArgs ...interface{}) {
	a._e(a.assert.WithinDuration(expected, actual, delta, msgAndArgs...), "WithinDuration", _errorString(msgAndArgs))
}

// WithinRange asserts that a time is within a time range (inclusive).
//
//   a.WithinRange(time.Now(), time.Now().Add(-time.Second), time.Now().Add(time.Second))
func (a *Assertions) WithinRange(actual time.Time, start time.Time, end time.Time, msgAndArgs ...interface{}) {
	a._e(a.assert.WithinRange(actual, start, end, msgAndArgs...), "WithinRange", _errorString(msgAndArgs))
}

// YAMLEq asserts that two YAML strings are equivalent.
func (a *Assertions) YAMLEq(expected string, actual string, msgAndArgs ...interface{}) {
	a._e(a.assert.YAMLEq(expected, actual, msgAndArgs...), "YAMLEq", _errorString(msgAndArgs))
}

// Zero asserts that i is the zero value for its type.
func (a *Assertions) Zero(i interface{}, msgAndArgs ...interface{}) {
	a._e(a.assert.Zero(i, msgAndArgs...), "Zero", _errorString(msgAndArgs))
}
