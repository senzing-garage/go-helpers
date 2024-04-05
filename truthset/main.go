package truthset

import "github.com/senzing-garage/go-helpers/record"

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

// Identfier of the  package found messages having the format "senzing-6404xxxx".
const ProductId = 6404

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

// A list of data sources.
var TruthsetDataSources = map[string]struct {
	Json string
}{
	"CUSTOMERS": {
		Json: `{"DSRC_CODE": "CUSTOMERS"}`,
	},
	"REFERENCE": {
		Json: `{"DSRC_CODE": "REFERENCE"}`,
	},
	"WATCHLIST": {
		Json: `{"DSRC_CODE": "WATCHLIST"}`,
	},
}

var TestRecordsWithoutRecordId = []record.Record{
	{
		DataSource: "CUSTOMERS",
		Json:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Kellar", "PRIMARY_NAME_FIRST": "Candace", "ADDR_LINE1": "1824 AspenOak Way", "ADDR_CITY": "Elmwood Park", "ADDR_STATE": "CA", "ADDR_POSTAL_CODE": "95865", "EMAIL_ADDRESS": "info@ca-state.gov"}`,
	},
	{
		DataSource: "CUSTOMERS",
		Json:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Sanders", "PRIMARY_NAME_FIRST": "Sandy", "ADDR_LINE1": "1376 BlueBell Rd", "ADDR_CITY": "Sacramento", "ADDR_STATE": "CA", "ADDR_POSTAL_CODE": "95823", "EMAIL_ADDRESS": "info@ca-state.gov"}`,
	},
}

var TestRecordsForReplacement = map[string]record.Record{
	"1004": {
		DataSource: "CUSTOMERS",
		Id:         "1004",
		Json:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1004", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Smith", "PRIMARY_NAME_FIRST": "B", "DATE_OF_BIRTH": "11/12/1980", "ADDR_TYPE": "HOME", "ADDR_LINE1": "1515 Adela Ln", "ADDR_CITY": "Las Vegas", "ADDR_STATE": "NV", "ADDR_POSTAL_CODE": "89132", "EMAIL_ADDRESS": "bsmith@work.com", "DATE": "1/5/15", "STATUS": "Inactive", "AMOUNT": "400"}`,
	},
}
