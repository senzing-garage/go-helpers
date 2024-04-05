//go:build linux

package engineconfigurationjson

import "context"

// ----------------------------------------------------------------------------
// Internal methods
// ----------------------------------------------------------------------------

func buildStruct(attributeMap map[string]string) SzConfiguration {
	var result SzConfiguration

	databaseUrl, ok := attributeMap["databaseUrl"]
	if !ok {
		return result
	}

	result = SzConfiguration{
		Pipeline: SzConfigurationPipeline{
			ConfigPath:   "/etc/opt/senzing",
			ResourcePath: "/opt/senzing/g2/resources",
			SupportPath:  "/opt/senzing/data",
		},
		Sql: SzConfigurationSql{
			Connection: databaseUrl,
		},
	}

	licenseStringBase64, ok := attributeMap["licenseStringBase64"]
	if ok {
		result.Pipeline.LicenseStringBase64 = licenseStringBase64
	}

	return result
}

func verifySenzingEngineConfigurationJson(ctx context.Context, engineConfigurationJson string) error {
	_ = engineConfigurationJson
	_ = ctx
	var err error = nil
	return err
}
