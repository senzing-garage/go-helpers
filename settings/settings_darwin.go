//go:build darwin

package settings

import (
	"context"
	"fmt"
	"os"
)

// ----------------------------------------------------------------------------
// Internal methods
// ----------------------------------------------------------------------------

func getConfigPath(senzingDirectory string) string {
	return senzingDirectory + "/er/etc"
}

func getResourcePath(senzingDirectory string) string {
	return senzingDirectory + "/er/resources"
}

func getSenzingDirectory(attributeMap map[string]string) string {
	result := "/opt/senzing"

	home, isSet := os.LookupEnv("HOME")
	if isSet {
		result = home + "/senzing"
	}

	senzingPath, ok := attributeMap["senzingPath"]
	if ok {
		result = senzingPath
	}

	return result
}

func getSupportPath(senzingDirectory string) string {
	return fmt.Sprintf("%s/data", senzingDirectory)
}

func verifySettings(ctx context.Context, settings string) error {
	_ = ctx
	_ = settings

	var err error

	return err
}
