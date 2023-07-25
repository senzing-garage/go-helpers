package jsonutil

import (
	"encoding/json"
	"fmt"
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
