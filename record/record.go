package record

import (
	"encoding/json"
)

// ----------------------------------------------------------------------------

type Record struct {
	DataSource string `json:"DATA_SOURCE"`
	Id         string `json:"RECORD_ID"`
	Json       string
}

// ----------------------------------------------------------------------------

// Returns a valid Record or an error if validation fails
func NewRecord(line string) (*Record, error) {
	var record Record
	err := json.Unmarshal([]byte(line), &record)
	if err == nil {
		record.Json = line
		_, validationErr := ValidateRecord(record)
		if validationErr == nil {
			return &record, nil
		} else {
			return &record, validationErr
		}
	}
	return &record, szerrors.Error(3000)
}

// ----------------------------------------------------------------------------

// A string is only a valid Record, if it is a well formed JSON-line
// and it has a DataSource field
// and it has an Id field
func Validate(line string) (bool, error) {
	var record Record
	valid := json.Unmarshal([]byte(line), &record) == nil
	if valid {
		return ValidateRecord(record)
	}
	//TODO: should we return the actual parse error???
	return valid, szerrors.Error(3000)
}

// ----------------------------------------------------------------------------

// A Record is only valid if it has a DataSource field
// and it has an Id field
func ValidateRecord(record Record) (bool, error) {

	if record.DataSource == "" {
		return false, szerrors.Error(3001)
	}
	if record.Id == "" {
		return false, szerrors.Error(3002)
	}
	return true, nil
}
