/*
 */
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/senzing-garage/go-helpers/jsonutil"
	"github.com/senzing-garage/go-helpers/settings"
)

func main() {
	if len(os.Args) > 2 && os.Args[1] == "-j" {
		args := os.Args[2:]

		for _, jsonText := range args {
			normalized, err := jsonutil.Normalize(jsonText)
			if err != nil {
				log.Fatal(err)
			}
			normSorted, _ := jsonutil.NormalizeAndSort(jsonText)
			fmt.Println("- - - - - - - - - - - - - - - - - - - - ")
			fmt.Println(normalized)
			fmt.Println()
			fmt.Println(normSorted)
			fmt.Println()
		}
		return
	}
	ctx := context.TODO()

	// ------------------------------------------------------------------------
	// --- Build JSON from environment variables.
	// ------------------------------------------------------------------------

	iniParams, err := settings.BuildSimpleSettingsUsingEnvVars()
	if err != nil {
		panic(err)
	}
	fmt.Println(iniParams)

	// ------------------------------------------------------------------------
	// --- Verify parameters
	// ------------------------------------------------------------------------

	err = settings.VerifySettings(ctx, iniParams)
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

	iniParams2, err := settings.BuildSimpleSettingsUsingMap(attributeMap)
	if err != nil {
		panic(err)
	}
	fmt.Println(iniParams2)

}
