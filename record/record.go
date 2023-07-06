package record

import (
	"encoding/json"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// The Record data structure conforms to the [Generic Entity Specification].
//
// [Generic Entity Specification]: https://senzing.zendesk.com/hc/en-us/articles/231925448-Generic-Entity-Specification-JSON-CSV-Mapping
type Record struct {
	DataSource string `json:"DATA_SOURCE"`
	Id         string `json:"RECORD_ID"`
	Json       string
}

// ----------------------------------------------------------------------------
// Functions
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
	return &record, szerrors.NewError(3000)
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
	return valid, szerrors.NewError(3000)
}

// ----------------------------------------------------------------------------

// A Record is only valid if it has a DataSource field
// and it has an Id field
func ValidateRecord(record Record) (bool, error) {

	if record.DataSource == "" {
		return false, szerrors.NewError(3001)
	}
	if record.Id == "" {
		return false, szerrors.NewError(3002)
	}
	return true, nil
}
