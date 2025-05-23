package env_test

import (
	"testing"

	"github.com/senzing-garage/go-helpers/env"
	"github.com/stretchr/testify/require"
)

// ----------------------------------------------------------------------------
// Interface methods - test
// ----------------------------------------------------------------------------

func TestHelpers_GetEnv(test *testing.T) {
	expected := "EXPECTED_VALUE"
	test.Setenv("TEST_ENV_VAR", expected)

	actual := env.GetEnv("TEST_ENV_VAR", "DEFAULT_VALUE")
	require.Equal(test, expected, actual)
}

func TestHelpers_GetEnv_default(test *testing.T) {
	test.Parallel()

	expected := "DEFAULT_VALUE"
	actual := env.GetEnv("NO_ENV_VAR", expected)
	require.Equal(test, expected, actual)
}
