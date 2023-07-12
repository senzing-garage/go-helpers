/*
 */
package main

import (
	"fmt"

	"github.com/senzing/go-common/g2engineconfigurationjson"
)

func main() {

	// ------------------------------------------------------------------------
	// --- Using a bare message generator
	// ------------------------------------------------------------------------

	iniParams, err := g2engineconfigurationjson.BuildSimpleSystemConfigurationJson("")
	if err != nil {
		panic(err)
	}
	fmt.Println(iniParams)

}
