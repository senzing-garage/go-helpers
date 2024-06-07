//go:build linux

package settings

import "context"

// ----------------------------------------------------------------------------
// Internal methods
// ----------------------------------------------------------------------------

func buildStruct(attributeMap map[string]string) SzConfiguration {
	var result SzConfiguration

	databaseURL, ok := attributeMap["databaseUrl"]
	if !ok {
		return result
	}

	result = SzConfiguration{
		Pipeline: SzConfigurationPipeline{
			ConfigPath:   "/etc/opt/senzing",
			ResourcePath: "/opt/senzing/g2/resources",
			SupportPath:  "/opt/senzing/data",
		},
		SQL: SzConfigurationSQL{
			Connection: databaseURL,
		},
	}

	licenseStringBase64, ok := attributeMap["licenseStringBase64"]
	if ok {
		result.Pipeline.LicenseStringBase64 = licenseStringBase64
	}

	return result
}

func verifySettings(ctx context.Context, settings string) error {
	_ = settings
	_ = ctx
	var err error
	return err
}
