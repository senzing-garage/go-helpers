package jsonutil

import (
	"encoding/json"
	"fmt"
)

func NormalizeJson(jsonText string) (string, error) {
	var parsedJson map[string]any

	err := json.Unmarshal([]byte(jsonText), &parsedJson)

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	normalizedJson, err := json.Marshal(parsedJson)

	return string(normalizedJson), err
}
