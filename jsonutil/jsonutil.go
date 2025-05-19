package jsonutil

import (
	"bytes"
	"encoding/json"
	"sort"
	"strings"

	"github.com/senzing-garage/go-helpers/wraperror"
)

const (
	Null = "null"
)

// ----------------------------------------------------------------------------
// Public functions
// ----------------------------------------------------------------------------

/*
Flattens the tuple result of a string representing JSON text and an error that may have occurred into a single
string result.  If the error is not nil, then this returns a string representation of the error, otherwise it
returns the specified JSON text.

Input
  - jsonText: The JSON text to be flattened.
  - error: The error that may have occurred or nil.

Output
  - The flattened text
*/
func Flatten(jsonText string, err error) string {
	if err != nil {
		newErr := wraperror.Errorf(err, "%s", err.Error())

		return newErr.Error()
	}

	return jsonText
}

/*
Checks if the specified text is in fact JSON.

Input
  - unknownString: The text to check if it is JSON

Output
  - true if the text is JSON, otherwise false
*/
func IsJSON(unknownText string) bool {
	return json.Valid([]byte(unknownText))
}

/*
Normalizes the specified JSON text using the Go encoding/json marshaller to ensure it is formatted consistently.
This should work with any JSON literal: objects, arrays, null, integers, booleans, decimal numbers, etc....

Input
  - jsonText: The JSON text to be nornalized.

Output
  - The JSON text that is the normalized representation of the specified text.
  - An error if a failure occurred in interpretting/normalizing the specified text.
*/
func Normalize(jsonText string) (string, error) {
	var parsedJSON *any

	// Unmarshall the text and let it allocate whatever object it wants to hold the result.

	err := json.Unmarshal([]byte(jsonText), &parsedJSON)
	// Check for an unmarshalling error.
	if err != nil {
		return jsonText, wraperror.Errorf(err, "Unmarshal")
	}

	// Check for a null literal which is unmarshalled as a nil pointer.

	if parsedJSON == nil {
		return Null, nil
	}

	// Marshall the parsed object back to text (bytes) and return the text and potential error.

	normalizedJSON, err := json.Marshal(*parsedJSON)

	return string(normalizedJSON), wraperror.Error(err)
}

/*
Normalizes the specified JSON text using the Go encoding/json marshaller to ensure it is formatted consistently,
but also sorts any JSON arrays in a consistent manner.  This should work with any JSON literal: objects, arrays,
null, integers, booleans, decimal numbers, etc....

Input
  - jsonText: The JSON text to be normalized and sorted.

Output
  - The JSON text that is the normalized representation of the specified text.
  - An error if a failure occurred in interpretting/normalizing the specified text.
*/
func NormalizeAndSort(jsonText string) (string, error) {
	var parsedJSON *any

	// unmarshall the text and let it allocate whatever object it wants to hold the result
	err := json.Unmarshal([]byte(jsonText), &parsedJSON)
	// check for an unmarshalling error
	if err != nil {
		return jsonText, wraperror.Errorf(err, "Unmarshal")
	}

	// check for a null literal which is unmarshalled as a nil pointer
	if parsedJSON == nil {
		return Null, nil
	}

	// sort the parsed JSON value
	sortValue(*parsedJSON)

	// marshall the parsed object back to text (bytes) and return the text and potential error
	normalizedJSON, err := json.Marshal(*parsedJSON)

	return string(normalizedJSON), wraperror.Error(err)
}

/*
PrettyPrint creates a multi-line, indented string representation of the submitted JSON.

Input
  - jsonText: The JSON text to be "prettied".
  - padding: Indentation padding.

Output
  - PrettyPrinted JSON.
*/
func PrettyPrint(jsonText string, padding string) string {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(jsonText), "", padding); err != nil {
		panic(err)
	}

	return prettyJSON.String()
}

/*
Given JSON text, this will unmarshal it and recursively descend JSON arrays and JSON objects
to redact the values for any JSON key/value pairs whose keys match those specified
for redaction.  The values will be replaced with JSON null values and the JSON will be marshalled
back into text and returned.  This should work with any JSON literal: objects, arrays, null,
integers, booleans, decimal numbers, etc.... However, this method has no effect on simple numeric,
boolean, or null values.

Input
  - jsonText: The JSON text to be redacted.
  - redactKeys: The JSON key/vaue pair to be redacted.

Output
  - The JSON text representing the redacted JSON.
  - An error if a failure occurred in unmarshalling the specified text.
*/
func Redact(jsonText string, redactKeys ...string) (string, error) {
	redactMap := map[string]any{}
	for _, redactKey := range redactKeys {
		redactMap[redactKey] = nil
	}

	return RedactWithMap(jsonText, redactMap)
}

/*
Given JSON text, this will unmarshal it and recursively descend JSON arrays and JSON objects
to redact the values for any JSON key/value pairs whose keys match those specified
for redaction.  The values will be replaced with the corresponding values from the redaction
map and the JSON will be marshalled back into text and returned.  This should work with any
JSON literal: objects, arrays, null, integers, booleans, decimal numbers, etc.... However,
this method has no effect on simple numeric, boolean, or null values.  NOTE: the redacted
values should be values that can be marshalled back into JSON, if nil, then a JSON null will
be used.

Input
  - jsonText: The JSON text to be redacted.
  - redactMap: The map of JSON key/value pairs to be redacted to values to be used for redaction.

Output
  - The JSON text representing the redacted JSON.
  - An error if a failure occurred in unmarshalling the specified text or marshalling the
    redacted JSON.
*/
func RedactWithMap(jsonText string, redactMap map[string]any) (string, error) {
	var parsedJSON *any

	// unmarshall the text and let it allocate whatever object it wants to hold the result
	err := json.Unmarshal([]byte(jsonText), &parsedJSON)
	// check for an unmarshalling error
	if err != nil {
		return jsonText, wraperror.Errorf(err, "Unmarshal")
	}

	// check for a null literal which is unmarshalled as a nil pointer
	if parsedJSON == nil {
		return Null, nil
	}

	// sort the parsed JSON value
	redactValue(*parsedJSON, redactMap)

	// marshall the parsed object back to text (bytes) and return the text and potential error
	redactedJSON, err := json.Marshal(*parsedJSON)

	return string(redactedJSON), wraperror.Error(err)
}

/*
Given a string, return the reversed version of the string.

Needed because of GitHub issues like:
- https://github.com/golang/go/issues/14777
- https://github.com/golang/go/issues/63683

Input
  - aString: The string to be reversed

Output
  - The input string in reversed order.
*/
func ReverseString(aString string) string {
	runes := []rune(aString)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

/*
Given JSON text, this will unmarshal it and recursively descend JSON arrays and JSON objects
to remove the JSON key/value pairs whose keys match those specified for removal.  The JSON
will then be marshalled back into text and returned.  This should work with any JSON literal:
objects, arrays, null, integers, booleans, decimal numbers, etc.... However, this method has
no effect on simple numeric, boolean, or null values.

Input
  - jsonText: The JSON text to be redacted.
  - removeKeys: The JSON keys to be removed.

Output
  - The JSON text representing the modified JSON.
  - An error if a failure occurred in unmarshalling the specified text.
*/
func Strip(jsonText string, removeKeys ...string) (string, error) {
	stripMap := map[string]any{}
	for _, removeKey := range removeKeys {
		stripMap[removeKey] = nil
	}

	var parsedJSON *any

	// unmarshall the text and let it allocate whatever object it wants to hold the result
	err := json.Unmarshal([]byte(jsonText), &parsedJSON)
	// check for an unmarshalling error
	if err != nil {
		return jsonText, wraperror.Errorf(err, "Unmarshal")
	}

	// check for a null literal which is unmarshalled as a nil pointer
	if parsedJSON == nil {
		return Null, nil
	}

	// sort the parsed JSON value
	stripFieldsFromValue(*parsedJSON, stripMap)

	// marshall the parsed object back to text (bytes) and return the text and potential error
	modifiedJSON, err := json.Marshal(*parsedJSON)

	return string(modifiedJSON), wraperror.Error(err)
}

/*
Given JSON text, [Truncate] will run:

  - [jsonutil.Strip] to remove JSON key/value pairs specified by "removeKeys"
  - [jsonutil.NormalizeAndSort] to sort the JSON by its keys
  - [encoding/json.Compact] to trim whitespace
  - [encoding/json.Indent] to create a multi-line JSON representation

Then [Truncate] will concatenate "lines" number of lines from the
multi-line JSON representation into a single line of truncated JSON.
If the JSON has been truncated,
the string will be suffixed with an ellipsis ("...").
Thus, the returned string may not be syntactically correct JSON.

Input
  - jsonText: The JSON text to be truncated.
  - lines: The number of pretty-printed JSON lines to concatenate.
  - removeKeys: The JSON keys to be removed.

Output
  - The text representing the modified JSON.
    If there is any error, the original jsonText will be returned.
*/
func Truncate(jsonText string, lines int, removeKeys ...string) string {
	var result string

	removedJSON, err := Strip(jsonText, removeKeys...)
	if err != nil {
		return jsonText
	}

	normalizedAndSortedJSON, err := NormalizeAndSort(removedJSON)
	if err != nil {
		return jsonText
	}

	var compactJSON bytes.Buffer

	err = json.Compact(&compactJSON, []byte(normalizedAndSortedJSON))
	if err != nil {
		return jsonText
	}

	var indentedJSON bytes.Buffer

	err = json.Indent(&indentedJSON, compactJSON.Bytes(), "", "")
	if err != nil {
		return jsonText
	}

	indentedSlices := strings.Split(indentedJSON.String(), "\n")

	resultSlices := make([]string, 0, len(indentedSlices))
	for _, resultSlice := range indentedSlices {
		resultSlices = append(resultSlices, strings.Replace(resultSlice, ": ", ":", 1))
	}

	switch {
	case lines <= 0:
		result = strings.Join(resultSlices, "")
	case lines >= len(resultSlices):
		result = strings.Join(resultSlices, "")
	default:
		result = strings.Join(resultSlices[0:lines], "") + "..."
	}

	return result
}

// ----------------------------------------------------------------------------
// Internal functions
// ----------------------------------------------------------------------------

func redactArray(jsonArray []any, redactMap map[string]any) {
	// sort each element in the array
	for _, jsonValue := range jsonArray {
		redactValue(jsonValue, redactMap)
	}
}

func redactObject(jsonObject map[string]any, redactMap map[string]any) {
	// sort each value in the object
	for key, jsonValue := range jsonObject {
		redactedValue, redacted := redactMap[key]
		if redacted {
			jsonObject[key] = redactedValue
		} else {
			redactValue(jsonValue, redactMap)
		}
	}
}

func redactValue(jsonValue any, redactMap map[string]any) {
	switch typedJSON := jsonValue.(type) {
	case map[string]any:
		redactObject(typedJSON, redactMap)
	case []any:
		redactArray(typedJSON, redactMap)
	}
}

func sortArray(jsonArray []any) {
	// sort each element in the array
	for _, jsonValue := range jsonArray {
		sortValue(jsonValue)
	}

	// now sort the array itself
	sort.Slice(jsonArray, func(iIndex, jIndex int) bool {
		// special case JSON null values
		if (jsonArray[iIndex] == nil) && (jsonArray[jIndex] == nil) {
			return false
		}

		if (jsonArray[iIndex] == nil) && (jsonArray[jIndex] != nil) {
			return true
		}

		if (jsonArray[iIndex] != nil) && (jsonArray[jIndex] == nil) {
			return false
		}

		// otherwise marshal the value and compare the text
		json1, _ := json.Marshal(jsonArray[iIndex])
		json2, _ := json.Marshal(jsonArray[jIndex])

		return (bytes.Compare(json1, json2) < 0)
	})
}

func sortObject(jsonObject map[string]any) {
	// sort each value in the object
	for _, jsonValue := range jsonObject {
		sortValue(jsonValue)
	}
}

func sortValue(jsonValue any) {
	switch typedJSON := jsonValue.(type) {
	case map[string]any:
		sortObject(typedJSON)
	case []any:
		sortArray(typedJSON)
	}
}

func stripFieldsFromArray(jsonArray []any, stripMap map[string]any) {
	// sort each element in the array
	for _, jsonValue := range jsonArray {
		stripFieldsFromValue(jsonValue, stripMap)
	}
}

func stripFieldsFromObject(jsonObject map[string]any, stripMap map[string]any) {
	// sort each value in the object
	for key, jsonValue := range jsonObject {
		_, remove := stripMap[key]
		if remove {
			delete(jsonObject, key)
		} else {
			stripFieldsFromValue(jsonValue, stripMap)
		}
	}
}

func stripFieldsFromValue(jsonValue any, stripMap map[string]any) {
	switch typedJSON := jsonValue.(type) {
	case map[string]any:
		stripFieldsFromObject(typedJSON, stripMap)
	case []any:
		stripFieldsFromArray(typedJSON, stripMap)
	}
}
