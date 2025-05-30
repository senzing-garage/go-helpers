// The settingsparser package helps parse the _ENGINE_CONFIGURATION_JSON.
package settingsparser

import (
	"context"
	"errors"

	"github.com/senzing-garage/go-helpers/wraperror"
)

// ----------------------------------------------------------------------------
// Types - interface
// ----------------------------------------------------------------------------

type SettingsParser interface {
	GetConfigPath(ctx context.Context) (string, error)
	GetDatabaseURIs(ctx context.Context) ([]string, error)
	GetLicenseStringBase64(ctx context.Context) (string, error)
	GetResourcePath(ctx context.Context) (string, error)
	GetSettings(ctx context.Context) string
	GetSupportPath(ctx context.Context) (string, error)
	RedactedJSON(ctx context.Context) (string, error)
}

// ----------------------------------------------------------------------------
// Types - struct
// ----------------------------------------------------------------------------

type EngineConfigurationPipeline struct {
	ConfigPath          string `json:"CONFIGPATH"`
	LicenseStringBase64 string `json:"LICENSESTRINGBASE64"`
	ResourcePath        string `json:"RESOURCEPATH"`
	SupportPath         string `json:"SUPPORTPATH"`
}

type EngineConfigurationSQL struct {
	Backend    string `json:"BACKEND"`
	Connection string `json:"CONNECTION"`
}

type EngineConfiguration struct {
	Pipeline EngineConfigurationPipeline `json:"PIPELINE"`
	SQL      EngineConfigurationSQL      `json:"SQL"`
}

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

// Identfier of the  package found messages having the format "SZSDK6401xxxx".
const ComponentID = 6401

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

var errForPackage = errors.New("settingsparser")

// ----------------------------------------------------------------------------
// Constructor  methods
// ----------------------------------------------------------------------------

func New(settings string) (SettingsParser, error) {
	var err error

	if !isJSON(settings) {
		return nil, wraperror.Errorf(errForPackage, "incorrect JSON syntax in %s", settings)
	}

	result := &BasicSettingsParser{
		Settings: settings,
	}

	return result, err
}
