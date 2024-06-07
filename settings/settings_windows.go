//go:build windows

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

	databaseURL, ok := attributeMap["databaseUrl"]
	if !ok {
		return result
	}

	// Construct directories based on senzingDirectory.

	senzingDirectory, ok := attributeMap["senzingDirectory"]
	if !ok {
		senzingDirectory = `C:\Program Files\Senzing\g2`
	}
	configPath := fmt.Sprintf("%s%cetc", senzingDirectory, os.PathSeparator)
	resourcePath := fmt.Sprintf("%s%cresources", senzingDirectory, os.PathSeparator)
	supportPath := fmt.Sprintf("%s%cdata", senzingDirectory, os.PathSeparator)

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

	licenseStringBase64, inMap := attributeMap["licenseStringBase64"]
	if inMap {
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
