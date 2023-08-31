//go:build linux

package g2engineconfigurationjson

import (
	"fmt"
)

// ----------------------------------------------------------------------------
// Examples for godoc documentation
// ----------------------------------------------------------------------------

func ExampleBuildSimpleSystemConfigurationJsonUsingMap() {
	aMap := map[string]string{
		"databaseUrl": "postgresql://username:password@hostname:5432/G2",
	}
	result, err := BuildSimpleSystemConfigurationJsonUsingMap(aMap)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	// Output: {"PIPELINE":{"CONFIGPATH":"/etc/opt/senzing","RESOURCEPATH":"/opt/senzing/g2/resources","SUPPORTPATH":"/opt/senzing/data"},"SQL":{"CONNECTION":"postgresql://username:password@hostname:5432:G2/"}}
}

func ExampleBuildSimpleSystemConfigurationJsonUsingMap_db2() {
	aMap := map[string]string{
		"databaseUrl": "db2://username:password@hostname:50000/G2",
	}
	result, err := BuildSimpleSystemConfigurationJsonUsingMap(aMap)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	// Output: {"PIPELINE":{"CONFIGPATH":"/etc/opt/senzing","RESOURCEPATH":"/opt/senzing/g2/resources","SUPPORTPATH":"/opt/senzing/data"},"SQL":{"CONNECTION":"db2://username:password@G2"}}
}

func ExampleBuildSimpleSystemConfigurationJsonUsingMap_db2WithSchema() {
	aMap := map[string]string{
		"databaseUrl": "db2://username:password@hostname:50000/G2/?schema=schemaname",
	}
	result, err := BuildSimpleSystemConfigurationJsonUsingMap(aMap)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	// Output: {"PIPELINE":{"CONFIGPATH":"/etc/opt/senzing","RESOURCEPATH":"/opt/senzing/g2/resources","SUPPORTPATH":"/opt/senzing/data"},"SQL":{"CONNECTION":"db2://username:password@G2/?schema=schemaname"}}
}

func ExampleBuildSimpleSystemConfigurationJsonUsingMap_oci() {
	aMap := map[string]string{
		"databaseUrl": "oci://username:password@hostname:1521/G2",
	}
	result, err := BuildSimpleSystemConfigurationJsonUsingMap(aMap)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	// Output: {"PIPELINE":{"CONFIGPATH":"/etc/opt/senzing","RESOURCEPATH":"/opt/senzing/g2/resources","SUPPORTPATH":"/opt/senzing/data"},"SQL":{"CONNECTION":"oci://username:password@G2"}}
}

func ExampleBuildSimpleSystemConfigurationJsonUsingMap_mssql() {
	aMap := map[string]string{
		"databaseUrl": "mssql://username:password@hostname:1433/G2",
	}
	result, err := BuildSimpleSystemConfigurationJsonUsingMap(aMap)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	// Output: {"PIPELINE":{"CONFIGPATH":"/etc/opt/senzing","RESOURCEPATH":"/opt/senzing/g2/resources","SUPPORTPATH":"/opt/senzing/data"},"SQL":{"CONNECTION":"mssql://username:password@G2"}}
}

func ExampleBuildSimpleSystemConfigurationJsonUsingMap_mysql() {
	aMap := map[string]string{
		"databaseUrl": "mysql://username:password@hostname:3306/G2",
	}
	result, err := BuildSimpleSystemConfigurationJsonUsingMap(aMap)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	// Output: {"PIPELINE":{"CONFIGPATH":"/etc/opt/senzing","RESOURCEPATH":"/opt/senzing/g2/resources","SUPPORTPATH":"/opt/senzing/data"},"SQL":{"CONNECTION":"mysql://username:password@hostname:3306/?schema=G2"}}
}

func ExampleBuildSimpleSystemConfigurationJsonUsingMap_postgresql() {
	aMap := map[string]string{
		"databaseUrl": "postgresql://username:password@hostname:5432/G2",
	}
	result, err := BuildSimpleSystemConfigurationJsonUsingMap(aMap)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	// Output: {"PIPELINE":{"CONFIGPATH":"/etc/opt/senzing","RESOURCEPATH":"/opt/senzing/g2/resources","SUPPORTPATH":"/opt/senzing/data"},"SQL":{"CONNECTION":"postgresql://username:password@hostname:5432:G2/"}}
}

func ExampleBuildSimpleSystemConfigurationJsonUsingMap_postgresqlWithSchema() {
	aMap := map[string]string{
		"databaseUrl": "postgresql://username:password@hostname:5432/G2/?schema=schemaname",
	}
	result, err := BuildSimpleSystemConfigurationJsonUsingMap(aMap)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	// Output: {"PIPELINE":{"CONFIGPATH":"/etc/opt/senzing","RESOURCEPATH":"/opt/senzing/g2/resources","SUPPORTPATH":"/opt/senzing/data"},"SQL":{"CONNECTION":"postgresql://username:password@hostname:5432:G2/?schema=schemaname"}}
}
func ExampleBuildSimpleSystemConfigurationJsonUsingMap_sqlite() {
	aMap := map[string]string{
		"databaseUrl": "sqlite3://na:na@/var/opt/senzing/sqlite/G2C.db",
	}
	result, err := BuildSimpleSystemConfigurationJsonUsingMap(aMap)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	// Output: {"PIPELINE":{"CONFIGPATH":"/etc/opt/senzing","RESOURCEPATH":"/opt/senzing/g2/resources","SUPPORTPATH":"/opt/senzing/data"},"SQL":{"CONNECTION":"sqlite3://na:na@/var/opt/senzing/sqlite/G2C.db"}}
}
