//go:build linux

package g2engineconfigurationjson

// ----------------------------------------------------------------------------
// Internal methods
// ----------------------------------------------------------------------------

// func buildStruct(specificDatabaseUrl string, licenseStringBase64 string, senzingDirectory string) G2Configuration {
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
