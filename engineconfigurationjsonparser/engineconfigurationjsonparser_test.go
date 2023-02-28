package engineconfigurationjsonparser

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	engineConfigurationJsonParserSingleton *EngineConfigurationJsonParserImpl
)

// ----------------------------------------------------------------------------
// Internal functions
// ----------------------------------------------------------------------------

func getTestObject(ctx context.Context, test *testing.T) *EngineConfigurationJsonParserImpl {
	return getParser(ctx)
}

func getParser(ctx context.Context) *EngineConfigurationJsonParserImpl {
	if engineConfigurationJsonParserSingleton == nil {
		engineConfigurationJsonParserSingleton = &EngineConfigurationJsonParserImpl{
			EnableConfigurationJson: `
			{
				"PIPELINE": {
					"CONFIGPATH": "/etc/opt/senzing",
					"LICENSESTRINGBASE64": "${SENZING_LICENSE_BASE64_ENCODED}",
					"RESOURCEPATH": "/opt/senzing/g2/resources",
					"SUPPORTPATH": "/opt/senzing/data"
				},
				"SQL": {
					"CONNECTION": "postgresql://username:password@db.example.com:5432:G2"
				}
			}
			`,
		}
	}
	return engineConfigurationJsonParserSingleton
}

func testError(test *testing.T, err error) {
	if err != nil {
		assert.FailNow(test, err.Error())
	}
}

// ----------------------------------------------------------------------------
// Test interface functions
// ----------------------------------------------------------------------------

func TestEngineConfigurationJsonParserImpl_GetConfigPath(test *testing.T) {
	ctx := context.TODO()
	parser := getTestObject(ctx, test)
	actual, err := parser.GetConfigPath(ctx)
	testError(test, err)
	assert.Equal(test, "/etc/opt/senzing", actual)
}

// ----------------------------------------------------------------------------
// Examples for godoc documentation
// ----------------------------------------------------------------------------

func ExampleEngineConfigurationJsonParserImpl_GetConfigPath() {
	// For more information, visit https://github.com/Senzing/go-common/blob/main/engineconfigurationjsonparser/engineconfigurationjsonparser_test.go
	ctx := context.TODO()
	parser := getParser(ctx)
	configPath, err := parser.GetConfigPath(ctx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(configPath)
	// Output: /etc/opt/senzing
}

func ExampleEngineConfigurationJsonParserImpl_GetDatabaseUrls() {
	// For more information, visit https://github.com/Senzing/go-common/blob/main/engineconfigurationjsonparser/engineconfigurationjsonparser_test.go
	ctx := context.TODO()
	parser := getParser(ctx)
	configPath, err := parser.GetDatabaseUrls(ctx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(configPath)
	// Output: [postgresql://username:password@db.example.com:5432:G2]
}

func ExampleEngineConfigurationJsonParserImpl_GetResourcePath() {
	// For more information, visit https://github.com/Senzing/go-common/blob/main/engineconfigurationjsonparser/engineconfigurationjsonparser_test.go
	ctx := context.TODO()
	parser := getParser(ctx)
	configPath, err := parser.GetResourcePath(ctx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(configPath)
	// Output: /opt/senzing/g2/resources
}

func ExampleEngineConfigurationJsonParserImpl_GetSupportPath() {
	// For more information, visit https://github.com/Senzing/go-common/blob/main/engineconfigurationjsonparser/engineconfigurationjsonparser_test.go
	ctx := context.TODO()
	parser := getParser(ctx)
	configPath, err := parser.GetSupportPath(ctx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(configPath)
	// Output: /opt/senzing/data
}
