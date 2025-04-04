package settingsparser_test

import (
	"context"
	"testing"

	"github.com/senzing-garage/go-helpers/settingsparser"
	"github.com/stretchr/testify/require"
)

var (
	settingsParserSingleton settingsparser.SettingsParser
)

// ----------------------------------------------------------------------------
// Test interface methods
// ----------------------------------------------------------------------------

func TestSettingsParser_GetConfigPath(test *testing.T) {
	ctx := test.Context()
	parser := getTestObject(ctx, test)
	actual, err := parser.GetConfigPath(ctx)
	require.NoError(test, err)
	require.Equal(test, "/etc/opt/senzing", actual)
}

func TestSettingsParser_GetDatabaseURIs(test *testing.T) {
	ctx := test.Context()
	parser := &settingsparser.BasicSettingsParser{
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
	require.NoError(test, err)
	require.Equal(test, []string{"postgresql://username:password@db.example.com:5432:G2"}, actual)
}

func TestSettingsParser_GetDatabaseURIs_Multi(test *testing.T) {
	ctx := test.Context()
	parser := &settingsparser.BasicSettingsParser{
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
	require.NoError(test, err)
	require.Len(test, actual, 3)
	require.Contains(test, actual, "postgresql://username:password@db-1.example.com:5432:G2")
	require.Contains(test, actual, "postgresql://username:password@db-2.example.com:5432:G2")
	require.Contains(test, actual, "postgresql://username:password@db-3.example.com:5432:G2")
}

func TestSettingsParser_GetLicenseStringBase64(test *testing.T) {
	ctx := test.Context()
	parser := &settingsparser.BasicSettingsParser{
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
	require.NoError(test, err)
	require.Equal(test, "${SENZING_LICENSE_BASE64_ENCODED}", actual)
}

func TestSettingsParser_GetSettings(test *testing.T) {
	ctx := test.Context()
	expected := `{"PIPELINE":{"CONFIGPATH":"/etc/opt/senzing","RESOURCEPATH":"/opt/senzing/er/resources","SUPPORTPATH":"/opt/senzing/data"},"SQL":{"CONNECTION":"postgresql://username:password@hostname:5432:G2/"}}`
	parser := &settingsparser.BasicSettingsParser{
		Settings: expected,
	}
	actual := parser.GetSettings(ctx)
	require.Equal(test, expected, actual)
}

func TestSettingsParser_New(test *testing.T) {
	ctx := test.Context()
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

	parser, err := settingsparser.New(settings)
	require.NoError(test, err)
	actual, err := parser.GetDatabaseURIs(ctx)
	require.NoError(test, err)
	require.Len(test, actual, 3)
	require.Contains(test, actual, "postgresql://username:password@db-1.example.com:5432:G2")
	require.Contains(test, actual, "postgresql://username:password@db-2.example.com:5432:G2")
	require.Contains(test, actual, "postgresql://username:password@db-3.example.com:5432:G2")
}

func TestSettingsParser_New_badJSON(test *testing.T) {
	settings := "}{"
	_, err := settingsparser.New(settings)
	require.Error(test, err)
}

// ----------------------------------------------------------------------------
// Internal functions
// ----------------------------------------------------------------------------

func getTestObject(ctx context.Context, t *testing.T) settingsparser.SettingsParser {
	t.Helper()
	return getParser(ctx)
}

func getParser(ctx context.Context) settingsparser.SettingsParser {
	_ = ctx

	if settingsParserSingleton == nil {
		settingsParserSingleton = &settingsparser.BasicSettingsParser{
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
