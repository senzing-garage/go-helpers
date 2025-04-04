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
	databaseURI         string
	licenseStringBase64 string
	name                string
	resourcePath        string
	notReversible       bool
	senzingDirectory    string
	supportPath         string
}

var testCasesForMultiPlatform = []testCaseMetadata{
	{
		name:          "mssql-001",
		databaseURL:   "mssql://username:password@hostname:1433/G2",
		databaseURI:   "mssql://username:password@G2",
		notReversible: true,
	},
	{
		name:        "mssql-002",
		databaseURL: "mssql://username:password@hostname:1433/G2/?TrustServerCertificate=True&driver=mssqldriver",
		databaseURI: "mssql://username:password@hostname:1433:G2/?TrustServerCertificate=True&driver=mssqldriver",
	},
	{
		name:        "mssql-003",
		databaseURL: "mssql://username:password@hostname:1433/G2/?driver=mssqldriver",
		databaseURI: "mssql://username:password@hostname:1433:G2/?driver=mssqldriver",
	},
	{
		name:          "mssql-004",
		databaseURL:   "mssql://username:password@hostname:1433/G2?driver=mssqldriver",
		databaseURI:   "mssql://username:password@hostname:1433:G2?driver=mssqldriver",
		notReversible: true,
	},
	{
		name:        "mssql-005",
		databaseURL: "mssql://sa:Passw0rd@localhost:1433/G2/?TrustServerCertificate=True&driver=libmsodbcsql-18.4.so.1.1",
		databaseURI: "mssql://sa:Passw0rd@localhost:1433:G2/?TrustServerCertificate=True&driver=libmsodbcsql-18.4.so.1.1",
	},
	{
		name:        "mysql-001",
		databaseURL: "mysql://username:password@hostname:3306/G2",
		databaseURI: "mysql://username:password@hostname:3306/?schema=G2",
	},
	{
		name:        "mysql-002",
		databaseURL: "mysql://mysql:mysql@127.0.0.1:3306/G2",
		databaseURI: "mysql://mysql:mysql@127.0.0.1:3306/?schema=G2",
	},
	{
		name:          "oci-001",
		databaseURL:   "oci://username:password@hostname:1521/G2",
		databaseURI:   "oci://username:password@//hostname:1521/G2",
		notReversible: true, // FIXME: The BuildSenzingDatabaseURL() regex needs to change to make this reversible.
	},
	{
		name:        "oci-002",
		databaseURL: "oci://username:password@hostname:1521/G2/?noTimezoneCheck=true&sysdba=true",
		databaseURI: "oci://username:password@//hostname:1521/G2/?noTimezoneCheck=true&sysdba=true",
	},
	{
		name:          "oci-003",
		databaseURL:   "oci://username:password@hostname:1521/G2?noTimezoneCheck=true&sysdba=true",
		databaseURI:   "oci://username:password@//hostname:1521/G2?noTimezoneCheck=true&sysdba=true",
		notReversible: true,
	},
	{
		name:        "oci-004",
		databaseURL: "oci://sys:Passw0rd@oracle:1521/G2/?noTimezoneCheck=true&sysdba=true",
		databaseURI: "oci://sys:Passw0rd@//oracle:1521/G2/?noTimezoneCheck=true&sysdba=true",
	},
	{
		name:        "postgresql-001",
		databaseURL: "postgresql://username:password@hostname:5432/G2",
		databaseURI: "postgresql://username:password@hostname:5432:G2/",
	},
	{
		name:        "postgresql-002",
		databaseURL: "postgresql://username:password@hostname:5432/G2/?schema=schemaname",
		databaseURI: "postgresql://username:password@hostname:5432:G2/?schema=schemaname",
	},
	{
		name:        "postgresql-003",
		databaseURL: "postgresql://postgres:postgres@localhost:5432/G2/?sslmode=disable",
		databaseURI: "postgresql://postgres:postgres@localhost:5432:G2/?sslmode=disable",
	},
}

var testCases = append(testCasesForMultiPlatform, testCasesForOsArch...)

// ----------------------------------------------------------------------------
// Test Public functions
// ----------------------------------------------------------------------------

func TestBuildSenzingDatabaseURI(test *testing.T) {
	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			result, err := BuildSenzingDatabaseURI(testCase.databaseURL)
			testError(test, err)
			assert.Equal(test, testCase.databaseURI, result)
		})
	}
}

func TestBuildSenzingDatabaseURL(test *testing.T) {
	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			result, err := BuildSenzingDatabaseURL(testCase.databaseURI)
			if testCase.notReversible {
				assert.Error(test, err)
			} else {
				testError(test, err)
				assert.Equal(test, testCase.databaseURL, result)
			}
		})
	}
}

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
				databaseURLs, err := parsedSettings.GetDatabaseURIs(ctx)
				testError(test, err)
				parsedDatabaseURL, err := url.Parse(databaseURLs[0])
				testError(test, err)
				assert.Equal(test, testCase.databaseURLPath, parsedDatabaseURL.Path)
			})
		}
	}
}

func TestGetSenzingPath(test *testing.T) {
	actual := GetSenzingPath()
	assert.Equal(test, getSenzingPath(), actual)
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
	actual, err := BuildSenzingDatabaseURI("::::")
	require.Error(test, err)
	assert.Empty(test, actual)
}

func Test_buildSpecificDatabaseURL_badDatabaseURLProtocol(test *testing.T) {
	actual, err := BuildSenzingDatabaseURI("xyzzy://something")
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
