package jsonutil

import (
	"encoding/json"
	"fmt"
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
func IsJson(unknownText string) bool {
	var jsonString json.RawMessage
	return json.Unmarshal([]byte(unknownText), &jsonString) == nil
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
func NormalizeJson(jsonText string) (string, error) {
	var parsedJson *any = nil

	// unmarshall the text and let it allocate whatever object it wants to hold the result
	err := json.Unmarshal([]byte(jsonText), &parsedJson)

	// check for an unmarshalling error
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	// check for a null literal which is unmarshalled as a nil pointer
	if parsedJson == nil {
		return "null", nil
	}

	// marshall the parsed object back to text (bytes) and return the text and potential error
	normalizedJson, err := json.Marshal(*parsedJson)

	return string(normalizedJson), err
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
func NormalizeAndSortJson(jsonText string) (string, error) {
	var parsedJson *any = nil

	// unmarshall the text and let it allocate whatever object it wants to hold the result
	err := json.Unmarshal([]byte(jsonText), &parsedJson)

	// check for an unmarshalling error
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	// check for a null literal which is unmarshalled as a nil pointer
	if parsedJson == nil {
		return "null", nil
	}

	// sort the parsed JSON value
	sortJsonValue(*parsedJson)

	// marshall the parsed object back to text (bytes) and return the text and potential error
	normalizedJson, err := json.Marshal(*parsedJson)

	return string(normalizedJson), err
}

func sortJsonValue(jsonValue any) {
	switch typedJson := jsonValue.(type) {
	case map[string]any:
		sortJsonObject(typedJson)
	case []any:
		sortJsonArray(typedJson)
	default:
		// do nothing for JSON values that are not objects or arrays
		// these values are already "sorted"
	}

}

func sortJsonObject(jsonObject map[string]any) {
	// sort each value in the object
	for _, jsonValue := range jsonObject {
		sortJsonValue(jsonValue)
	}
}

func sortJsonArray(jsonArray []any) {
	// sort each element in the array
	for _, jsonValue := range jsonArray {
		sortJsonValue(jsonValue)
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
