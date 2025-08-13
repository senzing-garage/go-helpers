package record

import (
	"encoding/json"

	"github.com/senzing-garage/go-helpers/wraperror"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// The Record data structure conforms to the [Generic Entity Specification].
//
// [Generic Entity Specification]: https://senzing.zendesk.com/hc/en-us/articles/231925448-Generic-Entity-Specification-JSON-CSV-Mapping
type Record struct {
	DataSource string `json:"DATA_SOURCE"`
	ID         string `json:"RECORD_ID"`
	JSON       string `json:"JSON"`
}

// ----------------------------------------------------------------------------
// Functions
// ----------------------------------------------------------------------------

// Returns a valid Record or an error if validation fails.
func NewRecord(line string) (*Record, error) {
	var (
		err    error
		record Record
	)

	err = json.Unmarshal([]byte(line), &record)
	if err == nil {
		record.JSON = line

		_, validationErr := ValidateRecord(record)
		if validationErr == nil {
			return &record, nil
		}

		return &record, validationErr
	}

	err = szerrors.NewError(3000)

	return &record, wraperror.Errorf(err, wraperror.NoMessage)
}

// ----------------------------------------------------------------------------

// A string is only a valid Record, if it is a well formed JSON-line
// and it has a DataSource field
// and it has an Id field.
func Validate(line string) (bool, error) {
	var record Record

	valid := json.Unmarshal([]byte(line), &record) == nil
	if valid {
		return ValidateRecord(record)
	}
	// IMPROVE: should we return the actual parse error???

	err := szerrors.NewError(3000)

	return valid, wraperror.Errorf(err, wraperror.NoMessage)
}

// ----------------------------------------------------------------------------

// A Record is only valid if it has a DataSource field
// and it has an Id field.
func ValidateRecord(record Record) (bool, error) {
	var err error
	if record.DataSource == "" {
		err = szerrors.NewError(3001)

		return false, wraperror.Errorf(err, "DataSource")
	}

	if record.ID == "" {
		err = szerrors.NewError(3002)

		return false, wraperror.Errorf(err, "Record.ID")
	}

	return true, wraperror.Errorf(err, wraperror.NoMessage)
}
