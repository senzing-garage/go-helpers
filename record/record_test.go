package record_test

import (
	"testing"

	"github.com/senzing-garage/go-helpers/record"
)

// ----------------------------------------------------------------------------
// Test the NewRecord function
func TestNewRecord_good(test *testing.T) {
	jsonLine := `{"DATA_SOURCE": "ICIJ", "RECORD_ID": "24000001", "ENTITY_TYPE": "ADDRESS", "RECORD_TYPE": "ADDRESS", "icij_source": "BAHAMAS", "icij_type": "ADDRESS", "COUNTRIES": [{"COUNTRY_OF_ASSOCIATION": "BHS"}], "ADDR_FULL": "ANNEX FREDERICK & SHIRLEY STS, P.O. BOX N-4805, NASSAU, BAHAMAS", "REL_ANCHOR_DOMAIN": "ICIJ_ID", "REL_ANCHOR_KEY": "24000001"}`

	record, err := record.NewRecord(jsonLine)

	switch {
	case err != nil:
		test.Errorf("FAILED, received err: %s", err.Error())
	case record.ID != "24000001":
		test.Errorf("FAILED, Id incorrect")
	case record.DataSource != "ICIJ":
		test.Errorf("FAILED, DataSource incorrect")
	case record.JSON != jsonLine:
		test.Errorf("FAILED, jsonLine incorrect")
	default:
		test.Log("SUCCEEDED. record created.")
	}
}

func TestNewRecord_invalidJson(test *testing.T) {
	jsonLine := `{"DATA_SOURCE": "ICIJ", "RECORD_ID": "24000005B" "ENTITY_TYPE": "ADDRESS", "RECORD_TYPE": "ADDRESS", "icij_source": "BAHAMAS", "icij_type": "ADDRESS", "COUNTRIES": [{"COUNTRY_OF_ASSOCIATION": "BHS"}], "ADDR_FULL": "LYFORD CAY HOUSE, 3RD FLOOR, LYFORD CAY, P.O. BOX N-3024, NASSAU, BAHAMAS", "REL_ANCHOR_DOMAIN": "ICIJ_ID", "REL_ANCHOR_KEY": "24000005"}`
	_, err := record.NewRecord(jsonLine)

	if err != nil {
		test.Logf("SUCCEEDED, received err: %s", err.Error())
	} else {
		test.Error("FAILED, expected err.")
	}
}

func TestNewRecord_noRecordId(test *testing.T) {
	jsonLine := `{"DATA_SOURCE": "ICIJ", "ENTITY_TYPE": "ADDRESS", "RECORD_TYPE": "ADDRESS", "icij_source": "BAHAMAS", "icij_type": "ADDRESS", "COUNTRIES": [{"COUNTRY_OF_ASSOCIATION": "BHS"}], "ADDR_FULL": "ANNEX FREDERICK & SHIRLEY STS, P.O. BOX N-4805, NASSAU, BAHAMAS", "REL_ANCHOR_DOMAIN": "ICIJ_ID", "REL_ANCHOR_KEY": "24000001"}`
	_, err := record.NewRecord(jsonLine)

	if err != nil {
		test.Logf("SUCCEEDED, received err: %s", err.Error())
	} else {
		test.Error("FAILED, expected err.")
	}
}

func TestNewRecord_noDataSource(test *testing.T) {
	jsonLine := `{"RECORD_ID": "24000001", "ENTITY_TYPE": "ADDRESS", "RECORD_TYPE": "ADDRESS", "icij_source": "BAHAMAS", "icij_type": "ADDRESS", "COUNTRIES": [{"COUNTRY_OF_ASSOCIATION": "BHS"}], "ADDR_FULL": "ANNEX FREDERICK & SHIRLEY STS, P.O. BOX N-4805, NASSAU, BAHAMAS", "REL_ANCHOR_DOMAIN": "ICIJ_ID", "REL_ANCHOR_KEY": "24000001"}`
	_, err := record.NewRecord(jsonLine)

	if err != nil {
		test.Logf("SUCCEEDED, received err: %s", err.Error())
	} else {
		test.Error("FAILED, expected err.")
	}
}

// ----------------------------------------------------------------------------
// Test the Validate function
func TestValidate_good(test *testing.T) {
	jsonLine := `{"DATA_SOURCE": "ICIJ", "RECORD_ID": "24000001", "ENTITY_TYPE": "ADDRESS", "RECORD_TYPE": "ADDRESS", "icij_source": "BAHAMAS", "icij_type": "ADDRESS", "COUNTRIES": [{"COUNTRY_OF_ASSOCIATION": "BHS"}], "ADDR_FULL": "ANNEX FREDERICK & SHIRLEY STS, P.O. BOX N-4805, NASSAU, BAHAMAS", "REL_ANCHOR_DOMAIN": "ICIJ_ID", "REL_ANCHOR_KEY": "24000001"}`

	val, err := record.Validate(jsonLine)
	if err != nil {
		test.Errorf("FAILED, received err: %s", err.Error())
	}

	if !val {
		test.Error("FAILED, expected JSON to validate.")
	} else {
		test.Log("SUCCEEDED. valid record.")
	}
}

func TestValidate_noRecordId(test *testing.T) {
	jsonLine := `{"DATA_SOURCE": "ICIJ", "ENTITY_TYPE": "ADDRESS", "RECORD_TYPE": "ADDRESS", "icij_source": "BAHAMAS", "icij_type": "ADDRESS", "COUNTRIES": [{"COUNTRY_OF_ASSOCIATION": "BHS"}], "ADDR_FULL": "ANNEX FREDERICK & SHIRLEY STS, P.O. BOX N-4805, NASSAU, BAHAMAS", "REL_ANCHOR_DOMAIN": "ICIJ_ID", "REL_ANCHOR_KEY": "24000001"}`
	val, err := record.Validate(jsonLine)

	if err != nil {
		test.Logf("SUCCEEDED, received err: %s", err.Error())
	} else {
		test.Error("FAILED, expected err.")
	}

	if val {
		test.Error("FAILED, did NOT expected JSON to validate.")
	} else {
		test.Log("SUCCEEDED, invalid record")
	}
}

func TestValidate_noDatasource(test *testing.T) {
	jsonLine := `{"RECORD_ID": "24000001", "ENTITY_TYPE": "ADDRESS", "RECORD_TYPE": "ADDRESS", "icij_source": "BAHAMAS", "icij_type": "ADDRESS", "COUNTRIES": [{"COUNTRY_OF_ASSOCIATION": "BHS"}], "ADDR_FULL": "ANNEX FREDERICK & SHIRLEY STS, P.O. BOX N-4805, NASSAU, BAHAMAS", "REL_ANCHOR_DOMAIN": "ICIJ_ID", "REL_ANCHOR_KEY": "24000001"}`
	val, err := record.Validate(jsonLine)

	if err != nil {
		test.Logf("SUCCEEDED, received err: %s", err.Error())
	} else {
		test.Error("FAILED, expected err.")
	}

	if val {
		test.Error("FAILED, did NOT expected JSON to validate.")
	} else {
		test.Log("SUCCEEDED, invalid record")
	}
}

func TestValidate_invalidJson(test *testing.T) {
	jsonLine := `{"DATA_SOURCE": "ICIJ", "RECORD_ID": "24000005B" "ENTITY_TYPE": "ADDRESS", "RECORD_TYPE": "ADDRESS", "icij_source": "BAHAMAS", "icij_type": "ADDRESS", "COUNTRIES": [{"COUNTRY_OF_ASSOCIATION": "BHS"}], "ADDR_FULL": "LYFORD CAY HOUSE, 3RD FLOOR, LYFORD CAY, P.O. BOX N-3024, NASSAU, BAHAMAS", "REL_ANCHOR_DOMAIN": "ICIJ_ID", "REL_ANCHOR_KEY": "24000005"}`
	val, err := record.Validate(jsonLine)

	if err != nil {
		test.Logf("SUCCEEDED, received err: %s", err.Error())
	} else {
		test.Error("FAILED, expected err.")
	}

	if val {
		test.Error("FAILED, did NOT expected JSON to validate.")
	} else {
		test.Log("SUCCEEDED, invalid record")
	}
}
