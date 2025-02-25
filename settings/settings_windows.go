//go:build windows

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
	return fmt.Sprintf("%s%cer%cetc", senzingDirectory, os.PathSeparator, os.PathSeparator)
}

func getResourcePath(senzingDirectory string) string {
	return fmt.Sprintf("%s%cer%cresources", senzingDirectory, os.PathSeparator, os.PathSeparator)
}

func getSenzingDirectory(attributeMap map[string]string) string {
	result := `C:\Program Files\senzing\er`
	homeDrive, isHomeDriveSet := os.LookupEnv("HOMEDRIVE")
	homeDir, isHomeDirSet := os.LookupEnv("HOMEDIR")
	if isHomeDriveSet && isHomeDirSet {
		result = fmt.Sprintf("%s%s/senzing", homeDrive, homeDir)
	}
	senzingPath, ok := attributeMap["senzingPath"]
	if ok {
		result = senzingPath
	}
	return result
}

func getSupportPath(senzingDirectory string) string {
	return fmt.Sprintf("%s%cer%cdata", senzingDirectory, os.PathSeparator, os.PathListSeparator)
}

func verifySettings(ctx context.Context, settings string) error {
	_ = ctx
	_ = settings
	var err error
	return err
}
