// The engineconfigurationjsonparser package helps parse the _ENGINE_CONFIGURATION_JSON.
package engineconfigurationjsonparser

import "context"

// ----------------------------------------------------------------------------
// Types - interface
// ----------------------------------------------------------------------------

type EngineConfigurationJsonParser interface {
	GetDatabaseUrls(ctx context.Context) ([]string, error)
	GetConfigPath(ctx context.Context) (string, error)
	GetResourcePath(ctx context.Context) (string, error)
	GetSupportPath(ctx context.Context) (string, error)
}

// ----------------------------------------------------------------------------
// Types - struct
// ----------------------------------------------------------------------------

type G2ConfigurationPipeline struct {
	ConfigPath   string `json:"CONFIGPATH"`
	ResourcePath string `json:"RESOURCEPATH"`
	SupportPath  string `json:"SUPPORTPATH"`
}

type G2ConfigurationSql struct {
	Connection string `json:"CONNECTION"`
}

type G2Configuration struct {
	Pipeline G2ConfigurationPipeline `json:"PIPELINE"`
	Sql      G2ConfigurationSql      `json:"SQL"`
}
