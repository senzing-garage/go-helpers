//go:build darwin

package g2engineconfigurationjson

// ----------------------------------------------------------------------------
// Internal methods
// ----------------------------------------------------------------------------

func buildStruct(specificDatabaseUrl string, licenseStringBase64 string, senzingDirectory string) G2Configuration {
	var result G2Configuration

	if len(licenseStringBase64) > 0 {
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
