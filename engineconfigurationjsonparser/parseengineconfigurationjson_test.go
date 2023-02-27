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
			EnableConfigurationJson: `{"customer":"Senzing Public Test License","contract":"EVALUATION - support@senzing.com","issueDate":"2022-11-29","licenseType":"EVAL (Solely for non-productive use)","licenseLevel":"STANDARD","billing":"MONTHLY","expireDate":"2023-11-29","recordLimit":50000}`,
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
	assert.Equal(test, "", actual)
}

// ----------------------------------------------------------------------------
// Examples for godoc documentation
// ----------------------------------------------------------------------------

func ExampleEngineConfigurationJsonParserImpl_GetConfigPath() {
	ctx := context.TODO()
	parser := getParser(ctx)
	configPath, err := parser.GetConfigPath(ctx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(configPath)
	// Output: {"PIPELINE":{"CONFIGPATH":"/etc/opt/senzing","RESOURCEPATH":"/opt/senzing/g2/resources","SUPPORTPATH":"/opt/senzing/data"},"SQL":{"CONNECTION":"postgresql://username:password@hostname:5432:G2/"}}
}
