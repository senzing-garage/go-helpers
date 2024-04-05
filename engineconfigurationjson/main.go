// The engineconfigurationjson package helps configure SENZING_ENGINE_CONFIGURATION_JSON.
package engineconfigurationjson

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

type SzConfigurationPipeline struct {
	ConfigPath          string `json:"CONFIGPATH"`
	LicenseStringBase64 string `json:"LICENSESTRINGBASE64,omitempty"`
	ResourcePath        string `json:"RESOURCEPATH"`
	SupportPath         string `json:"SUPPORTPATH"`
}

type SzConfigurationSql struct {
	Connection string `json:"CONNECTION"`
}

type SzConfiguration struct {
	Pipeline SzConfigurationPipeline `json:"PIPELINE"`
	Sql      SzConfigurationSql      `json:"SQL"`
}

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

// Identfier of the  package found messages having the format "senzing-6402xxxx".
const ProductId = 6402
