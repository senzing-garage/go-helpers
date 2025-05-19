/*
Package settingsparser is used to generate the JSON document used to configure a Senzing client.
*/
package settingsparser

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/senzing-garage/go-helpers/wraperror"
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
		return "", wraperror.Errorf(err, "GetConfigPath")
	}

	// return engineConfiguration.Pipeline.ConfigPath, nil

	return engineConfiguration.Pipeline.ConfigPath, wraperror.Errorf(err, "")
}

/*
The GetConfigPath method returns the PIPELINE.CONFIGPATH value of _ENGINE_CONFIGURATION_JSON.

Input
  - ctx: A context to control lifecycle.

Output
  - A string containing the value of a PIPELINE.CONFIGPATH.
*/
func (parser *BasicSettingsParser) GetDatabaseURIs(ctx context.Context) ([]string, error) {
	_ = ctx

	var result []string

	engineConfiguration := &EngineConfiguration{}

	err := json.Unmarshal([]byte(parser.Settings), &engineConfiguration)
	if err != nil {
		return result, wraperror.Errorf(err, "GetDatabaseURIs")
	}

	result = append(result, engineConfiguration.SQL.Connection)

	// Handle multi-database case.

	backend := engineConfiguration.SQL.Backend
	if (len(backend) > 0) && (backend != "SQL") {
		multiDatabaseURIs, err := getMultiDatabaseURIs(ctx, parser.Settings, backend)
		if err != nil {
			return result, wraperror.Errorf(err, "getMultiDatabaseURIs")
		}

		result = append(result, multiDatabaseURIs...)
	}

	// IMPROVE:  Implement multi-database list.

	// return result, nil

	return result, wraperror.Errorf(err, "")
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
		return "", wraperror.Errorf(err, "Unmarshal")
	}

	// return engineConfiguration.Pipeline.LicenseStringBase64, nil

	return engineConfiguration.Pipeline.LicenseStringBase64, wraperror.Errorf(err, "")
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
		return "", wraperror.Errorf(err, "Unmarshal")
	}

	return engineConfiguration.Pipeline.ResourcePath, wraperror.Errorf(err, wraperror.NoMessage)
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
		return "", wraperror.Errorf(err, "Unmarshal")
	}

	return engineConfiguration.Pipeline.SupportPath, wraperror.Errorf(err, wraperror.NoMessage)
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

	databaseURIs, err := parser.GetDatabaseURIs(ctx)
	if err != nil {
		return "", wraperror.Errorf(err, "GetDatabaseURIs")
	}

	// For each database URL in the string, replace it with a redacted database URL.

	for _, databaseURI := range databaseURIs {
		redactedURL, err := redactURL(databaseURI)
		if err == nil {
			result = strings.ReplaceAll(result, databaseURI, redactedURL)
		}
	}

	// Remove whitespace.

	result = strings.Join(strings.Fields(result), "")

	return result, wraperror.Errorf(err, wraperror.NoMessage)
}

// ----------------------------------------------------------------------------
// Private functions
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
		if strings.HasPrefix(aURL, "postgresql://") {
			index := strings.LastIndex(aURL, ":")
			aURL := aURL[:index] + "/" + aURL[index+1:]

			parsedURL, err = url.Parse(aURL)
			if err != nil {
				return "", wraperror.Errorf(err, "could not parse for postgresql:// %s", aURL)
			}
		}
	}

	return parsedURL.Redacted(), wraperror.Error(err)
}

func getMultiDatabaseURIs(ctx context.Context, settings string, backend string) ([]string, error) {
	var err error

	_ = ctx

	result := []string{}

	if (len(backend) > 0) && (backend != "SQL") { //nolint:nestif
		var dictionary map[string]interface{}

		var databaseJSONKeys []string

		err = json.Unmarshal([]byte(settings), &dictionary)
		if err != nil {
			return result, wraperror.Errorf(err, "Unmarshal")
		}

		// Determine JSON keys for database definitions.

		backendMap := dictionary[backend]

		backendMapTyped, isOK := backendMap.(map[string]interface{})
		if !isOK {
			panic(fmt.Sprintf("failed type assertion for %v.(map[string]interface{})", backendMap))
		}

		for _, value := range backendMapTyped {
			valueString, isOK := value.(string)
			if !isOK {
				panic(fmt.Sprintf("failed type assertion for %v.(string)", value))
			}

			if !contains(databaseJSONKeys, valueString) {
				databaseJSONKeys = append(databaseJSONKeys, valueString)
			}
		}

		// Add each database.

		for _, databaseJSONKey := range databaseJSONKeys {
			databaseJSON, isOK := dictionary[databaseJSONKey].(map[string]interface{})
			if !isOK {
				panic(fmt.Sprintf("failed type assertion for dictionary[%s].(map[string]interface{}", databaseJSONKey))
			}

			databaseName, isOK := databaseJSON["DB_1"].(string)
			if !isOK {
				panic(`failed type assertion for databaseJSON["DB_1"].(string)`)
			}

			if !contains(result, databaseName) {
				result = append(result, databaseName)
			}
		}
	}

	return result, wraperror.Errorf(err, wraperror.NoMessage)
}
