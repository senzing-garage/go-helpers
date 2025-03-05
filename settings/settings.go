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

	"github.com/senzing-garage/go-helpers/settingsparser"
)

// ----------------------------------------------------------------------------
// Public functions
// ----------------------------------------------------------------------------

func BuildSenzingDatabaseURI(databaseURL string) (string, error) {
	result := ""
	parsedURL, err := url.Parse(databaseURL)
	if err != nil {
		return "", err
	}
	switch parsedURL.Scheme {
	case "db2":
		result = fmt.Sprintf(
			"%s://%s@%s",
			parsedURL.Scheme,
			parsedURL.User,
			string(parsedURL.Path[1:]),
		)
		if len(parsedURL.RawQuery) > 0 {
			result = fmt.Sprintf("%s?%s", result, parsedURL.Query().Encode())
		}
	case "mssql":
		if len(parsedURL.RawQuery) > 0 {
			result = fmt.Sprintf(
				"%s://%s@%s:%s?%s",
				parsedURL.Scheme,
				parsedURL.User,
				parsedURL.Host,
				string(parsedURL.Path[1:]),
				parsedURL.Query().Encode(),
			)
		} else {
			result = fmt.Sprintf(
				"%s://%s@%s",
				parsedURL.Scheme,
				parsedURL.User,
				string(parsedURL.Path[1:]),
			)
		}
	case "mysql":
		result = fmt.Sprintf(
			"%s://%s@%s/?schema=%s%s",
			parsedURL.Scheme,
			parsedURL.User,
			parsedURL.Host,
			string(parsedURL.Path[1:]),
			parsedURL.RawQuery,
		)
	case "oci":
		result = fmt.Sprintf(
			"%s://%s@%s",
			parsedURL.Scheme,
			parsedURL.User,
			string(parsedURL.Path[1:]),
		)
	case "postgresql":
		result = fmt.Sprintf(
			"%s://%s@%s:%s",
			parsedURL.Scheme,
			parsedURL.User,
			parsedURL.Host,
			string(parsedURL.Path[1:]),
		)
		if len(parsedURL.RawQuery) > 0 {
			result = fmt.Sprintf("%s?%s", result, parsedURL.Query().Encode())
		} else {
			result = fmt.Sprintf("%s/", result)
		}
	case "sqlite3":
		result = fmt.Sprintf(
			"%s://%s@%s/%s",
			parsedURL.Scheme,
			parsedURL.User,
			parsedURL.Host,
			string(parsedURL.Path[1:]),
		)
		if len(parsedURL.RawQuery) > 0 {
			result = fmt.Sprintf("%s?%s", result, parsedURL.Query().Encode())
		}
	default:
		result = ""
		err = fmt.Errorf("unknown database schema: %s in %s", parsedURL.Scheme, databaseURL)
	}
	return result, err
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
		return senzingEngineConfigurationJSON, err
	}

	// If SENZING_ENGINE_CONFIGURATION_JSON is set, use it.
	// This is a legacy environment variable and won't be documented.

	senzingEngineConfigurationJSON, isSet = os.LookupEnv("SENZING_ENGINE_CONFIGURATION_JSON")
	if isSet {
		return senzingEngineConfigurationJSON, err
	}

	// If SENZING_PATH is set, use it.

	senzingPath, isSet := os.LookupEnv("SENZING_PATH")
	if isSet {
		attributeMap["senzingPath"] = senzingPath
	}

	// Add database URL.

	senzingDatabaseURL, inMap := attributeMap["databaseURL"]
	if !inMap {
		senzingDatabaseURL, err = getOsEnv("SENZING_TOOLS_DATABASE_URL")
		if err != nil {
			return "", err
		}
	}
	specificDatabaseURL, err := BuildSenzingDatabaseURI(senzingDatabaseURL)
	if err != nil {
		return "", err
	}

	attributeMap["databaseURL"] = specificDatabaseURL

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

	// Construct structure.

	resultStruct := buildStruct(attributeMap)

	// Transform structure to JSON.

	resultBuffer := &bytes.Buffer{}
	jsonEncoder := json.NewEncoder(resultBuffer)
	jsonEncoder.SetEscapeHTML(false)
	err = jsonEncoder.Encode(resultStruct)
	if err != nil {
		return "", err
	}

	return resultBuffer.String(), err
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

	databaseURLs, err := parser.GetDatabaseURLs(ctx)
	if err != nil {
		return err
	}
	for _, value := range databaseURLs {
		if len(value) == 0 {
			return fmt.Errorf("SQL.CONNECTION empty in Senzing engine configuration JSON.\nFor more information, visit https://garage.senzing.com/go-helpers/errors")
		}
	}

	// Check Config path.

	configPath, err := parser.GetConfigPath(ctx)
	if err != nil {
		return err
	}
	configFiles := []string{
		"cfgVariant.json",
		"defaultGNRCP.config",
	}
	for _, configFile := range configFiles {
		targetFile := fmt.Sprintf("%s/%s", configPath, configFile)
		if _, err := os.Stat(targetFile); err != nil {
			return fmt.Errorf("CONFIGPATH: Could not find %s\nFor more information, visit https://garage.senzing.com/go-helpers/errors", targetFile)
		}
	}

	// Check Resource path.

	resourcePath, err := parser.GetResourcePath(ctx)
	if err != nil {
		return err
	}
	resourceFiles := []string{
		"templates/g2config.json",
	}
	for _, resourceFile := range resourceFiles {
		targetFile := fmt.Sprintf("%s/%s", resourcePath, resourceFile)
		if _, err := os.Stat(targetFile); err != nil {
			return fmt.Errorf("RESOURCEPATH: Could not find %s\nFor more information, visit https://garage.senzing.com/go-helpers/errors", targetFile)
		}
	}

	// Check Support path.

	supportPath, err := parser.GetSupportPath(ctx)
	if err != nil {
		return err
	}
	supportFiles := []string{
		"anyTransRule.ibm",
		"g2SifterRules.ibm",
	}
	for _, supportFile := range supportFiles {
		targetFile := fmt.Sprintf("%s/%s", supportPath, supportFile)
		if _, err := os.Stat(targetFile); err != nil {
			return fmt.Errorf("SUPPORTPATH: Could not find %s\nFor more information, visit https://garage.senzing.com/go-helpers/errors", targetFile)
		}
	}

	// Os / Arch specific calls

	err = verifySettings(ctx, settings)

	return err
}

// ----------------------------------------------------------------------------
// Private functions
// ----------------------------------------------------------------------------

func buildStruct(attributeMap map[string]string) SzConfiguration {
	var result SzConfiguration

	databaseURL, ok := attributeMap["databaseURL"]
	if !ok {
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
			Connection: databaseURL,
		},
	}

	licenseStringBase64, ok := attributeMap["licenseStringBase64"]
	if ok {
		result.Pipeline.LicenseStringBase64 = licenseStringBase64
	}

	return result
}

func getOsEnv(variableName string) (string, error) {
	var err error
	result, isSet := os.LookupEnv(variableName)
	if !isSet {
		err = fmt.Errorf("environment variable not set: %s", variableName)
	}
	return result, err
}

func mapWithDefault(aMap map[string]string, key string, defaultValue string) string {
	result, ok := aMap[key]
	if ok {
		return result
	}
	return defaultValue
}
