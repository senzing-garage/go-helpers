package jsonutil

import (
	"errors"
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
// Test IsJson function
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

// ----------------------------------------------------------------------------
// Test NormalizeJson function
// ----------------------------------------------------------------------------
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
	assert.Equal(test, expected, actual, "JSON object (formatted) not normalized as expected")
}

// ----------------------------------------------------------------------------
// Test NormalizeAndSortJson function
// ----------------------------------------------------------------------------
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
	assert.Equal(test, expected, actual, "JSON string array not normalized as expected")
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
	assert.Equal(test, expected, actual, "JSON object (formatted) not normalized as expected")
}

// ----------------------------------------------------------------------------
// Test RedactJson function
// ----------------------------------------------------------------------------
func TestRedactJson_Basic0(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": "abc", "phoo": true, "lum": null}`
	var expected = `{"bar":"abc","foo":123,"lum":null,"phoo":true}`
	actual, err := RedactJson(jsonText)
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON object (basic 0) not redacted as expected")
}

func TestRedactJson_Basic1(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": "abc", "phoo": true, "lum": null}`
	var expected = `{"bar":"abc","foo":null,"lum":null,"phoo":true}`
	actual, err := RedactJson(jsonText, "foo")
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON object (basic 1) not redacted as expected")
}

func TestRedactJson_Basic2(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": "abc", "phoo": true, "lum": null}`
	var expected = `{"bar":null,"foo":null,"lum":null,"phoo":true}`
	actual, err := RedactJson(jsonText, "foo", "bar")
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON object (basic 2) not redacted as expected")
}

func TestRedactJson_StringArray(test *testing.T) {
	var jsonText = `["foo", "bar", "phoo", "lum"]`
	var expected = `["foo","bar","phoo","lum"]`
	actual, err := RedactJson(jsonText, "foo", "bar")
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON string array not redacted as expected")
}

func TestRedactJson_Compound0(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": [4, 6, 2], "phoo": true, "lum": {"phoox": false, "bax": 5}}`
	var expected = `{"bar":[4,6,2],"foo":123,"lum":{"bax":5,"phoox":false},"phoo":true}`
	actual, err := RedactJson(jsonText)
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON object (compound 0) not redacted as expected")
}

func TestRedactJson_Compound1(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": [4, 6, 2], "phoo": true, "lum": {"phoox": false, "bax": 5}}`
	var expected = `{"bar":[4,6,2],"foo":123,"lum":{"bax":5,"phoox":null},"phoo":true}`
	actual, err := RedactJson(jsonText, "phoox")
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON object (compound 1) not redacted as expected")
}

func TestRedactJson_Compound2(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": [4, 6, 2], "phoo": true, "lum": {"phoox": false, "bax": 5}}`
	var expected = `{"bar":null,"foo":123,"lum":{"bax":5,"phoox":null},"phoo":true}`
	actual, err := RedactJson(jsonText, "phoox", "bar")
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON object (compound 2) not redacted as expected")
}

func TestRedactJson_Compound3(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": [4, 6, 2], "phoox": true, "lum": {"phoox": false, "bax": 5}}`
	var expected = `{"bar":null,"foo":123,"lum":{"bax":null,"phoox":null},"phoox":null}`
	actual, err := RedactJson(jsonText, "phoox", "bar", "bax")
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON object (compound 3) not redacted as expected")
}

func TestRedactJson_Compound4(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": [4, 6, 2], "phoo": true, "lum": {"phoox": false, "bax": 5}}`
	var expected = `{"bar":null,"foo":null,"lum":{"bax":null,"phoox":null},"phoo":true}`
	actual, err := RedactJson(jsonText, "phoox", "bar", "bax", "foo")
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON object (compound 4) not redacted as expected")
}

func TestRedactJson_Compound5(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": [4, 6, 2], "phoo": true, "lum": {"phoox": false, "bax": 5}}`
	var expected = `{"bar":null,"foo":null,"lum":null,"phoo":null}`
	actual, err := RedactJson(jsonText, "phoox", "bar", "foo", "phoo", "lum")
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON object (compound 5) not redacted as expected")
}

func TestRedactJson_Integer(test *testing.T) {
	var jsonText = "123"
	var expected = "123"
	actual, err := RedactJson(jsonText, "123", "foo")
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON integer number not redacted as expected")
}

func TestRedactJson_Decimal(test *testing.T) {
	var jsonText = "123.4"
	var expected = "123.4"
	actual, err := RedactJson(jsonText, "123.4", "foo")
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON decimal number not redacted as expected")
}

func TestRedactJson_String(test *testing.T) {
	var jsonText = `"Hello"`
	var expected = `"Hello"`
	actual, err := RedactJson(jsonText, "Hello", "foo")
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON string not redacted as expected")
}

func TestRedactJson_Boolean(test *testing.T) {
	var jsonText = "true"
	var expected = "true"
	actual, err := RedactJson(jsonText, "true", "foo")
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON boolean not redacted as expected")
}

func TestRedactJson_Null(test *testing.T) {
	var jsonText = "null"
	var expected = "null"
	actual, err := RedactJson(jsonText, "null", "foo")
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON null not redacted as expected")
}

func TestRedactJson_MixedArray0(test *testing.T) {
	var jsonText = `[123, 123.5, "Hello", null, true, {"foo": 5, "bar": 6}]`
	var expected = `[123,123.5,"Hello",null,true,{"bar":6,"foo":5}]`
	actual, err := RedactJson(jsonText)
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON array (mixed 0) not redacted as expected")
}

func TestRedactJson_MixedArray1(test *testing.T) {
	var jsonText = `[123, 123.5, "Hello", null, true, {"foo": 5, "bar": 6}]`
	var expected = `[123,123.5,"Hello",null,true,{"bar":null,"foo":5}]`
	actual, err := RedactJson(jsonText, "bar")
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON array (mixed 1) not redacted as expected")
}

func TestRedactJson_MixedArray2(test *testing.T) {
	var jsonText = `[123, 123.5, "Hello", null, true, {"foo": 5, "bar": 6}]`
	var expected = `[123,123.5,"Hello",null,true,{"bar":null,"foo":null}]`
	actual, err := RedactJson(jsonText, "foo", "bar")
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON array (mixed 2) not redacted as expected")
}

func TestRedactJson_MixedArray3(test *testing.T) {
	var jsonText = `[123, 123.5, "Hello", null, true, {"foo": 5, "bar": 6}]`
	var expected = `[123,123.5,"Hello",null,true,{"bar":null,"foo":null}]`
	actual, err := RedactJson(jsonText, "foo", "bar", "Hello")
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON array (mixed 3) not redacted as expected")
}

func TestRedactJson_BadJson(test *testing.T) {
	var jsonText = `{foo: 123, bar: "abc", phoo: true, lum: {"phoox": 3, "bax": 5}}`
	actual, err := RedactJson(jsonText, "foo", "bar")
	assert.NotNil(test, err, "Invalid JSON text was redacted without an error: "+actual)
}

func TestRedactJson_Formatted0(test *testing.T) {
	var jsonText = `
	{
		"foo": 123,
		"bar": true,
		"phoo": [ {"c": 4, "a": 2}, {"a": 1, "c": [9, 0, 8]}, {"a": 1, "b": 5}]
	}`
	var expected = `{"bar":true,"foo":123,"phoo":[{"a":2,"c":4},{"a":1,"c":[9,0,8]},{"a":1,"b":5}]}`
	actual, err := RedactJson(jsonText)
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON object (formatted 0) not redacted as expected")
}

func TestRedactJson_Formatted1(test *testing.T) {
	var jsonText = `
	{
		"foo": 123,
		"bar": true,
		"phoo": [ {"c": 4, "a": 2}, {"a": 1, "c": [9, 0, 8]}, {"a": 1, "b": 5}]
	}`
	var expected = `{"bar":true,"foo":null,"phoo":[{"a":2,"c":4},{"a":1,"c":[9,0,8]},{"a":1,"b":5}]}`
	actual, err := RedactJson(jsonText, "foo")
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON object (formatted 1) not redacted as expected")
}

func TestRedactJson_Formatted2(test *testing.T) {
	var jsonText = `
	{
		"foo": 123,
		"bar": true,
		"phoo": [ {"c": 4, "a": 2}, {"a": 1, "c": [9, 0, 8]}, {"a": 1, "b": 5}]
	}`
	var expected = `{"bar":true,"foo":null,"phoo":[{"a":null,"c":4},{"a":null,"c":[9,0,8]},{"a":null,"b":5}]}`
	actual, err := RedactJson(jsonText, "foo", "a")
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON object (formatted 2) not redacted as expected")
}

func TestRedactJson_Formatted3(test *testing.T) {
	var jsonText = `
	{
		"foo": 123,
		"bar": true,
		"phoo": [ {"c": 4, "a": 2}, {"a": 1, "c": [9, 0, 8]}, {"a": 1, "b": 5}]
	}`
	var expected = `{"bar":true,"foo":null,"phoo":[{"a":null,"c":null},{"a":null,"c":null},{"a":null,"b":5}]}`
	actual, err := RedactJson(jsonText, "foo", "a", "c")
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON object (formatted 3) not redacted as expected")
}

func TestRedactJson_Formatted4(test *testing.T) {
	var jsonText = `
	{
		"foo": 123,
		"bar": true,
		"phoo": [ {"c": 4, "a": 2}, {"a": 1, "c": [9, 0, 8]}, {"a": 1, "b": 5}]
	}`
	var expected = `{"bar":null,"foo":null,"phoo":null}`
	actual, err := RedactJson(jsonText, "foo", "bar", "c", "phoo")
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON object (formatted 4) not redacted as expected")
}

// ----------------------------------------------------------------------------
// Test RedactJsonWithMap function
// ----------------------------------------------------------------------------
func TestRedactJsonWithMap_Basic0(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": "abc", "phoo": true, "lum": null}`
	var expected = `{"bar":"abc","foo":123,"lum":null,"phoo":true}`
	actual, err := RedactJsonWithMap(jsonText, map[string]any{})
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON object (basic 0) not redacted with map as expected")
}

func TestRedactJsonWithMap_Basic1(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": "abc", "phoo": true, "lum": null}`
	var expected = `{"bar":"abc","foo":"","lum":null,"phoo":true}`
	actual, err := RedactJsonWithMap(jsonText, map[string]any{"foo": ""})
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON object (basic 1) not redacted with map as expected")
}

func TestRedactJsonWithMap_Basic2(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": "abc", "phoo": true, "lum": null}`
	var expected = `{"bar":"-","foo":"","lum":null,"phoo":true}`
	actual, err := RedactJsonWithMap(jsonText, map[string]any{"foo": "", "bar": "-"})
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON object (basic 2) not redacted with map as expected")
}

func TestRedactJsonWithMap_StringArray(test *testing.T) {
	var jsonText = `["foo", "bar", "phoo", "lum"]`
	var expected = `["foo","bar","phoo","lum"]`
	actual, err := RedactJsonWithMap(jsonText, map[string]any{"foo": "", "bar": "-"})
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON string array not redacted with map as expected")
}

func TestRedactJsonWithMap_Compound0(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": [4, 6, 2], "phoo": true, "lum": {"phoox": false, "bax": 5}}`
	var expected = `{"bar":[4,6,2],"foo":123,"lum":{"bax":5,"phoox":false},"phoo":true}`
	actual, err := RedactJsonWithMap(jsonText, map[string]any{})
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON object (compound 0) not redacted with map as expected")
}

func TestRedactJsonWithMap_Compound1(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": [4, 6, 2], "phoo": true, "lum": {"phoox": false, "bax": 5}}`
	var expected = `{"bar":[4,6,2],"foo":123,"lum":{"bax":5,"phoox":""},"phoo":true}`
	actual, err := RedactJsonWithMap(jsonText, map[string]any{"phoox": ""})
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON object (compound 1) not redacted with map as expected")
}

func TestRedactJsonWithMap_Compound2(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": [4, 6, 2], "phoox": true, "lum": {"phoox": false, "bax": 5}}`
	var expected = `{"bar":"-","foo":123,"lum":{"bax":5,"phoox":""},"phoox":""}`
	actual, err := RedactJsonWithMap(jsonText, map[string]any{"phoox": "", "bar": "-"})
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON object (compound 2) not redacted with map as expected")
}

func TestRedactJsonWithMap_Compound3(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": [4, 6, 2], "phoo": true, "lum": {"phoox": false, "bax": 5}}`
	var expected = `{"bar":"-","foo":123,"lum":{"bax":null,"phoox":""},"phoo":true}`
	actual, err := RedactJsonWithMap(jsonText, map[string]any{"phoox": "", "bar": "-", "bax": nil})
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON object (compound 3) not redacted with map as expected")
}

func TestRedactJsonWithMap_Compound4(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": [4, 6, 2], "phoo": true, "lum": {"phoox": false, "bax": 5}}`
	var expected = `{"bar":"-","foo":"xxx","lum":{"bax":null,"phoox":""},"phoo":true}`
	actual, err := RedactJsonWithMap(jsonText, map[string]any{"phoox": "", "bar": "-", "bax": nil, "foo": "xxx"})
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON object (compound 4) not redacted with map as expected")
}

func TestRedactJsonWithMap_Compound5(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": [4, 6, 2], "phoo": true, "lum": {"phoox": false, "bax": 5}}`
	var expected = `{"bar":"-","foo":null,"lum":false,"phoo":"xxx"}`
	actual, err := RedactJsonWithMap(jsonText,
		map[string]any{"phoox": "", "bar": "-", "foo": nil, "phoo": "xxx", "lum": false})
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON object (compound 5) not redacted with map as expected")
}

func TestRedactJsonWithMap_Integer(test *testing.T) {
	var jsonText = "123"
	var expected = "123"
	actual, err := RedactJsonWithMap(jsonText, map[string]any{"123": "", "foo": "-"})
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON integer number not redacted with map as expected")
}

func TestRedactJsonWithMap_Decimal(test *testing.T) {
	var jsonText = "123.4"
	var expected = "123.4"
	actual, err := RedactJsonWithMap(jsonText, map[string]any{"123.4": "", "foo": "-"})
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON decimal number not redacted with map as expected")
}

func TestRedactJsonWithMap_String(test *testing.T) {
	var jsonText = `"Hello"`
	var expected = `"Hello"`
	actual, err := RedactJsonWithMap(jsonText, map[string]any{"Hello": "", "foo": "-"})
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON string not redacted with map as expected")
}

func TestRedactJsonWithMap_Boolean(test *testing.T) {
	var jsonText = "true"
	var expected = "true"
	actual, err := RedactJsonWithMap(jsonText, map[string]any{"true": false, "foo": "-"})
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON boolean not redacted with map as expected")
}

func TestRedactJsonWithMap_Null(test *testing.T) {
	var jsonText = "null"
	var expected = "null"
	actual, err := RedactJsonWithMap(jsonText, map[string]any{"null": nil, "foo": "-"})
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON null not redacted with map as expected")
}

func TestRedactJsonWithMap_MixedArray0(test *testing.T) {
	var jsonText = `[123, 123.5, "Hello", null, true, {"foo": 5, "bar": 6}]`
	var expected = `[123,123.5,"Hello",null,true,{"bar":6,"foo":5}]`
	actual, err := RedactJsonWithMap(jsonText, map[string]any{})
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON array (mixed 0) not redacted with map as expected")
}

func TestRedactJsonWithMap_MixedArray1(test *testing.T) {
	var jsonText = `[123, 123.5, "Hello", null, true, {"foo": 5, "bar": 6}]`
	var expected = `[123,123.5,"Hello",null,true,{"bar":"","foo":5}]`
	actual, err := RedactJsonWithMap(jsonText, map[string]any{"bar": ""})
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON array (mixed 1) not redacted with map as expected")
}

func TestRedactJsonWithMap_MixedArray2(test *testing.T) {
	var jsonText = `[123, 123.5, "Hello", null, true, {"foo": 5, "bar": 6}]`
	var expected = `[123,123.5,"Hello",null,true,{"bar":"-","foo":""}]`
	actual, err := RedactJsonWithMap(jsonText, map[string]any{"foo": "", "bar": "-"})
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON array (mixed 2) not redacted with map as expected")
}

func TestRedactJsonWithMap_MixedArray3(test *testing.T) {
	var jsonText = `[123, 123.5, "Hello", null, true, {"foo": 5, "bar": 6}]`
	var expected = `[123,123.5,"Hello",null,true,{"bar":"-","foo":""}]`
	actual, err := RedactJsonWithMap(jsonText, map[string]any{"foo": "", "bar": "-", "Hello": false})
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON array (mixed 3) not redacted with map as expected")
}

func TestRedactJsonWithMap_BadJson(test *testing.T) {
	var jsonText = `{foo: 123, bar: "abc", phoo: true, lum: {"phoox": 3, "bax": 5}}`
	actual, err := RedactJsonWithMap(jsonText, map[string]any{"foo": "", "bar": "-"})
	assert.NotNil(test, err, "Invalid JSON text was redacted with map without an error: "+actual)
}

func TestRedactJsonWithMap_Formatted0(test *testing.T) {
	var jsonText = `
	{
		"foo": 123,
		"bar": true,
		"phoo": [ {"c": 4, "a": 2}, {"a": 1, "c": [9, 0, 8]}, {"a": 1, "b": 5}]
	}`
	var expected = `{"bar":true,"foo":123,"phoo":[{"a":2,"c":4},{"a":1,"c":[9,0,8]},{"a":1,"b":5}]}`
	actual, err := RedactJsonWithMap(jsonText, map[string]any{})
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON object (formatted 0) not redacted with map as expected")
}

func TestRedactJsonWithMap_Formatted1(test *testing.T) {
	var jsonText = `
	{
		"foo": 123,
		"bar": true,
		"phoo": [ {"c": 4, "a": 2}, {"a": 1, "c": [9, 0, 8]}, {"a": 1, "b": 5}]
	}`
	var expected = `{"bar":true,"foo":"","phoo":[{"a":2,"c":4},{"a":1,"c":[9,0,8]},{"a":1,"b":5}]}`
	actual, err := RedactJsonWithMap(jsonText, map[string]any{"foo": ""})
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON object (formatted 1) not redacted with map as expected")
}

func TestRedactJsonWithMap_Formatted2(test *testing.T) {
	var jsonText = `
	{
		"foo": 123,
		"bar": true,
		"phoo": [ {"c": 4, "a": 2}, {"a": 1, "c": [9, 0, 8]}, {"a": 1, "b": 5}]
	}`
	var expected = `{"bar":true,"foo":"","phoo":[{"a":"-","c":4},{"a":"-","c":[9,0,8]},{"a":"-","b":5}]}`
	actual, err := RedactJsonWithMap(jsonText, map[string]any{"foo": "", "a": "-"})
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON object (formatted 2) not redacted with map as expected")
}

func TestRedactJsonWithMap_Formatted3(test *testing.T) {
	var jsonText = `
	{
		"foo": 123,
		"bar": true,
		"phoo": [ {"c": 4, "a": 2}, {"a": 1, "c": [9, 0, 8]}, {"a": 1, "b": 5}]
	}`
	var expected = `{"bar":true,"foo":"","phoo":[{"a":"-","c":"xxx"},{"a":"-","c":"xxx"},{"a":"-","b":5}]}`
	actual, err := RedactJsonWithMap(jsonText, map[string]any{"foo": "", "a": "-", "c": "xxx"})
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON object (formatted 3) not redacted with map as expected")
}

func TestRedactJsonWithMap_Formatted4(test *testing.T) {
	var jsonText = `
	{
		"foo": 123,
		"bar": true,
		"phoo": [ {"c": 4, "a": 2}, {"a": 1, "c": [9, 0, 8]}, {"a": 1, "b": 5}]
	}`
	var expected = `{"bar":"-","foo":"","phoo":false}`
	actual, err := RedactJsonWithMap(jsonText, map[string]any{"foo": "", "bar": "-", "c": "xxx", "phoo": false})
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON object (formatted 4) not redacted with map as expected")
}

func TestFlatten_NoError(test *testing.T) {
	actual := Flatten(`{"foo": 5, "bar": 6}`, nil)
	var expected = `{"foo": 5, "bar": 6}`
	assert.Equal(test, expected, actual, "Flattening without an error did not work as expected: "+actual)
}

func TestFlatten_WithError(test *testing.T) {
	err := errors.New("Failed")
	actual := Flatten(`{"foo": 5, "bar": 6}`, err)
	var expected = "Failed"
	assert.Equal(test, expected, actual, "Flattening with an error did not work as expected: "+actual)
}

// ----------------------------------------------------------------------------
// Example functions
// ----------------------------------------------------------------------------
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

func ExampleRedactJson() {
	// For more information, visit https://github.com/Senzing/go-common/blob/main/jsonutil/jsonutil_test.go
	var jsonText = `
	{
		"givenName": "Joe",
		"surname": "Schmoe",
		"age": 35,
		"member": true,
		"ssn": "111-22-3333"
	}`

	redactedJson, err := RedactJson(jsonText, "ssn")
	if err != nil {
		fmt.Println("An error occurred: " + err.Error())
	}

	fmt.Println(redactedJson)
	// Output: {"age":35,"givenName":"Joe","member":true,"ssn":null,"surname":"Schmoe"}
}

func ExampleRedactJsonWithMap() {
	// For more information, visit https://github.com/Senzing/go-common/blob/main/jsonutil/jsonutil_test.go
	var jsonText = `
	{
		"givenName": "Joe",
		"surname": "Schmoe",
		"age": 35,
		"member": true,
		"ssn": "111-22-3333"
	}`

	redactedJson, err := RedactJsonWithMap(jsonText, map[string]any{"ssn": "***-**-****"})
	if err != nil {
		fmt.Println("An error occurred: " + err.Error())
	}

	fmt.Println(redactedJson)
	// Output: {"age":35,"givenName":"Joe","member":true,"ssn":"***-**-****","surname":"Schmoe"}
}

func ExampleFlatten_noError() {
	// For more information, visit https://github.com/Senzing/go-common/blob/main/jsonutil/jsonutil_test.go
	var jsonText = `{ "name": "Joe Schmoe", "ssn": "111-22-3333" }`

	redactedJson := Flatten(RedactJsonWithMap(jsonText, map[string]any{"ssn": "***-**-****"}))

	fmt.Println(redactedJson)
	// Output: {"name":"Joe Schmoe","ssn":"***-**-****"}
}

func ExampleFlatten_withError() {
	// For more information, visit https://github.com/Senzing/go-common/blob/main/jsonutil/jsonutil_test.go
	var jsonText = `{ "name": "Joe Schmoe" "ssn": "111-22-3333" }` // missing a comma

	redactedJson := Flatten(RedactJsonWithMap(jsonText, map[string]any{"ssn": "***-**-****"}))

	fmt.Println(redactedJson)
	// Output: invalid character '"' after object key:value pair
}
