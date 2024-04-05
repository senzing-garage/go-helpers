//go:build linux

package engineconfigurationjson

var testCasesForOsArch = []testCaseMetadata{
	{
		name:            "sqlite3-001",
		databaseUrl:     "sqlite3://na:na@/var/opt/senzing/sqlite/G2C.db",
		databaseUrlPath: "/var/opt/senzing/sqlite/G2C.db",
	},
	{
		name:            "sqlite3-002",
		databaseUrl:     `sqlite3://na:na@hostname/var/opt/senzing/sqlite/G2C.db`,
		databaseUrlPath: "/var/opt/senzing/sqlite/G2C.db",
	},
}
