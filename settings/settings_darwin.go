//go:build darwin

package settings

import (
	"context"
	"fmt"
	"os"
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
	home, isSet := os.LookupEnv("HOME")
	if isSet {
		senzingDirectory = fmt.Sprintf("%s/senzing", home)
	}
	senzingPath, ok := attributeMap["senzingPath"]
	if ok {
		senzingDirectory = senzingPath
	}

	configPath := fmt.Sprintf("%s/er/etc", senzingDirectory)
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
	_ = ctx
	_ = settings
	var err error
	return err
}
