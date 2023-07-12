/*
Package g2engineconfigurationjson is used to generate the JSON document used to configure a Senzing client.
*/
package g2engineconfigurationjson

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
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
The BuildSimpleSystemConfigurationJson method returns a JSON document for use with Senzing's Init(...) methods.
The configuration is for a "system install" with a single database.

If the senzingDatabaseUrl parameter is an empty string and the environment variable SENZING_TOOLS_ENGINE_CONFIGURATION_JSON is set,
the value of SENZING_TOOLS_ENGINE_CONFIGURATION_JSON will be returned.

If the senzingDatabaseUrl parameter is an empty string and the environment variable SENZING_ENGINE_CONFIGURATION_JSON is set,
the value of SENZING_ENGINE_CONFIGURATION_JSON will be returned.

If the senzingDatabaseUrl parameter is an empty string and the environment variable SENZING_TOOLS_DATABASE_URL is set,
the value of SENZING_TOOLS_DATABASE_URL will  be used as the senzingDatabaseUrl.

Input
  - senzingDatabaseUrl: A Database URL.
    If empty, the SENZING_ENGINE_CONFIGURATION_JSON and SENZING_TOOLS_DATABASE_URL environment variables will be used in calculating the result.

Output
  - A string containing a JSON document use when calling Senzing's Init(...) methods.
    See the example output.
*/
func BuildSimpleSystemConfigurationJson(senzingDatabaseUrl string) (string, error) {
	specificDatabaseUrl, err := buildSpecificDatabaseUrl(senzingDatabaseUrl)
	if err != nil {
		return "", err
	}
	attributeMap := map[string]string{
		"databaseUrl": specificDatabaseUrl,
	}
	return BuildSimpleSystemConfigurationJsonViaMap(attributeMap)
}

/*
The BuildSimpleSystemConfigurationJsonWithLicense method returns a JSON document for use with Senzing's Init(...) methods.
The configuration is for a "system install" with a single database.

If the senzingDatabaseUrl parameter is an empty string and the environment variable SENZING_TOOLS_ENGINE_CONFIGURATION_JSON is set,
the value of SENZING_TOOLS_ENGINE_CONFIGURATION_JSON will be returned.

If the senzingDatabaseUrl parameter is an empty string and the environment variable SENZING_ENGINE_CONFIGURATION_JSON is set,
the value of SENZING_ENGINE_CONFIGURATION_JSON will be returned.

If the senzingDatabaseUrl parameter is an empty string and the environment variable SENZING_TOOLS_DATABASE_URL is set,
the value of SENZING_TOOLS_DATABASE_URL will  be used as the senzingDatabaseUrl.

Input
  - senzingDatabaseUrl: A Database URL.
    If empty, the SENZING_ENGINE_CONFIGURATION_JSON and SENZING_TOOLS_DATABASE_URL environment variables will be used in calculating the result.
  - licenseStringBase64: A Base64 of a Senzing license. (Optional, may be blank string)
  - senzingDirectory: Location of the Senzing API binaries. (Optional, may be blank string)

Output
  - A string containing a JSON document use when calling Senzing's Init(...) methods.
    See the example output.
*/
func BuildSimpleSystemConfigurationJsonViaMap(attributeMap map[string]string) (string, error) {
	var err error = nil

	// If SENZING_TOOLS_ENGINE_CONFIGURATION_JSON is set, use it.

	senzingEngineConfigurationJson, isSet := os.LookupEnv("SENZING_TOOLS_ENGINE_CONFIGURATION_JSON")
	if isSet {
		return senzingEngineConfigurationJson, err
	}

	// If SENZING_ENGINE_CONFIGURATION_JSON is set, use it. This is a legacy environment variable.

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
		"licenseStringBase64": "8BD296A26F2034AAB436045...",
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
				attributeMap[mapKey] = environmentValue
			}
		}
	}

	// Construct structure.

	resultStruct := buildStruct(attributeMap)

	// Transform structure to JSON.

	resultBytes, _ := json.Marshal(resultStruct)
	return string(resultBytes), err
}
