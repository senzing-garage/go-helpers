package g2engineconfigurationjson

import (
	"context"
	"net/url"
	"testing"

	"github.com/senzing/go-common/engineconfigurationjsonparser"
	"github.com/stretchr/testify/assert"
)

type testCaseMetadata struct {
	configPath          string
	databaseUrl         string
	databaseUrlPath     string
	licenseStringBase64 string
	name                string
	resourcePath        string
	senzingDirectory    string
	supportPath         string
}

var testCasesForMultiPlatform = []testCaseMetadata{

	{
		name:        "db2-001",
		databaseUrl: "db2://username:password@hostname:50000/G2",
	},
	{
		name:        "db2-002",
		databaseUrl: "db2://username:password@hostname:50000/G2/?schema=schemaname",
	},
	{
		name:        "oci-001",
		databaseUrl: "oci://username:password@hostname:1521/G2",
	},
	{
		name:        "mssql-001",
		databaseUrl: "mssql://username:password@hostname:1433/G2",
	},
	{
		name:        "mysql-001",
		databaseUrl: "mysql://username:password@hostname:3306/G2",
	},
	{
		name:        "oci-001",
		databaseUrl: "oci://username:password@hostname:1521/G2",
	},
	{
		name:        "postgresql-001",
		databaseUrl: "postgresql://username:password@hostname:5432/G2",
	},
	{
		name:        "postgresql-002",
		databaseUrl: "postgresql://username:password@hostname:5432/G2/?schema=schemaname",
	},
}

var testCases = append(testCasesForMultiPlatform, testCasesForOsArch...)

// ----------------------------------------------------------------------------
// Internal functions
// ----------------------------------------------------------------------------

func testError(test *testing.T, err error) {
	if err != nil {
		assert.FailNow(test, err.Error())
	}
}

func buildMap(testCase testCaseMetadata) map[string]string {
	result := map[string]string{}
	if len(testCase.configPath) > 0 {
		result["configPath"] = testCase.configPath
	}
	if len(testCase.databaseUrl) > 0 {
		result["databaseUrl"] = testCase.databaseUrl
	}
	if len(testCase.licenseStringBase64) > 0 {
		result["licenseStringBase64"] = testCase.licenseStringBase64
	}
	if len(testCase.resourcePath) > 0 {
		result["resourcePath"] = testCase.resourcePath
	}
	if len(testCase.senzingDirectory) > 0 {
		result["senzingDirectory"] = testCase.senzingDirectory
	}
	if len(testCase.supportPath) > 0 {
		result["supportPath"] = testCase.supportPath
	}
	return result

}

// ----------------------------------------------------------------------------
// Test interface functions
// ----------------------------------------------------------------------------

func TestBuildSimpleSystemConfigurationJsonUsingEnvVars(test *testing.T) {
	_, err := BuildSimpleSystemConfigurationJsonUsingEnvVars()
	testError(test, err)
}

func TestBuildSimpleSystemConfigurationJsonUsingMap(test *testing.T) {
	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			aMap := buildMap(testCase)
			_, err := BuildSimpleSystemConfigurationJsonUsingMap(aMap)
			testError(test, err)
		})
	}
}

func TestVerifySenzingEngineConfigurationJson(test *testing.T) {
	ctx := context.TODO()
	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			aMap := buildMap(testCase)
			testJson, err := BuildSimpleSystemConfigurationJsonUsingMap(aMap)
			testError(test, err)
			err = VerifySenzingEngineConfigurationJson(ctx, testJson)
			testError(test, err)
		})
	}
}

func TestBuildSimpleSystemConfigurationJsonUsingMap_ParseResult(test *testing.T) {
	ctx := context.TODO()
	for _, testCase := range testCases {
		if len(testCase.databaseUrlPath) > 0 {
			test.Run(testCase.name, func(test *testing.T) {
				aMap := buildMap(testCase)
				engineConfigurationJson, err := BuildSimpleSystemConfigurationJsonUsingMap(aMap)
				testError(test, err)
				parsedEngineConfigurationJson, err := engineconfigurationjsonparser.New(engineConfigurationJson)
				testError(test, err)
				databaseUrls, err := parsedEngineConfigurationJson.GetDatabaseUrls(ctx)
				testError(test, err)
				parsedDatabaseUrl, err := url.Parse(databaseUrls[0])
				testError(test, err)
				assert.Equal(test, testCase.databaseUrlPath, parsedDatabaseUrl.Path)
			})
		}
	}
}
