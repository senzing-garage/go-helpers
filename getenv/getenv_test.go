package getenv_test

import (
	"os"
	"testing"

	"github.com/senzing-garage/go-helpers/getenv"
	"github.com/stretchr/testify/require"
)

// ----------------------------------------------------------------------------
// Interface methods - test
// ----------------------------------------------------------------------------

func TestHelpers_GetEnv(test *testing.T) {
	expected := "EXPECTED_VALUE"
	os.Setenv("TEST_ENV_VAR", expected)
	actual := getenv.GetEnv("TEST_ENV_VAR", "DEFAULT_VALUE")
	require.Equal(test, expected, actual)
}

func TestHelpers_GetEnv_default(test *testing.T) {
	expected := "DEFAULT_VALUE"
	actual := getenv.GetEnv("NO_ENV_VAR", expected)
	require.Equal(test, expected, actual)
}
