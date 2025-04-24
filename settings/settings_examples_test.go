//go:build linux

package settings_test

import (
	"context"
	"fmt"

	"github.com/senzing-garage/go-helpers/settings"
)

// ----------------------------------------------------------------------------
// Examples for godoc documentation
// ----------------------------------------------------------------------------

func ExampleBuildSenzingDatabaseURI() {
	databaseURL := "postgresql://username:password@hostname:5432/G2/?schema=schemaname"

	result, err := settings.BuildSenzingDatabaseURI(databaseURL)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)
	// Output: postgresql://username:password@hostname:5432:G2/?schema=schemaname
}

func ExampleBuildSenzingDatabaseURL() {
	databaseURI := "postgresql://username:password@hostname:5432:G2/?schema=schemaname"

	result, err := settings.BuildSenzingDatabaseURL(databaseURI)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)

	// Output: postgresql://username:password@hostname:5432/G2/?schema=schemaname
}

func ExampleBuildSimpleSettingsUsingEnvVars() {
	result, err := settings.BuildSimpleSettingsUsingEnvVars()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)
	// Output: {"PIPELINE":{"CONFIGPATH":"/etc/opt/senzing","RESOURCEPATH":"/opt/senzing/er/resources","SUPPORTPATH":"/opt/senzing/data"},"SQL":{"CONNECTION":"sqlite3://na:na@nowhere/tmp/sqlite/G2C.db"}}
}

func ExampleBuildSimpleSettingsUsingMap() {
	aMap := map[string]string{
		"databaseURL": "postgresql://username:password@hostname:5432/G2",
	}

	result, err := settings.BuildSimpleSettingsUsingMap(aMap)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)
	// Output: {"PIPELINE":{"CONFIGPATH":"/etc/opt/senzing","RESOURCEPATH":"/opt/senzing/er/resources","SUPPORTPATH":"/opt/senzing/data"},"SQL":{"CONNECTION":"postgresql://username:password@hostname:5432:G2/"}}
}

func ExampleBuildSimpleSettingsUsingMap_oci() {
	aMap := map[string]string{
		"databaseURL": "oci://username:password@hostname:1521/G2",
	}

	result, err := settings.BuildSimpleSettingsUsingMap(aMap)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)
	// Output: {"PIPELINE":{"CONFIGPATH":"/etc/opt/senzing","RESOURCEPATH":"/opt/senzing/er/resources","SUPPORTPATH":"/opt/senzing/data"},"SQL":{"CONNECTION":"oci://username:password@//hostname:1521/G2"}}
}

func ExampleBuildSimpleSettingsUsingMap_mssql() {
	aMap := map[string]string{
		"databaseURL": "mssql://username:password@hostname:1433/G2",
	}

	result, err := settings.BuildSimpleSettingsUsingMap(aMap)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)
	// Output: {"PIPELINE":{"CONFIGPATH":"/etc/opt/senzing","RESOURCEPATH":"/opt/senzing/er/resources","SUPPORTPATH":"/opt/senzing/data"},"SQL":{"CONNECTION":"mssql://username:password@G2"}}
}

func ExampleBuildSimpleSettingsUsingMap_mysql() {
	aMap := map[string]string{
		"databaseURL": "mysql://username:password@hostname:3306/G2",
	}

	result, err := settings.BuildSimpleSettingsUsingMap(aMap)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)
	// Output: {"PIPELINE":{"CONFIGPATH":"/etc/opt/senzing","RESOURCEPATH":"/opt/senzing/er/resources","SUPPORTPATH":"/opt/senzing/data"},"SQL":{"CONNECTION":"mysql://username:password@hostname:3306/?schema=G2"}}
}

func ExampleBuildSimpleSettingsUsingMap_postgresql() {
	aMap := map[string]string{
		"databaseURL": "postgresql://username:password@hostname:5432/G2",
	}

	result, err := settings.BuildSimpleSettingsUsingMap(aMap)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)
	// Output: {"PIPELINE":{"CONFIGPATH":"/etc/opt/senzing","RESOURCEPATH":"/opt/senzing/er/resources","SUPPORTPATH":"/opt/senzing/data"},"SQL":{"CONNECTION":"postgresql://username:password@hostname:5432:G2/"}}
}

func ExampleBuildSimpleSettingsUsingMap_postgresqlWithSchema() {
	aMap := map[string]string{
		"databaseURL": "postgresql://username:password@hostname:5432/G2/?schema=schemaname",
	}

	result, err := settings.BuildSimpleSettingsUsingMap(aMap)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)
	// Output: {"PIPELINE":{"CONFIGPATH":"/etc/opt/senzing","RESOURCEPATH":"/opt/senzing/er/resources","SUPPORTPATH":"/opt/senzing/data"},"SQL":{"CONNECTION":"postgresql://username:password@hostname:5432:G2/?schema=schemaname"}}
}

func ExampleBuildSimpleSettingsUsingMap_sqlite() {
	aMap := map[string]string{
		"databaseURL": "sqlite3://na:na@/var/opt/senzing/sqlite/G2C.db",
	}

	result, err := settings.BuildSimpleSettingsUsingMap(aMap)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)
	// Output: {"PIPELINE":{"CONFIGPATH":"/etc/opt/senzing","RESOURCEPATH":"/opt/senzing/er/resources","SUPPORTPATH":"/opt/senzing/data"},"SQL":{"CONNECTION":"sqlite3://na:na@/var/opt/senzing/sqlite/G2C.db"}}
}

func ExampleGetSenzingPath() {
	result := settings.GetSenzingPath()
	fmt.Println(result)
	// Output: /opt/senzing
}

func ExampleVerifySettings() {
	ctx := context.TODO()

	testJSON, err := settings.BuildSimpleSettingsUsingEnvVars()
	if err != nil {
		fmt.Println(err)
	}

	err = settings.VerifySettings(ctx, testJSON)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("JSON is valid.")
	}

	// Output: JSON is valid.
}
