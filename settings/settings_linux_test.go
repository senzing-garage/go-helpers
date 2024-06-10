//go:build linux

package settings

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
