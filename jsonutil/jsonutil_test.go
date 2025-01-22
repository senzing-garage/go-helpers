package jsonutil

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
	actual := IsJSON(jsonText)
	assert.Equal(test, expected, actual, "JSON object (basic) not recognized as JSON")
}

func TestIsJson_Compound(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": "abc", "phoo": true, "lum": {"phoox": 3, "bax": 5}}`
	var expected = true
	actual := IsJSON(jsonText)
	assert.Equal(test, expected, actual, "JSON object (compound) not recognized as JSON")
}

func TestIsJson_Integer(test *testing.T) {
	var jsonText = "123"
	var expected = true
	actual := IsJSON(jsonText)
	assert.Equal(test, expected, actual, "JSON integer not recognized as JSON")
}

func TestIsJson_Decimal(test *testing.T) {
	var jsonText = "123.4"
	var expected = true
	actual := IsJSON(jsonText)
	assert.Equal(test, expected, actual, "JSON decimal number not recognized as JSON")
}

func TestIsJson_String(test *testing.T) {
	var jsonText = `"Hello"`
	var expected = true
	actual := IsJSON(jsonText)
	assert.Equal(test, expected, actual, "JSON string not recognized as JSON")
}

func TestIsJson_Boolean(test *testing.T) {
	var jsonText = "true"
	var expected = true
	actual := IsJSON(jsonText)
	assert.Equal(test, expected, actual, "JSON boolean not recognized as JSON")
}

func TestIsJson_Null(test *testing.T) {
	var jsonText = "null"
	var expected = true
	actual := IsJSON(jsonText)
	assert.Equal(test, expected, actual, "JSON null not recognized as JSON")
}

func TestIsJson_Array(test *testing.T) {
	var jsonText = `[123, 123.5, "Hello", true, {"foo": 5, "bar": 6}]`
	var expected = true
	actual := IsJSON(jsonText)
	assert.Equal(test, expected, actual, "JSON array not recognized as JSON")
}

func TestIsJson_BadJson(test *testing.T) {
	var jsonText = `{foo: 123, bar: "abc", phoo: true, lum: {"phoox": 3, "bax": 5}}`
	var expected = false
	actual := IsJSON(jsonText)
	assert.Equal(test, expected, actual, "Invalid JSON text incorrectly recognized as JSON")
}

func TestIsJson_Formatted(test *testing.T) {
	var jsonText = `
	{
		"foo": 123,
		"bar": true
	}`
	var expected = true
	actual := IsJSON(jsonText)
	assert.Equal(test, expected, actual, "Formatted JSON object not recognized as JSON")
}

// ----------------------------------------------------------------------------
// Test NormalizeJson function
// ----------------------------------------------------------------------------
func TestNormalize_Basic(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": "abc", "phoo": true, "lum": null}`
	var expected = `{"bar":"abc","foo":123,"lum":null,"phoo":true}`
	actual, err := Normalize(jsonText)
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON object (basic) not normalized as expected")
}

func TestNormalize_Compound(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": "abc", "phoo": true, "lum": {"phoox": null, "bax": 5}}`
	var expected = `{"bar":"abc","foo":123,"lum":{"bax":5,"phoox":null},"phoo":true}`
	actual, err := Normalize(jsonText)
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON object (compound) not normalized as expected")
}

func TestNormalize_Integer(test *testing.T) {
	var jsonText = "123"
	var expected = "123"
	actual, err := Normalize(jsonText)
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON integer number not normalized as expected")
}

func TestNormalize_Decimal(test *testing.T) {
	var jsonText = "123.4"
	var expected = "123.4"
	actual, err := Normalize(jsonText)
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON decimal number not normalized as expected")
}

func TestNormalize_String(test *testing.T) {
	var jsonText = `"Hello"`
	var expected = `"Hello"`
	actual, err := Normalize(jsonText)
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON string not normalized as expected")
}

func TestNormalize_Boolean(test *testing.T) {
	var jsonText = "true"
	var expected = "true"
	actual, err := Normalize(jsonText)
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON boolean not normalized as expected")
}

func TestNormalize_Null(test *testing.T) {
	var jsonText = "null"
	var expected = "null"
	actual, err := Normalize(jsonText)
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON null not normalized as expected")
}

func TestNormalize_Array(test *testing.T) {
	var jsonText = `[123, 123.5, "Hello", true, {"foo": 5, "bar": 6}]`
	var expected = `[123,123.5,"Hello",true,{"bar":6,"foo":5}]`
	actual, err := Normalize(jsonText)
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON array not normalized as expected")
}

func TestNormalize_BadJson(test *testing.T) {
	var jsonText = `{foo: 123, bar: "abc", phoo: true, lum: {"phoox": 3, "bax": 5}}`
	actual, err := Normalize(jsonText)
	require.Error(test, err, "Invalid JSON text was normalized without an error: "+actual)
}

func TestNormalize_Formatted(test *testing.T) {
	var jsonText = `
	{
		"foo": 123,
		"bar": true
	}`
	var expected = `{"bar":true,"foo":123}`
	actual, err := Normalize(jsonText)
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON object (formatted) not normalized as expected")
}

// ----------------------------------------------------------------------------
// Test NormalizeAndSortJson function
// ----------------------------------------------------------------------------
func TestNormalizeAndSort_Basic(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": "abc", "phoo": true, "lum": null}`
	var expected = `{"bar":"abc","foo":123,"lum":null,"phoo":true}`
	actual, err := NormalizeAndSort(jsonText)
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON object (basic) not normalized as expected")
}

func TestNormalizeAndSort_StringArray(test *testing.T) {
	var jsonText = `["foo", "bar", "phoo", "lum"]`
	var expected = `["bar","foo","lum","phoo"]`
	actual, err := NormalizeAndSort(jsonText)
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON string array not normalized as expected")
}

func TestNormalizeAndSort_Compound(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": [4, 6, 2], "phoo": true, "lum": {"phoox": null, "bax": 5}}`
	var expected = `{"bar":[2,4,6],"foo":123,"lum":{"bax":5,"phoox":null},"phoo":true}`
	actual, err := NormalizeAndSort(jsonText)
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON object (compound) not normalized as expected")
}

func TestNormalizeAndSort_Integer(test *testing.T) {
	var jsonText = "123"
	var expected = "123"
	actual, err := NormalizeAndSort(jsonText)
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON integer number not normalized as expected")
}

func TestNormalizeAndSort_Decimal(test *testing.T) {
	var jsonText = "123.4"
	var expected = "123.4"
	actual, err := NormalizeAndSort(jsonText)
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON decimal number not normalized as expected")
}

func TestNormalizeAndSort_String(test *testing.T) {
	var jsonText = `"Hello"`
	var expected = `"Hello"`
	actual, err := NormalizeAndSort(jsonText)
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON string not normalized as expected")
}

func TestNormalizeAndSort_Boolean(test *testing.T) {
	var jsonText = "true"
	var expected = "true"
	actual, err := NormalizeAndSort(jsonText)
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON boolean not normalized as expected")
}

func TestNormalizeAndSort_Null(test *testing.T) {
	var jsonText = "null"
	var expected = "null"
	actual, err := NormalizeAndSort(jsonText)
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON null not normalized as expected")
}

func TestNormalizeAndSort_MixedArray(test *testing.T) {
	var jsonText = `[123, 123.5, "Hello", null, true, {"foo": 5, "bar": 6}]`
	var expected = `[null,"Hello",123,123.5,true,{"bar":6,"foo":5}]`
	actual, err := NormalizeAndSort(jsonText)
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON array not normalized as expected")
}

func TestNormalizeAndSort_BadJson(test *testing.T) {
	var jsonText = `{foo: 123, bar: "abc", phoo: true, lum: {"phoox": 3, "bax": 5}}`
	actual, err := NormalizeAndSort(jsonText)
	require.Error(test, err, "Invalid JSON text was normalized without an error: "+actual)
}

func TestNormalizeAndSort_Formatted(test *testing.T) {
	var jsonText = `
	{
		"foo": 123,
		"bar": true,
		"phoo": [ {"c": 4, "a": 2}, {"a": 1, "c": [9, 0, 8]}, {"a": 1, "b": 5}]
	}`
	var expected = `{"bar":true,"foo":123,"phoo":[{"a":1,"b":5},{"a":1,"c":[0,8,9]},{"a":2,"c":4}]}`
	actual, err := NormalizeAndSort(jsonText)
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON object (formatted) not normalized as expected")
}

// ----------------------------------------------------------------------------
// Test PrettyPrint function
// ----------------------------------------------------------------------------

func TestPrettyPrint_x(test *testing.T) {
	var jsonText = `
	{
		"foo": 123,
		"bar": true,
		"phoo": [ {"c": 4, "a": 2}, {"a": 1, "c": [9, 0, 8]}, {"a": 1, "b": 5}]
	}`
	var expected = `{"bar":true,"foo":123,"phoo":[{"a":1,"b":5},{"a":1,"c":[0,8,9]},{"a":2,"c":4}]}`
	actual := PrettyPrint(jsonText)
	assert.Equal(test, expected, actual, "JSON object (formatted) not normalized as expected")
}

// ----------------------------------------------------------------------------
// Test RedactJson function
// ----------------------------------------------------------------------------
func TestRedact_Basic0(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": "abc", "phoo": true, "lum": null}`
	var expected = `{"bar":"abc","foo":123,"lum":null,"phoo":true}`
	actual, err := Redact(jsonText)
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON object (basic 0) not redacted as expected")
}

func TestRedact_Basic1(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": "abc", "phoo": true, "lum": null}`
	var expected = `{"bar":"abc","foo":null,"lum":null,"phoo":true}`
	actual, err := Redact(jsonText, "foo")
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON object (basic 1) not redacted as expected")
}

func TestRedact_Basic2(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": "abc", "phoo": true, "lum": null}`
	var expected = `{"bar":null,"foo":null,"lum":null,"phoo":true}`
	actual, err := Redact(jsonText, "foo", "bar")
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON object (basic 2) not redacted as expected")
}

func TestRedact_StringArray(test *testing.T) {
	var jsonText = `["foo", "bar", "phoo", "lum"]`
	var expected = `["foo","bar","phoo","lum"]`
	actual, err := Redact(jsonText, "foo", "bar")
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON string array not redacted as expected")
}

func TestRedact_Compound0(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": [4, 6, 2], "phoo": true, "lum": {"phoox": false, "bax": 5}}`
	var expected = `{"bar":[4,6,2],"foo":123,"lum":{"bax":5,"phoox":false},"phoo":true}`
	actual, err := Redact(jsonText)
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON object (compound 0) not redacted as expected")
}

func TestRedact_Compound1(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": [4, 6, 2], "phoo": true, "lum": {"phoox": false, "bax": 5}}`
	var expected = `{"bar":[4,6,2],"foo":123,"lum":{"bax":5,"phoox":null},"phoo":true}`
	actual, err := Redact(jsonText, "phoox")
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON object (compound 1) not redacted as expected")
}

func TestRedact_Compound2(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": [4, 6, 2], "phoo": true, "lum": {"phoox": false, "bax": 5}}`
	var expected = `{"bar":null,"foo":123,"lum":{"bax":5,"phoox":null},"phoo":true}`
	actual, err := Redact(jsonText, "phoox", "bar")
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON object (compound 2) not redacted as expected")
}

func TestRedact_Compound3(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": [4, 6, 2], "phoox": true, "lum": {"phoox": false, "bax": 5}}`
	var expected = `{"bar":null,"foo":123,"lum":{"bax":null,"phoox":null},"phoox":null}`
	actual, err := Redact(jsonText, "phoox", "bar", "bax")
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON object (compound 3) not redacted as expected")
}

func TestRedact_Compound4(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": [4, 6, 2], "phoo": true, "lum": {"phoox": false, "bax": 5}}`
	var expected = `{"bar":null,"foo":null,"lum":{"bax":null,"phoox":null},"phoo":true}`
	actual, err := Redact(jsonText, "phoox", "bar", "bax", "foo")
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON object (compound 4) not redacted as expected")
}

func TestRedact_Compound5(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": [4, 6, 2], "phoo": true, "lum": {"phoox": false, "bax": 5}}`
	var expected = `{"bar":null,"foo":null,"lum":null,"phoo":null}`
	actual, err := Redact(jsonText, "phoox", "bar", "foo", "phoo", "lum")
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON object (compound 5) not redacted as expected")
}

func TestRedact_Integer(test *testing.T) {
	var jsonText = "123"
	var expected = "123"
	actual, err := Redact(jsonText, "123", "foo")
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON integer number not redacted as expected")
}

func TestRedact_Decimal(test *testing.T) {
	var jsonText = "123.4"
	var expected = "123.4"
	actual, err := Redact(jsonText, "123.4", "foo")
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON decimal number not redacted as expected")
}

func TestRedact_String(test *testing.T) {
	var jsonText = `"Hello"`
	var expected = `"Hello"`
	actual, err := Redact(jsonText, "Hello", "foo")
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON string not redacted as expected")
}

func TestRedact_Boolean(test *testing.T) {
	var jsonText = "true"
	var expected = "true"
	actual, err := Redact(jsonText, "true", "foo")
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON boolean not redacted as expected")
}

func TestRedact_Null(test *testing.T) {
	var jsonText = "null"
	var expected = "null"
	actual, err := Redact(jsonText, "null", "foo")
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON null not redacted as expected")
}

func TestRedact_MixedArray0(test *testing.T) {
	var jsonText = `[123, 123.5, "Hello", null, true, {"foo": 5, "bar": 6}]`
	var expected = `[123,123.5,"Hello",null,true,{"bar":6,"foo":5}]`
	actual, err := Redact(jsonText)
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON array (mixed 0) not redacted as expected")
}

func TestRedact_MixedArray1(test *testing.T) {
	var jsonText = `[123, 123.5, "Hello", null, true, {"foo": 5, "bar": 6}]`
	var expected = `[123,123.5,"Hello",null,true,{"bar":null,"foo":5}]`
	actual, err := Redact(jsonText, "bar")
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON array (mixed 1) not redacted as expected")
}

func TestRedact_MixedArray2(test *testing.T) {
	var jsonText = `[123, 123.5, "Hello", null, true, {"foo": 5, "bar": 6}]`
	var expected = `[123,123.5,"Hello",null,true,{"bar":null,"foo":null}]`
	actual, err := Redact(jsonText, "foo", "bar")
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON array (mixed 2) not redacted as expected")
}

func TestRedact_MixedArray3(test *testing.T) {
	var jsonText = `[123, 123.5, "Hello", null, true, {"foo": 5, "bar": 6}]`
	var expected = `[123,123.5,"Hello",null,true,{"bar":null,"foo":null}]`
	actual, err := Redact(jsonText, "foo", "bar", "Hello")
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON array (mixed 3) not redacted as expected")
}

func TestRedact_BadJson(test *testing.T) {
	var jsonText = `{foo: 123, bar: "abc", phoo: true, lum: {"phoox": 3, "bax": 5}}`
	actual, err := Redact(jsonText, "foo", "bar")
	require.Error(test, err, "Invalid JSON text was redacted without an error: "+actual)
}

func TestRedact_Formatted0(test *testing.T) {
	var jsonText = `
	{
		"foo": 123,
		"bar": true,
		"phoo": [ {"c": 4, "a": 2}, {"a": 1, "c": [9, 0, 8]}, {"a": 1, "b": 5}]
	}`
	var expected = `{"bar":true,"foo":123,"phoo":[{"a":2,"c":4},{"a":1,"c":[9,0,8]},{"a":1,"b":5}]}`
	actual, err := Redact(jsonText)
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON object (formatted 0) not redacted as expected")
}

func TestRedact_Formatted1(test *testing.T) {
	var jsonText = `
	{
		"foo": 123,
		"bar": true,
		"phoo": [ {"c": 4, "a": 2}, {"a": 1, "c": [9, 0, 8]}, {"a": 1, "b": 5}]
	}`
	var expected = `{"bar":true,"foo":null,"phoo":[{"a":2,"c":4},{"a":1,"c":[9,0,8]},{"a":1,"b":5}]}`
	actual, err := Redact(jsonText, "foo")
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON object (formatted 1) not redacted as expected")
}

func TestRedact_Formatted2(test *testing.T) {
	var jsonText = `
	{
		"foo": 123,
		"bar": true,
		"phoo": [ {"c": 4, "a": 2}, {"a": 1, "c": [9, 0, 8]}, {"a": 1, "b": 5}]
	}`
	var expected = `{"bar":true,"foo":null,"phoo":[{"a":null,"c":4},{"a":null,"c":[9,0,8]},{"a":null,"b":5}]}`
	actual, err := Redact(jsonText, "foo", "a")
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON object (formatted 2) not redacted as expected")
}

func TestRedact_Formatted3(test *testing.T) {
	var jsonText = `
	{
		"foo": 123,
		"bar": true,
		"phoo": [ {"c": 4, "a": 2}, {"a": 1, "c": [9, 0, 8]}, {"a": 1, "b": 5}]
	}`
	var expected = `{"bar":true,"foo":null,"phoo":[{"a":null,"c":null},{"a":null,"c":null},{"a":null,"b":5}]}`
	actual, err := Redact(jsonText, "foo", "a", "c")
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON object (formatted 3) not redacted as expected")
}

func TestRedact_Formatted4(test *testing.T) {
	var jsonText = `
	{
		"foo": 123,
		"bar": true,
		"phoo": [ {"c": 4, "a": 2}, {"a": 1, "c": [9, 0, 8]}, {"a": 1, "b": 5}]
	}`
	var expected = `{"bar":null,"foo":null,"phoo":null}`
	actual, err := Redact(jsonText, "foo", "bar", "c", "phoo")
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON object (formatted 4) not redacted as expected")
}

// ----------------------------------------------------------------------------
// Test RedactWithMap function
// ----------------------------------------------------------------------------
func TestRedactWithMap_Basic0(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": "abc", "phoo": true, "lum": null}`
	var expected = `{"bar":"abc","foo":123,"lum":null,"phoo":true}`
	actual, err := RedactWithMap(jsonText, map[string]any{})
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON object (basic 0) not redacted with map as expected")
}

func TestRedactWithMap_Basic1(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": "abc", "phoo": true, "lum": null}`
	var expected = `{"bar":"abc","foo":"","lum":null,"phoo":true}`
	actual, err := RedactWithMap(jsonText, map[string]any{"foo": ""})
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON object (basic 1) not redacted with map as expected")
}

func TestRedactWithMap_Basic2(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": "abc", "phoo": true, "lum": null}`
	var expected = `{"bar":"-","foo":"","lum":null,"phoo":true}`
	actual, err := RedactWithMap(jsonText, map[string]any{"foo": "", "bar": "-"})
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON object (basic 2) not redacted with map as expected")
}

func TestRedactWithMap_StringArray(test *testing.T) {
	var jsonText = `["foo", "bar", "phoo", "lum"]`
	var expected = `["foo","bar","phoo","lum"]`
	actual, err := RedactWithMap(jsonText, map[string]any{"foo": "", "bar": "-"})
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON string array not redacted with map as expected")
}

func TestRedactWithMap_Compound0(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": [4, 6, 2], "phoo": true, "lum": {"phoox": false, "bax": 5}}`
	var expected = `{"bar":[4,6,2],"foo":123,"lum":{"bax":5,"phoox":false},"phoo":true}`
	actual, err := RedactWithMap(jsonText, map[string]any{})
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON object (compound 0) not redacted with map as expected")
}

func TestRedactWithMap_Compound1(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": [4, 6, 2], "phoo": true, "lum": {"phoox": false, "bax": 5}}`
	var expected = `{"bar":[4,6,2],"foo":123,"lum":{"bax":5,"phoox":""},"phoo":true}`
	actual, err := RedactWithMap(jsonText, map[string]any{"phoox": ""})
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON object (compound 1) not redacted with map as expected")
}

func TestRedactWithMap_Compound2(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": [4, 6, 2], "phoox": true, "lum": {"phoox": false, "bax": 5}}`
	var expected = `{"bar":"-","foo":123,"lum":{"bax":5,"phoox":""},"phoox":""}`
	actual, err := RedactWithMap(jsonText, map[string]any{"phoox": "", "bar": "-"})
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON object (compound 2) not redacted with map as expected")
}

func TestRedactWithMap_Compound3(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": [4, 6, 2], "phoo": true, "lum": {"phoox": false, "bax": 5}}`
	var expected = `{"bar":"-","foo":123,"lum":{"bax":null,"phoox":""},"phoo":true}`
	actual, err := RedactWithMap(jsonText, map[string]any{"phoox": "", "bar": "-", "bax": nil})
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON object (compound 3) not redacted with map as expected")
}

func TestRedactWithMap_Compound4(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": [4, 6, 2], "phoo": true, "lum": {"phoox": false, "bax": 5}}`
	var expected = `{"bar":"-","foo":"xxx","lum":{"bax":null,"phoox":""},"phoo":true}`
	actual, err := RedactWithMap(jsonText, map[string]any{"phoox": "", "bar": "-", "bax": nil, "foo": "xxx"})
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON object (compound 4) not redacted with map as expected")
}

func TestRedactWithMap_Compound5(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": [4, 6, 2], "phoo": true, "lum": {"phoox": false, "bax": 5}}`
	var expected = `{"bar":"-","foo":null,"lum":false,"phoo":"xxx"}`
	actual, err := RedactWithMap(jsonText,
		map[string]any{"phoox": "", "bar": "-", "foo": nil, "phoo": "xxx", "lum": false})
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON object (compound 5) not redacted with map as expected")
}

func TestRedactWithMap_Integer(test *testing.T) {
	var jsonText = "123"
	var expected = "123"
	actual, err := RedactWithMap(jsonText, map[string]any{"123": "", "foo": "-"})
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON integer number not redacted with map as expected")
}

func TestRedactWithMap_Decimal(test *testing.T) {
	var jsonText = "123.4"
	var expected = "123.4"
	actual, err := RedactWithMap(jsonText, map[string]any{"123.4": "", "foo": "-"})
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON decimal number not redacted with map as expected")
}

func TestRedactWithMap_String(test *testing.T) {
	var jsonText = `"Hello"`
	var expected = `"Hello"`
	actual, err := RedactWithMap(jsonText, map[string]any{"Hello": "", "foo": "-"})
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON string not redacted with map as expected")
}

func TestRedactWithMap_Boolean(test *testing.T) {
	var jsonText = "true"
	var expected = "true"
	actual, err := RedactWithMap(jsonText, map[string]any{"true": false, "foo": "-"})
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON boolean not redacted with map as expected")
}

func TestRedactWithMap_Null(test *testing.T) {
	var jsonText = "null"
	var expected = "null"
	actual, err := RedactWithMap(jsonText, map[string]any{"null": nil, "foo": "-"})
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON null not redacted with map as expected")
}

func TestRedactWithMap_MixedArray0(test *testing.T) {
	var jsonText = `[123, 123.5, "Hello", null, true, {"foo": 5, "bar": 6}]`
	var expected = `[123,123.5,"Hello",null,true,{"bar":6,"foo":5}]`
	actual, err := RedactWithMap(jsonText, map[string]any{})
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON array (mixed 0) not redacted with map as expected")
}

func TestRedactWithMap_MixedArray1(test *testing.T) {
	var jsonText = `[123, 123.5, "Hello", null, true, {"foo": 5, "bar": 6}]`
	var expected = `[123,123.5,"Hello",null,true,{"bar":"","foo":5}]`
	actual, err := RedactWithMap(jsonText, map[string]any{"bar": ""})
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON array (mixed 1) not redacted with map as expected")
}

func TestRedactWithMap_MixedArray2(test *testing.T) {
	var jsonText = `[123, 123.5, "Hello", null, true, {"foo": 5, "bar": 6}]`
	var expected = `[123,123.5,"Hello",null,true,{"bar":"-","foo":""}]`
	actual, err := RedactWithMap(jsonText, map[string]any{"foo": "", "bar": "-"})
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON array (mixed 2) not redacted with map as expected")
}

func TestRedactWithMap_MixedArray3(test *testing.T) {
	var jsonText = `[123, 123.5, "Hello", null, true, {"foo": 5, "bar": 6}]`
	var expected = `[123,123.5,"Hello",null,true,{"bar":"-","foo":""}]`
	actual, err := RedactWithMap(jsonText, map[string]any{"foo": "", "bar": "-", "Hello": false})
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON array (mixed 3) not redacted with map as expected")
}

func TestRedactWithMap_BadJson(test *testing.T) {
	var jsonText = `{foo: 123, bar: "abc", phoo: true, lum: {"phoox": 3, "bax": 5}}`
	actual, err := RedactWithMap(jsonText, map[string]any{"foo": "", "bar": "-"})
	require.Error(test, err, "Invalid JSON text was redacted with map without an error: "+actual)
}

func TestRedactWithMap_Formatted0(test *testing.T) {
	var jsonText = `
	{
		"foo": 123,
		"bar": true,
		"phoo": [ {"c": 4, "a": 2}, {"a": 1, "c": [9, 0, 8]}, {"a": 1, "b": 5}]
	}`
	var expected = `{"bar":true,"foo":123,"phoo":[{"a":2,"c":4},{"a":1,"c":[9,0,8]},{"a":1,"b":5}]}`
	actual, err := RedactWithMap(jsonText, map[string]any{})
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON object (formatted 0) not redacted with map as expected")
}

func TestRedactWithMap_Formatted1(test *testing.T) {
	var jsonText = `
	{
		"foo": 123,
		"bar": true,
		"phoo": [ {"c": 4, "a": 2}, {"a": 1, "c": [9, 0, 8]}, {"a": 1, "b": 5}]
	}`
	var expected = `{"bar":true,"foo":"","phoo":[{"a":2,"c":4},{"a":1,"c":[9,0,8]},{"a":1,"b":5}]}`
	actual, err := RedactWithMap(jsonText, map[string]any{"foo": ""})
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON object (formatted 1) not redacted with map as expected")
}

func TestRedactWithMap_Formatted2(test *testing.T) {
	var jsonText = `
	{
		"foo": 123,
		"bar": true,
		"phoo": [ {"c": 4, "a": 2}, {"a": 1, "c": [9, 0, 8]}, {"a": 1, "b": 5}]
	}`
	var expected = `{"bar":true,"foo":"","phoo":[{"a":"-","c":4},{"a":"-","c":[9,0,8]},{"a":"-","b":5}]}`
	actual, err := RedactWithMap(jsonText, map[string]any{"foo": "", "a": "-"})
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON object (formatted 2) not redacted with map as expected")
}

func TestRedactWithMap_Formatted3(test *testing.T) {
	var jsonText = `
	{
		"foo": 123,
		"bar": true,
		"phoo": [ {"c": 4, "a": 2}, {"a": 1, "c": [9, 0, 8]}, {"a": 1, "b": 5}]
	}`
	var expected = `{"bar":true,"foo":"","phoo":[{"a":"-","c":"xxx"},{"a":"-","c":"xxx"},{"a":"-","b":5}]}`
	actual, err := RedactWithMap(jsonText, map[string]any{"foo": "", "a": "-", "c": "xxx"})
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON object (formatted 3) not redacted with map as expected")
}

func TestRedactWithMap_Formatted4(test *testing.T) {
	var jsonText = `
	{
		"foo": 123,
		"bar": true,
		"phoo": [ {"c": 4, "a": 2}, {"a": 1, "c": [9, 0, 8]}, {"a": 1, "b": 5}]
	}`
	var expected = `{"bar":"-","foo":"","phoo":false}`
	actual, err := RedactWithMap(jsonText, map[string]any{"foo": "", "bar": "-", "c": "xxx", "phoo": false})
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON object (formatted 4) not redacted with map as expected")
}

// ----------------------------------------------------------------------------
// Test RemoveJson function
// ----------------------------------------------------------------------------
func TestStrip_Basic0(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": "abc", "phoo": true, "lum": null}`
	var expected = `{"bar":"abc","foo":123,"lum":null,"phoo":true}`
	actual, err := Strip(jsonText)
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON object (basic 0) not redacted as expected")
}

func TestStrip_Basic1(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": "abc", "phoo": true, "lum": null}`
	var expected = `{"bar":"abc","lum":null,"phoo":true}`
	actual, err := Strip(jsonText, "foo")
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON object (basic 1) not stripped as expected")
}

func TestStrip_Basic2(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": "abc", "phoo": true, "lum": null}`
	var expected = `{"lum":null,"phoo":true}`
	actual, err := Strip(jsonText, "foo", "bar")
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON object (basic 2) not stripped as expected")
}

func TestStrip_StringArray(test *testing.T) {
	var jsonText = `["foo", "bar", "phoo", "lum"]`
	var expected = `["foo","bar","phoo","lum"]`
	actual, err := Strip(jsonText, "foo", "bar")
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON string array not stripped as expected")
}

func TestStrip_Compound0(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": [4, 6, 2], "phoo": true, "lum": {"phoox": false, "bax": 5}}`
	var expected = `{"bar":[4,6,2],"foo":123,"lum":{"bax":5,"phoox":false},"phoo":true}`
	actual, err := Strip(jsonText)
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON object (compound 0) not stripped as expected")
}

func TestStrip_Compound1(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": [4, 6, 2], "phoo": true, "lum": {"phoox": false, "bax": 5}}`
	var expected = `{"bar":[4,6,2],"foo":123,"lum":{"bax":5},"phoo":true}`
	actual, err := Strip(jsonText, "phoox")
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON object (compound 1) not stripped as expected")
}

func TestStrip_Compound2(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": [4, 6, 2], "phoo": true, "lum": {"phoox": false, "bax": 5}}`
	var expected = `{"foo":123,"lum":{"bax":5},"phoo":true}`
	actual, err := Strip(jsonText, "phoox", "bar")
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON object (compound 2) not stripped as expected")
}

func TestStrip_Compound3(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": [4, 6, 2], "phoox": true, "lum": {"phoox": false, "bax": 5}}`
	var expected = `{"foo":123,"lum":{}}`
	actual, err := Strip(jsonText, "phoox", "bar", "bax")
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON object (compound 3) not stripped as expected")
}

func TestStrip_Compound4(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": [4, 6, 2], "phoo": true, "lum": {"phoox": false, "bax": 5}}`
	var expected = `{"lum":{},"phoo":true}`
	actual, err := Strip(jsonText, "phoox", "bar", "bax", "foo")
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON object (compound 4) not stripped as expected")
}

func TestStrip_Compound5(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": [4, 6, 2], "phoo": true, "lum": {"phoox": false, "bax": 5}}`
	var expected = `{}`
	actual, err := Strip(jsonText, "phoox", "bar", "foo", "phoo", "lum")
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON object (compound 5) not stripped as expected")
}

func TestStrip_Integer(test *testing.T) {
	var jsonText = "123"
	var expected = "123"
	actual, err := Strip(jsonText, "123", "foo")
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON integer number not stripped as expected")
}

func TestStrip_Decimal(test *testing.T) {
	var jsonText = "123.4"
	var expected = "123.4"
	actual, err := Strip(jsonText, "123.4", "foo")
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON decimal number not stripped as expected")
}

func TestStrip_String(test *testing.T) {
	var jsonText = `"Hello"`
	var expected = `"Hello"`
	actual, err := Strip(jsonText, "Hello", "foo")
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON string not stripped as expected")
}

func TestStrip_Boolean(test *testing.T) {
	var jsonText = "true"
	var expected = "true"
	actual, err := Strip(jsonText, "true", "foo")
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON boolean not stripped as expected")
}

func TestStrip_Null(test *testing.T) {
	var jsonText = "null"
	var expected = "null"
	actual, err := Strip(jsonText, "null", "foo")
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON null not stripped as expected")
}

func TestStrip_MixedArray0(test *testing.T) {
	var jsonText = `[123, 123.5, "Hello", null, true, {"foo": 5, "bar": 6}]`
	var expected = `[123,123.5,"Hello",null,true,{"bar":6,"foo":5}]`
	actual, err := Strip(jsonText)
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON array (mixed 0) not stripped as expected")
}

func TestStrip_MixedArray1(test *testing.T) {
	var jsonText = `[123, 123.5, "Hello", null, true, {"foo": 5, "bar": 6}]`
	var expected = `[123,123.5,"Hello",null,true,{"foo":5}]`
	actual, err := Strip(jsonText, "bar")
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON array (mixed 1) not stripped as expected")
}

func TestStrip_MixedArray2(test *testing.T) {
	var jsonText = `[123, 123.5, "Hello", null, true, {"foo": 5, "bar": 6}]`
	var expected = `[123,123.5,"Hello",null,true,{}]`
	actual, err := Strip(jsonText, "foo", "bar")
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON array (mixed 2) not stripped as expected")
}

func TestStrip_MixedArray3(test *testing.T) {
	var jsonText = `[123, 123.5, "Hello", null, true, {"foo": 5, "bar": 6}]`
	var expected = `[123,123.5,"Hello",null,true,{}]`
	actual, err := Strip(jsonText, "foo", "bar", "Hello")
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON array (mixed 3) not stripped as expected")
}

func TestStrip_BadJson(test *testing.T) {
	var jsonText = `{foo: 123, bar: "abc", phoo: true, lum: {"phoox": 3, "bax": 5}}`
	actual, err := Strip(jsonText, "foo", "bar")
	require.Error(test, err, "Invalid JSON text was stripped without an error: "+actual)
}

func TestStrip_Formatted0(test *testing.T) {
	var jsonText = `
	{
		"foo": 123,
		"bar": true,
		"phoo": [ {"c": 4, "a": 2}, {"a": 1, "c": [9, 0, 8]}, {"a": 1, "b": 5}]
	}`
	var expected = `{"bar":true,"foo":123,"phoo":[{"a":2,"c":4},{"a":1,"c":[9,0,8]},{"a":1,"b":5}]}`
	actual, err := Strip(jsonText)
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON object (formatted 0) not stripped as expected")
}

func TestStrip_Formatted1(test *testing.T) {
	var jsonText = `
	{
		"foo": 123,
		"bar": true,
		"phoo": [ {"c": 4, "a": 2}, {"a": 1, "c": [9, 0, 8]}, {"a": 1, "b": 5}]
	}`
	var expected = `{"bar":true,"phoo":[{"a":2,"c":4},{"a":1,"c":[9,0,8]},{"a":1,"b":5}]}`
	actual, err := Strip(jsonText, "foo")
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON object (formatted 1) not stripped as expected")
}

func TestStrip_Formatted2(test *testing.T) {
	var jsonText = `
	{
		"foo": 123,
		"bar": true,
		"phoo": [ {"c": 4, "a": 2}, {"a": 1, "c": [9, 0, 8]}, {"a": 1, "b": 5}]
	}`
	var expected = `{"bar":true,"phoo":[{"c":4},{"c":[9,0,8]},{"b":5}]}`
	actual, err := Strip(jsonText, "foo", "a")
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON object (formatted 2) not stripped as expected")
}

func TestStrip_Formatted3(test *testing.T) {
	var jsonText = `
	{
		"foo": 123,
		"bar": true,
		"phoo": [ {"c": 4, "a": 2}, {"a": 1, "c": [9, 0, 8]}, {"a": 1, "b": 5}]
	}`
	var expected = `{"bar":true,"phoo":[{},{},{"b":5}]}`
	actual, err := Strip(jsonText, "foo", "a", "c")
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON object (formatted 3) not stripped as expected")
}

func TestStrip_Formatted4(test *testing.T) {
	var jsonText = `
	{
		"foo": 123,
		"bar": true,
		"phoo": [ {"c": 4, "a": 2}, {"a": 1, "c": [9, 0, 8]}, {"a": 1, "b": 5}]
	}`
	var expected = `{}`
	actual, err := Strip(jsonText, "foo", "bar", "c", "phoo")
	testError(test, err)
	assert.Equal(test, expected, actual, "JSON object (formatted 4) not stripped as expected")
}

// ----------------------------------------------------------------------------
// Test Flatten function
// ----------------------------------------------------------------------------
func TestFlatten_NoError(test *testing.T) {
	actual := Flatten(`{"foo": 5, "bar": 6}`, nil)
	var expected = `{"foo": 5, "bar": 6}`
	assert.Equal(test, expected, actual, "Flattening without an error did not work as expected: "+actual)
}

func TestFlatten_WithError(test *testing.T) {
	err := errors.New("failed")
	actual := Flatten(`{"foo": 5, "bar": 6}`, err)
	var expected = `{"error":"failed","text":"{\"foo\": 5, \"bar\": 6}"}`
	assert.Equal(test, expected, actual, "Flattening with an error did not work as expected: "+actual)
}
