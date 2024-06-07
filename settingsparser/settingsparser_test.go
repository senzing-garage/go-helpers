package settingsparser

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	engineConfigurationJSONParserSingleton *BasicEngineConfigurationJSONParser
)

// ----------------------------------------------------------------------------
// Internal functions
// ----------------------------------------------------------------------------

func getTestObject(ctx context.Context, test *testing.T) *BasicEngineConfigurationJSONParser {
	_ = test
	return getParser(ctx)
}

func getParser(ctx context.Context) *BasicEngineConfigurationJSONParser {
	_ = ctx
	if engineConfigurationJSONParserSingleton == nil {
		engineConfigurationJSONParserSingleton = &BasicEngineConfigurationJSONParser{
			EngineConfigurationJSON: `
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
	return engineConfigurationJSONParserSingleton
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

func TestEngineConfigurationJsonParserImpl_GetDatabaseUrls(test *testing.T) {
	ctx := context.TODO()
	parser := &BasicEngineConfigurationJSONParser{
		EngineConfigurationJSON: `
        {
            "PIPELINE": {
                "CONFIGPATH": "/etc/opt/senzing",
                "LICENSESTRINGBASE64": "${SENZING_LICENSE_BASE64_ENCODED}",
                "RESOURCEPATH": "/opt/senzing/g2/resources",
                "SUPPORTPATH": "/opt/senzing/data"
            },
            "SQL": {
                "BACKEND": "SQL",
                "CONNECTION": "postgresql://username:password@db.example.com:5432:G2"
            }
        }
        `,
	}
	actual, err := parser.GetDatabaseUrls(ctx)
	testError(test, err)
	assert.Equal(test, []string{"postgresql://username:password@db.example.com:5432:G2"}, actual)
}

func TestEngineConfigurationJsonParserImpl_GetDatabaseUrls_Multi(test *testing.T) {
	ctx := context.TODO()
	parser := &BasicEngineConfigurationJSONParser{
		EngineConfigurationJSON: `
        {
            "PIPELINE": {
                "CONFIGPATH": "/etc/opt/senzing",
                "LICENSESTRINGBASE64": "${SENZING_LICENSE_BASE64_ENCODED}",
                "RESOURCEPATH": "/opt/senzing/g2/resources",
                "SUPPORTPATH": "/opt/senzing/data"
            },
            "SQL": {
                "BACKEND": "HYBRID",
                "CONNECTION": "postgresql://username:password@db-1.example.com:5432:G2"
            },
            "C1": {
                "CLUSTER_SIZE": "1",
                "DB_1": "postgresql://username:password@db-2.example.com:5432:G2"
            },
            "C2": {
                "CLUSTER_SIZE": "1",
                "DB_1": "postgresql://username:password@db-3.example.com:5432:G2"
            },
            "HYBRID": {
                "RES_FEAT": "C1",
                "RES_FEAT_EKEY": "C1",
                "RES_FEAT_LKEY": "C1",
                "RES_FEAT_STAT": "C1",
                "LIB_FEAT": "C2",
                "LIB_FEAT_HKEY": "C2"
            }
        }
        `,
	}
	actual, err := parser.GetDatabaseUrls(ctx)
	testError(test, err)
	assert.Len(test, actual, 3)
	assert.True(test, contains(actual, "postgresql://username:password@db-1.example.com:5432:G2"))
	assert.True(test, contains(actual, "postgresql://username:password@db-2.example.com:5432:G2"))
	assert.True(test, contains(actual, "postgresql://username:password@db-3.example.com:5432:G2"))
}

func TestEngineConfigurationJsonParserImpl_New(test *testing.T) {
	ctx := context.TODO()

	enginConfigurationJSON := `
        {
            "PIPELINE": {
                "CONFIGPATH": "/etc/opt/senzing",
                "LICENSESTRINGBASE64": "${SENZING_LICENSE_BASE64_ENCODED}",
                "RESOURCEPATH": "/opt/senzing/g2/resources",
                "SUPPORTPATH": "/opt/senzing/data"
            },
            "SQL": {
                "BACKEND": "HYBRID",
                "CONNECTION": "postgresql://username:password@db-1.example.com:5432:G2"
            },
            "C1": {
                "CLUSTER_SIZE": "1",
                "DB_1": "postgresql://username:password@db-2.example.com:5432:G2"
            },
            "C2": {
                "CLUSTER_SIZE": "1",
                "DB_1": "postgresql://username:password@db-3.example.com:5432:G2"
            },
            "HYBRID": {
                "RES_FEAT": "C1",
                "RES_FEAT_EKEY": "C1",
                "RES_FEAT_LKEY": "C1",
                "RES_FEAT_STAT": "C1",
                "LIB_FEAT": "C2",
                "LIB_FEAT_HKEY": "C2"
            }
        }
        `

	parser, err := New(enginConfigurationJSON)
	testError(test, err)
	actual, err := parser.GetDatabaseUrls(ctx)
	testError(test, err)
	assert.Len(test, actual, 3)
	assert.True(test, contains(actual, "postgresql://username:password@db-1.example.com:5432:G2"))
	assert.True(test, contains(actual, "postgresql://username:password@db-2.example.com:5432:G2"))
	assert.True(test, contains(actual, "postgresql://username:password@db-3.example.com:5432:G2"))
}

// ----------------------------------------------------------------------------
// Examples for godoc documentation
// ----------------------------------------------------------------------------

func ExampleBasicEngineConfigurationJSONParser_GetConfigPath() {
	// For more information, visit https://github.com/senzing-garage/go-helpers/blob/main/engineconfigurationjsonparser/engineconfigurationjsonparser_test.go
	ctx := context.TODO()
	parser := getParser(ctx)
	configPath, err := parser.GetConfigPath(ctx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(configPath)
	// Output: /etc/opt/senzing
}

func ExampleBasicEngineConfigurationJSONParser_GetDatabaseUrls() {
	// For more information, visit https://github.com/senzing-garage/go-helpers/blob/main/engineconfigurationjsonparser/engineconfigurationjsonparser_test.go
	ctx := context.TODO()
	parser := getParser(ctx)
	configPath, err := parser.GetDatabaseUrls(ctx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(configPath)
	// Output: [postgresql://username:password@db.example.com:5432:G2]
}

func ExampleBasicEngineConfigurationJSONParser_GetResourcePath() {
	// For more information, visit https://github.com/senzing-garage/go-helpers/blob/main/engineconfigurationjsonparser/engineconfigurationjsonparser_test.go
	ctx := context.TODO()
	parser := getParser(ctx)
	configPath, err := parser.GetResourcePath(ctx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(configPath)
	// Output: /opt/senzing/g2/resources
}

func ExampleBasicEngineConfigurationJSONParser_GetSupportPath() {
	// For more information, visit https://github.com/senzing-garage/go-helpers/blob/main/engineconfigurationjsonparser/engineconfigurationjsonparser_test.go
	ctx := context.TODO()
	parser := getParser(ctx)
	configPath, err := parser.GetSupportPath(ctx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(configPath)
	// Output: /opt/senzing/data
}

func ExampleBasicEngineConfigurationJSONParser_RedactedJSON_single() {
	ctx := context.TODO()
	parser := &BasicEngineConfigurationJSONParser{
		EngineConfigurationJSON: `
        {
            "PIPELINE": {
                "CONFIGPATH": "/etc/opt/senzing",
                "LICENSESTRINGBASE64": "${SENZING_LICENSE_BASE64_ENCODED}",
                "RESOURCEPATH": "/opt/senzing/g2/resources",
                "SUPPORTPATH": "/opt/senzing/data"
            },
            "SQL": {
                "BACKEND": "SQL",
                "CONNECTION": "postgresql://username:password@db.example.com:5432:G2"
            }
        }
        `,
	}

	actual, err := parser.RedactedJSON(ctx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(actual)
	// Output: {"PIPELINE":{"CONFIGPATH":"/etc/opt/senzing","LICENSESTRINGBASE64":"${SENZING_LICENSE_BASE64_ENCODED}","RESOURCEPATH":"/opt/senzing/g2/resources","SUPPORTPATH":"/opt/senzing/data"},"SQL":{"BACKEND":"SQL","CONNECTION":"postgresql://username:xxxxx@db.example.com:5432/G2"}}
}

func ExampleBasicEngineConfigurationJSONParser_RedactedJSON_multiple() {
	ctx := context.TODO()

	engineConfigurationJSON := `
        {
            "PIPELINE": {
                "CONFIGPATH": "/etc/opt/senzing",
                "LICENSESTRINGBASE64": "${SENZING_LICENSE_BASE64_ENCODED}",
                "RESOURCEPATH": "/opt/senzing/g2/resources",
                "SUPPORTPATH": "/opt/senzing/data"
            },
            "SQL": {
                "BACKEND": "HYBRID",
                "CONNECTION": "postgresql://username:password@db-1.example.com:5432:G2"
            },
            "C1": {
                "CLUSTER_SIZE": "1",
                "DB_1": "postgresql://username:password@db-2.example.com:5432:G2"
            },
            "C2": {
                "CLUSTER_SIZE": "1",
                "DB_1": "postgresql://username:password@db-3.example.com:5432:G2"
            },
            "HYBRID": {
                "RES_FEAT": "C1",
                "RES_FEAT_EKEY": "C1",
                "RES_FEAT_LKEY": "C1",
                "RES_FEAT_STAT": "C1",
                "LIB_FEAT": "C2",
                "LIB_FEAT_HKEY": "C2"
            }
        }
        `

	parser, err := New(engineConfigurationJSON)
	if err != nil {
		fmt.Println(err)
	}

	actual, err := parser.RedactedJSON(ctx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(actual)
	// Output: {"PIPELINE":{"CONFIGPATH":"/etc/opt/senzing","LICENSESTRINGBASE64":"${SENZING_LICENSE_BASE64_ENCODED}","RESOURCEPATH":"/opt/senzing/g2/resources","SUPPORTPATH":"/opt/senzing/data"},"SQL":{"BACKEND":"HYBRID","CONNECTION":"postgresql://username:xxxxx@db-1.example.com:5432/G2"},"C1":{"CLUSTER_SIZE":"1","DB_1":"postgresql://username:xxxxx@db-2.example.com:5432/G2"},"C2":{"CLUSTER_SIZE":"1","DB_1":"postgresql://username:xxxxx@db-3.example.com:5432/G2"},"HYBRID":{"RES_FEAT":"C1","RES_FEAT_EKEY":"C1","RES_FEAT_LKEY":"C1","RES_FEAT_STAT":"C1","LIB_FEAT":"C2","LIB_FEAT_HKEY":"C2"}}
}
