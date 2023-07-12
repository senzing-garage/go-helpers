// The g2engineconfigurationjson package helps configure SENZING_ENGINE_CONFIGURATION_JSON.
package g2engineconfigurationjson

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

type G2ConfigurationPipeline struct {
	ConfigPath          string `json:"CONFIGPATH"`
	LicenseStringBase64 string `json:"LICENSESTRINGBASE64,omitempty"`
	ResourcePath        string `json:"RESOURCEPATH"`
	SupportPath         string `json:"SUPPORTPATH"`
}

type G2ConfigurationSql struct {
	Connection string `json:"CONNECTION"`
}

type G2Configuration struct {
	Pipeline G2ConfigurationPipeline `json:"PIPELINE"`
	Sql      G2ConfigurationSql      `json:"SQL"`
}

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

// Identfier of the  package found messages having the format "senzing-6402xxxx".
const ProductId = 6402
