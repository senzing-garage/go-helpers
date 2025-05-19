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

func ExampleErrorf_nested() {
	err := function1()
	fmt.Println(err.Error())
	// Output: {"function": "wraperror_test.function1", "text": "result from function2 with dog and cat", "error": {"function": "wraperror_test.function2", "text": "result from function3", "error": "testError"}}
}

func ExampleErrorf_noMessage() {
	err := errors.New("test error")
	newErr := wraperror.Errorf(err, wraperror.NoMessage)
	fmt.Println(newErr.Error())
	// Output: {"function": "wraperror_test.ExampleErrorf_noMessage", "error": "test error"}
}
