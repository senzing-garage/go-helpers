/*
Package settingsparser is used to generate the JSON document used to configure a Senzing client.
*/
package settingsparser

import (
	"context"
	"encoding/json"
	"net/url"
	"strconv"
	"strings"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// BasicSettingsParser is the default implementation of the SettingsParser interface.
type BasicSettingsParser struct {
	Settings string
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

/*
The GetConfigPath method returns the PIPELINE.CONFIGPATH value of _ENGINE_CONFIGURATION_JSON.

Input
  - ctx: A context to control lifecycle.

Output
  - A string containing the value of a PIPELINE.CONFIGPATH.
*/
func (parser *BasicSettingsParser) GetConfigPath(ctx context.Context) (string, error) {
	_ = ctx
	engineConfiguration := &EngineConfiguration{}
	err := json.Unmarshal([]byte(parser.Settings), &engineConfiguration)
	if err != nil {
		return "", err
	}
	return engineConfiguration.Pipeline.ConfigPath, err
}

/*
The GetConfigPath method returns the PIPELINE.CONFIGPATH value of _ENGINE_CONFIGURATION_JSON.

Input
  - ctx: A context to control lifecycle.

Output
  - A string containing the value of a PIPELINE.CONFIGPATH.
*/
func (parser *BasicSettingsParser) GetDatabaseURLs(ctx context.Context) ([]string, error) {
	_ = ctx
	var result []string
	engineConfiguration := &EngineConfiguration{}
	err := json.Unmarshal([]byte(parser.Settings), &engineConfiguration)
	if err != nil {
		return result, err
	}
	result = append(result, engineConfiguration.SQL.Connection)

	// Handle multi-database case.

	backend := engineConfiguration.SQL.Backend
	if len(backend) > 0 && backend != "SQL" {
		var dictionary map[string]interface{}
		var databaseJSONKeys []string
		err = json.Unmarshal([]byte(parser.Settings), &dictionary)
		if err != nil {
			return result, err
		}

		// Determine JSON keys for database definitions.

		backendMap := dictionary[backend]
		for _, value := range backendMap.(map[string]interface{}) {
			valueString := value.(string)
			if !contains(databaseJSONKeys, valueString) {
				databaseJSONKeys = append(databaseJSONKeys, valueString)
			}
		}

		// Add each database.

		for _, databaseJSONKey := range databaseJSONKeys {
			databaseJSON := dictionary[databaseJSONKey].(map[string]interface{})
			databaseName := databaseJSON["DB_1"].(string)
			if !contains(result, databaseName) {
				result = append(result, databaseName)
			}
		}
	}

	// TODO:  Implement multi-database list.

	return result, err
}

/*
The GetLicenseStringBase64 method returns the PIPELINE.LICENSESTRINGBASE64 value of _ENGINE_CONFIGURATION_JSON.

Input
  - ctx: A context to control lifecycle.

Output
  - A string containing the value of a PIPELINE.LICENSESTRINGBASE64.
*/
func (parser *BasicSettingsParser) GetLicenseStringBase64(ctx context.Context) (string, error) {
	_ = ctx
	engineConfiguration := &EngineConfiguration{}
	err := json.Unmarshal([]byte(parser.Settings), &engineConfiguration)
	if err != nil {
		return "", err
	}
	return engineConfiguration.Pipeline.LicenseStringBase64, err
}

/*
The GetResourcePath method returns the PIPELINE.RESOURCEPATH value of _ENGINE_CONFIGURATION_JSON.

Input
  - ctx: A context to control lifecycle.

Output
  - A string containing the value of a PIPELINE.RESOURCEPATH.
*/
func (parser *BasicSettingsParser) GetResourcePath(ctx context.Context) (string, error) {
	_ = ctx
	engineConfiguration := &EngineConfiguration{}
	err := json.Unmarshal([]byte(parser.Settings), &engineConfiguration)
	if err != nil {
		return "", err
	}
	return engineConfiguration.Pipeline.ResourcePath, err
}

/*
The GetSettings returns the entire _ENGINE_CONFIGURATION_JSON.

Input
  - ctx: A context to control lifecycle.

Output
  - A string containing the value of _ENGINE_CONFIGURATION_JSON..
*/
func (parser *BasicSettingsParser) GetSettings(ctx context.Context) string {
	_ = ctx
	return parser.Settings
}

/*
The GetSupportPath method returns the PIPELINE.SUPPORTPATH value of _ENGINE_CONFIGURATION_JSON.

Input
  - ctx: A context to control lifecycle.

Output
  - A string containing the value of a PIPELINE.SUPPORTPATH.
*/
func (parser *BasicSettingsParser) GetSupportPath(ctx context.Context) (string, error) {
	_ = ctx
	engineConfiguration := &EngineConfiguration{}
	err := json.Unmarshal([]byte(parser.Settings), &engineConfiguration)
	if err != nil {
		return "", err
	}
	return engineConfiguration.Pipeline.SupportPath, err
}

/*
The RedactedJSON method returns the JSON string with passwords redacted.

Input
  - ctx: A context to control lifecycle.

Output
  - The Senzing engine configuration JSON string with database URLs having redacted passwords.
*/
func (parser *BasicSettingsParser) RedactedJSON(ctx context.Context) (string, error) {
	result := parser.Settings

	// Get list of database URLs in the Senzing engine configuration json.

	databaseURLs, err := parser.GetDatabaseURLs(ctx)
	if err != nil {
		return "", err
	}

	// For each database URL in the string, replace it with a redacted database URL.

	for _, databaseURL := range databaseURLs {
		redactedURL, err := redactURL(databaseURL)
		if err == nil {
			result = strings.ReplaceAll(result, databaseURL, redactedURL)
		}
	}

	// Remove whitespace.

	result = strings.Join(strings.Fields(result), "")
	return result, err
}

// ----------------------------------------------------------------------------
// Internal methods
// ----------------------------------------------------------------------------

func contains(haystack []string, needle string) bool {
	for _, value := range haystack {
		if value == needle {
			return true
		}
	}
	return false
}

func isJSON(unknownString string) bool {
	unknownStringUnescaped, err := strconv.Unquote(unknownString)
	if err != nil {
		unknownStringUnescaped = unknownString
	}
	var jsonString json.RawMessage
	return json.Unmarshal([]byte(unknownStringUnescaped), &jsonString) == nil
}

func redactURL(aURL string) (string, error) {
	parsedURL, err := url.Parse(aURL)
	if err != nil {
		if strings.HasPrefix(aURL, "postgresql") {
			index := strings.LastIndex(aURL, ":")
			aURL := aURL[:index] + "/" + aURL[index+1:]
			parsedURL, err = url.Parse(aURL)
		}
		if err != nil {
			return "", err
		}
	}
	return parsedURL.Redacted(), nil
}
