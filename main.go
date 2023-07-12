/*
 */
package main

import (
	"fmt"

	"github.com/senzing/go-common/g2engineconfigurationjson"
)

func main() {

	// ------------------------------------------------------------------------
	// --- Using a bare message generator accepting databaseUrl
	// ------------------------------------------------------------------------

	iniParams, err := g2engineconfigurationjson.BuildSimpleSystemConfigurationJson("")
	if err != nil {
		panic(err)
	}
	fmt.Println(iniParams)

	// ------------------------------------------------------------------------
	// --- Using a generator accepting databaseUrl, licenseStringBase64,
	// ---    and senzingDirectory
	// ------------------------------------------------------------------------

	attributeMap := map[string]string{
		"licenseStringBase64": "SENZING_TOOLS_LICENSE_STRING_BASE64",
		"senzingDirectory":    "/path/to/senzing",
		"configPath":          "/another/path/for/config",
		"resourcePath":        "/yet/another/path/to/resources",
		"supportPath":         "/final/path/to/support",
	}

	iniParams2, err := g2engineconfigurationjson.BuildSimpleSystemConfigurationJsonViaMap(attributeMap)
	if err != nil {
		panic(err)
	}
	fmt.Println(iniParams2)

}
