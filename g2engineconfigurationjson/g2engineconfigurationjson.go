/*
Package g2engineconfigurationjson is used to generate the JSON document used to configure a Senzing client.
*/
package g2engineconfigurationjson

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"os"

	"github.com/senzing/go-common/engineconfigurationjsonparser"
)

// ----------------------------------------------------------------------------
// Internal methods
// ----------------------------------------------------------------------------

func getOsEnv(variableName string) (string, error) {
	var err error = nil
	result, isSet := os.LookupEnv(variableName)
	if !isSet {
		err = fmt.Errorf("environment variable not set: %s", variableName)
	}
	return result, err
}

func buildSpecificDatabaseUrl(databaseUrl string) (string, error) {
	result := ""
	parsedUrl, err := url.Parse(databaseUrl)
	if err != nil {
		return "", err
	}

	switch parsedUrl.Scheme {
	case "db2":
		result = fmt.Sprintf(
			"%s://%s@%s",
			parsedUrl.Scheme,
			parsedUrl.User,
			string(parsedUrl.Path[1:]),
		)
		if len(parsedUrl.RawQuery) > 0 {
			result = fmt.Sprintf("%s?%s", result, parsedUrl.RawQuery)
		}
	case "mssql":
		result = fmt.Sprintf(
			"%s://%s@%s",
			parsedUrl.Scheme,
			parsedUrl.User,
			string(parsedUrl.Path[1:]),
		)
	case "mysql":
		result = fmt.Sprintf(
			"%s://%s@%s/?schema=%s%s",
			parsedUrl.Scheme,
			parsedUrl.User,
			parsedUrl.Host,
			string(parsedUrl.Path[1:]),
			parsedUrl.RawQuery,
		)
	case "oci":
		result = fmt.Sprintf(
			"%s://%s@%s",
			parsedUrl.Scheme,
			parsedUrl.User,
			string(parsedUrl.Path[1:]),
		)
	case "postgresql":
		result = fmt.Sprintf(
			"%s://%s@%s:%s",
			parsedUrl.Scheme,
			parsedUrl.User,
			parsedUrl.Host,
			string(parsedUrl.Path[1:]),
		)
		if len(parsedUrl.RawQuery) > 0 {
			result = fmt.Sprintf("%s?%s", result, parsedUrl.RawQuery)
		} else {
			result = fmt.Sprintf("%s/", result)
		}
	case "sqlite3":
		result = fmt.Sprintf(
			"%s://%s@%s/%s",
			parsedUrl.Scheme,
			parsedUrl.User,
			parsedUrl.Host,
			string(parsedUrl.Path[1:]),
		)
	default:
		result = ""
	}

	return result, err
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

/*
The BuildSimpleSystemConfigurationJson method is a convenience method
for invoking BuildSimpleSystemConfigurationJsonUsingMap() passing in only the
"databaseUrl" mapped value.

Input
  - senzingDatabaseUrl: A Database URL.

Output
  - A string containing a JSON document use when calling Senzing's Init(...) methods.
    See the example output.

Deprecated: Use BuildSimpleSystemConfigurationJsonUsingEnvVars() or BuildSimpleSystemConfigurationJsonUsingMap() instead.
*/
func BuildSimpleSystemConfigurationJson(senzingDatabaseUrl string) (string, error) {
	specificDatabaseUrl, err := buildSpecificDatabaseUrl(senzingDatabaseUrl)
	if err != nil {
		return "", err
	}
	attributeMap := map[string]string{
		"databaseUrl": specificDatabaseUrl,
	}
	return BuildSimpleSystemConfigurationJsonUsingMap(attributeMap)
}

/*
The BuildSimpleSystemConfigurationJsonUsingEnvVars method is a convenience method
for invoking BuildSimpleSystemConfigurationJsonUsingMap without any mapped values.
In other words, only environment variables will be used.

See BuildSimpleSystemConfigurationJsonUsingMap() for information on the environment variables used.

Output
  - A string containing a JSON document use when calling Senzing's Init(...) methods.
    See the example output.
*/
func BuildSimpleSystemConfigurationJsonUsingEnvVars() (string, error) {
	attributeMap := map[string]string{}
	return BuildSimpleSystemConfigurationJsonUsingMap(attributeMap)
}

/*
The BuildSimpleSystemConfigurationJsonUsingMap method returns a JSON document for use with Senzing's Init(...) methods.

If the environment variable SENZING_TOOLS_ENGINE_CONFIGURATION_JSON is set,
the value of SENZING_TOOLS_ENGINE_CONFIGURATION_JSON will be returned unchanged.

If the SENZING_TOOLS_ENGINE_CONFIGURATION_JSON environment variable is not found,
the precedence used in calculating the values of the returned JSON are:

 1. Key/value in attributeMap
 2. Environment variable
 3. Default or a calculated value

The keys and corresponding environment variables are:

	Key						Environment variable
	---------------------  	----------------------------------
	databaseUrl 			SENZING_TOOLS_DATABASE_URL
	licenseStringBase64 	SENZING_TOOLS_LICENSE_STRING_BASE64
	senzingDirectory    	SENZING_TOOLS_SENZING_DIRECTORY
	configPath          	SENZING_TOOLS_CONFIG_PATH
	resourcePath        	SENZING_TOOLS_RESOURCE_PATH
	supportPath         	SENZING_TOOLS_SUPPORT_PATH

Input
  - attributeMap: A mapping of a keys to desired values.
    If key doesn't exist, an environment variable will be used when constructing output JSON.
    If environment variable doesn't exist, a default or calculated value will be used when constructing output JSON.

Output
  - A string containing a JSON document use when calling Senzing's Init(...) methods.
    See the example output.
*/
func BuildSimpleSystemConfigurationJsonUsingMap(attributeMap map[string]string) (string, error) {
	var err error = nil

	// If SENZING_TOOLS_ENGINE_CONFIGURATION_JSON is set, use it.

	senzingEngineConfigurationJson, isSet := os.LookupEnv("SENZING_TOOLS_ENGINE_CONFIGURATION_JSON")
	if isSet {
		return senzingEngineConfigurationJson, err
	}

	// If SENZING_ENGINE_CONFIGURATION_JSON is set, use it.
	// This is a legacy environment variable and won't be documented.

	senzingEngineConfigurationJson, isSet = os.LookupEnv("SENZING_ENGINE_CONFIGURATION_JSON")
	if isSet {
		return senzingEngineConfigurationJson, err
	}

	// Add database URL.

	_, inMap := attributeMap["databaseUrl"]
	if !inMap {
		senzingDatabaseUrl, err := getOsEnv("SENZING_TOOLS_DATABASE_URL")
		if err != nil {
			return "", err
		}
		specificDatabaseUrl, err := buildSpecificDatabaseUrl(senzingDatabaseUrl)
		if err != nil {
			return "", err
		}
		attributeMap["databaseUrl"] = specificDatabaseUrl
	}

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

	resultBytes, _ := json.Marshal(resultStruct)
	return string(resultBytes), err
}

/*
The VerifySenzingEngineConfigurationJson method inspects the Senzing engine configuration JSON to see if it is misconfigured.

Errors are documented at https://hub.senzing.com/go-common/errors.

Input
  - ctx: A context to control lifecycle.
  - senzingEngineConfigurationJson: A JSON string. See https://github.com/Senzing/knowledge-base/blob/main/lists/environment-variables.md#senzing_tools_engine_configuration_json
*/
func VerifySenzingEngineConfigurationJson(ctx context.Context, senzingEngineConfigurationJson string) error {
	var err error = nil
	parser := engineconfigurationjsonparser.EngineConfigurationJsonParserImpl{
		EngineConfigurationJson: senzingEngineConfigurationJson,
	}

	// Check database URLs.

	databaseUrls, err := parser.GetDatabaseUrls(ctx)
	if err != nil {
		return err
	}
	for _, value := range databaseUrls {
		if len(value) == 0 {
			return fmt.Errorf("SQL.CONNECTION empty in Senzing engine configuration JSON.\nFor more information, visit https://hub.senzing.com/go-common/errors")
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
			return fmt.Errorf("CONFIGPATH: Could not find %s\nFor more information, visit https://hub.senzing.com/go-common/errors", targetFile)
		}
	}

	// Check Resource path.

	resourcePath, err := parser.GetResourcePath(ctx)
	if err != nil {
		return err
	}
	resourceFiles := []string{
		"templates/cfgVariant.json",
	}
	for _, resourceFile := range resourceFiles {
		targetFile := fmt.Sprintf("%s/%s", resourcePath, resourceFile)
		if _, err := os.Stat(targetFile); err != nil {
			return fmt.Errorf("RESOURCEPATH: Could not find %s\nFor more information, visit https://hub.senzing.com/go-common/errors", targetFile)
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
	for _, resourceFile := range supportFiles {
		targetFile := fmt.Sprintf("%s/%s", supportPath, resourceFile)
		if _, err := os.Stat(targetFile); err != nil {
			return fmt.Errorf("SUPPORTPATH: Could not find %s\nFor more information, visit https://hub.senzing.com/go-common/errors", targetFile)
		}
	}

	// Os / Arch specific calls

	err = verifySenzingEngineConfigurationJson(ctx, senzingEngineConfigurationJson)

	return err
}
