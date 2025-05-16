/*
Package settings is used to generate the JSON document used to configure a Senzing client.
*/
package settings

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"regexp"
	"strings"

	"github.com/senzing-garage/go-helpers/settingsparser"
	"github.com/senzing-garage/go-helpers/wraperror"
)

const (
	pathPattern = "/%s/"
)

// ----------------------------------------------------------------------------
// Public functions
// ----------------------------------------------------------------------------

/*
The BuildSenzingDatabaseURI method returns a database URI that is recognized by the Senzing binaries.

Input
  - databaseURL: A parseable URL.

Output
  - A string containing a database URI that can be used in the Senzing engine configuration JSON document.
*/
func BuildSenzingDatabaseURI(databaseURL string) (string, error) {
	var (
		err    error
		result string
	)

	parsedURL, err := url.Parse(databaseURL)
	if err != nil {
		return result, wraperror.Errorf(err, "url.Parse")
	}

	switch parsedURL.Scheme {
	case "mssql":
		return buildURIForMssql(parsedURL)
	case "mysql":
		return buildURIForMysql(parsedURL)
	case "oci":
		return buildURIForOci(parsedURL)
	case "postgresql":
		return buildURIForPostgresql(parsedURL)
	case "sqlite3":
		return buildURIForSqlite3(parsedURL)
	default:
		err = wraperror.Errorf(
			errForPackage,
			"unknown database schema: %s in %s",
			parsedURL.Scheme,
			databaseURL,
		)
	}

	return result, wraperror.Error(err)
}

/*
The BuildSenzingDatabaseURL method returns a parseable database URL based on a Senzing database URI.

Input
  - databaseURI: A string containing a database URI that is used in the Senzing engine configuration JSON document.

Output
  - databaseURL: A parseable URL.
*/
func BuildSenzingDatabaseURL(databaseURI string) (string, error) {
	var (
		err    error
		result string
	)

	switch {
	case strings.HasPrefix(databaseURI, "mssql://"):
		return buildURLForMssql(databaseURI)
	case strings.HasPrefix(databaseURI, "mysql://"):
		return buildURLForMysql(databaseURI)
	case strings.HasPrefix(databaseURI, "oci://"):
		return buildURLForOci(databaseURI)
	case strings.HasPrefix(databaseURI, "postgresql://"):
		return buildURLForPostgresql(databaseURI)
	case strings.HasPrefix(databaseURI, "sqlite3://"):
		return buildURLForSqlite3(databaseURI)
	default:
		err = wraperror.Errorf(errForPackage, "unknown database schema: %s", databaseURI)
	}

	return result, wraperror.Error(err)
}

/*
The BuildSimpleSettingsUsingEnvVars method is a convenience method
for invoking BuildSimpleSettingsUsingMap without any mapped values.
In other words, only environment variables will be used.

See BuildSimpleSettingsUsingMap() for information on the environment variables used.

Output
  - A string containing a JSON document use when calling Senzing's Init(...) methods.
    See the example output.
*/
func BuildSimpleSettingsUsingEnvVars() (string, error) {
	attributeMap := map[string]string{}

	return BuildSimpleSettingsUsingMap(attributeMap)
}

/*
The BuildSimpleSettingsUsingMap method returns a JSON document for use with Senzing's Init(...) methods.

If the environment variable SENZING_TOOLS_ENGINE_CONFIGURATION_JSON is set,
the value of SENZING_TOOLS_ENGINE_CONFIGURATION_JSON will be returned unchanged.

If the SENZING_TOOLS_ENGINE_CONFIGURATION_JSON environment variable is not found,
the precedence used in calculating the values of the returned JSON are:

 1. Key/value in attributeMap
 2. Environment variable
 3. Default or a calculated value

The keys and corresponding environment variables are:

	Key                     Environment variable
	---------------------   ----------------------------------
	configPath              SENZING_TOOLS_CONFIG_PATH
	databaseURL             SENZING_TOOLS_DATABASE_URL
	licenseStringBase64     SENZING_TOOLS_LICENSE_STRING_BASE64
	resourcePath            SENZING_TOOLS_RESOURCE_PATH
	senzingDirectory        SENZING_TOOLS_SENZING_DIRECTORY
	senzingPath             SENZING_PATH
	supportPath             SENZING_TOOLS_SUPPORT_PATH

Input
  - attributeMap: A mapping of a keys to desired values.
    If key doesn't exist, an environment variable will be used when constructing output JSON.
    If environment variable doesn't exist, a default or calculated value will be used when constructing output JSON.

Output
  - A string containing a JSON document use when calling Senzing's Init(...) methods.
    See the example output.
*/
func BuildSimpleSettingsUsingMap(attributeMap map[string]string) (string, error) {
	var err error

	// If SENZING_TOOLS_ENGINE_CONFIGURATION_JSON is set, use it.

	senzingEngineConfigurationJSON, isSet := os.LookupEnv("SENZING_TOOLS_ENGINE_CONFIGURATION_JSON")
	if isSet {
		return senzingEngineConfigurationJSON, wraperror.Errorf(err, "os.LookupEnv.1")
	}

	// If SENZING_ENGINE_CONFIGURATION_JSON is set, use it.
	// This is a legacy environment variable and won't be documented.

	senzingEngineConfigurationJSON, isSet = os.LookupEnv("SENZING_ENGINE_CONFIGURATION_JSON")
	if isSet {
		return senzingEngineConfigurationJSON, wraperror.Errorf(err, "os.LookupEnv.2")
	}

	// If SENZING_PATH is set, use it.

	err = buildAttributeMap(attributeMap)
	if err != nil {
		return senzingEngineConfigurationJSON, wraperror.Errorf(err, "buildAttributeMap")
	}

	// Construct structure.

	resultStruct := buildStruct(attributeMap)

	// Transform structure to JSON.

	resultBuffer := &bytes.Buffer{}
	jsonEncoder := json.NewEncoder(resultBuffer)
	jsonEncoder.SetEscapeHTML(false)

	err = jsonEncoder.Encode(resultStruct)
	if err != nil {
		return "", wraperror.Errorf(err, "jsonEncoder.Encode")
	}

	return resultBuffer.String(), wraperror.Error(err)
}

/*
The GetSenzingPath method returns the path to the Senzing binaries.
If set, This is the value of the SENZING_PATH environment variable.
If not set, it is a default value.

Output
  - A string containing the path to the Senzing binaries.
*/
func GetSenzingPath() string {
	attributeMap := map[string]string{}
	result := getSenzingDirectory(attributeMap)

	return result
}

/*
The VerifySettings method inspects the Senzing engine configuration JSON to see if it is misconfigured.

Errors are documented at https://garage.senzing.com/go-helpers/errors.

Input

  - ctx: A context to control lifecycle.

  - settings: A JSON string. See https://github.com/senzing-garage/knowledge-base/blob/main/lists/environment-variables.md#senzing_tools_engine_configuration_json
*/
func VerifySettings(ctx context.Context, settings string) error {
	var err error

	parser := settingsparser.BasicSettingsParser{
		Settings: settings,
	}

	// Check database URLs.

	databaseURIs, err := parser.GetDatabaseURIs(ctx)
	if err != nil {
		return wraperror.Errorf(err, "GetDatabaseURIs")
	}

	for _, value := range databaseURIs {
		if len(value) == 0 {
			return wraperror.Errorf(
				errForPackage,
				"SQL.CONNECTION empty in Senzing engine configuration JSON. For more information, visit https://garage.senzing.com/go-helpers/errors",
			)
		}
	}

	// Check Config path.

	configPath, err := parser.GetConfigPath(ctx)
	if err != nil {
		return wraperror.Errorf(err, "GetConfigPath")
	}

	err = checkConfigPath(configPath)
	if err != nil {
		return wraperror.Errorf(err, "checkConfigPath")
	}

	// Check Resource path.

	resourcePath, err := parser.GetResourcePath(ctx)
	if err != nil {
		return wraperror.Errorf(err, "GetResourcePath")
	}

	err = checkResourcePath(resourcePath)
	if err != nil {
		return wraperror.Errorf(err, "checkResourcePath")
	}

	// Check Support path.

	supportPath, err := parser.GetSupportPath(ctx)
	if err != nil {
		return wraperror.Errorf(err, "GetSupportPath")
	}

	err = checkSupportPath(supportPath)
	if err != nil {
		return wraperror.Errorf(err, "checkResourcePath")
	}

	// Os / Arch specific calls

	err = verifySettings(ctx, settings)

	return err
}

// ----------------------------------------------------------------------------
// Private functions
// ----------------------------------------------------------------------------

func buildAttributeMap(attributeMap map[string]string) error {
	var err error

	senzingPath, isSet := os.LookupEnv("SENZING_PATH")
	if isSet {
		attributeMap["senzingPath"] = senzingPath
	}

	// Add database URL.

	senzingDatabaseURL, inMap := attributeMap["databaseURL"]
	if !inMap {
		senzingDatabaseURL, err = getOsEnv("SENZING_TOOLS_DATABASE_URL")
		if err != nil {
			return wraperror.Errorf(err, "getOsEnv")
		}
	}

	senzingDatabaseURI, err := BuildSenzingDatabaseURI(senzingDatabaseURL)
	if err != nil {
		return wraperror.Errorf(err, "BuildSenzingDatabaseURI")
	}

	attributeMap["databaseURL"] = senzingDatabaseURI

	// Add Environment Variables to the map, if not already specified in the map.

	keys := map[string]string{
		"licenseStringBase64": "SENZING_TOOLS_LICENSE_STRING_BASE64",
		"senzingDirectory":    "SENZING_TOOLS_SENZING_DIRECTORY",
		"configPath":          "SENZING_TOOLS_CONFIG_PATH",
		"resourcePath":        "SENZING_TOOLS_RESOURCE_PATH",
		"supportPath":         "SENZING_TOOLS_SUPPORT_PATH",
	}

	for mapKey, environmentVariable := range keys {
		_, inMap := attributeMap[mapKey]
		if !inMap {
			environmentValue, isSet := os.LookupEnv(environmentVariable)
			if isSet {
				if len(environmentValue) > 0 {
					attributeMap[mapKey] = environmentValue
				}
			}
		}
	}

	return err
}

func buildStruct(attributeMap map[string]string) SzConfiguration {
	var result SzConfiguration

	databaseURI, isOK := attributeMap["databaseURL"]
	if !isOK {
		return result
	}

	senzingDirectory := getSenzingDirectory(attributeMap)

	// Apply attributeMap.

	result = SzConfiguration{
		Pipeline: SzConfigurationPipeline{
			ConfigPath:   mapWithDefault(attributeMap, "configPath", getConfigPath(senzingDirectory)),
			ResourcePath: mapWithDefault(attributeMap, "resourcePath", getResourcePath(senzingDirectory)),
			SupportPath:  mapWithDefault(attributeMap, "supportPath", getSupportPath(senzingDirectory)),
		},
		SQL: SzConfigurationSQL{
			Connection: databaseURI,
		},
	}

	licenseStringBase64, isOK := attributeMap["licenseStringBase64"]
	if isOK {
		result.Pipeline.LicenseStringBase64 = licenseStringBase64
	}

	return result
}

func buildURIForMssql(parsedURL *url.URL) (string, error) {
	var (
		err    error
		result string
	)

	if len(parsedURL.RawQuery) > 0 {
		result = fmt.Sprintf(
			"%s://%s@%s:%s?%s",
			parsedURL.Scheme,
			parsedURL.User,
			parsedURL.Host,
			parsedURL.Path[1:],
			parsedURL.Query().Encode(),
		)
	} else {
		result = fmt.Sprintf(
			"%s://%s@%s",
			parsedURL.Scheme,
			parsedURL.User,
			parsedURL.Path[1:],
		)
	}

	return result, err
}

func buildURIForMysql(parsedURL *url.URL) (string, error) {
	var err error

	result := fmt.Sprintf(
		"%s://%s@%s/?schema=%s%s",
		parsedURL.Scheme,
		parsedURL.User,
		parsedURL.Host,
		parsedURL.Path[1:],
		parsedURL.RawQuery,
	)

	return result, err
}

func buildURIForOci(parsedURL *url.URL) (string, error) {
	var err error

	result := fmt.Sprintf(
		"%s://%s@//%s/%s",
		parsedURL.Scheme,
		parsedURL.User,
		parsedURL.Host,
		parsedURL.Path[1:],
	)
	if len(parsedURL.RawQuery) > 0 {
		result = fmt.Sprintf("%s?%s", result, parsedURL.Query().Encode())
	}

	return result, err
}

func buildURIForPostgresql(parsedURL *url.URL) (string, error) {
	var err error

	result := fmt.Sprintf(
		"%s://%s@%s:%s",
		parsedURL.Scheme,
		parsedURL.User,
		parsedURL.Host,
		parsedURL.Path[1:],
	)
	if len(parsedURL.RawQuery) > 0 {
		result = fmt.Sprintf("%s?%s", result, parsedURL.Query().Encode())
	} else {
		result += "/"
	}

	return result, err
}

func buildURIForSqlite3(parsedURL *url.URL) (string, error) {
	var err error

	result := fmt.Sprintf(
		"%s://%s@%s/%s",
		parsedURL.Scheme,
		parsedURL.User,
		parsedURL.Host,
		parsedURL.Path[1:],
	)
	if len(parsedURL.RawQuery) > 0 {
		result = fmt.Sprintf("%s?%s", result, parsedURL.Query().Encode())
	}

	return result, err
}

func buildURL(aMap map[string]string) *url.URL {
	var username string

	var password string

	result := &url.URL{}

	for key, value := range aMap {
		switch key {
		case "Scheme":
			result.Scheme = value
		case "Opaque":
			result.Opaque = value
		case "Host":
			result.Host = value
		case "Path":
			result.Path = value
		case "RawPath":
			result.RawPath = value
		case "RawQuery":
			result.RawQuery = value
		case "Fragment":
			result.Fragment = value
		case "RawFragment":
			result.RawFragment = value
		case "username":
			username = value
		case "password":
			password = value
		}
	}

	// Create url.Userinfo

	if len(password) > 0 {
		result.User = url.UserPassword(username, password)
	} else if len(username) > 0 {
		result.User = url.User(username)
	}

	return result
}

func buildURLForMssql(databaseURI string) (string, error) {
	var (
		err    error
		result string
	)

	regExp := regexp.MustCompile(
		`(?P<Scheme>.+)://(?P<username>.+):(?P<password>.+)@(?P<Host>.+):(?P<database>.+)/\?(?P<RawQuery>.+)`,
	)
	regExpMatches := regExp.FindStringSubmatch(databaseURI)
	regExpFieldNames := regExp.SubexpNames()

	aMap := mapNamesToMatches(regExpFieldNames, regExpMatches)
	if !hasRequiredKeys(aMap) {
		return result, wraperror.Errorf(errForPackage, "cannot reconstruct mssql from %s", databaseURI)
	}

	resultURL := buildURL(aMap)

	database, ok := aMap["database"]
	if ok {
		resultURL.Path = fmt.Sprintf(pathPattern, database)
	}

	result = resultURL.String()

	return result, wraperror.Error(err)
}

func buildURLForMysql(databaseURI string) (string, error) {
	var (
		err    error
		result string
	)

	regExp := regexp.MustCompile(
		`(?P<Scheme>.+)://(?P<username>.+):(?P<password>.+)@(?P<Host>.+)/\?schema=(?P<database>.+)`,
	)
	regExpMatches := regExp.FindStringSubmatch(databaseURI)
	regExpFieldNames := regExp.SubexpNames()

	aMap := mapNamesToMatches(regExpFieldNames, regExpMatches)
	if !hasRequiredKeys(aMap) {
		return result, wraperror.Errorf(errForPackage, "cannot reconstruct mysql from %s", databaseURI)
	}

	resultURL := buildURL(aMap)

	database, ok := aMap["database"]
	if ok {
		localPathPattern := "/%s"
		if strings.HasSuffix(databaseURI, "/") {
			localPathPattern = "/%s/"
		}

		resultURL.Path = fmt.Sprintf(localPathPattern, database)
	}

	result = resultURL.String()

	return result, wraperror.Error(err)
}

func buildURLForOci(databaseURI string) (string, error) {
	var (
		err    error
		result string
	)

	regExp := regexp.MustCompile(
		`(?P<Scheme>.+)://(?P<username>.+):(?P<password>.+)@//(?P<Host>.+)/(?P<database>.+)/\?((?P<RawQuery>.+))?`,
	)

	// (?P<Scheme>.+)://(?P<username>.+):(?P<password>.+)@//(?P<Host>.+)/(?P<database>.+)(/?(?P<RawQuery>))?`)
	regExpMatches := regExp.FindStringSubmatch(databaseURI)
	regExpFieldNames := regExp.SubexpNames()

	aMap := mapNamesToMatches(regExpFieldNames, regExpMatches)
	if !hasRequiredKeys(aMap) {
		return result, wraperror.Errorf(
			errForPackage,
			"settings.buildURLForOci cannot reconstruct oci from %s",
			databaseURI,
		)
	}

	resultURL := buildURL(aMap)

	database, ok := aMap["database"]
	if ok {
		resultURL.Path = fmt.Sprintf(pathPattern, database)
	}

	result = resultURL.String()

	return result, wraperror.Error(err)
}

func buildURLForPostgresql(databaseURI string) (string, error) {
	var err error

	index := strings.LastIndex(databaseURI, ":")
	result := strings.TrimSuffix(databaseURI[:index]+"/"+databaseURI[index+1:], "/")

	return result, wraperror.Error(err)
}

func buildURLForSqlite3(databaseURI string) (string, error) {
	var err error

	return databaseURI, wraperror.Error(err)
}

func checkConfigPath(configPath string) error {
	var err error

	configFiles := []string{
		"cfgVariant.json",
		"defaultGNRCP.config",
	}
	for _, configFile := range configFiles {
		targetFile := fmt.Sprintf("%s/%s", configPath, configFile)
		if _, err := os.Stat(targetFile); err != nil {
			return wraperror.Errorf(
				err,
				"CONFIGPATH: Could not find %s. For more information, visit https://garage.senzing.com/go-helpers/errors ",
				targetFile,
			)
		}
	}

	return err
}

func checkResourcePath(resourcePath string) error {
	var err error

	resourceFiles := []string{
		"templates/g2config.json",
	}
	for _, resourceFile := range resourceFiles {
		targetFile := fmt.Sprintf("%s/%s", resourcePath, resourceFile)
		if _, err := os.Stat(targetFile); err != nil {
			return wraperror.Errorf(
				err,
				"RESOURCEPATH: Could not find %s. For more information, visit https://garage.senzing.com/go-helpers/errors",
				targetFile,
			)
		}
	}

	return err
}

func checkSupportPath(supportPath string) error {
	var err error

	supportFiles := []string{
		"anyTransRule.ibm",
		"g2SifterRules.ibm",
	}
	for _, supportFile := range supportFiles {
		targetFile := fmt.Sprintf("%s/%s", supportPath, supportFile)
		if _, err := os.Stat(targetFile); err != nil {
			return wraperror.Errorf(
				err,
				"SUPPORTPATH: Could not find %s. For more information, visit https://garage.senzing.com/go-helpers/errors ",
				targetFile,
			)
		}
	}

	return err
}

func getOsEnv(variableName string) (string, error) {
	var (
		err    error
		result string
	)

	result, isSet := os.LookupEnv(variableName)
	if !isSet {
		return result, wraperror.Errorf(err, "environment variable not set: %s", variableName)
	}

	return result, err
}

func hasRequiredKeys(aMap map[string]string) bool {
	result := true

	requiredKeys := []string{
		"database",
		"Host",
		"Scheme",
		"username",
	}
	for _, requiredKey := range requiredKeys {
		_, ok := aMap[requiredKey]
		if !ok {
			return false
		}
	}

	return result
}

func mapNamesToMatches(names []string, matches []string) map[string]string {
	result := map[string]string{}

	for i, match := range matches {
		result[names[i]] = match
	}

	return result
}

func mapWithDefault(aMap map[string]string, key string, defaultValue string) string {
	result, ok := aMap[key]
	if ok {
		return result
	}

	return defaultValue
}
