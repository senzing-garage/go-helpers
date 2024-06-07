/*
Package engineconfigurationjsonparser is used to generate the JSON document used to configure a Senzing client.
*/
package settingsparser

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// BasicEngineConfigurationJSONParser is the default implementation of the EngineConfigurationJsonParser interface.
type BasicEngineConfigurationJSONParser struct {
	EngineConfigurationJSON string
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

// ----------------------------------------------------------------------------
// Constructor  methods
// ----------------------------------------------------------------------------

func New(engineConfigurationJSON string) (EngineConfigurationJSONParser, error) {
	var err error
	if !isJSON(engineConfigurationJSON) {
		return nil, fmt.Errorf("incorrect JSON syntax in %s", engineConfigurationJSON)
	}
	result := &BasicEngineConfigurationJSONParser{
		EngineConfigurationJSON: engineConfigurationJSON,
	}
	return result, err
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
func (parser *BasicEngineConfigurationJSONParser) GetConfigPath(ctx context.Context) (string, error) {
	_ = ctx
	engineConfiguration := &EngineConfiguration{}

	err := json.Unmarshal([]byte(parser.EngineConfigurationJSON), &engineConfiguration)
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
func (parser *BasicEngineConfigurationJSONParser) GetDatabaseUrls(ctx context.Context) ([]string, error) {
	_ = ctx
	var result []string

	engineConfiguration := &EngineConfiguration{}
	err := json.Unmarshal([]byte(parser.EngineConfigurationJSON), &engineConfiguration)
	if err != nil {
		return result, err
	}
	result = append(result, engineConfiguration.SQL.Connection)

	// Handle multi-database case.

	backend := engineConfiguration.SQL.Backend
	if len(backend) > 0 && backend != "SQL" {
		var dictionary map[string]interface{}
		var databaseJSONKeys []string
		err = json.Unmarshal([]byte(parser.EngineConfigurationJSON), &dictionary)
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
The GetResourcePath method returns the PIPELINE.RESOURCEPATH value of _ENGINE_CONFIGURATION_JSON.

Input
  - ctx: A context to control lifecycle.

Output
  - A string containing the value of a PIPELINE.RESOURCEPATH.
*/
func (parser *BasicEngineConfigurationJSONParser) GetResourcePath(ctx context.Context) (string, error) {
	_ = ctx
	engineConfiguration := &EngineConfiguration{}
	err := json.Unmarshal([]byte(parser.EngineConfigurationJSON), &engineConfiguration)
	if err != nil {
		return "", err
	}
	return engineConfiguration.Pipeline.ResourcePath, err
}

/*
The GetSupportPath method returns the PIPELINE.SUPPORTPATH value of _ENGINE_CONFIGURATION_JSON.

Input
  - ctx: A context to control lifecycle.

Output
  - A string containing the value of a PIPELINE.SUPPORTPATH.
*/
func (parser *BasicEngineConfigurationJSONParser) GetSupportPath(ctx context.Context) (string, error) {
	_ = ctx
	engineConfiguration := &EngineConfiguration{}
	err := json.Unmarshal([]byte(parser.EngineConfigurationJSON), &engineConfiguration)
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
func (parser *BasicEngineConfigurationJSONParser) RedactedJSON(ctx context.Context) (string, error) {
	result := parser.EngineConfigurationJSON

	// Get list of database URLs in the Senzing engine configuration json.

	databaseUrls, err := parser.GetDatabaseUrls(ctx)
	if err != nil {
		return "", err
	}

	// For each database URL in the string, replace it with a redacted database URL.

	for _, databaseURL := range databaseUrls {
		redactedURL, err := redactURL(databaseURL)
		if err == nil {
			result = strings.ReplaceAll(result, databaseURL, redactedURL)
		}
	}

	// Remove whitespace.

	result = strings.Join(strings.Fields(result), "")
	return result, err
}
