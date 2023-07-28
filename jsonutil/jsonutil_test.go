package jsonutil

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ----------------------------------------------------------------------------
// Internal functions
// ----------------------------------------------------------------------------

func testError(test *testing.T, err error) {
	if err != nil {
		assert.FailNow(test, err.Error())
	}
}

// ----------------------------------------------------------------------------
// Test interface functions
// ----------------------------------------------------------------------------
func TestIsJson_Basic(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": "abc", "phoo": true, "lum": 20.5}`
	var expected = true
	actual := IsJson(jsonText)
	assert.Equal(test, expected, actual, "JSON object (basic) not recognized as JSON")
}

func TestIsJson_Compound(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": "abc", "phoo": true, "lum": {"phoox": 3, "bax": 5}}`
	var expected = true
	actual := IsJson(jsonText)
	assert.Equal(test, expected, actual, "JSON object (compound) not recognized as JSON")
}

func TestIsJson_Integer(test *testing.T) {
	var jsonText = "123"
	var expected = true
	actual := IsJson(jsonText)
	assert.Equal(test, expected, actual, "JSON integer not recognized as JSON")
}

func TestIsJson_Decimal(test *testing.T) {
	var jsonText = "123.4"
	var expected = true
	actual := IsJson(jsonText)
	assert.Equal(test, expected, actual, "JSON decimal number not recognized as JSON")
}

func TestIsJson_String(test *testing.T) {
	var jsonText = `"Hello"`
	var expected = true
	actual := IsJson(jsonText)
	assert.Equal(test, expected, actual, "JSON string not recognized as JSON")
}

func TestIsJson_Boolean(test *testing.T) {
	var jsonText = "true"
	var expected = true
	actual := IsJson(jsonText)
	assert.Equal(test, expected, actual, "JSON boolean not recognized as JSON")
}

func TestIsJson_Null(test *testing.T) {
	var jsonText = "null"
	var expected = true
	actual := IsJson(jsonText)
	assert.Equal(test, expected, actual, "JSON null not recognized as JSON")
}

func TestIsJson_Array(test *testing.T) {
	var jsonText = `[123, 123.5, "Hello", true, {"foo": 5, "bar": 6}]`
	var expected = true
	actual := IsJson(jsonText)
	assert.Equal(test, expected, actual, "JSON array not recognized as JSON")
}

func TestIsJson_BadJson(test *testing.T) {
	var jsonText = `{foo: 123, bar: "abc", phoo: true, lum: {"phoox": 3, "bax": 5}}`
	var expected = false
	actual := IsJson(jsonText)
	assert.Equal(test, expected, actual, "Invalid JSON text incorrectly recognized as JSON")
}

func TestIsJson_Formatted(test *testing.T) {
	var jsonText = `
	{
		"foo": 123,
		"bar": true
	}`
	var expected = true
	actual := IsJson(jsonText)
	assert.Equal(test, expected, actual, "Formatted JSON object not recognized as JSON")
}

func TestNormalizeJson_Basic(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": "abc", "phoo": true, "lum": null}`
	var expected = `{"bar":"abc","foo":123,"lum":null,"phoo":true}`
	actual, err := NormalizeJson(jsonText)
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON object (basic) not normalized as expected")
}

func TestNormalizeJson_Compound(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": "abc", "phoo": true, "lum": {"phoox": null, "bax": 5}}`
	var expected = `{"bar":"abc","foo":123,"lum":{"bax":5,"phoox":null},"phoo":true}`
	actual, err := NormalizeJson(jsonText)
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON object (compound) not normalized as expected")
}

func TestNormalizeJson_Integer(test *testing.T) {
	var jsonText = "123"
	var expected = "123"
	actual, err := NormalizeJson(jsonText)
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON integer number not normalized as expected")
}

func TestNormalizeJson_Decimal(test *testing.T) {
	var jsonText = "123.4"
	var expected = "123.4"
	actual, err := NormalizeJson(jsonText)
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON decimal number not normalized as expected")
}

func TestNormalizeJson_String(test *testing.T) {
	var jsonText = `"Hello"`
	var expected = `"Hello"`
	actual, err := NormalizeJson(jsonText)
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON string not normalized as expected")
}

func TestNormalizeJson_Boolean(test *testing.T) {
	var jsonText = "true"
	var expected = "true"
	actual, err := NormalizeJson(jsonText)
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON boolean not normalized as expected")
}

func TestNormalizeJson_Null(test *testing.T) {
	var jsonText = "null"
	var expected = "null"
	actual, err := NormalizeJson(jsonText)
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON null not normalized as expected")
}

func TestNormalizeJson_Array(test *testing.T) {
	var jsonText = `[123, 123.5, "Hello", true, {"foo": 5, "bar": 6}]`
	var expected = `[123,123.5,"Hello",true,{"bar":6,"foo":5}]`
	actual, err := NormalizeJson(jsonText)
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON array not normalized as expected")
}

func TestNormalizeJson_BadJson(test *testing.T) {
	var jsonText = `{foo: 123, bar: "abc", phoo: true, lum: {"phoox": 3, "bax": 5}}`
	actual, err := NormalizeJson(jsonText)
	assert.NotNil(test, err, "Invalid JSON text was normalized without an error: "+actual)
}

func TestNormalizeJson_Formatted(test *testing.T) {
	var jsonText = `
	{
		"foo": 123,
		"bar": true
	}`
	var expected = `{"bar":true,"foo":123}`
	actual, err := NormalizeJson(jsonText)
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON null not normalized as expected")
}

func TestNormalizeAndSortJson_Basic(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": "abc", "phoo": true, "lum": null}`
	var expected = `{"bar":"abc","foo":123,"lum":null,"phoo":true}`
	actual, err := NormalizeAndSortJson(jsonText)
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON object (basic) not normalized as expected")
}

func TestNormalizeAndSortJson_StringArray(test *testing.T) {
	var jsonText = `["foo", "bar", "phoo", "lum"]`
	var expected = `["bar","foo","lum","phoo"]`
	actual, err := NormalizeAndSortJson(jsonText)
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON object (basic) not normalized as expected")
}

func TestNormalizeAndSortJson_Compound(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": [4, 6, 2], "phoo": true, "lum": {"phoox": null, "bax": 5}}`
	var expected = `{"bar":[2,4,6],"foo":123,"lum":{"bax":5,"phoox":null},"phoo":true}`
	actual, err := NormalizeAndSortJson(jsonText)
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON object (compound) not normalized as expected")
}

func TestNormalizeAndSortJson_Integer(test *testing.T) {
	var jsonText = "123"
	var expected = "123"
	actual, err := NormalizeAndSortJson(jsonText)
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON integer number not normalized as expected")
}

func TestNormalizeAndSortJson_Decimal(test *testing.T) {
	var jsonText = "123.4"
	var expected = "123.4"
	actual, err := NormalizeAndSortJson(jsonText)
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON decimal number not normalized as expected")
}

func TestNormalizeAndSortJson_String(test *testing.T) {
	var jsonText = `"Hello"`
	var expected = `"Hello"`
	actual, err := NormalizeAndSortJson(jsonText)
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON string not normalized as expected")
}

func TestNormalizeAndSortJson_Boolean(test *testing.T) {
	var jsonText = "true"
	var expected = "true"
	actual, err := NormalizeAndSortJson(jsonText)
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON boolean not normalized as expected")
}

func TestNormalizeAndSortJson_Null(test *testing.T) {
	var jsonText = "null"
	var expected = "null"
	actual, err := NormalizeAndSortJson(jsonText)
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON null not normalized as expected")
}

func TestNormalizeAndSortJson_MixedArray(test *testing.T) {
	var jsonText = `[123, 123.5, "Hello", null, true, {"foo": 5, "bar": 6}]`
	var expected = `[null,"Hello",123,123.5,true,{"bar":6,"foo":5}]`
	actual, err := NormalizeAndSortJson(jsonText)
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON array not normalized as expected")
}

func TestNormalizeAndSortJson_BadJson(test *testing.T) {
	var jsonText = `{foo: 123, bar: "abc", phoo: true, lum: {"phoox": 3, "bax": 5}}`
	actual, err := NormalizeAndSortJson(jsonText)
	assert.NotNil(test, err, "Invalid JSON text was normalized without an error: "+actual)
}

func TestNormalizeAndSortJson_Formatted(test *testing.T) {
	var jsonText = `
	{
		"foo": 123,
		"bar": true,
		"phoo": [ {"c": 4, "a": 2}, {"a": 1, "c": [9, 0, 8]}, {"a": 1, "b": 5}]
	}`
	var expected = `{"bar":true,"foo":123,"phoo":[{"a":1,"b":5},{"a":1,"c":[0,8,9]},{"a":2,"c":4}]}`
	actual, err := NormalizeAndSortJson(jsonText)
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON null not normalized as expected")
}

func ExampleIsJson() {
	// For more information, visit https://github.com/Senzing/go-common/blob/main/jsonutil/jsonutil_test.go
	var jsonText = `{"givenName": "Joe","surname": "Schmoe","age": 35,"member": true}`

	validJson := IsJson(jsonText)

	if validJson {
		fmt.Println(jsonText + " is valid JSON")
	} else {
		fmt.Println(jsonText + " is NOT valid JSON")
	}
	// Output: {"givenName": "Joe","surname": "Schmoe","age": 35,"member": true} is valid JSON
}

func ExampleNormalizeJson() {
	// For more information, visit https://github.com/Senzing/go-common/blob/main/jsonutil/jsonutil_test.go
	var jsonText = `
	{
		"givenName": "Joe",
		"surname": "Schmoe",
		"age": 35,
		"member": true
	}`

	normalizedJson, err := NormalizeJson(jsonText)
	if err != nil {
		fmt.Println("An error occurred: " + err.Error())
	}

	fmt.Println(normalizedJson)
	// Output: {"age":35,"givenName":"Joe","member":true,"surname":"Schmoe"}
}

func ExampleNormalizeAndSortJson() {
	// For more information, visit https://github.com/Senzing/go-common/blob/main/jsonutil/jsonutil_test.go
	var jsonText = `
	{
		"givenName": "Joe",
		"surname": "Schmoe",
		"age": 35,
		"member": true,
		"nicknames": ["Joseph", "Joey"]
	}`

	normalizedJson, err := NormalizeAndSortJson(jsonText)
	if err != nil {
		fmt.Println("An error occurred: " + err.Error())
	}

	fmt.Println(normalizedJson)
	// Output: {"age":35,"givenName":"Joe","member":true,"nicknames":["Joey","Joseph"],"surname":"Schmoe"}
}
