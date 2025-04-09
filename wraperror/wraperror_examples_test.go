package wraperror_test

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/senzing-garage/go-helpers/wraperror"
)

func ExampleErrorf() {
	err := errors.New("test error")
	newErr := wraperror.Errorf(err, "wrapped with error: %w", err)
	fmt.Println(newErr.Error())
	// Output: wrapped with error: test error
}
