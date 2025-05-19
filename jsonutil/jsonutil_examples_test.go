package jsonutil_test

import (
	"fmt"

	"github.com/senzing-garage/go-helpers/jsonutil"
)

// ----------------------------------------------------------------------------
// Example functions
// ----------------------------------------------------------------------------

func ExampleFlatten_noError() {
	// For more information, visit https://github.com/senzing-garage/go-helpers/blob/main/jsonutil/jsonutil_test.go
	jsonText := `{ "name": "Joe Schmoe", "ssn": "111-22-3333" }`
	redactedJSON := jsonutil.Flatten(jsonutil.RedactWithMap(jsonText, map[string]any{"ssn": "***-**-****"}))
	fmt.Println(redactedJSON)
	// Output: {"name":"Joe Schmoe","ssn":"***-**-****"}
}

func ExampleFlatten_withError() {
	// For more information, visit https://github.com/senzing-garage/go-helpers/blob/main/jsonutil/jsonutil_test.go
	jsonText := `{ "name": "Joe Schmoe" "ssn": "111-22-3333" }` // missing a comma
	redactedJSON := jsonutil.Flatten(jsonutil.RedactWithMap(jsonText, map[string]any{"ssn": "***-**-****"}))
	fmt.Println(redactedJSON)
	// Output: {"function": "jsonutil.Flatten", "text": "{"function": "jsonutil.RedactWithMap", "text": "Unmarshal", "error": "invalid character '"' after object key:value pair"}", "error": "{"function": "jsonutil.RedactWithMap", "text": "Unmarshal", "error": "invalid character '"' after object key:value pair"}"}
}

func ExampleIsJSON() {
	// For more information, visit https://github.com/senzing-garage/go-helpers/blob/main/jsonutil/jsonutil_test.go
	jsonText := `{"givenName": "Joe","surname": "Schmoe","age": 35,"member": true}`

	validJSON := jsonutil.IsJSON(jsonText)
	if validJSON {
		fmt.Println(jsonText + " is valid JSON")
	} else {
		fmt.Println(jsonText + " is NOT valid JSON")
	}
	// Output: {"givenName": "Joe","surname": "Schmoe","age": 35,"member": true} is valid JSON
}

func ExampleNormalize() {
	// For more information, visit https://github.com/senzing-garage/go-helpers/blob/main/jsonutil/jsonutil_test.go
	jsonText := `
	{
		"givenName": "Jack",
		"surname": "Smith",
		"age": 43,
		"member": true
	}`

	normalizedJSON, err := jsonutil.Normalize(jsonText)
	if err != nil {
		fmt.Println("An error occurred: " + err.Error())
	}

	fmt.Println(normalizedJSON)
	// Output: {"age":43,"givenName":"Jack","member":true,"surname":"Smith"}
}

func ExampleNormalizeAndSort() {
	// For more information, visit https://github.com/senzing-garage/go-helpers/blob/main/jsonutil/jsonutil_test.go
	jsonText := `
	{
		"givenName": "Jane",
		"surname": "Doe",
		"age": 29,
		"member": true,
		"nicknames": ["Joseph", "Joey"]
	}`

	normalizedJSON, err := jsonutil.NormalizeAndSort(jsonText)
	if err != nil {
		fmt.Println("An error occurred: " + err.Error())
	}

	fmt.Println(normalizedJSON)
	// Output: {"age":29,"givenName":"Jane","member":true,"nicknames":["Joey","Joseph"],"surname":"Doe"}
}

func ExamplePrettyPrint() {
	// For more information, visit https://github.com/senzing-garage/go-helpers/blob/main/jsonutil/jsonutil_test.go
	jsonText := `{"givenName": "Don","surname": "Juan","age": 52,"member": true,"ssn": "111-22-3333"}`

	fmt.Println(jsonutil.PrettyPrint(jsonText, "    "))
	// Output:
	// {
	//     "givenName": "Don",
	//     "surname": "Juan",
	//     "age": 52,
	//     "member": true,
	//     "ssn": "111-22-3333"
	// }
}

func ExampleRedact() {
	// For more information, visit https://github.com/senzing-garage/go-helpers/blob/main/jsonutil/jsonutil_test.go
	jsonText := `
	{
		"givenName": "Bill",
		"surname": "Jackson",
		"age": 46,
		"member": true,
		"ssn": "111-22-3333"
	}`

	redactedJSON, err := jsonutil.Redact(jsonText, "ssn")
	if err != nil {
		fmt.Println("An error occurred: " + err.Error())
	}

	fmt.Println(redactedJSON)
	// Output: {"age":46,"givenName":"Bill","member":true,"ssn":null,"surname":"Jackson"}
}

func ExampleRedactWithMap() {
	// For more information, visit https://github.com/senzing-garage/go-helpers/blob/main/jsonutil/jsonutil_test.go
	jsonText := `
	{
		"givenName": "Roger",
		"surname": "That",
		"age": 65,
		"member": true,
		"ssn": "111-22-3333"
	}`

	redactedJSON, err := jsonutil.RedactWithMap(jsonText, map[string]any{"ssn": "***-**-****"})
	if err != nil {
		fmt.Println("An error occurred: " + err.Error())
	}

	fmt.Println(redactedJSON)
	// Output: {"age":65,"givenName":"Roger","member":true,"ssn":"***-**-****","surname":"That"}
}

func ExampleReverseString() {
	// For more information, visit https://github.com/senzing-garage/go-helpers/blob/main/jsonutil/jsonutil_test.go
	jsonText := `{"alpha": "beta"}`
	reversedJSON := jsonutil.ReverseString(jsonText)
	fmt.Println(reversedJSON)
	// Output: }"ateb" :"ahpla"{
}

func ExampleStrip() {
	// For more information, visit https://github.com/senzing-garage/go-helpers/blob/main/jsonutil/jsonutil_test.go
	jsonText := `
	{
		"givenName": "Joe",
		"surname": "Schmoe",
		"age": 35,
		"member": true,
		"ssn": "111-22-3333"
	}`

	redactedJSON, err := jsonutil.Strip(jsonText, "ssn")
	if err != nil {
		fmt.Println("An error occurred: " + err.Error())
	}

	fmt.Println(redactedJSON)
	// Output: {"age":35,"givenName":"Joe","member":true,"surname":"Schmoe"}
}

func ExampleTruncate() {
	// For more information, visit https://github.com/senzing-garage/go-helpers/blob/main/jsonutil/jsonutil_test.go
	jsonText := `
	{
		"givenName": "Joe",
		"surname": "Schmoe",
		"age": 35,
		"member": true,
		"ssn": "111-22-3333"
	}`

	fmt.Println(jsonutil.Truncate(jsonText, 5, "age"))
	// Output: {"givenName":"Joe","member":true,"ssn":"111-22-3333","surname":"Schmoe"...
}
