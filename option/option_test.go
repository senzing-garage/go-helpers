package option

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ----------------------------------------------------------------------------
// Tests
// ----------------------------------------------------------------------------

func TestOsLookupEnvBool(test *testing.T) {
	assert.True(test, OsLookupEnvBool("NOT_AN_ENVIRONMENT_VARIABLE", true))
}

func TestOsLookupEnvInt(test *testing.T) {
	assert.Equal(test, 10, OsLookupEnvInt("NOT_AN_ENVIRONMENT_VARIABLE", 10))
}

func TestOsLookupEnvString(test *testing.T) {
	assert.Equal(test, "default", OsLookupEnvString("NOT_AN_ENVIRONMENT_VARIABLE", "default"))
}

func TestSetDefault(test *testing.T) {
	assert.Equal(test, "NOT a default", OptionDatabaseUrl.SetDefault("NOT a default").Default)
}

// ----------------------------------------------------------------------------
// Examples
// ----------------------------------------------------------------------------

func ExampleOsLookupEnvBool() {
	fmt.Println(OsLookupEnvBool("NOT_AN_ENVIRONMENT_VARIABLE", true))
	// Output: true
}

func ExampleOsLookupEnvInt() {
	fmt.Println(OsLookupEnvInt("NOT_AN_ENVIRONMENT_VARIABLE", 10))
	// Output: 10
}

func ExampleOsLookupEnvString() {
	fmt.Println(OsLookupEnvString("NOT_AN_ENVIRONMENT_VARIABLE", "default"))
	// Output: default
}

func ExampleContextVariable_SetDefault() {
	fmt.Println(OptionDatabaseUrl.SetDefault("NOT a default").Default)
	// Output: NOT a default
}
