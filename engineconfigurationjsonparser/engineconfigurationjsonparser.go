/*
Package engineconfigurationjsonparser is used to generate the JSON document used to configure a Senzing client.
*/
package engineconfigurationjsonparser

import (
	"context"
	"encoding/json"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// EngineConfigurationJsonParserImpl is the default implementation of the EngineConfigurationJsonParser interface.
type EngineConfigurationJsonParserImpl struct {
	EnableConfigurationJson string
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
	err := json.Unmarshal([]byte(parser.EnableConfigurationJson), &engineConfiguration)
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
	err := json.Unmarshal([]byte(parser.EnableConfigurationJson), &engineConfiguration)
	if err != nil {
		return result, err
	}
	backend := engineConfiguration.Sql.Backend
	if len(backend) == 0 || backend == "SQL" {
		result = append(result, engineConfiguration.Sql.Connection)
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
	err := json.Unmarshal([]byte(parser.EnableConfigurationJson), &engineConfiguration)
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
	err := json.Unmarshal([]byte(parser.EnableConfigurationJson), &engineConfiguration)
	if err != nil {
		return "", err
	}
	return engineConfiguration.Pipeline.SupportPath, err
}
