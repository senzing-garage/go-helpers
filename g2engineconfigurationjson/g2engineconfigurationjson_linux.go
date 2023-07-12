//go:build linux

package g2engineconfigurationjson

// ----------------------------------------------------------------------------
// Internal methods
// ----------------------------------------------------------------------------

// func buildStruct(specificDatabaseUrl string, licenseStringBase64 string, senzingDirectory string) G2Configuration {
func buildStruct(attributeMap map[string]string) G2Configuration {
	var result G2Configuration

	databaseUrl, inMap := attributeMap["databaseUrl"]
	if !inMap {
		return result
	}

	licenseStringBase64, inMap := attributeMap["licenseStringBase64"]
	if inMap {
		result = G2Configuration{
			Pipeline: G2ConfigurationPipeline{
				ConfigPath:          "/etc/opt/senzing",
				LicenseStringBase64: licenseStringBase64,
				ResourcePath:        "/opt/senzing/g2/resources",
				SupportPath:         "/opt/senzing/data",
			},
			Sql: G2ConfigurationSql{
				Connection: databaseUrl,
			},
		}
	} else {
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
	}

	return result
}
