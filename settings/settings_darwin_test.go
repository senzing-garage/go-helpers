//go:build darwin

package settings

import (
	"fmt"
	"os"
)

var testCasesForOsArch = []testCaseMetadata{
	{
		name:            "sqlite3-001",
		databaseURL:     "sqlite3://na:na@/var/opt/senzing/sqlite/G2C.db",
		databaseURLPath: "/var/opt/senzing/sqlite/G2C.db",
	},
	{
		name:            "sqlite3-002",
		databaseURL:     `sqlite3://na:na@hostname/var/opt/senzing/sqlite/G2C.db`,
		databaseURLPath: "/var/opt/senzing/sqlite/G2C.db",
	},
}

func getSenzingPath() string {
	var result string
	home, isSet := os.LookupEnv("HOME")
	if isSet {
		result = fmt.Sprintf("%s/senzing", home)
	}
	return result
}
