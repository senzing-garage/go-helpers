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

	databaseURL, ok := attributeMap["databaseURL"]
	if !ok {
		return result
	}

	// Determine defaultDirectory.

	senzingDirectory := `C:\Program Files\senzing\er`
	homeDrive, isHomeDriveSet := os.LookupEnv("HOMEDRIVE")
	homeDir, isHomeDirSet := os.LookupEnv("HOMEDIR")
	if isHomeDriveSet && isHomeDirSet {
		senzingDirectory = fmt.Sprintf("%s%s/senzing", homeDrive, homeDir)
	}

	senzingPath, ok := attributeMap["senzingPath"]
	if ok {
		senzingDirectory = senzingPath
	}

	// Construct directories based on senzingDirectory.

	configPath := fmt.Sprintf("%s%cer%cetc", senzingDirectory, os.PathSeparator, os.PathSeparator)
	resourcePath := fmt.Sprintf("%s%cer%cresources", senzingDirectory, os.PathSeparator, os.PathSeparator)
	supportPath := fmt.Sprintf("%s%cer%cdata", senzingDirectory, os.PathSeparator, os.PathListSeparator)

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
