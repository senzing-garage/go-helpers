package wraperror_test

import (
	"fmt"
)

// func ExampleErrorf() {
// 	err := errors.New("test error")
// 	newErr := wraperror.Errorf(err, "wrapped with error: %w", err)
// 	fmt.Println(newErr.Error())
// 	// Output: wrapped with error: test error
// }

func ExampleErrorf_nested() {
	err := function1()
	fmt.Println(err.Error())
	// Output: {"function": "wraperror_test.function1", "text": "function2 with dog and cat", "error": {"function": "wraperror_test.function2", "text": "function3", "error": "testError"}}

}
