//go:build windows

package settings

var testCasesForOsArch = []testCaseMetadata{
	{
		name:            "sqlite3-001",
		databaseURL:     "sqlite3://na:na@nowhere/C:\\Temp\\sqlite\\G2C.db",
		databaseURLPath: "/C:\\Temp\\sqlite\\G2C.db",
	},
	{
		name:            "sqlite3-002",
		databaseURL:     "sqlite3://na:na@nowhere/C:\\Temp\\sqlite\\G2C.db",
		databaseURLPath: `/C:\Temp\sqlite\G2C.db`,
	},
	{
		name:            "sqlite3-003",
		databaseURL:     `sqlite3://na:na@nowhere/C:\Temp\sqlite\G2C.db`,
		databaseURLPath: "/C:\\Temp\\sqlite\\G2C.db",
	},
	{
		name:            "sqlite3-004",
		databaseURL:     `sqlite3://na:na@nowhere/C:\Temp\sqlite\G2C.db`,
		databaseURLPath: `/C:\Temp\sqlite\G2C.db`,
	},
}
