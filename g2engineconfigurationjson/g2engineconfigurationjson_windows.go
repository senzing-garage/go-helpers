//go:build windows

package g2engineconfigurationjson

// ----------------------------------------------------------------------------
// Internal methods
// ----------------------------------------------------------------------------

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
				ConfigPath:          "/path/to/config",
				LicenseStringBase64: licenseStringBase64,
				ResourcePath:        "/path/to/resources",
				SupportPath:         "/path/to/data",
			},
			Sql: G2ConfigurationSql{
				Connection: specificDatabaseUrl,
			},
		}
	} else {
		result = G2Configuration{
			Pipeline: G2ConfigurationPipeline{
				ConfigPath:   "/path/to/config",
				ResourcePath: "/path/to/resources",
				SupportPath:  "/path/to/data",
			},
			Sql: G2ConfigurationSql{
				Connection: specificDatabaseUrl,
			},
		}
	}

	return result
}
