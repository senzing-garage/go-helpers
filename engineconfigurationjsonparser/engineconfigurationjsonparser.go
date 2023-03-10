/*
Package engineconfigurationjsonparser is used to generate the JSON document used to configure a Senzing client.
*/
package engineconfigurationjsonparser

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

// EngineConfigurationJsonParserImpl is the default implementation of the EngineConfigurationJsonParser interface.
type EngineConfigurationJsonParserImpl struct {
	EngineConfigurationJson string
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

func isJson(unknownString string) bool {
	unknownStringUnescaped, err := strconv.Unquote(unknownString)
	if err != nil {
		unknownStringUnescaped = unknownString
	}
	var jsonString json.RawMessage
	return json.Unmarshal([]byte(unknownStringUnescaped), &jsonString) == nil
}

func redactUrl(aUrl string) (string, error) {
	parsedUrl, err := url.Parse(aUrl)
	if err != nil {
		if strings.HasPrefix(aUrl, "postgresql") {
			index := strings.LastIndex(aUrl, ":")
			aUrl := aUrl[:index] + "/" + aUrl[index+1:]
			parsedUrl, err = url.Parse(aUrl)
		}
		if err != nil {
			return "", err
		}
	}
	return parsedUrl.Redacted(), nil
}

// ----------------------------------------------------------------------------
// Constructor  methods
// ----------------------------------------------------------------------------

func New(engineConfigurationJson string) (EngineConfigurationJsonParser, error) {
	var err error = nil
	if !isJson(engineConfigurationJson) {
		return nil, fmt.Errorf("incorrect JSON syntax in %s", engineConfigurationJson)
	}
	result := &EngineConfigurationJsonParserImpl{
		EngineConfigurationJson: engineConfigurationJson,
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
func (parser *EngineConfigurationJsonParserImpl) GetConfigPath(ctx context.Context) (string, error) {
	engineConfiguration := &EngineConfiguration{}

	err := json.Unmarshal([]byte(parser.EngineConfigurationJson), &engineConfiguration)
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
func (parser *EngineConfigurationJsonParserImpl) GetDatabaseUrls(ctx context.Context) ([]string, error) {
	var result []string

	engineConfiguration := &EngineConfiguration{}
	err := json.Unmarshal([]byte(parser.EngineConfigurationJson), &engineConfiguration)
	if err != nil {
		return result, err
	}
	result = append(result, engineConfiguration.Sql.Connection)

	// Handle multi-database case.

	backend := engineConfiguration.Sql.Backend
	if len(backend) > 0 && backend != "SQL" {
		var dictionary map[string]interface{}
		var databaseJsonKeys []string
		err = json.Unmarshal([]byte(parser.EngineConfigurationJson), &dictionary)
		if err != nil {
			return result, err
		}

		// Determine JSON keys for database definitions.

		backendMap := dictionary[backend]
		for _, value := range backendMap.(map[string]interface{}) {
			valueString := value.(string)
			if !contains(databaseJsonKeys, valueString) {
				databaseJsonKeys = append(databaseJsonKeys, valueString)
			}
		}

		// Add each database.

		for _, databaseJsonKey := range databaseJsonKeys {
			databaseJson := dictionary[databaseJsonKey].(map[string]interface{})
			databaseName := databaseJson["DB_1"].(string)
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
func (parser *EngineConfigurationJsonParserImpl) GetResourcePath(ctx context.Context) (string, error) {
	engineConfiguration := &EngineConfiguration{}
	err := json.Unmarshal([]byte(parser.EngineConfigurationJson), &engineConfiguration)
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
func (parser *EngineConfigurationJsonParserImpl) GetSupportPath(ctx context.Context) (string, error) {
	engineConfiguration := &EngineConfiguration{}
	err := json.Unmarshal([]byte(parser.EngineConfigurationJson), &engineConfiguration)
	if err != nil {
		return "", err
	}
	return engineConfiguration.Pipeline.SupportPath, err
}

/*
The RedactedJson method returns the JSON string with passwords redacted.

Input
  - ctx: A context to control lifecycle.

Output
  - The Senzing engine configuration JSON string with database URLs having redacted passwords.
*/
func (parser *EngineConfigurationJsonParserImpl) RedactedJson(ctx context.Context) (string, error) {
	result := parser.EngineConfigurationJson

	// Get list of database URLs in the Senzing engine configuration json.

	databaseUrls, err := parser.GetDatabaseUrls(ctx)
	if err != nil {
		return "", err
	}

	// For each database URL in the string, replace it with a redacted database URL.

	for _, databaseUrl := range databaseUrls {
		redactedUrl, err := redactUrl(databaseUrl)
		if err == nil {
			result = strings.Replace(result, databaseUrl, redactedUrl, -1)
		}
	}

	// Remove whitespace.

	result = strings.Join(strings.Fields(result), "")
	return result, err
}
