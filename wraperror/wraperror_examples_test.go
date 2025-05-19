package wraperror_test

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/senzing-garage/go-helpers/wraperror"
)

func ExampleError() {
	err := errors.New("test error")
	newErr := wraperror.Error(err)
	fmt.Println(newErr.Error())
	// Output: {"function": "wraperror_test.ExampleError", "error": "test error"}
}

func ExampleErrorf() {
	err := errors.New("test error")
	newErr := wraperror.Errorf(err, "wrap an error")
	fmt.Println(newErr.Error())
	// Output: {"function": "wraperror_test.ExampleErrorf", "text": "wrap an error", "error": "test error"}
}

func ExampleErrorf_noMessage() {
	err := errors.New("test error")
	newErr := wraperror.Errorf(err, wraperror.NoMessage)
	fmt.Println(newErr.Error())
	// Output: {"function": "wraperror_test.ExampleErrorf_noMessage", "error": "test error"}
}

func ExampleErrorf_withVariables() {
	err := errors.New("test error")
	newErr := wraperror.Errorf(err, "wrap an error with %d %s", 1, "message")
	fmt.Println(newErr.Error())
	// Output: {"function": "wraperror_test.ExampleErrorf_withVariables", "text": "wrap an error with 1 message", "error": "test error"}
}
