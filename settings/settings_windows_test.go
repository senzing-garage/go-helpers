//go:build windows

package settings

import (
	"fmt"
	"os"
)

var testCasesForOsArch = []testCaseMetadata{
	{
		name:            "sqlite3-001",
		databaseURL:     "sqlite3://na:na@nowhere/C:\\Temp\\sqlite\\G2C.db",
		databaseURLPath: "/C:\\Temp\\sqlite\\G2C.db",
		databaseURI:     "sqlite3://na:na@nowhere/C:\\Temp\\sqlite\\G2C.db",
		isReversible:    true,
	},
	{
		name:            "sqlite3-002",
		databaseURL:     "sqlite3://na:na@nowhere/C:\\Temp\\sqlite\\G2C.db",
		databaseURLPath: `/C:\Temp\sqlite\G2C.db`,
		databaseURI:     "sqlite3://na:na@nowhere/C:\\Temp\\sqlite\\G2C.db",
		isReversible:    true,
	},
	{
		name:            "sqlite3-003",
		databaseURL:     `sqlite3://na:na@nowhere/C:\Temp\sqlite\G2C.db`,
		databaseURLPath: "/C:\\Temp\\sqlite\\G2C.db",
		databaseURI:     `sqlite3://na:na@nowhere/C:\Temp\sqlite\G2C.db`,
		isReversible:    true,
	},
	{
		name:            "sqlite3-004",
		databaseURL:     `sqlite3://na:na@nowhere/C:\Temp\sqlite\G2C.db`,
		databaseURLPath: `/C:\Temp\sqlite\G2C.db`,
		databaseURI:     `sqlite3://na:na@nowhere/C:\Temp\sqlite\G2C.db`,
		isReversible:    true,
	},
}

func getSenzingPath() string {
	var result string
	homeDrive, isHomeDriveSet := os.LookupEnv("HOMEDRIVE")
	homePath, isHomeDirSet := os.LookupEnv("HOMEPATH")
	if isHomeDriveSet && isHomeDirSet {
		result = fmt.Sprintf("%s%s\\Senzing", homeDrive, homePath)
	}
	return result
}
