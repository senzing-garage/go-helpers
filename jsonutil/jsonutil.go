package jsonutil

import (
	"encoding/json"
	"sort"
	"strings"
)

/*
Checks if the specified text is in fact JSON.

Input
  - unknownString: The text to check if it is JSON

Output
  - true if the text is JSON, otherwise false
*/
func IsJSON(unknownText string) bool {
	var jsonString json.RawMessage
	return json.Unmarshal([]byte(unknownText), &jsonString) == nil
}

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
		errorMap := map[string]any{}
		errorMap["text"] = jsonText
		errorMap["error"] = err.Error()
		errorJSON, err := json.Marshal(errorMap)
		if err != nil {
			panic(err)
		}
		return string(errorJSON)
	}
	return jsonText
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

	// unmarshall the text and let it allocate whatever object it wants to hold the result
	err := json.Unmarshal([]byte(jsonText), &parsedJSON)

	// check for an unmarshalling error
	if err != nil {
		return jsonText, err
	}

	// check for a null literal which is unmarshalled as a nil pointer
	if parsedJSON == nil {
		return "null", nil
	}

	// marshall the parsed object back to text (bytes) and return the text and potential error
	normalizedJSON, err := json.Marshal(*parsedJSON)

	return string(normalizedJSON), err
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
		return jsonText, err
	}

	// check for a null literal which is unmarshalled as a nil pointer
	if parsedJSON == nil {
		return "null", nil
	}

	// sort the parsed JSON value
	sortValue(*parsedJSON)

	// marshall the parsed object back to text (bytes) and return the text and potential error
	normalizedJSON, err := json.Marshal(*parsedJSON)

	return string(normalizedJSON), err
}

func PrettyPrint(jsonText string) string {
	return jsonText

}

func sortValue(jsonValue any) {
	switch typedJSON := jsonValue.(type) {
	case map[string]any:
		sortObject(typedJSON)
	case []any:
		sortArray(typedJSON)
	default:
		// do nothing for JSON values that are not objects or arrays
		// these values are already "sorted"
	}

}

func sortObject(jsonObject map[string]any) {
	// sort each value in the object
	for _, jsonValue := range jsonObject {
		sortValue(jsonValue)
	}
}

func sortArray(jsonArray []any) {
	// sort each element in the array
	for _, jsonValue := range jsonArray {
		sortValue(jsonValue)
	}

	// now sort the array itself
	sort.Slice(jsonArray, func(i, j int) bool {
		// special case JSON null values
		if (jsonArray[i] == nil) && (jsonArray[j] == nil) {
			return false
		}
		if (jsonArray[i] == nil) && (jsonArray[j] != nil) {
			return true
		}
		if (jsonArray[i] != nil) && (jsonArray[j] == nil) {
			return false
		}

		// otherwise marshal the value and compare the text
		json1, _ := json.Marshal(jsonArray[i])
		json2, _ := json.Marshal(jsonArray[j])
		return (strings.Compare(string(json1), string(json2)) < 0)
	})
}

/*
Given JSON text, this will unmarshal it and recursively descend JSON arrays and JSON objects
to remove the JSON object properties whose names match those specified for removal.  The JSON
will then be marshalled back into text and returned.  This should work with any JSON literal:
objects, arrays, null, integers, booleans, decimal numbers, etc.... However, this method has
no effect on simple numeric, boolean, or null values.

Input
  - jsonText: The JSON text to be redacted.
  - removeProps: The JSON properties to be removed.

Output
  - The JSON text representing the modified JSON.
  - An error if a failure occurred in unmarshalling the specified text.
*/
func Strip(jsonText string, removeProps ...string) (string, error) {
	stripMap := map[string]any{}
	for _, jsonProp := range removeProps {
		stripMap[jsonProp] = nil
	}

	var parsedJSON *any

	// unmarshall the text and let it allocate whatever object it wants to hold the result
	err := json.Unmarshal([]byte(jsonText), &parsedJSON)

	// check for an unmarshalling error
	if err != nil {
		return jsonText, err
	}

	// check for a null literal which is unmarshalled as a nil pointer
	if parsedJSON == nil {
		return "null", nil
	}

	// sort the parsed JSON value
	stripFieldsFromValue(*parsedJSON, stripMap)

	// marshall the parsed object back to text (bytes) and return the text and potential error
	modifiedJSON, err := json.Marshal(*parsedJSON)

	return string(modifiedJSON), err
}

func stripFieldsFromValue(jsonValue any, stripMap map[string]any) {
	switch typedJSON := jsonValue.(type) {
	case map[string]any:
		stripFieldsFromObject(typedJSON, stripMap)
	case []any:
		stripFieldsFromArray(typedJSON, stripMap)
	default:
		// do nothing for JSON values that are not objects or arrays
		// these values are already "sorted"
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

func stripFieldsFromArray(jsonArray []any, stripMap map[string]any) {
	// sort each element in the array
	for _, jsonValue := range jsonArray {
		stripFieldsFromValue(jsonValue, stripMap)
	}
}

/*
Given JSON text, this will unmarshal it and recursively descend JSON arrays and JSON objects
to redact the values for any JSON object properties whose property names match those specified
for redaction.  The values will be replaced with JSON null values and the JSON will be marshalled
back into text and returned.  This should work with any JSON literal: objects, arrays, null,
integers, booleans, decimal numbers, etc.... However, this method has no effect on simple numeric,
boolean, or null values.

Input
  - jsonText: The JSON text to be redacted.
  - redactProps: The JSON properties to be redacted.

Output
  - The JSON text representing the redacted JSON.
  - An error if a failure occurred in unmarshalling the specified text.
*/
func Redact(jsonText string, redactProps ...string) (string, error) {
	redactMap := map[string]any{}
	for _, jsonProp := range redactProps {
		redactMap[jsonProp] = nil
	}

	return RedactWithMap(jsonText, redactMap)
}

/*
Given JSON text, this will unmarshal it and recursively descend JSON arrays and JSON objects
to redact the values for any JSON object properties whose property names match those specified
for redaction.  The values will be replaced with the corresponding values from the redaction
map and the JSON will be marshalled back into text and returned.  This should work with any
JSON literal: objects, arrays, null, integers, booleans, decimal numbers, etc.... However,
this method has no effect on simple numeric, boolean, or null values.  NOTE: the redacted
values should be values that can be marshalled back into JSON, if nil, then a JSON null will
be used.

Input
  - jsonText: The JSON text to be redacted.
  - redactMap: The map of JSON properties to be redacted to values to be used for redaction.

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
		return jsonText, err
	}

	// check for a null literal which is unmarshalled as a nil pointer
	if parsedJSON == nil {
		return "null", nil
	}

	// sort the parsed JSON value
	redactValue(*parsedJSON, redactMap)

	// marshall the parsed object back to text (bytes) and return the text and potential error
	redactedJSON, err := json.Marshal(*parsedJSON)

	return string(redactedJSON), err

}

func redactValue(jsonValue any, redactMap map[string]any) {
	switch typedJSON := jsonValue.(type) {
	case map[string]any:
		redactObject(typedJSON, redactMap)
	case []any:
		redactArray(typedJSON, redactMap)
	default:
		// do nothing for JSON values that are not objects or arrays
		// these values are already "sorted"
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

func redactArray(jsonArray []any, redactMap map[string]any) {
	// sort each element in the array
	for _, jsonValue := range jsonArray {
		redactValue(jsonValue, redactMap)
	}
}
