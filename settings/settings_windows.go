//go:build windows

package engineconfigurationjson

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

	databaseUrl, ok := attributeMap["databaseUrl"]
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
		Sql: SzConfigurationSql{
			Connection: databaseUrl,
		},
	}

	licenseStringBase64, inMap := attributeMap["licenseStringBase64"]
	if inMap {
		result.Pipeline.LicenseStringBase64 = licenseStringBase64
	}

	return result
}

func verifySenzingEngineConfigurationJson(ctx context.Context, engineConfigurationJson string) error {
	_ = ctx
	_ = engineConfigurationJson
	var err error = nil
	return err
}
