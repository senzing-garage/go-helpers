package settings

import (
	"context"
	"net/url"
	"testing"

	"github.com/senzing-garage/go-helpers/settingsparser"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testCaseMetadata struct {
	configPath          string
	databaseURL         string
	databaseURLPath     string
	licenseStringBase64 string
	name                string
	resourcePath        string
	senzingDirectory    string
	supportPath         string
}

var testCasesForMultiPlatform = []testCaseMetadata{

	{
		name:        "db2-001",
		databaseURL: "db2://username:password@hostname:50000/G2",
	},
	{
		name:        "db2-002",
		databaseURL: "db2://username:password@hostname:50000/G2/?schema=schemaname",
	},
	{
		name:        "oci-001",
		databaseURL: "oci://username:password@hostname:1521/G2",
	},
	{
		name:        "mssql-001",
		databaseURL: "mssql://username:password@hostname:1433/G2",
	},
	{
		name:        "mysql-001",
		databaseURL: "mysql://username:password@hostname:3306/G2",
	},
	{
		name:        "oci-001",
		databaseURL: "oci://username:password@hostname:1521/G2",
	},
	{
		name:        "postgresql-001",
		databaseURL: "postgresql://username:password@hostname:5432/G2",
	},
	{
		name:        "postgresql-002",
		databaseURL: "postgresql://username:password@hostname:5432/G2/?schema=schemaname",
	},
}

var testCases = append(testCasesForMultiPlatform, testCasesForOsArch...)

// ----------------------------------------------------------------------------
// Test Public functions
// ----------------------------------------------------------------------------

func TestBuildSimpleSettingsUsingEnvVars(test *testing.T) {
	_, err := BuildSimpleSettingsUsingEnvVars()
	testError(test, err)
}

func TestBuildSimpleSettingsUsingMap(test *testing.T) {
	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			aMap := buildMap(testCase)
			_, err := BuildSimpleSettingsUsingMap(aMap)
			testError(test, err)
		})
	}
}

func TestBuildSimpleSettingsUsingMap_using_SENZING_TOOLS_ENGINE_CONFIGURATION_JSON(test *testing.T) {
	expected := "test value"
	test.Setenv("SENZING_TOOLS_ENGINE_CONFIGURATION_JSON", expected)
	actual, err := BuildSimpleSettingsUsingMap(map[string]string{})
	require.NoError(test, err)
	assert.Equal(test, expected, actual)
}

func TestBuildSimpleSettingsUsingMap_using_SENZING_ENGINE_CONFIGURATION_JSON(test *testing.T) {
	expected := "test value"
	test.Setenv("SENZING_ENGINE_CONFIGURATION_JSON", expected)
	actual, err := BuildSimpleSettingsUsingMap(map[string]string{})
	require.NoError(test, err)
	assert.Equal(test, expected, actual)
}

func TestBuildSimpleSettingsUsingMap_using_SENZING_TOOLS_LICENSE_STRING_BASE64(test *testing.T) {
	ctx := context.TODO()
	expected := "A1B2C3D4"
	test.Setenv("SENZING_TOOLS_LICENSE_STRING_BASE64", expected)
	actual, err := BuildSimpleSettingsUsingMap(map[string]string{})
	require.NoError(test, err)
	parsedActual, err := settingsparser.New(actual)
	require.NoError(test, err)
	licenseStringBase64, err := parsedActual.GetLicenseStringBase64(ctx)
	require.NoError(test, err)
	assert.Equal(test, expected, licenseStringBase64)
}

func TestBuildSimpleSettingsUsingMap_ParseResult(test *testing.T) {
	ctx := context.TODO()
	for _, testCase := range testCases {
		if len(testCase.databaseURLPath) > 0 {
			test.Run(testCase.name, func(test *testing.T) {
				aMap := buildMap(testCase)
				settings, err := BuildSimpleSettingsUsingMap(aMap)
				testError(test, err)
				parsedSettings, err := settingsparser.New(settings)
				testError(test, err)
				databaseURLs, err := parsedSettings.GetDatabaseURLs(ctx)
				testError(test, err)
				parsedDatabaseURL, err := url.Parse(databaseURLs[0])
				testError(test, err)
				assert.Equal(test, testCase.databaseURLPath, parsedDatabaseURL.Path)
			})
		}
	}
}

func TestVerifySettings(test *testing.T) {
	ctx := context.TODO()
	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			aMap := buildMap(testCase)
			testJSON, err := BuildSimpleSettingsUsingMap(aMap)
			testError(test, err)
			err = VerifySettings(ctx, testJSON)
			testError(test, err)
		})
	}
}

// ----------------------------------------------------------------------------
// Test private functions
// ----------------------------------------------------------------------------

func Test_buildSpecificDatabaseURL_badDatabaseURL(test *testing.T) {
	actual, err := buildSpecificDatabaseURL("::::")
	require.Error(test, err)
	assert.Empty(test, actual)
}

func Test_buildSpecificDatabaseURL_badDatabaseURLProtocol(test *testing.T) {
	actual, err := buildSpecificDatabaseURL("xyzzy://something")
	require.Error(test, err)
	assert.Empty(test, actual)
}

func Test_getOsEnv_badEnvVarName(test *testing.T) {
	actual, err := getOsEnv("SENZING_ENVIRONMENT_VARIABLE_DOES_NOT_EXIST")
	require.Error(test, err)
	assert.Empty(test, actual)
}

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
	if len(testCase.databaseURL) > 0 {
		result["databaseURL"] = testCase.databaseURL
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
