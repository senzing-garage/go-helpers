package settingsparser

import (
	"context"
	"fmt"
)

// ----------------------------------------------------------------------------
// Examples for godoc documentation
// ----------------------------------------------------------------------------

func ExampleBasicSettingsParser_GetConfigPath() {
	// For more information, visit https://github.com/senzing-garage/go-helpers/blob/main/settingsparser/settingsparser_test.go
	ctx := context.TODO()
	parser := getParser(ctx)
	configPath, err := parser.GetConfigPath(ctx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(configPath)
	// Output: /etc/opt/senzing
}

func ExampleBasicSettingsParser_GetDatabaseURLs() {
	// For more information, visit https://github.com/senzing-garage/go-helpers/blob/main/settingsparser/settingsparser_test.go
	ctx := context.TODO()
	parser := getParser(ctx)
	configPath, err := parser.GetDatabaseURLs(ctx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(configPath)
	// Output: [postgresql://username:password@db.example.com:5432:G2]
}

func ExampleBasicSettingsParser_GetResourcePath() {
	// For more information, visit https://github.com/senzing-garage/go-helpers/blob/main/settingsparser/settingsparser_test.go
	ctx := context.TODO()
	parser := getParser(ctx)
	configPath, err := parser.GetResourcePath(ctx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(configPath)
	// Output: /opt/senzing/er/resources
}

func ExampleBasicSettingsParser_GetSupportPath() {
	// For more information, visit https://github.com/senzing-garage/go-helpers/blob/main/settingsparser/settingsparser_test.go
	ctx := context.TODO()
	parser := getParser(ctx)
	configPath, err := parser.GetSupportPath(ctx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(configPath)
	// Output: /opt/senzing/data
}

func ExampleBasicSettingsParser_RedactedJSON_single() {
	ctx := context.TODO()
	parser := &BasicSettingsParser{
		Settings: `
        {
            "PIPELINE": {
                "CONFIGPATH": "/etc/opt/senzing",
                "LICENSESTRINGBASE64": "${SENZING_LICENSE_BASE64_ENCODED}",
                "RESOURCEPATH": "/opt/senzing/er/resources",
                "SUPPORTPATH": "/opt/senzing/data"
            },
            "SQL": {
                "BACKEND": "SQL",
                "CONNECTION": "postgresql://username:password@db.example.com:5432:G2"
            }
        }
        `,
	}

	actual, err := parser.RedactedJSON(ctx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(actual)
	// Output: {"PIPELINE":{"CONFIGPATH":"/etc/opt/senzing","LICENSESTRINGBASE64":"${SENZING_LICENSE_BASE64_ENCODED}","RESOURCEPATH":"/opt/senzing/er/resources","SUPPORTPATH":"/opt/senzing/data"},"SQL":{"BACKEND":"SQL","CONNECTION":"postgresql://username:xxxxx@db.example.com:5432/G2"}}
}

func ExampleBasicSettingsParser_RedactedJSON_multiple() {
	ctx := context.TODO()

	settings := `
        {
            "PIPELINE": {
                "CONFIGPATH": "/etc/opt/senzing",
                "LICENSESTRINGBASE64": "${SENZING_LICENSE_BASE64_ENCODED}",
                "RESOURCEPATH": "/opt/senzing/er/resources",
                "SUPPORTPATH": "/opt/senzing/data"
            },
            "SQL": {
                "BACKEND": "HYBRID",
                "CONNECTION": "postgresql://username:password@db-1.example.com:5432:G2"
            },
            "C1": {
                "CLUSTER_SIZE": "1",
                "DB_1": "postgresql://username:password@db-2.example.com:5432:G2"
            },
            "C2": {
                "CLUSTER_SIZE": "1",
                "DB_1": "postgresql://username:password@db-3.example.com:5432:G2"
            },
            "HYBRID": {
                "RES_FEAT": "C1",
                "RES_FEAT_EKEY": "C1",
                "RES_FEAT_LKEY": "C1",
                "RES_FEAT_STAT": "C1",
                "LIB_FEAT": "C2",
                "LIB_FEAT_HKEY": "C2"
            }
        }
        `

	parser, err := New(settings)
	if err != nil {
		fmt.Println(err)
	}

	actual, err := parser.RedactedJSON(ctx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(actual)
	// Output: {"PIPELINE":{"CONFIGPATH":"/etc/opt/senzing","LICENSESTRINGBASE64":"${SENZING_LICENSE_BASE64_ENCODED}","RESOURCEPATH":"/opt/senzing/er/resources","SUPPORTPATH":"/opt/senzing/data"},"SQL":{"BACKEND":"HYBRID","CONNECTION":"postgresql://username:xxxxx@db-1.example.com:5432/G2"},"C1":{"CLUSTER_SIZE":"1","DB_1":"postgresql://username:xxxxx@db-2.example.com:5432/G2"},"C2":{"CLUSTER_SIZE":"1","DB_1":"postgresql://username:xxxxx@db-3.example.com:5432/G2"},"HYBRID":{"RES_FEAT":"C1","RES_FEAT_EKEY":"C1","RES_FEAT_LKEY":"C1","RES_FEAT_STAT":"C1","LIB_FEAT":"C2","LIB_FEAT_HKEY":"C2"}}
}
