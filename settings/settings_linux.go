//go:build linux

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

	senzingDirectory := "/opt/senzing"
	senzingPath, ok := attributeMap["senzingPath"]
	if ok {
		senzingDirectory = senzingPath
	}

	configPath := "/etc/opt/senzing"
	resourcePath := fmt.Sprintf("%s/er/resources", senzingDirectory)
	supportPath := fmt.Sprintf("%s/data", senzingDirectory)

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
	_ = settings
	_ = ctx
	var err error
	return err
}
