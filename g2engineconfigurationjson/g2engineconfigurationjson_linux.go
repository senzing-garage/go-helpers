//go:build linux

package g2engineconfigurationjson

import "context"

// ----------------------------------------------------------------------------
// Internal methods
// ----------------------------------------------------------------------------

func buildStruct(attributeMap map[string]string) G2Configuration {
	var result G2Configuration

	databaseUrl, ok := attributeMap["databaseUrl"]
	if !ok {
		return result
	}

	result = G2Configuration{
		Pipeline: G2ConfigurationPipeline{
			ConfigPath:   "/etc/opt/senzing",
			ResourcePath: "/opt/senzing/g2/resources",
			SupportPath:  "/opt/senzing/data",
		},
		Sql: G2ConfigurationSql{
			Connection: databaseUrl,
		},
	}

	licenseStringBase64, ok := attributeMap["licenseStringBase64"]
	if ok {
		result.Pipeline.LicenseStringBase64 = licenseStringBase64
	}

	return result
}

func verifySenzingEngineConfigurationJson(ctx context.Context, senzingEngineConfigurationJson string) error {
	var err error = nil
	return err
}
