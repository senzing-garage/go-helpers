package env_test

import (
	"fmt"

	"github.com/senzing-garage/go-helpers/env"
)

func ExampleGetEnv() {
	testEnvVar := env.GetEnv("MY_TEST_ENV_VAR", "DEFAULT_VALUE")
	fmt.Println(testEnvVar)
	// Output: DEFAULT_VALUE
}
