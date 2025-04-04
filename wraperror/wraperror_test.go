package wraperror_test

import (
	"errors"
	"testing"

	"github.com/senzing-garage/go-helpers/wraperror"
	"github.com/stretchr/testify/require"
)

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

func TestHelpers_Errorf(test *testing.T) {
	test.Parallel()

	var err error
	newError := wraperror.Errorf(err, "not an error: %w", err)
	require.NoError(test, newError)
}

func TestHelpers_Errorf_isError(test *testing.T) {
	test.Parallel()

	err := errors.New("new error") //nolint:err113
	newError := wraperror.Errorf(err, "is an error: %w", err)
	require.Error(test, newError)
}
