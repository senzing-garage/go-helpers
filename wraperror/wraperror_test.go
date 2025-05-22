package wraperror_test

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/senzing-garage/go-helpers/wraperror"
	"github.com/stretchr/testify/require"
)

var errForTest = errors.New("testError")

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

func TestHelpers_Error(test *testing.T) {
	test.Parallel()

	var err error
	newError := wraperror.Error(err)
	require.NoError(test, newError)
}

func TestHelpers_Error_isError(test *testing.T) {
	test.Parallel()

	expected := `{"function": "wraperror_test.TestHelpers_Error_isError", "error": "new error"}`
	err := errors.New("new error")
	actual := wraperror.Error(err)
	require.Error(test, actual)
	require.JSONEq(test, expected, actual.Error())
}

func TestHelpers_Error_isError_withJSON(test *testing.T) {
	test.Parallel()

	expected := `{"function": "wraperror_test.TestHelpers_Error_isError_withJSON", "error": {"function": "wraperror_test.TestHelpers_Error_isError", "error": "new error"}}`
	err := errors.New(`{"function": "wraperror_test.TestHelpers_Error_isError", "error": "new error"}`)
	actual := wraperror.Error(err)
	require.Error(test, actual)
	require.JSONEq(test, expected, actual.Error())
}

func TestHelpers_Error_isError_withNesting(test *testing.T) {
	test.Parallel()

	expected := `{"function": "wraperror_test.TestHelpers_Error_isError_withNesting", "error": {"function": "wraperror_test.function1", "text": "result from function2 with dog and cat", "error": {"function": "wraperror_test.function2", "text": "result from function3", "error": "testError"}}}`
	err := function1()
	actual := wraperror.Error(err)
	require.Error(test, actual)
	require.JSONEq(test, expected, actual.Error())
}

func TestHelpers_Errorf(test *testing.T) {
	test.Parallel()

	var err error
	newError := wraperror.Errorf(err, "not an error")
	require.NoError(test, newError)
}

func TestHelpers_Errorf_isError(test *testing.T) {
	test.Parallel()

	expected := `{"error":"new error", "function":"wraperror_test.TestHelpers_Errorf_isError", "text":"is an error"}`
	err := errors.New("new error")
	actual := wraperror.Errorf(err, "is an error")
	require.Error(test, actual)
	require.JSONEq(test, expected, actual.Error())
}

func TestHelpers_Errorf_isError_withJSON(test *testing.T) {
	test.Parallel()

	expected := `{"function": "wraperror_test.TestHelpers_Errorf_isError_withJSON", "text": "with JSON", "error": {"function": "wraperror_test.TestHelpers_Errorf_isError", "error": "new error"}}`
	err := errors.New(`{"function": "wraperror_test.TestHelpers_Errorf_isError", "error": "new error"}`)
	actual := wraperror.Errorf(err, "with JSON")
	require.Error(test, actual)
	require.JSONEq(test, expected, actual.Error())
}

func TestHelpers_Errorf_isError_withNesting(test *testing.T) {
	test.Parallel()

	expected := `{"function": "wraperror_test.TestHelpers_Errorf_isError_withNesting", "text": "with nesting", "error": {"function": "wraperror_test.function1", "text": "result from function2 with dog and cat", "error": {"function": "wraperror_test.function2", "text": "result from function3", "error": "testError"}}}`
	err := function1()
	actual := wraperror.Errorf(err, "with nesting")
	require.Error(test, actual)
	require.JSONEq(test, expected, actual.Error())
}

func TestHelpers_Errorf_isError_withMessage(test *testing.T) {
	test.Parallel()

	expected := `{"function":"wraperror_test.TestHelpers_Errorf_isError_withMessage", "text":"error", "error":{"id":"SZSDK60014001","reason":"SENZ3121|JSON Parsing Failure [code=12,offset=15]"}}`
	err := errors.New(`{"id":"SZSDK60014001","reason":"SENZ3121|JSON Parsing Failure [code=12,offset=15]"}`)
	actual := wraperror.Errorf(err, "error")
	require.Error(test, actual)
	require.JSONEq(test, expected, actual.Error())
}

func TestHelpers_Errorf_isError_withMessage_Double(test *testing.T) {
	test.Parallel()

	expected := `{"function":"wraperror_test.TestHelpers_Errorf_isError_withMessage_Double","text":"error 2","error":{"function":"wraperror_test.TestHelpers_Errorf_isError_withMessage_Double","text":"error 1","error":{"id":"SZSDK60014001","reason":"SENZ3121|JSON Parsing Failure [code=12,offset=15]"}}}`
	err := errors.New(`{"id":"SZSDK60014001","reason":"SENZ3121|JSON Parsing Failure [code=12,offset=15]"}`)
	err1 := wraperror.Errorf(err, "error 1")
	actual := wraperror.Errorf(err1, "error 2")
	require.Error(test, actual)
	require.JSONEq(test, expected, actual.Error())
}

func TestHelpers_Errorf_isError_withMessage_Triple(test *testing.T) {
	test.Parallel()

	expected := `{"function":"wraperror_test.TestHelpers_Errorf_isError_withMessage_Triple","text":"error 3","error":{"function":"wraperror_test.TestHelpers_Errorf_isError_withMessage_Triple","text":"error 2","error":{"function":"wraperror_test.TestHelpers_Errorf_isError_withMessage_Triple","text":"error 1","error":{"id":"SZSDK60014001","reason":"SENZ3121|JSON Parsing Failure [code=12,offset=15]"}}}}`
	err := errors.New(`{"id":"SZSDK60014001","reason":"SENZ3121|JSON Parsing Failure [code=12,offset=15]"}`)
	err1 := wraperror.Errorf(err, "error 1")
	err2 := wraperror.Errorf(err1, "error 2")
	actual := wraperror.Errorf(err2, "error 3")
	require.Error(test, actual)
	require.JSONEq(test, expected, actual.Error())
}

func TestHelpers_Errorf_isError_withVariables(test *testing.T) {
	test.Parallel()

	expected := `{"error":"new error", "function":"wraperror_test.TestHelpers_Errorf_isError_withVariables", "text":"is an error with aString and 99"}`
	aString := "aString"
	aNumber := 99
	err := errors.New("new error")
	actual := wraperror.Errorf(err, "is an error with %s and %d", aString, aNumber)
	require.Error(test, actual)
	require.JSONEq(test, expected, actual.Error())
}

// ----------------------------------------------------------------------------
// Private functions
// ----------------------------------------------------------------------------

func function1() error {
	err := function2()

	return wraperror.Errorf(err, "result from function2 with %s and %s", "dog", "cat")
}

func function2() error {
	err := function3()

	return wraperror.Errorf(err, "result from function3")
}

func function3() error {
	return errForTest
}
