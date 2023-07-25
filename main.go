/*
 */
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/senzing/go-common/g2engineconfigurationjson"
	"github.com/senzing/go-common/jsonutil"
)

func main() {
	args := os.Args[1:]

	for _, name := range args {
		normalized, err := jsonutil.NormalizeJson(name)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(normalized)
		}
	}

	ctx := context.TODO()

	// ------------------------------------------------------------------------
	// --- Build JSON from environment variables.
	// ------------------------------------------------------------------------

	iniParams, err := g2engineconfigurationjson.BuildSimpleSystemConfigurationJsonUsingEnvVars()
	if err != nil {
		panic(err)
	}
	fmt.Println(iniParams)

	// ------------------------------------------------------------------------
	// --- Verify parameters
	// ------------------------------------------------------------------------

	err = g2engineconfigurationjson.VerifySenzingEngineConfigurationJson(ctx, iniParams)
	if err != nil {
		panic(err)
	}

	// ------------------------------------------------------------------------
	// --- Build JSON from map of key/values.
	// ------------------------------------------------------------------------

	attributeMap := map[string]string{
		"licenseStringBase64": "8BD296A26F2034AAB436045...",
		"senzingDirectory":    "/path/to/senzing",
		"configPath":          "/another/path/for/config",
		"resourcePath":        "/yet/another/path/to/resources",
		"supportPath":         "/final/path/to/support",
	}

	iniParams2, err := g2engineconfigurationjson.BuildSimpleSystemConfigurationJsonUsingMap(attributeMap)
	if err != nil {
		panic(err)
	}
	fmt.Println(iniParams2)

}
