package wraperror_test

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/senzing-garage/go-helpers/wraperror"
	"github.com/stretchr/testify/require"
)

var testError = errors.New("testError")

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

func TestHelpers_Errorf(test *testing.T) {
	test.Parallel()
	var err error
	newError := wraperror.Errorf(err, "not an error")
	require.NoError(test, newError)
}

func TestHelpers_Errorf_isError(test *testing.T) {
	test.Parallel()
	expected := `{"error":"new error", "function":"wraperror_test.TestHelpers_Errorf_isError", "text":"is an error"}`
	err := errors.New("new error") //nolint:err113
	actual := wraperror.Errorf(err, "is an error")
	require.Error(test, actual)
	require.JSONEq(test, actual.Error(), expected)
}

func TestHelpers_Errorf_isError_withVariables(test *testing.T) {
	test.Parallel()
	expected := `{"error":"new error", "function":"wraperror_test.TestHelpers_Errorf_isError_withVariables", "text":"is an error with aString and 99"}`
	aString := "aString"
	aNumber := 99
	err := errors.New("new error") //nolint:err113
	actual := wraperror.Errorf(err, "is an error with %s and %d", aString, aNumber)
	require.Error(test, actual)
	require.JSONEq(test, actual.Error(), expected)
}

func TestHelpers_Errorf_isError_withNesting(test *testing.T) {
	test.Parallel()
	expected := `{"function": "wraperror_test.function1", "text": "function2 with dog and cat", "error": {"function": "wraperror_test.function2", "text": "function3", "error": "testError"}}`
	actual := function1()
	require.Error(test, actual)
	require.JSONEq(test, actual.Error(), expected)
}

// ----------------------------------------------------------------------------
// Private functions
// ----------------------------------------------------------------------------

func function1() error {
	err := function2()
	return wraperror.Errorf(err, "function2 with %s and %s", "dog", "cat")
}

func function2() error {
	err := function3()
	return wraperror.Errorf(err, "function3")
}

func function3() error {
	return testError
}
