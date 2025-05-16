package wraperror_test

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/senzing-garage/go-helpers/wraperror"
)

func ExampleErrorf() {
	err := errors.New("test error")
	newErr := wraperror.Errorf(err, "wrap an error")
	fmt.Println(newErr.Error())
	// Output: {"function": "wraperror_test.ExampleErrorf", "text": "wrap an error", "error": "test error"}
}

func ExampleErrorf_nested() {
	err := function1()
	fmt.Println(err.Error())
	// Output: {"function": "wraperror_test.function1", "text": "function2 with dog and cat", "error": {"function": "wraperror_test.function2", "text": "function3", "error": "testError"}}

}
