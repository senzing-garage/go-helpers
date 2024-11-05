//go:build darwin

package settings

import (
	"context"
	"fmt"
)

// ----------------------------------------------------------------------------
// Internal methods
// ----------------------------------------------------------------------------

func mapWithDefault(aMap map[string]string, key string, defaultValue string) string {
	result, ok := aMap[key]
	if ok {
		return result
	}
	return defaultValue
}

func buildStruct(attributeMap map[string]string) SzConfiguration {
	var result SzConfiguration

	databaseURL, ok := attributeMap["databaseURL"]
	if !ok {
		return result
	}

	// Determine defaultDirectory.

	defaultDirectory := "/opt/senzing/er"
	senzingDirectory, ok := attributeMap["senzingDirectory"]
	if ok {
		defaultDirectory = senzingDirectory
	}

	configPath := fmt.Sprintf("%s/etc", defaultDirectory)
	resourcePath := fmt.Sprintf("%s/resources", defaultDirectory)
	supportPath := fmt.Sprintf("%s/data", defaultDirectory)

	// Apply attributeMap.

	result = SzConfiguration{
		Pipeline: SzConfigurationPipeline{
			ConfigPath:   mapWithDefault(attributeMap, "configPath", configPath),
			ResourcePath: mapWithDefault(attributeMap, "resourcePath", resourcePath),
			SupportPath:  mapWithDefault(attributeMap, "supportPath", supportPath),
		},
		SQL: SzConfigurationSQL{
			Connection: databaseURL,
		},
	}

	licenseStringBase64, ok := attributeMap["licenseStringBase64"]
	if ok {
		result.Pipeline.LicenseStringBase64 = licenseStringBase64
	}

	return result
}

func verifySettings(ctx context.Context, settings string) error {
	_ = ctx
	_ = settings
	var err error = nil
	return err
}
