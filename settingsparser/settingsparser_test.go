package settingsparser

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	settingsParserSingleton SettingsParser
)

// ----------------------------------------------------------------------------
// Test interface methods
// ----------------------------------------------------------------------------

func TestSettingsParser_GetConfigPath(test *testing.T) {
	ctx := context.TODO()
	parser := getTestObject(ctx, test)
	actual, err := parser.GetConfigPath(ctx)
	testError(test, err)
	assert.Equal(test, "/etc/opt/senzing", actual)
}

func TestSettingsParser_GetDatabaseURIs(test *testing.T) {
	ctx := context.TODO()
	parser := &BasicSettingsParser{
		Settings: `
        {
            "PIPELINE": {
                "CONFIGPATH": "/etc/opt/senzing",
                "LICENSESTRINGBASE64": "${SENZING_LICENSE_BASE64_ENCODED}",
                "RESOURCEPATH": "/opt/senzing/er/resources",
                "SUPPORTPATH": "/opt/senzing/data"
            },
            "SQL": {
                "BACKEND": "SQL",
                "CONNECTION": "postgresql://username:password@db.example.com:5432:G2"
            }
        }
        `,
	}
	actual, err := parser.GetDatabaseURIs(ctx)
	testError(test, err)
	assert.Equal(test, []string{"postgresql://username:password@db.example.com:5432:G2"}, actual)
}

func TestSettingsParser_GetDatabaseURIs_Multi(test *testing.T) {
	ctx := context.TODO()
	parser := &BasicSettingsParser{
		Settings: `
        {
            "PIPELINE": {
                "CONFIGPATH": "/etc/opt/senzing",
                "LICENSESTRINGBASE64": "${SENZING_LICENSE_BASE64_ENCODED}",
                "RESOURCEPATH": "/opt/senzing/er/resources",
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
	actual, err := parser.GetDatabaseURIs(ctx)
	testError(test, err)
	assert.Len(test, actual, 3)
	assert.True(test, contains(actual, "postgresql://username:password@db-1.example.com:5432:G2"))
	assert.True(test, contains(actual, "postgresql://username:password@db-2.example.com:5432:G2"))
	assert.True(test, contains(actual, "postgresql://username:password@db-3.example.com:5432:G2"))
}

func TestSettingsParser_GetLicenseStringBase64(test *testing.T) {
	ctx := context.TODO()
	parser := &BasicSettingsParser{
		Settings: `
        {
            "PIPELINE": {
                "CONFIGPATH": "/etc/opt/senzing",
                "LICENSESTRINGBASE64": "${SENZING_LICENSE_BASE64_ENCODED}",
                "RESOURCEPATH": "/opt/senzing/er/resources",
                "SUPPORTPATH": "/opt/senzing/data"
            },
            "SQL": {
                "BACKEND": "SQL",
                "CONNECTION": "postgresql://username:password@db.example.com:5432:G2"
            }
        }
        `,
	}
	actual, err := parser.GetLicenseStringBase64(ctx)
	testError(test, err)
	assert.Equal(test, "${SENZING_LICENSE_BASE64_ENCODED}", actual)
}

func TestSettingsParser_GetSettings(test *testing.T) {
	ctx := context.TODO()
	expected := `{"PIPELINE":{"CONFIGPATH":"/etc/opt/senzing","RESOURCEPATH":"/opt/senzing/er/resources","SUPPORTPATH":"/opt/senzing/data"},"SQL":{"CONNECTION":"postgresql://username:password@hostname:5432:G2/"}}`
	parser := &BasicSettingsParser{
		Settings: expected,
	}
	actual := parser.GetSettings(ctx)
	assert.Equal(test, expected, actual)
}

func TestSettingsParser_New(test *testing.T) {
	ctx := context.TODO()
	settings := `
        {
            "PIPELINE": {
                "CONFIGPATH": "/etc/opt/senzing",
                "LICENSESTRINGBASE64": "${SENZING_LICENSE_BASE64_ENCODED}",
                "RESOURCEPATH": "/opt/senzing/er/resources",
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

	parser, err := New(settings)
	testError(test, err)
	actual, err := parser.GetDatabaseURIs(ctx)
	testError(test, err)
	assert.Len(test, actual, 3)
	assert.True(test, contains(actual, "postgresql://username:password@db-1.example.com:5432:G2"))
	assert.True(test, contains(actual, "postgresql://username:password@db-2.example.com:5432:G2"))
	assert.True(test, contains(actual, "postgresql://username:password@db-3.example.com:5432:G2"))
}

func TestSettingsParser_New_badJSON(test *testing.T) {
	settings := "}{"
	_, err := New(settings)
	require.Error(test, err)
}

// ----------------------------------------------------------------------------
// Internal functions
// ----------------------------------------------------------------------------

func getTestObject(ctx context.Context, test *testing.T) SettingsParser {
	_ = test
	return getParser(ctx)
}

func getParser(ctx context.Context) SettingsParser {
	_ = ctx

	if settingsParserSingleton == nil {
		settingsParserSingleton = &BasicSettingsParser{
			Settings: `
            {
                "PIPELINE": {
                    "CONFIGPATH": "/etc/opt/senzing",
                    "LICENSESTRINGBASE64": "${SENZING_LICENSE_BASE64_ENCODED}",
                    "RESOURCEPATH": "/opt/senzing/er/resources",
                    "SUPPORTPATH": "/opt/senzing/data"
                },
                "SQL": {
                    "CONNECTION": "postgresql://username:password@db.example.com:5432:G2"
                }
            }
            `,
		}
	}

	return settingsParserSingleton
}

func testError(test *testing.T, err error) {
	if err != nil {
		assert.FailNow(test, err.Error())
	}
}
