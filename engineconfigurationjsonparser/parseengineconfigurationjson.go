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
// Internal methods
// ----------------------------------------------------------------------------

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

/*
The BuildSimpleSystemConfigurationJson method returns a JSON document for use with Senzing's Init(...) methods.
The configuration is for a "system install" with a single database.

Input
  - senzingDatabaseUrl: A Database URL.
    If empty, the SENZING_ENGINE_CONFIGURATION_JSON and SENZING_TOOLS_DATABASE_URL environment variables will be used in calculating the result.

Output
  - A string containing a JSON document use when calling Senzing's Init(...) methods.
    See the example output.
*/
func (parser *EngineConfigurationJsonParserImpl) GetConfigPath(ctx context.Context) (string, error) {

	engineConfiguration := &EngineConfiguration{}
	err := json.Unmarshal([]byte(parser.EnableConfigurationJson), &engineConfiguration)
	if err != nil {
		return "", err
	}

	return engineConfiguration.Pipeline.ConfigPath, err
}
