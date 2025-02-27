//go:build linux

package settings

import (
	"context"
	"fmt"
)

// ----------------------------------------------------------------------------
// Internal methods
// ----------------------------------------------------------------------------

func getConfigPath(senzingDirectory string) string {
	_ = senzingDirectory
	return "/etc/opt/senzing"
}

func getResourcePath(senzingDirectory string) string {
	return fmt.Sprintf("%s/er/resources", senzingDirectory)
}

func getSenzingDirectory(attributeMap map[string]string) string {
	result := "/opt/senzing"
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
