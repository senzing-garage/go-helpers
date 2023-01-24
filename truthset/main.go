package truthset

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

// A list of data sources.
var TruthsetDataSources = map[string]struct {
	Data string
}{
	"CUSTOMERS": {
		Data: `{"DSRC_CODE": "CUSTOMERS"}`,
	},
	"REFERENCE": {
		Data: `{"DSRC_CODE": "REFERENCE"}`,
	},
	"WATCHLIST": {
		Data: `{"DSRC_CODE": "WATCHLIST"}`,
	},
}

// Must match value in sys_cfg.config_data_id.
var TestConfigDataId = 3436584709

var TestRecordsWithoutRecordId = []struct {
	DataSource string
	Data       string
	LoadId     string
}{
	{
		DataSource: "CUSTOMERS",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Kellar", "PRIMARY_NAME_FIRST": "Candace", "ADDR_LINE1": "1824 AspenOak Way", "ADDR_CITY": "Elmwood Park", "ADDR_STATE": "CA", "ADDR_POSTAL_CODE": "95865", "EMAIL_ADDRESS": "info@ca-state.gov"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD_WITHOUT_ID",
	},
	{
		DataSource: "CUSTOMERS",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Sanders", "PRIMARY_NAME_FIRST": "Sandy", "ADDR_LINE1": "1376 BlueBell Rd", "ADDR_CITY": "Sacramento", "ADDR_STATE": "CA", "ADDR_POSTAL_CODE": "95823", "EMAIL_ADDRESS": "info@ca-state.gov"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD_WITHOUT_ID",
	},
}

var TestRecordsForReplacement = []struct {
	DataSource string
	Id         string
	Data       string
	LoadId     string
}{
	{
		DataSource: "CUSTOMERS",
		Id:         "1004",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1004", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Smith", "PRIMARY_NAME_FIRST": "B", "DATE_OF_BIRTH": "11/12/1980", "ADDR_TYPE": "HOME", "ADDR_LINE1": "1515 Adela Ln", "ADDR_CITY": "Las Vegas", "ADDR_STATE": "NV", "ADDR_POSTAL_CODE": "89132", "EMAIL_ADDRESS": "bsmith@work.com", "DATE": "1/5/15", "STATUS": "Inactive", "AMOUNT": "400"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD_REPLACE",
	},
}
