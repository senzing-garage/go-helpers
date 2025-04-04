package getenv_test

import (
	"fmt"

	"github.com/senzing-garage/go-helpers/getenv"
)

func ExampleGetEnv() {
	testEnvVar := getenv.GetEnv("MY_TEST_ENV_VAR", "DEFAULT_VALUE")
	fmt.Println(testEnvVar)
	// Output: DEFAULT_VALUE
}
