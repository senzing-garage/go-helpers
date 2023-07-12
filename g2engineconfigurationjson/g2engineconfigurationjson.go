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
	var err error = nil

	if len(senzingDatabaseUrl) == 0 {

		// If SENZING_TOOLS_ENGINE_CONFIGURATION_JSON is set, use it.

		senzingEngineConfigurationJson, err := getOsEnv("SENZING_TOOLS_ENGINE_CONFIGURATION_JSON")
		if err == nil {
			return senzingEngineConfigurationJson, err
		}

		// If SENZING_ENGINE_CONFIGURATION_JSON is set, use it.

		senzingEngineConfigurationJson, err = getOsEnv("SENZING_ENGINE_CONFIGURATION_JSON")
		if err == nil {
			return senzingEngineConfigurationJson, err
		}

		senzingDatabaseUrl, err = getOsEnv("SENZING_TOOLS_DATABASE_URL")
		if err != nil {
			return "", err
		}
	}

	// Construct structure.

	specificDatabaseUrl, specificDatabaseUrlErr := buildSpecificDatabaseUrl(senzingDatabaseUrl)
	if specificDatabaseUrlErr != nil {
		return "", specificDatabaseUrlErr
	}
	licenseStringBase64, _ := os.LookupEnv("SENZING_TOOLS_LICENSE_STRING_BASE64")
	resultStruct := buildStruct(specificDatabaseUrl, licenseStringBase64)

	// Transform structure to JSON.

	resultBytes, _ := json.Marshal(resultStruct)
	return string(resultBytes), err
}
