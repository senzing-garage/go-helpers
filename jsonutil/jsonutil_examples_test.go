package jsonutil

import (
	"fmt"
)

// ----------------------------------------------------------------------------
// Example functions
// ----------------------------------------------------------------------------

func ExampleIsJSON() {
	// For more information, visit https://github.com/senzing-garage/go-helpers/blob/main/jsonutil/jsonutil_test.go
	var jsonText = `{"givenName": "Joe","surname": "Schmoe","age": 35,"member": true}`
	validJSON := IsJSON(jsonText)
	if validJSON {
		fmt.Println(jsonText + " is valid JSON")
	} else {
		fmt.Println(jsonText + " is NOT valid JSON")
	}
	// Output: {"givenName": "Joe","surname": "Schmoe","age": 35,"member": true} is valid JSON
}

func ExampleNormalize() {
	// For more information, visit https://github.com/senzing-garage/go-helpers/blob/main/jsonutil/jsonutil_test.go
	var jsonText = `
	{
		"givenName": "Joe",
		"surname": "Schmoe",
		"age": 35,
		"member": true
	}`
	normalizedJSON, err := Normalize(jsonText)
	if err != nil {
		fmt.Println("An error occurred: " + err.Error())
	}
	fmt.Println(normalizedJSON)
	// Output: {"age":35,"givenName":"Joe","member":true,"surname":"Schmoe"}
}

func ExampleNormalizeAndSort() {
	// For more information, visit https://github.com/senzing-garage/go-helpers/blob/main/jsonutil/jsonutil_test.go
	var jsonText = `
	{
		"givenName": "Joe",
		"surname": "Schmoe",
		"age": 35,
		"member": true,
		"nicknames": ["Joseph", "Joey"]
	}`
	normalizedJSON, err := NormalizeAndSort(jsonText)
	if err != nil {
		fmt.Println("An error occurred: " + err.Error())
	}
	fmt.Println(normalizedJSON)
	// Output: {"age":35,"givenName":"Joe","member":true,"nicknames":["Joey","Joseph"],"surname":"Schmoe"}
}

func ExampleRedact() {
	// For more information, visit https://github.com/senzing-garage/go-helpers/blob/main/jsonutil/jsonutil_test.go
	var jsonText = `
	{
		"givenName": "Joe",
		"surname": "Schmoe",
		"age": 35,
		"member": true,
		"ssn": "111-22-3333"
	}`
	redactedJSON, err := Redact(jsonText, "ssn")
	if err != nil {
		fmt.Println("An error occurred: " + err.Error())
	}
	fmt.Println(redactedJSON)
	// Output: {"age":35,"givenName":"Joe","member":true,"ssn":null,"surname":"Schmoe"}
}

func ExampleRedactWithMap() {
	// For more information, visit https://github.com/senzing-garage/go-helpers/blob/main/jsonutil/jsonutil_test.go
	var jsonText = `
	{
		"givenName": "Joe",
		"surname": "Schmoe",
		"age": 35,
		"member": true,
		"ssn": "111-22-3333"
	}`
	redactedJSON, err := RedactWithMap(jsonText, map[string]any{"ssn": "***-**-****"})
	if err != nil {
		fmt.Println("An error occurred: " + err.Error())
	}
	fmt.Println(redactedJSON)
	// Output: {"age":35,"givenName":"Joe","member":true,"ssn":"***-**-****","surname":"Schmoe"}
}

func ExampleStrip() {
	// For more information, visit https://github.com/senzing-garage/go-helpers/blob/main/jsonutil/jsonutil_test.go
	var jsonText = `
	{
		"givenName": "Joe",
		"surname": "Schmoe",
		"age": 35,
		"member": true,
		"ssn": "111-22-3333"
	}`
	redactedJSON, err := Strip(jsonText, "ssn")
	if err != nil {
		fmt.Println("An error occurred: " + err.Error())
	}
	fmt.Println(redactedJSON)
	// Output: {"age":35,"givenName":"Joe","member":true,"surname":"Schmoe"}
}

func ExampleFlatten_noError() {
	// For more information, visit https://github.com/senzing-garage/go-helpers/blob/main/jsonutil/jsonutil_test.go
	var jsonText = `{ "name": "Joe Schmoe", "ssn": "111-22-3333" }`
	redactedJSON := Flatten(RedactWithMap(jsonText, map[string]any{"ssn": "***-**-****"}))
	fmt.Println(redactedJSON)
	// Output: {"name":"Joe Schmoe","ssn":"***-**-****"}
}

func ExampleFlatten_withError() {
	// For more information, visit https://github.com/senzing-garage/go-helpers/blob/main/jsonutil/jsonutil_test.go
	var jsonText = `{ "name": "Joe Schmoe" "ssn": "111-22-3333" }` // missing a comma
	redactedJSON := Flatten(RedactWithMap(jsonText, map[string]any{"ssn": "***-**-****"}))
	fmt.Println(redactedJSON)
	// Output: {"error":"invalid character '\"' after object key:value pair","text":"{ \"name\": \"Joe Schmoe\" \"ssn\": \"111-22-3333\" }"}
}
