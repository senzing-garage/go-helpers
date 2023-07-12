package g2engineconfigurationjson

// ----------------------------------------------------------------------------
// Internal methods
// ----------------------------------------------------------------------------

func buildStruct(specificDatabaseUrl string, licenseStringBase64 string) G2Configuration {

	var result G2Configuration

	if len(licenseStringBase64) > 0 {
		result = G2Configuration{
			Pipeline: G2ConfigurationPipeline{
				ConfigPath:          "/etc/opt/senzing",
				LicenseStringBase64: licenseStringBase64,
				ResourcePath:        "/opt/senzing/g2/resources",
				SupportPath:         "/opt/senzing/data",
			},
			Sql: G2ConfigurationSql{
				Connection: specificDatabaseUrl,
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
				Connection: specificDatabaseUrl,
			},
		}
	}

	return result
}
