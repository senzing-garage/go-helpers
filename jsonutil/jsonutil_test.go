package jsonutil_test

import (
	"errors"
	"testing"

	"github.com/senzing-garage/go-helpers/jsonutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// ----------------------------------------------------------------------------
// Test Flatten function
// ----------------------------------------------------------------------------

func TestFlatten_NoError(test *testing.T) {
	actual := jsonutil.Flatten(`{"foo": 5, "bar": 6}`, nil)

	var expected = `{"foo": 5, "bar": 6}`

	assert.Equal(test, expected, actual, "Flattening without an error did not work as expected: "+actual)
}

func TestFlatten_WithError(test *testing.T) {
	err := errors.New("failed")
	actual := jsonutil.Flatten(`{"foo": 5, "bar": 6}`, err)

	var expected = `{"error":"failed","text":"{\"foo\": 5, \"bar\": 6}"}`

	assert.Equal(test, expected, actual, "Flattening with an error did not work as expected: "+actual)
}

// ----------------------------------------------------------------------------
// Test IsJson function
// ----------------------------------------------------------------------------

func TestIsJson_Basic(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": "abc", "phoo": true, "lum": 20.5}`

	var expected = true

	actual := jsonutil.IsJSON(jsonText)
	assert.Equal(test, expected, actual, "JSON object (basic) not recognized as JSON")
}

func TestIsJson_Compound(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": "abc", "phoo": true, "lum": {"phoox": 3, "bax": 5}}`

	var expected = true

	actual := jsonutil.IsJSON(jsonText)
	assert.Equal(test, expected, actual, "JSON object (compound) not recognized as JSON")
}

func TestIsJson_Integer(test *testing.T) {
	var jsonText = "123"

	var expected = true

	actual := jsonutil.IsJSON(jsonText)
	assert.Equal(test, expected, actual, "JSON integer not recognized as JSON")
}

func TestIsJson_Decimal(test *testing.T) {
	var jsonText = "123.4"

	var expected = true

	actual := jsonutil.IsJSON(jsonText)
	assert.Equal(test, expected, actual, "JSON decimal number not recognized as JSON")
}

func TestIsJson_String(test *testing.T) {
	var jsonText = `"Hello"`

	var expected = true

	actual := jsonutil.IsJSON(jsonText)
	assert.Equal(test, expected, actual, "JSON string not recognized as JSON")
}

func TestIsJson_Boolean(test *testing.T) {
	var jsonText = "true"

	var expected = true

	actual := jsonutil.IsJSON(jsonText)
	assert.Equal(test, expected, actual, "JSON boolean not recognized as JSON")
}

func TestIsJson_Null(test *testing.T) {
	var jsonText = "null"

	var expected = true

	actual := jsonutil.IsJSON(jsonText)
	assert.Equal(test, expected, actual, "JSON null not recognized as JSON")
}

func TestIsJson_Array(test *testing.T) {
	var jsonText = `[123, 123.5, "Hello", true, {"foo": 5, "bar": 6}]`

	var expected = true

	actual := jsonutil.IsJSON(jsonText)
	assert.Equal(test, expected, actual, "JSON array not recognized as JSON")
}

func TestIsJson_BadJson(test *testing.T) {
	var jsonText = `{foo: 123, bar: "abc", phoo: true, lum: {"phoox": 3, "bax": 5}}`

	var expected = false

	actual := jsonutil.IsJSON(jsonText)
	assert.Equal(test, expected, actual, "Invalid JSON text incorrectly recognized as JSON")
}

func TestIsJson_Formatted(test *testing.T) {
	var jsonText = `
	{
		"foo": 123,
		"bar": true
	}`

	var expected = true

	actual := jsonutil.IsJSON(jsonText)
	assert.Equal(test, expected, actual, "Formatted JSON object not recognized as JSON")
}

// ----------------------------------------------------------------------------
// Test NormalizeJson function
// ----------------------------------------------------------------------------

func TestNormalize_Basic(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": "abc", "phoo": true, "lum": null}`

	var expected = `{"bar":"abc","foo":123,"lum":null,"phoo":true}`

	actual, err := jsonutil.Normalize(jsonText)
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON object (basic) not normalized as expected")
}

func TestNormalize_Compound(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": "abc", "phoo": true, "lum": {"phoox": null, "bax": 5}}`

	var expected = `{"bar":"abc","foo":123,"lum":{"bax":5,"phoox":null},"phoo":true}`

	actual, err := jsonutil.Normalize(jsonText)
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON object (compound) not normalized as expected")
}

func TestNormalize_Integer(test *testing.T) {
	var jsonText = "123"

	var expected = "123"

	actual, err := jsonutil.Normalize(jsonText)
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON integer number not normalized as expected")
}

func TestNormalize_Decimal(test *testing.T) {
	var jsonText = "123.4"

	var expected = "123.4"

	actual, err := jsonutil.Normalize(jsonText)
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON decimal number not normalized as expected")
}

func TestNormalize_String(test *testing.T) {
	var jsonText = `"Hello"`

	var expected = `"Hello"`

	actual, err := jsonutil.Normalize(jsonText)
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON string not normalized as expected")
}

func TestNormalize_Boolean(test *testing.T) {
	var jsonText = "true"

	var expected = "true"

	actual, err := jsonutil.Normalize(jsonText)
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON boolean not normalized as expected")
}

func TestNormalize_Null(test *testing.T) {
	var jsonText = "null"

	var expected = "null"

	actual, err := jsonutil.Normalize(jsonText)
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON null not normalized as expected")
}

func TestNormalize_Array(test *testing.T) {
	var jsonText = `[123, 123.5, "Hello", true, {"foo": 5, "bar": 6}]`

	var expected = `[123,123.5,"Hello",true,{"bar":6,"foo":5}]`

	actual, err := jsonutil.Normalize(jsonText)
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON array not normalized as expected")
}

func TestNormalize_BadJson(test *testing.T) {
	var jsonText = `{foo: 123, bar: "abc", phoo: true, lum: {"phoox": 3, "bax": 5}}`
	actual, err := jsonutil.Normalize(jsonText)
	require.Error(test, err, "Invalid JSON text was normalized without an error: "+actual)
}

func TestNormalize_Formatted(test *testing.T) {
	var jsonText = `
	{
		"foo": 123,
		"bar": true
	}`

	var expected = `{"bar":true,"foo":123}`

	actual, err := jsonutil.Normalize(jsonText)
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON object (formatted) not normalized as expected")
}

// ----------------------------------------------------------------------------
// Test NormalizeAndSortJson function
// ----------------------------------------------------------------------------

func TestNormalizeAndSort_Basic(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": "abc", "phoo": true, "lum": null}`

	var expected = `{"bar":"abc","foo":123,"lum":null,"phoo":true}`

	actual, err := jsonutil.NormalizeAndSort(jsonText)
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON object (basic) not normalized as expected")
}

func TestNormalizeAndSort_StringArray(test *testing.T) {
	var jsonText = `["foo", "bar", "phoo", "lum"]`

	var expected = `["bar","foo","lum","phoo"]`

	actual, err := jsonutil.NormalizeAndSort(jsonText)
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON string array not normalized as expected")
}

func TestNormalizeAndSort_Compound(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": [4, 6, 2], "phoo": true, "lum": {"phoox": null, "bax": 5}}`

	var expected = `{"bar":[2,4,6],"foo":123,"lum":{"bax":5,"phoox":null},"phoo":true}`

	actual, err := jsonutil.NormalizeAndSort(jsonText)
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON object (compound) not normalized as expected")
}

func TestNormalizeAndSort_Integer(test *testing.T) {
	var jsonText = "123"

	var expected = "123"

	actual, err := jsonutil.NormalizeAndSort(jsonText)
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON integer number not normalized as expected")
}

func TestNormalizeAndSort_Decimal(test *testing.T) {
	var jsonText = "123.4"

	var expected = "123.4"

	actual, err := jsonutil.NormalizeAndSort(jsonText)
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON decimal number not normalized as expected")
}

func TestNormalizeAndSort_String(test *testing.T) {
	var jsonText = `"Hello"`

	var expected = `"Hello"`

	actual, err := jsonutil.NormalizeAndSort(jsonText)
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON string not normalized as expected")
}

func TestNormalizeAndSort_Boolean(test *testing.T) {
	var jsonText = "true"

	var expected = "true"

	actual, err := jsonutil.NormalizeAndSort(jsonText)
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON boolean not normalized as expected")
}

func TestNormalizeAndSort_Null(test *testing.T) {
	var jsonText = "null"

	var expected = "null"

	actual, err := jsonutil.NormalizeAndSort(jsonText)
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON null not normalized as expected")
}

func TestNormalizeAndSort_MixedArray(test *testing.T) {
	var jsonText = `[123, 123.5, "Hello", null, true, {"foo": 5, "bar": 6}]`

	var expected = `[null,"Hello",123,123.5,true,{"bar":6,"foo":5}]`

	actual, err := jsonutil.NormalizeAndSort(jsonText)
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON array not normalized as expected")
}

func TestNormalizeAndSort_BadJson(test *testing.T) {
	var jsonText = `{foo: 123, bar: "abc", phoo: true, lum: {"phoox": 3, "bax": 5}}`
	actual, err := jsonutil.NormalizeAndSort(jsonText)
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

	actual, err := jsonutil.NormalizeAndSort(jsonText)
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON object (formatted) not normalized as expected")
}

// ----------------------------------------------------------------------------
// Test PrettyPrint function
// ----------------------------------------------------------------------------

func TestPrettyPrint(test *testing.T) {
	var jsonText = `{"bar":true,"foo":123,"phoo":[{"a":1,"b":5},{"a":1,"c":[0,8,9]},{"a":2,"c":4}]}`

	var expected = `{
	"bar": true,
	"foo": 123,
	"phoo": [
		{
			"a": 1,
			"b": 5
		},
		{
			"a": 1,
			"c": [
				0,
				8,
				9
			]
		},
		{
			"a": 2,
			"c": 4
		}
	]
}`

	actual := jsonutil.PrettyPrint(jsonText, "	")
	assert.Equal(test, expected, actual, "JSON object (formatted) not pretty printed as expected")
}

func TestPrettyPrint_UsingSpaces(test *testing.T) {
	var jsonText = `{"bar":true,"foo":123,"phoo":[{"a":1,"b":5},{"a":1,"c":[0,8,9]},{"a":2,"c":4}]}`

	var expected = `{
    "bar": true,
    "foo": 123,
    "phoo": [
        {
            "a": 1,
            "b": 5
        },
        {
            "a": 1,
            "c": [
                0,
                8,
                9
            ]
        },
        {
            "a": 2,
            "c": 4
        }
    ]
}`

	actual := jsonutil.PrettyPrint(jsonText, "    ")
	assert.Equal(test, expected, actual, "JSON object (formatted) not pretty printed as expected")
}

// ----------------------------------------------------------------------------
// Test RedactJson function
// ----------------------------------------------------------------------------

func TestRedact_Basic0(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": "abc", "phoo": true, "lum": null}`

	var expected = `{"bar":"abc","foo":123,"lum":null,"phoo":true}`

	actual, err := jsonutil.Redact(jsonText)
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON object (basic 0) not redacted as expected")
}

func TestRedact_Basic1(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": "abc", "phoo": true, "lum": null}`

	var expected = `{"bar":"abc","foo":null,"lum":null,"phoo":true}`

	actual, err := jsonutil.Redact(jsonText, "foo")
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON object (basic 1) not redacted as expected")
}

func TestRedact_Basic2(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": "abc", "phoo": true, "lum": null}`

	var expected = `{"bar":null,"foo":null,"lum":null,"phoo":true}`

	actual, err := jsonutil.Redact(jsonText, "foo", "bar")
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON object (basic 2) not redacted as expected")
}

func TestRedact_StringArray(test *testing.T) {
	var jsonText = `["foo", "bar", "phoo", "lum"]`

	var expected = `["foo","bar","phoo","lum"]`

	actual, err := jsonutil.Redact(jsonText, "foo", "bar")
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON string array not redacted as expected")
}

func TestRedact_Compound0(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": [4, 6, 2], "phoo": true, "lum": {"phoox": false, "bax": 5}}`

	var expected = `{"bar":[4,6,2],"foo":123,"lum":{"bax":5,"phoox":false},"phoo":true}`

	actual, err := jsonutil.Redact(jsonText)
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON object (compound 0) not redacted as expected")
}

func TestRedact_Compound1(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": [4, 6, 2], "phoo": true, "lum": {"phoox": false, "bax": 5}}`

	var expected = `{"bar":[4,6,2],"foo":123,"lum":{"bax":5,"phoox":null},"phoo":true}`

	actual, err := jsonutil.Redact(jsonText, "phoox")
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON object (compound 1) not redacted as expected")
}

func TestRedact_Compound2(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": [4, 6, 2], "phoo": true, "lum": {"phoox": false, "bax": 5}}`

	var expected = `{"bar":null,"foo":123,"lum":{"bax":5,"phoox":null},"phoo":true}`

	actual, err := jsonutil.Redact(jsonText, "phoox", "bar")
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON object (compound 2) not redacted as expected")
}

func TestRedact_Compound3(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": [4, 6, 2], "phoox": true, "lum": {"phoox": false, "bax": 5}}`

	var expected = `{"bar":null,"foo":123,"lum":{"bax":null,"phoox":null},"phoox":null}`

	actual, err := jsonutil.Redact(jsonText, "phoox", "bar", "bax")
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON object (compound 3) not redacted as expected")
}

func TestRedact_Compound4(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": [4, 6, 2], "phoo": true, "lum": {"phoox": false, "bax": 5}}`

	var expected = `{"bar":null,"foo":null,"lum":{"bax":null,"phoox":null},"phoo":true}`

	actual, err := jsonutil.Redact(jsonText, "phoox", "bar", "bax", "foo")
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON object (compound 4) not redacted as expected")
}

func TestRedact_Compound5(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": [4, 6, 2], "phoo": true, "lum": {"phoox": false, "bax": 5}}`

	var expected = `{"bar":null,"foo":null,"lum":null,"phoo":null}`

	actual, err := jsonutil.Redact(jsonText, "phoox", "bar", "foo", "phoo", "lum")
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON object (compound 5) not redacted as expected")
}

func TestRedact_Integer(test *testing.T) {
	var jsonText = "123"

	var expected = "123"

	actual, err := jsonutil.Redact(jsonText, "123", "foo")
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON integer number not redacted as expected")
}

func TestRedact_Decimal(test *testing.T) {
	var jsonText = "123.4"

	var expected = "123.4"

	actual, err := jsonutil.Redact(jsonText, "123.4", "foo")
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON decimal number not redacted as expected")
}

func TestRedact_String(test *testing.T) {
	var jsonText = `"Hello"`

	var expected = `"Hello"`

	actual, err := jsonutil.Redact(jsonText, "Hello", "foo")
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON string not redacted as expected")
}

func TestRedact_Boolean(test *testing.T) {
	var jsonText = "true"

	var expected = "true"

	actual, err := jsonutil.Redact(jsonText, "true", "foo")
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON boolean not redacted as expected")
}

func TestRedact_Null(test *testing.T) {
	var jsonText = "null"

	var expected = "null"

	actual, err := jsonutil.Redact(jsonText, "null", "foo")
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON null not redacted as expected")
}

func TestRedact_MixedArray0(test *testing.T) {
	var jsonText = `[123, 123.5, "Hello", null, true, {"foo": 5, "bar": 6}]`

	var expected = `[123,123.5,"Hello",null,true,{"bar":6,"foo":5}]`

	actual, err := jsonutil.Redact(jsonText)
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON array (mixed 0) not redacted as expected")
}

func TestRedact_MixedArray1(test *testing.T) {
	var jsonText = `[123, 123.5, "Hello", null, true, {"foo": 5, "bar": 6}]`

	var expected = `[123,123.5,"Hello",null,true,{"bar":null,"foo":5}]`

	actual, err := jsonutil.Redact(jsonText, "bar")
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON array (mixed 1) not redacted as expected")
}

func TestRedact_MixedArray2(test *testing.T) {
	var jsonText = `[123, 123.5, "Hello", null, true, {"foo": 5, "bar": 6}]`

	var expected = `[123,123.5,"Hello",null,true,{"bar":null,"foo":null}]`

	actual, err := jsonutil.Redact(jsonText, "foo", "bar")
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON array (mixed 2) not redacted as expected")
}

func TestRedact_MixedArray3(test *testing.T) {
	var jsonText = `[123, 123.5, "Hello", null, true, {"foo": 5, "bar": 6}]`

	var expected = `[123,123.5,"Hello",null,true,{"bar":null,"foo":null}]`

	actual, err := jsonutil.Redact(jsonText, "foo", "bar", "Hello")
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON array (mixed 3) not redacted as expected")
}

func TestRedact_BadJson(test *testing.T) {
	var jsonText = `{foo: 123, bar: "abc", phoo: true, lum: {"phoox": 3, "bax": 5}}`
	actual, err := jsonutil.Redact(jsonText, "foo", "bar")
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

	actual, err := jsonutil.Redact(jsonText)
	require.NoError(test, err)
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

	actual, err := jsonutil.Redact(jsonText, "foo")
	require.NoError(test, err)
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

	actual, err := jsonutil.Redact(jsonText, "foo", "a")
	require.NoError(test, err)
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

	actual, err := jsonutil.Redact(jsonText, "foo", "a", "c")
	require.NoError(test, err)
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

	actual, err := jsonutil.Redact(jsonText, "foo", "bar", "c", "phoo")
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON object (formatted 4) not redacted as expected")
}

// ----------------------------------------------------------------------------
// Test RedactWithMap function
// ----------------------------------------------------------------------------

func TestRedactWithMap_Basic0(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": "abc", "phoo": true, "lum": null}`

	var expected = `{"bar":"abc","foo":123,"lum":null,"phoo":true}`

	actual, err := jsonutil.RedactWithMap(jsonText, map[string]any{})
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON object (basic 0) not redacted with map as expected")
}

func TestRedactWithMap_Basic1(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": "abc", "phoo": true, "lum": null}`

	var expected = `{"bar":"abc","foo":"","lum":null,"phoo":true}`

	actual, err := jsonutil.RedactWithMap(jsonText, map[string]any{"foo": ""})
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON object (basic 1) not redacted with map as expected")
}

func TestRedactWithMap_Basic2(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": "abc", "phoo": true, "lum": null}`

	var expected = `{"bar":"-","foo":"","lum":null,"phoo":true}`

	actual, err := jsonutil.RedactWithMap(jsonText, map[string]any{"foo": "", "bar": "-"})
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON object (basic 2) not redacted with map as expected")
}

func TestRedactWithMap_StringArray(test *testing.T) {
	var jsonText = `["foo", "bar", "phoo", "lum"]`

	var expected = `["foo","bar","phoo","lum"]`

	actual, err := jsonutil.RedactWithMap(jsonText, map[string]any{"foo": "", "bar": "-"})
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON string array not redacted with map as expected")
}

func TestRedactWithMap_Compound0(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": [4, 6, 2], "phoo": true, "lum": {"phoox": false, "bax": 5}}`

	var expected = `{"bar":[4,6,2],"foo":123,"lum":{"bax":5,"phoox":false},"phoo":true}`

	actual, err := jsonutil.RedactWithMap(jsonText, map[string]any{})
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON object (compound 0) not redacted with map as expected")
}

func TestRedactWithMap_Compound1(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": [4, 6, 2], "phoo": true, "lum": {"phoox": false, "bax": 5}}`

	var expected = `{"bar":[4,6,2],"foo":123,"lum":{"bax":5,"phoox":""},"phoo":true}`

	actual, err := jsonutil.RedactWithMap(jsonText, map[string]any{"phoox": ""})
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON object (compound 1) not redacted with map as expected")
}

func TestRedactWithMap_Compound2(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": [4, 6, 2], "phoox": true, "lum": {"phoox": false, "bax": 5}}`

	var expected = `{"bar":"-","foo":123,"lum":{"bax":5,"phoox":""},"phoox":""}`

	actual, err := jsonutil.RedactWithMap(jsonText, map[string]any{"phoox": "", "bar": "-"})
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON object (compound 2) not redacted with map as expected")
}

func TestRedactWithMap_Compound3(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": [4, 6, 2], "phoo": true, "lum": {"phoox": false, "bax": 5}}`

	var expected = `{"bar":"-","foo":123,"lum":{"bax":null,"phoox":""},"phoo":true}`

	actual, err := jsonutil.RedactWithMap(jsonText, map[string]any{"phoox": "", "bar": "-", "bax": nil})
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON object (compound 3) not redacted with map as expected")
}

func TestRedactWithMap_Compound4(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": [4, 6, 2], "phoo": true, "lum": {"phoox": false, "bax": 5}}`

	var expected = `{"bar":"-","foo":"xxx","lum":{"bax":null,"phoox":""},"phoo":true}`

	actual, err := jsonutil.RedactWithMap(jsonText, map[string]any{"phoox": "", "bar": "-", "bax": nil, "foo": "xxx"})
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON object (compound 4) not redacted with map as expected")
}

func TestRedactWithMap_Compound5(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": [4, 6, 2], "phoo": true, "lum": {"phoox": false, "bax": 5}}`

	var expected = `{"bar":"-","foo":null,"lum":false,"phoo":"xxx"}`

	actual, err := jsonutil.RedactWithMap(jsonText,
		map[string]any{"phoox": "", "bar": "-", "foo": nil, "phoo": "xxx", "lum": false})
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON object (compound 5) not redacted with map as expected")
}

func TestRedactWithMap_Integer(test *testing.T) {
	var jsonText = "123"

	var expected = "123"

	actual, err := jsonutil.RedactWithMap(jsonText, map[string]any{"123": "", "foo": "-"})
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON integer number not redacted with map as expected")
}

func TestRedactWithMap_Decimal(test *testing.T) {
	var jsonText = "123.4"

	var expected = "123.4"

	actual, err := jsonutil.RedactWithMap(jsonText, map[string]any{"123.4": "", "foo": "-"})
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON decimal number not redacted with map as expected")
}

func TestRedactWithMap_String(test *testing.T) {
	var jsonText = `"Hello"`

	var expected = `"Hello"`

	actual, err := jsonutil.RedactWithMap(jsonText, map[string]any{"Hello": "", "foo": "-"})
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON string not redacted with map as expected")
}

func TestRedactWithMap_Boolean(test *testing.T) {
	var jsonText = "true"

	var expected = "true"

	actual, err := jsonutil.RedactWithMap(jsonText, map[string]any{"true": false, "foo": "-"})
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON boolean not redacted with map as expected")
}

func TestRedactWithMap_Null(test *testing.T) {
	var jsonText = "null"

	var expected = "null"

	actual, err := jsonutil.RedactWithMap(jsonText, map[string]any{"null": nil, "foo": "-"})
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON null not redacted with map as expected")
}

func TestRedactWithMap_MixedArray0(test *testing.T) {
	var jsonText = `[123, 123.5, "Hello", null, true, {"foo": 5, "bar": 6}]`

	var expected = `[123,123.5,"Hello",null,true,{"bar":6,"foo":5}]`

	actual, err := jsonutil.RedactWithMap(jsonText, map[string]any{})
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON array (mixed 0) not redacted with map as expected")
}

func TestRedactWithMap_MixedArray1(test *testing.T) {
	var jsonText = `[123, 123.5, "Hello", null, true, {"foo": 5, "bar": 6}]`

	var expected = `[123,123.5,"Hello",null,true,{"bar":"","foo":5}]`

	actual, err := jsonutil.RedactWithMap(jsonText, map[string]any{"bar": ""})
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON array (mixed 1) not redacted with map as expected")
}

func TestRedactWithMap_MixedArray2(test *testing.T) {
	var jsonText = `[123, 123.5, "Hello", null, true, {"foo": 5, "bar": 6}]`

	var expected = `[123,123.5,"Hello",null,true,{"bar":"-","foo":""}]`

	actual, err := jsonutil.RedactWithMap(jsonText, map[string]any{"foo": "", "bar": "-"})
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON array (mixed 2) not redacted with map as expected")
}

func TestRedactWithMap_MixedArray3(test *testing.T) {
	var jsonText = `[123, 123.5, "Hello", null, true, {"foo": 5, "bar": 6}]`

	var expected = `[123,123.5,"Hello",null,true,{"bar":"-","foo":""}]`

	actual, err := jsonutil.RedactWithMap(jsonText, map[string]any{"foo": "", "bar": "-", "Hello": false})
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON array (mixed 3) not redacted with map as expected")
}

func TestRedactWithMap_BadJson(test *testing.T) {
	var jsonText = `{foo: 123, bar: "abc", phoo: true, lum: {"phoox": 3, "bax": 5}}`
	actual, err := jsonutil.RedactWithMap(jsonText, map[string]any{"foo": "", "bar": "-"})
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

	actual, err := jsonutil.RedactWithMap(jsonText, map[string]any{})
	require.NoError(test, err)
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

	actual, err := jsonutil.RedactWithMap(jsonText, map[string]any{"foo": ""})
	require.NoError(test, err)
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

	actual, err := jsonutil.RedactWithMap(jsonText, map[string]any{"foo": "", "a": "-"})
	require.NoError(test, err)
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

	actual, err := jsonutil.RedactWithMap(jsonText, map[string]any{"foo": "", "a": "-", "c": "xxx"})
	require.NoError(test, err)
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

	actual, err := jsonutil.RedactWithMap(jsonText, map[string]any{"foo": "", "bar": "-", "c": "xxx", "phoo": false})
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON object (formatted 4) not redacted with map as expected")
}

// ----------------------------------------------------------------------------
// Test ReferseString function
// ----------------------------------------------------------------------------

type testReverseStringMetadata struct {
	name           string
	reversedString string
	theString      string
}

var testReverseStringCases = []testReverseStringMetadata{
	{
		name:           "001",
		theString:      "abc",
		reversedString: "cba",
	},
	{
		name:           "002",
		theString:      "{}",
		reversedString: "}{",
	},
	{
		name:           "003",
		theString:      `{"alpha": "beta"}`,
		reversedString: `}"ateb" :"ahpla"{`,
	},
}

func TestReverseString(test *testing.T) {
	for _, testCase := range testReverseStringCases {
		test.Run(testCase.name, func(test *testing.T) {
			actual := jsonutil.ReverseString(testCase.theString)
			assert.Equal(test, testCase.reversedString, actual)
		})
	}
}

// ----------------------------------------------------------------------------
// Test Strip function
// ----------------------------------------------------------------------------

func TestStrip_Basic0(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": "abc", "phoo": true, "lum": null}`

	var expected = `{"bar":"abc","foo":123,"lum":null,"phoo":true}`

	actual, err := jsonutil.Strip(jsonText)
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON object (basic 0) not redacted as expected")
}

func TestStrip_Basic1(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": "abc", "phoo": true, "lum": null}`

	var expected = `{"bar":"abc","lum":null,"phoo":true}`

	actual, err := jsonutil.Strip(jsonText, "foo")
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON object (basic 1) not stripped as expected")
}

func TestStrip_Basic2(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": "abc", "phoo": true, "lum": null}`

	var expected = `{"lum":null,"phoo":true}`

	actual, err := jsonutil.Strip(jsonText, "foo", "bar")
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON object (basic 2) not stripped as expected")
}

func TestStrip_StringArray(test *testing.T) {
	var jsonText = `["foo", "bar", "phoo", "lum"]`

	var expected = `["foo","bar","phoo","lum"]`

	actual, err := jsonutil.Strip(jsonText, "foo", "bar")
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON string array not stripped as expected")
}

func TestStrip_Compound0(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": [4, 6, 2], "phoo": true, "lum": {"phoox": false, "bax": 5}}`

	var expected = `{"bar":[4,6,2],"foo":123,"lum":{"bax":5,"phoox":false},"phoo":true}`

	actual, err := jsonutil.Strip(jsonText)
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON object (compound 0) not stripped as expected")
}

func TestStrip_Compound1(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": [4, 6, 2], "phoo": true, "lum": {"phoox": false, "bax": 5}}`

	var expected = `{"bar":[4,6,2],"foo":123,"lum":{"bax":5},"phoo":true}`

	actual, err := jsonutil.Strip(jsonText, "phoox")
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON object (compound 1) not stripped as expected")
}

func TestStrip_Compound2(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": [4, 6, 2], "phoo": true, "lum": {"phoox": false, "bax": 5}}`

	var expected = `{"foo":123,"lum":{"bax":5},"phoo":true}`

	actual, err := jsonutil.Strip(jsonText, "phoox", "bar")
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON object (compound 2) not stripped as expected")
}

func TestStrip_Compound3(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": [4, 6, 2], "phoox": true, "lum": {"phoox": false, "bax": 5}}`

	var expected = `{"foo":123,"lum":{}}`

	actual, err := jsonutil.Strip(jsonText, "phoox", "bar", "bax")
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON object (compound 3) not stripped as expected")
}

func TestStrip_Compound4(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": [4, 6, 2], "phoo": true, "lum": {"phoox": false, "bax": 5}}`

	var expected = `{"lum":{},"phoo":true}`

	actual, err := jsonutil.Strip(jsonText, "phoox", "bar", "bax", "foo")
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON object (compound 4) not stripped as expected")
}

func TestStrip_Compound5(test *testing.T) {
	var jsonText = `{"foo": 123, "bar": [4, 6, 2], "phoo": true, "lum": {"phoox": false, "bax": 5}}`

	var expected = `{}`

	actual, err := jsonutil.Strip(jsonText, "phoox", "bar", "foo", "phoo", "lum")
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON object (compound 5) not stripped as expected")
}

func TestStrip_Integer(test *testing.T) {
	var jsonText = "123"

	var expected = "123"

	actual, err := jsonutil.Strip(jsonText, "123", "foo")
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON integer number not stripped as expected")
}

func TestStrip_Decimal(test *testing.T) {
	var jsonText = "123.4"

	var expected = "123.4"

	actual, err := jsonutil.Strip(jsonText, "123.4", "foo")
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON decimal number not stripped as expected")
}

func TestStrip_String(test *testing.T) {
	var jsonText = `"Hello"`

	var expected = `"Hello"`

	actual, err := jsonutil.Strip(jsonText, "Hello", "foo")
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON string not stripped as expected")
}

func TestStrip_Boolean(test *testing.T) {
	var jsonText = "true"

	var expected = "true"

	actual, err := jsonutil.Strip(jsonText, "true", "foo")
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON boolean not stripped as expected")
}

func TestStrip_Null(test *testing.T) {
	var jsonText = "null"

	var expected = "null"

	actual, err := jsonutil.Strip(jsonText, "null", "foo")
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON null not stripped as expected")
}

func TestStrip_MixedArray0(test *testing.T) {
	var jsonText = `[123, 123.5, "Hello", null, true, {"foo": 5, "bar": 6}]`

	var expected = `[123,123.5,"Hello",null,true,{"bar":6,"foo":5}]`

	actual, err := jsonutil.Strip(jsonText)
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON array (mixed 0) not stripped as expected")
}

func TestStrip_MixedArray1(test *testing.T) {
	var jsonText = `[123, 123.5, "Hello", null, true, {"foo": 5, "bar": 6}]`

	var expected = `[123,123.5,"Hello",null,true,{"foo":5}]`

	actual, err := jsonutil.Strip(jsonText, "bar")
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON array (mixed 1) not stripped as expected")
}

func TestStrip_MixedArray2(test *testing.T) {
	var jsonText = `[123, 123.5, "Hello", null, true, {"foo": 5, "bar": 6}]`

	var expected = `[123,123.5,"Hello",null,true,{}]`

	actual, err := jsonutil.Strip(jsonText, "foo", "bar")
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON array (mixed 2) not stripped as expected")
}

func TestStrip_MixedArray3(test *testing.T) {
	var jsonText = `[123, 123.5, "Hello", null, true, {"foo": 5, "bar": 6}]`

	var expected = `[123,123.5,"Hello",null,true,{}]`

	actual, err := jsonutil.Strip(jsonText, "foo", "bar", "Hello")
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON array (mixed 3) not stripped as expected")
}

func TestStrip_BadJson(test *testing.T) {
	var jsonText = `{foo: 123, bar: "abc", phoo: true, lum: {"phoox": 3, "bax": 5}}`
	actual, err := jsonutil.Strip(jsonText, "foo", "bar")
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

	actual, err := jsonutil.Strip(jsonText)
	require.NoError(test, err)
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

	actual, err := jsonutil.Strip(jsonText, "foo")
	require.NoError(test, err)
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

	actual, err := jsonutil.Strip(jsonText, "foo", "a")
	require.NoError(test, err)
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

	actual, err := jsonutil.Strip(jsonText, "foo", "a", "c")
	require.NoError(test, err)
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

	actual, err := jsonutil.Strip(jsonText, "foo", "bar", "c", "phoo")
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON object (formatted 4) not stripped as expected")
}

// ----------------------------------------------------------------------------
// Test Truncate function
// ----------------------------------------------------------------------------

func TestTruncate_AllLines(test *testing.T) {
	var jsonText = `
	{
		"foo": 123,
		"bar": true,
		"phoo": [ {"c": 4, "a": 2}, {"a": 1, "c": [9, 0, 8]}, {"a": 1, "b": 5}]
	}`

	var expected = `{"bar":true,"foo":123,"phoo":[{"a":1,"b":5},{"a":1,"c":[0,8,9]},{"a":2,"c":4}]}`

	actual := jsonutil.Truncate(jsonText, 0)
	assert.Equal(test, expected, actual)
}

func TestTruncate_3_lines(test *testing.T) {
	var jsonText = `
	{
		"foo": 123,
		"bar": true,
		"phoo": [ {"c": 4, "a": 2}, {"a": 1, "c": [9, 0, 8]}, {"a": 1, "b": 5}]
	}`

	var expected = `{"foo":123}`

	actual := jsonutil.Truncate(jsonText, 3, "bar", "phoo")
	assert.Equal(test, expected, actual)
}

func TestTruncate_6_lines(test *testing.T) {
	var jsonText = `
	{
		"foo": 123,
		"bar": true,
		"phoo": [ {"c": 4, "a": 2}, {"a": 1, "c": [9, 0, 8]}, {"a": 1, "b": 5}]
	}`

	var expected = `{"foo":123,"phoo":[{"a":1,"b":5...`

	actual := jsonutil.Truncate(jsonText, 6, "bar")
	assert.Equal(test, expected, actual)
}

func TestTruncate_bad_JSON(test *testing.T) {
	var jsonText = `
	{
		"foo": 123,
		"bar": true,
		"phoo": [ {"c": 4, "a": 2}, {"a": 1, "c": [9, 0, 8]}, {"a": 1, "b": 5}
	}`

	var expected = jsonText

	actual := jsonutil.Truncate(jsonText, 3, "bar", "phoo")
	assert.Equal(test, expected, actual)
}
