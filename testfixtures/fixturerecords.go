package testfixtures

import (
	"strings"

	"github.com/senzing-garage/go-helpers/record"
)

const Max16Bit = 65536

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

/*
A very long record.
*/
var FixtureRecords = map[string]record.Record{
	"65536-periods": {
		DataSource: "TEST",
		ID:         "65536-periods",
		JSON: `{"DATA_SOURCE":"TEST","RECORD_ID":"65536-periods","NAME_FULL":"Nobody Really","PERIODS":"` + strings.Repeat(
			".",
			Max16Bit,
		) + `"}`,
	},
}
