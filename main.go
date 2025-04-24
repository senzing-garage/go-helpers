/*
 */
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/senzing-garage/go-helpers/jsonutil"
	"github.com/senzing-garage/go-helpers/settings"
)

func main() {
	if len(os.Args) > 2 && os.Args[1] == "-j" {
		args := os.Args[2:]

		for _, jsonText := range args {
			normalized, err := jsonutil.Normalize(jsonText)
			failOnError(err)

			normSorted, _ := jsonutil.NormalizeAndSort(jsonText)

			outputln("- - - - - - - - - - - - - - - - - - - - ")
			outputln(normalized)
			outputln()
			outputln(normSorted)
			outputln()
		}

		return
	}

	ctx := context.TODO()

	// ------------------------------------------------------------------------
	// --- Build JSON from environment variables.
	// ------------------------------------------------------------------------

	iniParams, err := settings.BuildSimpleSettingsUsingEnvVars()
	failOnError(err)
	outputln(iniParams)

	// ------------------------------------------------------------------------
	// --- Verify parameters
	// ------------------------------------------------------------------------

	err = settings.VerifySettings(ctx, iniParams)
	failOnError(err)

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
	failOnError(err)
	outputln(iniParams2)
}

// ----------------------------------------------------------------------------
// Private functions
// ----------------------------------------------------------------------------

func failOnError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func outputln(message ...any) {
	fmt.Println(message...) //nolint
}
