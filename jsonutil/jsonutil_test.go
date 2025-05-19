package jsonutil_test

import (
	"errors"
	"testing"

	"github.com/senzing-garage/go-helpers/jsonutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	badJSON                = `{foo: 123, bar: "abc", phoo: true, lum: {"phoox": 3, "bax": 5}}`
	jsonTextForArray       = `[123, 123.5, "Hello", true, {"foo": 5, "bar": 6}]`
	jsonTextForBoolean     = "true"
	jsonTextForDecimal     = "123.4"
	jsonTextForFormatted   = `{"bar":true,"foo":123,"phoo":[{"a":2,"c":4},{"a":1,"c":[9,0,8]},{"a":1,"b":5}]}`
	jsonTextForInteger     = "123"
	jsonTextForMixedArray  = `[123, 123.5, "Hello", null, true, {"foo": 5, "bar": 6}]`
	jsonTextForNull        = "null"
	jsonTextForRedact1     = `{"foo": 123, "bar": "abc", "phoo": true, "lum": null}`
	jsonTextForRedact2     = `{"foo": 123, "bar": [4, 6, 2], "phoo": true, "lum": {"phoox": false, "bax": 5}}`
	jsonTextForString      = `"Hello"`
	jsonTextForStringArray = `["foo", "bar", "phoo", "lum"]`
	jsonTextPretty         = `
	{
		"foo": 123,
		"bar": true,
		"phoo": [ {"c": 4, "a": 2}, {"a": 1, "c": [9, 0, 8]}, {"a": 1, "b": 5}]
	}`
)

var errFailed = errors.New("failed")

var testCasesForIsJSON = []struct {
	name     string
	jsonText string
	expected bool
}{
	{
		name:     "Basic",
		jsonText: `{"foo": 123, "bar": "abc", "phoo": true, "lum": 20.5}`,
		expected: true,
	},
	{
		name:     "Compound",
		jsonText: `{"foo": 123, "bar": "abc", "phoo": true, "lum": {"phoox": 3, "bax": 5}}`,
		expected: true,
	},
	{
		name:     "Integer",
		jsonText: jsonTextForInteger,
		expected: true,
	},
	{
		name:     "Decimal",
		jsonText: jsonTextForDecimal,
		expected: true,
	},
	{
		name:     "String",
		jsonText: jsonTextForString,
		expected: true,
	},
	{
		name:     "Boolean",
		jsonText: jsonTextForBoolean,
		expected: true,
	},
	{
		name:     "Null",
		jsonText: jsonTextForNull,
		expected: true,
	},
	{
		name:     "Array",
		jsonText: jsonTextForArray,
		expected: true,
	},
}

var testCasesForNormalize = []struct {
	name     string
	jsonText string
	expected string
}{
	{
		name:     "Basic",
		jsonText: jsonTextForRedact1,
		expected: `{"bar":"abc","foo":123,"lum":null,"phoo":true}`,
	}, {
		name:     "Compound",
		jsonText: `{"foo": 123, "bar": "abc", "phoo": true, "lum": {"phoox": null, "bax": 5}}`,
		expected: `{"bar":"abc","foo":123,"lum":{"bax":5,"phoox":null},"phoo":true}`,
	}, {
		name:     "Integer",
		jsonText: jsonTextForInteger,
		expected: "123",
	}, {
		name:     "Decimal",
		jsonText: jsonTextForDecimal,
		expected: "123.4",
	}, {
		name:     "String",
		jsonText: jsonTextForString,
		expected: `"Hello"`,
	}, {
		name:     "Boolean",
		jsonText: jsonTextForBoolean,
		expected: "true",
	}, {
		name:     "Null",
		jsonText: jsonTextForNull,
		expected: "null",
	}, {
		name:     "Array",
		jsonText: jsonTextForArray,
		expected: `[123,123.5,"Hello",true,{"bar":6,"foo":5}]`,
	},
}

var testCasesForNormalizeAndSort = []struct {
	name     string
	jsonText string
	expected string
}{
	{
		name:     "Basic",
		jsonText: jsonTextForRedact1,
		expected: `{"bar":"abc","foo":123,"lum":null,"phoo":true}`,
	}, {
		name:     "StringArray",
		jsonText: jsonTextForStringArray,
		expected: `["bar","foo","lum","phoo"]`,
	}, {
		name:     "Compound",
		jsonText: `{"foo": 123, "bar": [4, 6, 2], "phoo": true, "lum": {"phoox": null, "bax": 5}}`,
		expected: `{"bar":[2,4,6],"foo":123,"lum":{"bax":5,"phoox":null},"phoo":true}`,
	}, {
		name:     "Integer",
		jsonText: jsonTextForInteger,
		expected: "123",
	}, {
		name:     "Decimal",
		jsonText: jsonTextForDecimal,
		expected: "123.4",
	}, {
		name:     "String",
		jsonText: jsonTextForString,
		expected: `"Hello"`,
	}, {
		name:     "Boolean",
		jsonText: jsonTextForBoolean,
		expected: "true",
	}, {
		name:     "Null",
		jsonText: jsonTextForNull,
		expected: "null",
	}, {
		name:     "MixedArray",
		jsonText: jsonTextForMixedArray,
		expected: `[null,"Hello",123,123.5,true,{"bar":6,"foo":5}]`,
	},
}

var testCasesForRedact = []struct {
	name     string
	jsonText string
	redact   []string
	expected string
}{
	{
		name:     "Basic",
		jsonText: jsonTextForRedact1,
		redact:   []string{},
		expected: `{"bar":"abc","foo":123,"lum":null,"phoo":true}`,
	}, {
		name:     "Basic with nil",
		jsonText: jsonTextForRedact1,
		expected: `{"bar":"abc","foo":123,"lum":null,"phoo":true}`,
	}, {
		name:     "Basic1",
		jsonText: jsonTextForRedact1,
		redact:   []string{"foo"},
		expected: `{"bar":"abc","foo":null,"lum":null,"phoo":true}`,
	}, {
		name:     "Basic2",
		jsonText: jsonTextForRedact1,
		redact:   []string{"foo", "bar"},
		expected: `{"bar":null,"foo":null,"lum":null,"phoo":true}`,
	}, {
		name:     "StringArray",
		jsonText: jsonTextForStringArray,
		redact:   []string{"foo", "bar"},
		expected: `["foo","bar","phoo","lum"]`,
	}, {
		name:     "Compound 0",
		jsonText: jsonTextForRedact2,
		redact:   []string{},
		expected: `{"bar":[4,6,2],"foo":123,"lum":{"bax":5,"phoox":false},"phoo":true}`,
	}, {
		name:     "Compound 1",
		jsonText: jsonTextForRedact2,
		redact:   []string{"phoox"},
		expected: `{"bar":[4,6,2],"foo":123,"lum":{"bax":5,"phoox":null},"phoo":true}`,
	}, {
		name:     "Compound 2",
		jsonText: jsonTextForRedact2,
		redact:   []string{"phoox", "bar"},
		expected: `{"bar":null,"foo":123,"lum":{"bax":5,"phoox":null},"phoo":true}`,
	}, {
		name:     "Compound 3",
		jsonText: jsonTextForRedact2,
		redact:   []string{"phoox", "bar", "bax"},
		expected: `{"bar":null,"foo":123,"lum":{"bax":null,"phoox":null},"phoo":true}`,
	}, {
		name:     "Compound 4",
		jsonText: jsonTextForRedact2,
		redact:   []string{"phoox", "bar", "bax", "foo"},
		expected: `{"bar":null,"foo":null,"lum":{"bax":null,"phoox":null},"phoo":true}`,
	}, {
		name:     "Compound 5",
		jsonText: jsonTextForRedact2,
		redact:   []string{"phoox", "bar", "foo", "phoo", "lum"},
		expected: `{"bar":null,"foo":null,"lum":null,"phoo":null}`,
	}, {
		name:     "Integer",
		jsonText: jsonTextForInteger,
		redact:   []string{"123", "foo"},
		expected: "123",
	}, {
		name:     "Decimal",
		jsonText: jsonTextForDecimal,
		redact:   []string{"123.4", "foo"},
		expected: "123.4",
	}, {
		name:     "String",
		jsonText: jsonTextForString,
		redact:   []string{"Hello", "foo"},
		expected: `"Hello"`,
	}, {
		name:     "Boolean",
		jsonText: jsonTextForBoolean,
		redact:   []string{"true", "foo"},
		expected: "true",
	}, {
		name:     "Null",
		jsonText: jsonTextForNull,
		redact:   []string{"null", "foo"},
		expected: "null",
	}, {
		name:     "MixedArray0",
		jsonText: jsonTextForMixedArray,
		redact:   []string{},
		expected: `[123,123.5,"Hello",null,true,{"bar":6,"foo":5}]`,
	}, {
		name:     "MixedArray1",
		jsonText: jsonTextForMixedArray,
		redact:   []string{"bar"},
		expected: `[123,123.5,"Hello",null,true,{"bar":null,"foo":5}]`,
	}, {
		name:     "MixedArray2",
		jsonText: jsonTextForMixedArray,
		redact:   []string{"foo", "bar"},
		expected: `[123,123.5,"Hello",null,true,{"bar":null,"foo":null}]`,
	}, {
		name:     "MixedArray3",
		jsonText: jsonTextForMixedArray,
		redact:   []string{"foo", "bar", "Hello"},
		expected: `[123,123.5,"Hello",null,true,{"bar":null,"foo":null}]`,
	},
}

var testCasesForRedactWithMap = []struct {
	name     string
	jsonText string
	redact   map[string]any
	expected string
}{
	{
		name:     "Basic0",
		jsonText: jsonTextForRedact1,
		redact:   map[string]any{},
		expected: `{"bar":"abc","foo":123,"lum":null,"phoo":true}`,
	}, {
		name:     "Basic1",
		jsonText: jsonTextForRedact1,
		redact:   map[string]any{"foo": ""},
		expected: `{"bar":"abc","foo":"","lum":null,"phoo":true}`,
	}, {
		name:     "Basic2",
		jsonText: jsonTextForRedact1,
		redact:   map[string]any{"foo": "", "bar": "-"},
		expected: `{"bar":"-","foo":"","lum":null,"phoo":true}`,
	}, {
		name:     "StringArray",
		jsonText: jsonTextForStringArray,
		redact:   map[string]any{"foo": "", "bar": "-"},
		expected: `["foo","bar","phoo","lum"]`,
	}, {
		name:     "Compound0",
		jsonText: jsonTextForRedact2,
		redact:   map[string]any{},
		expected: `{"bar":[4,6,2],"foo":123,"lum":{"bax":5,"phoox":false},"phoo":true}`,
	}, {
		name:     "Compound1",
		jsonText: jsonTextForRedact2,
		redact:   map[string]any{"phoox": ""},
		expected: `{"bar":[4,6,2],"foo":123,"lum":{"bax":5,"phoox":""},"phoo":true}`,
	}, {
		name:     "Compound2",
		jsonText: jsonTextForRedact2,
		redact:   map[string]any{"phoox": "", "bar": "-"},
		expected: `{"bar":"-","foo":123,"lum":{"bax":5,"phoox":""},"phoo":true}`,
	}, {
		name:     "Compound3",
		jsonText: jsonTextForRedact2,
		redact:   map[string]any{"phoox": "", "bar": "-", "bax": nil},
		expected: `{"bar":"-","foo":123,"lum":{"bax":null,"phoox":""},"phoo":true}`,
	}, {
		name:     "Compound4",
		jsonText: jsonTextForRedact2,
		redact:   map[string]any{"phoox": "", "bar": "-", "bax": nil, "foo": "xxx"},
		expected: `{"bar":"-","foo":"xxx","lum":{"bax":null,"phoox":""},"phoo":true}`,
	}, {
		name:     "Compound5",
		jsonText: jsonTextForRedact2,
		redact:   map[string]any{"phoox": "", "bar": "-", "foo": nil, "phoo": "xxx", "lum": false},
		expected: `{"bar":"-","foo":null,"lum":false,"phoo":"xxx"}`,
	}, {
		name:     "Integer",
		jsonText: jsonTextForInteger,
		redact:   map[string]any{"123": "", "foo": "-"},
		expected: "123",
	}, {
		name:     "Decimal",
		jsonText: jsonTextForDecimal,
		redact:   map[string]any{"123.4": "", "foo": "-"},
		expected: "123.4",
	}, {
		name:     "String",
		jsonText: jsonTextForString,
		redact:   map[string]any{"Hello": "", "foo": "-"},
		expected: `"Hello"`,
	}, {
		name:     "Boolean",
		jsonText: jsonTextForBoolean,
		redact:   map[string]any{"true": false, "foo": "-"},
		expected: "true",
	}, {
		name:     "Null",
		jsonText: jsonTextForNull,
		redact:   map[string]any{"null": nil, "foo": "-"},
		expected: "null",
	}, {
		name:     "MixedArray0",
		jsonText: jsonTextForMixedArray,
		redact:   map[string]any{},
		expected: `[123,123.5,"Hello",null,true,{"bar":6,"foo":5}]`,
	}, {
		name:     "MixedArray1",
		jsonText: jsonTextForMixedArray,
		redact:   map[string]any{"bar": ""},
		expected: `[123,123.5,"Hello",null,true,{"bar":"","foo":5}]`,
	}, {
		name:     "MixedArray2",
		jsonText: jsonTextForMixedArray,
		redact:   map[string]any{"foo": "", "bar": "-"},
		expected: `[123,123.5,"Hello",null,true,{"bar":"-","foo":""}]`,
	}, {
		name:     "MixedArray3",
		jsonText: jsonTextForMixedArray,
		redact:   map[string]any{"foo": "", "bar": "-", "Hello": false},
		expected: `[123,123.5,"Hello",null,true,{"bar":"-","foo":""}]`,
	},
}

var testCasesForStrip = []struct {
	name     string
	jsonText string
	redact   []string
	expected string
}{
	{
		name:     "Basic0 with nil",
		jsonText: jsonTextForRedact1,
		expected: `{"bar":"abc","foo":123,"lum":null,"phoo":true}`,
	}, {
		name:     "Basic0",
		jsonText: jsonTextForRedact1,
		redact:   []string{},
		expected: `{"bar":"abc","foo":123,"lum":null,"phoo":true}`,
	}, {
		name:     "Basic1",
		jsonText: jsonTextForRedact1,
		redact:   []string{"foo"},
		expected: `{"bar":"abc","lum":null,"phoo":true}`,
	}, {
		name:     "Basic2",
		jsonText: jsonTextForRedact1,
		redact:   []string{"foo", "bar"},
		expected: `{"lum":null,"phoo":true}`,
	}, {
		name:     "StringArray",
		jsonText: jsonTextForStringArray,
		redact:   []string{"foo", "bar"},
		expected: `["foo","bar","phoo","lum"]`,
	}, {
		name:     "Compound0",
		jsonText: jsonTextForRedact2,
		redact:   []string{},
		expected: `{"bar":[4,6,2],"foo":123,"lum":{"bax":5,"phoox":false},"phoo":true}`,
	}, {
		name:     "Compound1",
		jsonText: jsonTextForRedact2,
		redact:   []string{"phoox"},
		expected: `{"bar":[4,6,2],"foo":123,"lum":{"bax":5},"phoo":true}`,
	}, {
		name:     "Compound2",
		jsonText: jsonTextForRedact2,
		redact:   []string{"phoox", "bar"},
		expected: `{"foo":123,"lum":{"bax":5},"phoo":true}`,
	}, {
		name:     "Compound3",
		jsonText: jsonTextForRedact2,
		redact:   []string{"phoox", "bar", "bax"},
		expected: `{"foo":123,"lum":{},"phoo":true}`,
	}, {
		name:     "Compound4",
		jsonText: jsonTextForRedact2,
		redact:   []string{"phoox", "bar", "bax", "foo"},
		expected: `{"lum":{},"phoo":true}`,
	}, {
		name:     "Compound5",
		jsonText: jsonTextForRedact2,
		redact:   []string{"phoox", "bar", "foo", "phoo", "lum"},
		expected: `{}`,
	}, {
		name:     "Integer",
		jsonText: jsonTextForInteger,
		redact:   []string{"123", "foo"},
		expected: "123",
	}, {
		name:     "Decimal",
		jsonText: jsonTextForDecimal,
		redact:   []string{"123.4", "foo"},
		expected: "123.4",
	}, {
		name:     "String",
		jsonText: jsonTextForString,
		redact:   []string{"Hello", "foo"},
		expected: `"Hello"`,
	}, {
		name:     "Boolean",
		jsonText: jsonTextForBoolean,
		redact:   []string{"true", "foo"},
		expected: "true",
	}, {
		name:     "Null",
		jsonText: jsonTextForNull,
		redact:   []string{"null", "foo"},
		expected: "null",
	}, {
		name:     "MixedArray0",
		jsonText: jsonTextForMixedArray,
		redact:   []string{},
		expected: `[123,123.5,"Hello",null,true,{"bar":6,"foo":5}]`,
	}, {
		name:     "MixedArray1",
		jsonText: jsonTextForMixedArray,
		redact:   []string{"bar"},
		expected: `[123,123.5,"Hello",null,true,{"foo":5}]`,
	}, {
		name:     "MixedArray2",
		jsonText: jsonTextForMixedArray,
		redact:   []string{"foo", "bar"},
		expected: `[123,123.5,"Hello",null,true,{}]`,
	}, {
		name:     "MixedArray3",
		jsonText: jsonTextForMixedArray,
		redact:   []string{"foo", "bar", "Hello"},
		expected: `[123,123.5,"Hello",null,true,{}]`,
	},
}

// ----------------------------------------------------------------------------
// Test Flatten function
// ----------------------------------------------------------------------------

func TestFlatten_NoError(test *testing.T) {
	test.Parallel()

	actual := jsonutil.Flatten(`{"foo": 5, "bar": 6}`, nil)
	expected := `{"foo": 5, "bar": 6}`
	assert.Equal(test, expected, actual, "Flattening without an error did not work as expected: "+actual)
}

func TestFlatten_WithError(test *testing.T) {
	test.Parallel()

	actual := jsonutil.Flatten(`{"foo": 5, "bar": 6}`, errFailed)
	expected := `{"function": "jsonutil.Flatten", "text": "failed", "error": "failed"}`
	assert.Equal(test, expected, actual, "Flattening with an error did not work as expected: "+actual)
}

// ----------------------------------------------------------------------------
// Test IsJson function
// ----------------------------------------------------------------------------

func TestIsJson(test *testing.T) {
	test.Parallel()

	for _, testCase := range testCasesForIsJSON {
		test.Run(testCase.name, func(test *testing.T) {
			test.Parallel()

			actual := jsonutil.IsJSON(testCase.jsonText)
			assert.Equal(test, testCase.expected, actual)
		})
	}
}

func TestIsJson_BadJson(test *testing.T) {
	test.Parallel()

	expected := false
	actual := jsonutil.IsJSON(badJSON)
	assert.Equal(test, expected, actual, "Invalid JSON text incorrectly recognized as JSON")
}

func TestIsJson_Formatted(test *testing.T) {
	test.Parallel()

	var (
		jsonText = `
	{
		"foo": 123,
		"bar": true
	}`
		expected = true
	)

	actual := jsonutil.IsJSON(jsonText)
	assert.Equal(test, expected, actual, "Formatted JSON object not recognized as JSON")
}

// ----------------------------------------------------------------------------
// Test NormalizeJson function
// ----------------------------------------------------------------------------

func TestNormalize(test *testing.T) {
	test.Parallel()

	for _, testCase := range testCasesForNormalize {
		test.Run(testCase.name, func(test *testing.T) {
			test.Parallel()

			actual, err := jsonutil.Normalize(testCase.jsonText)
			require.NoError(test, err)
			assert.Equal(test, testCase.expected, actual)
		})
	}
}

func TestNormalize_BadJson(test *testing.T) {
	test.Parallel()

	actual, err := jsonutil.Normalize(badJSON)
	require.Error(test, err, "Invalid JSON text was normalized without an error: "+actual)
}

func TestNormalize_Formatted(test *testing.T) {
	test.Parallel()

	var (
		jsonText = `
	{
		"foo": 123,
		"bar": true
	}`
		expected = `{"bar":true,"foo":123}`
	)

	actual, err := jsonutil.Normalize(jsonText)
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON object (formatted) not normalized as expected")
}

// ----------------------------------------------------------------------------
// Test NormalizeAndSortJson function
// ----------------------------------------------------------------------------

func TestNormalizeAndSort(test *testing.T) {
	test.Parallel()

	for _, testCase := range testCasesForNormalizeAndSort {
		test.Run(testCase.name, func(test *testing.T) {
			test.Parallel()

			actual, err := jsonutil.NormalizeAndSort(testCase.jsonText)
			require.NoError(test, err)
			assert.Equal(test, testCase.expected, actual)
		})
	}
}

func TestNormalizeAndSort_BadJson(test *testing.T) {
	test.Parallel()

	actual, err := jsonutil.NormalizeAndSort(badJSON)
	require.Error(test, err, "Invalid JSON text was normalized without an error: "+actual)
}

func TestNormalizeAndSort_Formatted(test *testing.T) {
	test.Parallel()

	var (
		jsonText = jsonTextPretty
		expected = `{"bar":true,"foo":123,"phoo":[{"a":1,"b":5},{"a":1,"c":[0,8,9]},{"a":2,"c":4}]}`
	)

	actual, err := jsonutil.NormalizeAndSort(jsonText)
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON object (formatted) not normalized as expected")
}

// ----------------------------------------------------------------------------
// Test PrettyPrint function
// ----------------------------------------------------------------------------

func TestPrettyPrint(test *testing.T) {
	test.Parallel()

	var (
		jsonText = `{"bar":true,"foo":123,"phoo":[{"a":1,"b":5},{"a":1,"c":[0,8,9]},{"a":2,"c":4}]}`
		expected = `{
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
	)

	actual := jsonutil.PrettyPrint(jsonText, "	")
	assert.Equal(test, expected, actual, "JSON object (formatted) not pretty printed as expected")
}

func TestPrettyPrint_UsingSpaces(test *testing.T) {
	test.Parallel()

	var (
		jsonText = `{"bar":true,"foo":123,"phoo":[{"a":1,"b":5},{"a":1,"c":[0,8,9]},{"a":2,"c":4}]}`
		expected = `{
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
	)

	actual := jsonutil.PrettyPrint(jsonText, "    ")
	assert.Equal(test, expected, actual, "JSON object (formatted) not pretty printed as expected")
}

// ----------------------------------------------------------------------------
// Test RedactJson function
// ----------------------------------------------------------------------------

func TestRedact(test *testing.T) {
	test.Parallel()

	for _, testCase := range testCasesForRedact {
		test.Run(testCase.name, func(test *testing.T) {
			test.Parallel()

			actual, err := jsonutil.Redact(testCase.jsonText, testCase.redact...)
			require.NoError(test, err)
			assert.Equal(test, testCase.expected, actual)
		})
	}
}

func TestRedact_BadJson(test *testing.T) {
	test.Parallel()

	actual, err := jsonutil.Redact(badJSON, "foo", "bar")
	require.Error(test, err, "Invalid JSON text was redacted without an error: "+actual)
}

func TestRedact_Formatted0(test *testing.T) {
	test.Parallel()

	jsonText := jsonTextPretty
	expected := jsonTextForFormatted

	actual, err := jsonutil.Redact(jsonText)
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON object (formatted 0) not redacted as expected")
}

func TestRedact_Formatted1(test *testing.T) {
	test.Parallel()

	jsonText := jsonTextPretty
	expected := `{"bar":true,"foo":null,"phoo":[{"a":2,"c":4},{"a":1,"c":[9,0,8]},{"a":1,"b":5}]}`

	actual, err := jsonutil.Redact(jsonText, "foo")
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON object (formatted 1) not redacted as expected")
}

func TestRedact_Formatted2(test *testing.T) {
	test.Parallel()

	jsonText := jsonTextPretty
	expected := `{"bar":true,"foo":null,"phoo":[{"a":null,"c":4},{"a":null,"c":[9,0,8]},{"a":null,"b":5}]}`

	actual, err := jsonutil.Redact(jsonText, "foo", "a")
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON object (formatted 2) not redacted as expected")
}

func TestRedact_Formatted3(test *testing.T) {
	test.Parallel()

	jsonText := jsonTextPretty
	expected := `{"bar":true,"foo":null,"phoo":[{"a":null,"c":null},{"a":null,"c":null},{"a":null,"b":5}]}`

	actual, err := jsonutil.Redact(jsonText, "foo", "a", "c")
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON object (formatted 3) not redacted as expected")
}

func TestRedact_Formatted4(test *testing.T) {
	test.Parallel()

	jsonText := jsonTextPretty
	expected := `{"bar":null,"foo":null,"phoo":null}`

	actual, err := jsonutil.Redact(jsonText, "foo", "bar", "c", "phoo")
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON object (formatted 4) not redacted as expected")
}

// ----------------------------------------------------------------------------
// Test RedactWithMap function
// ----------------------------------------------------------------------------

func TestRedactWithMap(test *testing.T) {
	test.Parallel()

	for _, testCase := range testCasesForRedactWithMap {
		test.Run(testCase.name, func(test *testing.T) {
			test.Parallel()

			actual, err := jsonutil.RedactWithMap(testCase.jsonText, testCase.redact)
			require.NoError(test, err)
			assert.Equal(test, testCase.expected, actual)
		})
	}
}

func TestRedactWithMap_BadJson(test *testing.T) {
	test.Parallel()

	actual, err := jsonutil.RedactWithMap(badJSON, map[string]any{"foo": "", "bar": "-"})
	require.Error(test, err, "Invalid JSON text was redacted with map without an error: "+actual)
}

func TestRedactWithMap_Formatted0(test *testing.T) {
	test.Parallel()

	jsonText := jsonTextPretty
	expected := jsonTextForFormatted

	actual, err := jsonutil.RedactWithMap(jsonText, map[string]any{})
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON object (formatted 0) not redacted with map as expected")
}

func TestRedactWithMap_Formatted1(test *testing.T) {
	test.Parallel()

	jsonText := jsonTextPretty
	expected := `{"bar":true,"foo":"","phoo":[{"a":2,"c":4},{"a":1,"c":[9,0,8]},{"a":1,"b":5}]}`

	actual, err := jsonutil.RedactWithMap(jsonText, map[string]any{"foo": ""})
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON object (formatted 1) not redacted with map as expected")
}

func TestRedactWithMap_Formatted2(test *testing.T) {
	test.Parallel()

	jsonText := jsonTextPretty
	expected := `{"bar":true,"foo":"","phoo":[{"a":"-","c":4},{"a":"-","c":[9,0,8]},{"a":"-","b":5}]}`

	actual, err := jsonutil.RedactWithMap(jsonText, map[string]any{"foo": "", "a": "-"})
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON object (formatted 2) not redacted with map as expected")
}

func TestRedactWithMap_Formatted3(test *testing.T) {
	test.Parallel()

	jsonText := jsonTextPretty
	expected := `{"bar":true,"foo":"","phoo":[{"a":"-","c":"xxx"},{"a":"-","c":"xxx"},{"a":"-","b":5}]}`

	actual, err := jsonutil.RedactWithMap(jsonText, map[string]any{"foo": "", "a": "-", "c": "xxx"})
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON object (formatted 3) not redacted with map as expected")
}

func TestRedactWithMap_Formatted4(test *testing.T) {
	test.Parallel()

	jsonText := jsonTextPretty
	expected := `{"bar":"-","foo":"","phoo":false}`

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
	test.Parallel()

	for _, testCase := range testReverseStringCases {
		test.Run(testCase.name, func(test *testing.T) {
			test.Parallel()

			actual := jsonutil.ReverseString(testCase.theString)
			assert.Equal(test, testCase.reversedString, actual)
		})
	}
}

// ----------------------------------------------------------------------------
// Test Strip function
// ----------------------------------------------------------------------------

func TestStrip(test *testing.T) {
	test.Parallel()

	for _, testCase := range testCasesForStrip {
		test.Run(testCase.name, func(test *testing.T) {
			test.Parallel()

			actual, err := jsonutil.Strip(testCase.jsonText, testCase.redact...)
			require.NoError(test, err)
			assert.Equal(test, testCase.expected, actual)
		})
	}
}

func TestStrip_BadJson(test *testing.T) {
	test.Parallel()

	actual, err := jsonutil.Strip(badJSON, "foo", "bar")
	require.Error(test, err, "Invalid JSON text was stripped without an error: "+actual)
}

func TestStrip_Formatted0(test *testing.T) {
	test.Parallel()

	jsonText := jsonTextPretty
	expected := jsonTextForFormatted

	actual, err := jsonutil.Strip(jsonText)
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON object (formatted 0) not stripped as expected")
}

func TestStrip_Formatted1(test *testing.T) {
	test.Parallel()

	jsonText := jsonTextPretty
	expected := `{"bar":true,"phoo":[{"a":2,"c":4},{"a":1,"c":[9,0,8]},{"a":1,"b":5}]}`

	actual, err := jsonutil.Strip(jsonText, "foo")
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON object (formatted 1) not stripped as expected")
}

func TestStrip_Formatted2(test *testing.T) {
	test.Parallel()

	jsonText := jsonTextPretty
	expected := `{"bar":true,"phoo":[{"c":4},{"c":[9,0,8]},{"b":5}]}`

	actual, err := jsonutil.Strip(jsonText, "foo", "a")
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON object (formatted 2) not stripped as expected")
}

func TestStrip_Formatted3(test *testing.T) {
	test.Parallel()

	jsonText := jsonTextPretty
	expected := `{"bar":true,"phoo":[{},{},{"b":5}]}`

	actual, err := jsonutil.Strip(jsonText, "foo", "a", "c")
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON object (formatted 3) not stripped as expected")
}

func TestStrip_Formatted4(test *testing.T) {
	test.Parallel()

	jsonText := jsonTextPretty
	expected := `{}`

	actual, err := jsonutil.Strip(jsonText, "foo", "bar", "c", "phoo")
	require.NoError(test, err)
	assert.Equal(test, expected, actual, "JSON object (formatted 4) not stripped as expected")
}

// ----------------------------------------------------------------------------
// Test Truncate function
// ----------------------------------------------------------------------------

func TestTruncate_AllLines(test *testing.T) {
	test.Parallel()

	jsonText := jsonTextPretty
	expected := `{"bar":true,"foo":123,"phoo":[{"a":1,"b":5},{"a":1,"c":[0,8,9]},{"a":2,"c":4}]}`

	actual := jsonutil.Truncate(jsonText, 0)
	assert.Equal(test, expected, actual)
}

func TestTruncate_3_lines(test *testing.T) {
	test.Parallel()

	jsonText := jsonTextPretty
	expected := `{"foo":123}`

	actual := jsonutil.Truncate(jsonText, 3, "bar", "phoo")
	assert.Equal(test, expected, actual)
}

func TestTruncate_6_lines(test *testing.T) {
	test.Parallel()

	jsonText := jsonTextPretty
	expected := `{"foo":123,"phoo":[{"a":1,"b":5...`

	actual := jsonutil.Truncate(jsonText, 6, "bar")
	assert.Equal(test, expected, actual)
}

func TestTruncate_bad_JSON(test *testing.T) {
	test.Parallel()

	jsonText := `
	{
		"foo": 123,
		"bar": true,
		"phoo": [ {"c": 4, "a": 2}, {"a": 1, "c": [9, 0, 8]}, {"a": 1, "b": 5}
	}`

	expected := jsonText

	actual := jsonutil.Truncate(jsonText, 3, "bar", "phoo")
	assert.Equal(test, expected, actual)
}
