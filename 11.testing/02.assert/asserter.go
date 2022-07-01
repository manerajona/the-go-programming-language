package asserter

import (
	"fmt"
	"github.com/stretchr/testify/assert"
)

// asserter is used to be able to retrieve the error reported by the called assertion
type asserter struct {
	err error
}

// Errorf is used by the called assertion to report an error
func (a *asserter) Errorf(format string, args ...interface{}) {
	a.err = fmt.Errorf(format, args...)
}

// ActualAssertion receives the assertion type, actual value and optionally the error message
type ActualAssertion func(t assert.TestingT, actual interface{}, msgAndArgs ...interface{}) bool

// AssertActual is a helper function to allow the step function to call
// assertion functions where you want to compare an actual value to a
// predefined state like nil, empty or true/false.
func AssertActual(a ActualAssertion, actual interface{}, msgAndArgs ...interface{}) error {
	var t asserter
	a(&t, actual, msgAndArgs...)
	return t.err
}

// ExpectedAndActualAssertion receives the assertion type, expected value, actual value and optionally the error message
type ExpectedAndActualAssertion func(t assert.TestingT, expected, actual interface{}, msgAndArgs ...interface{}) bool

// AssertExpectedAndActual is a helper function to allow the step function to call
// assertion functions where you want to compare an expected and an actual value.
func AssertExpectedAndActual(a ExpectedAndActualAssertion, expected, actual interface{}, msgAndArgs ...interface{}) error {
	var t asserter
	a(&t, expected, actual, msgAndArgs...)
	return t.err
}
