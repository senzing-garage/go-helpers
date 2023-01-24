package testrecords01

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

// A list of data sources.
var TestDataSources = []struct {
	Data string
}{
	{
		Data: `{"DSRC_CODE": "CUSTOMERS"}`,
	},
}

// Must match value in sys_cfg.config_data_id.
var TestConfigDataId = 2644872116

// A list of test records.
var TestRecords = []struct {
	DataSource string
	Id         string
	Data       string
	LoadId     string
}{
	{
		DataSource: "CUSTOMERS",
		Id:         "1001",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1001", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Smith", "PRIMARY_NAME_FIRST": "Robert", "DATE_OF_BIRTH": "12/11/1978", "ADDR_TYPE": "MAILING", "ADDR_LINE1": "123 Main Street, Las Vegas NV 89132", "PHONE_TYPE": "HOME", "PHONE_NUMBER": "702-919-1300", "EMAIL_ADDRESS": "bsmith@work.com", "DATE": "1/2/18", "STATUS": "Active", "AMOUNT": "100"}`,
		LoadId:     "TESTRECORDS01_LOAD",
	},
	{
		DataSource: "CUSTOMERS",
		Id:         "1002",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1002", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Smith", "PRIMARY_NAME_FIRST": "Bob", "DATE_OF_BIRTH": "11/12/1978", "ADDR_TYPE": "HOME", "ADDR_LINE1": "1515 Adela Lane", "ADDR_CITY": "Las Vegas", "ADDR_STATE": "NV", "ADDR_POSTAL_CODE": "89111", "PHONE_TYPE": "MOBILE", "PHONE_NUMBER": "702-919-1300", "DATE": "3/10/17", "STATUS": "Inactive", "AMOUNT": "200"}`,
		LoadId:     "TESTRECORDS01_LOAD",
	},
	{
		DataSource: "CUSTOMERS",
		Id:         "1003",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1003", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Smith", "PRIMARY_NAME_FIRST": "Bob", "PRIMARY_NAME_MIDDLE": "J", "DATE_OF_BIRTH": "12/11/1978", "EMAIL_ADDRESS": "bsmith@work.com", "DATE": "4/9/16", "STATUS": "Inactive", "AMOUNT": "300"}`,
		LoadId:     "TESTRECORDS01_LOAD",
	},
	{
		DataSource: "CUSTOMERS",
		Id:         "1004",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1004", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Smith", "PRIMARY_NAME_FIRST": "B", "DATE_OF_BIRTH": "11/12/1979", "ADDR_TYPE": "HOME", "ADDR_LINE1": "1515 Adela Ln", "ADDR_CITY": "Las Vegas", "ADDR_STATE": "NV", "ADDR_POSTAL_CODE": "89132", "EMAIL_ADDRESS": "bsmith@work.com", "DATE": "1/5/15", "STATUS": "Inactive", "AMOUNT": "400"}`,
		LoadId:     "TESTRECORDS01_LOAD",
	},
	{
		DataSource: "CUSTOMERS",
		Id:         "1005",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1005", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Smith", "PRIMARY_NAME_FIRST": "Robbie", "DRIVERS_LICENSE_NUMBER": "112233", "DRIVERS_LICENSE_STATE": "NV", "ADDR_TYPE": "MAILING", "ADDR_LINE1": "123 E Main St", "ADDR_CITY": "Henderson", "ADDR_STATE": "NV", "ADDR_POSTAL_CODE": "89132", "DATE": "7/16/19", "STATUS": "Active", "AMOUNT": "500"}`,
		LoadId:     "TESTRECORDS01_LOAD",
	},
	{
		DataSource: "CUSTOMERS",
		Id:         "1039",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1039", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Smith", "PRIMARY_NAME_FIRST": "John", "GENDER": "M", "DATE_OF_BIRTH": "10/10/70", "ADDR_TYPE": "HOME", "ADDR_LINE1": "3212 W. 32nd St Palm Harbor, FL 60527", "DATE": "1/28/18", "STATUS": "Active", "AMOUNT": "900"}`,
		LoadId:     "TESTRECORDS01_LOAD",
	},
	{
		DataSource: "CUSTOMERS",
		Id:         "1040",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1040", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Smith", "PRIMARY_NAME_FIRST": "John", "DATE_OF_BIRTH": "3/15/90", "ADDR_TYPE": "HOME", "ADDR_LINE1": "3212 W. 32nd St Palm Harbor, FL 60527", "DATE": "1/29/18", "STATUS": "Active", "AMOUNT": "100"}`,
		LoadId:     "TESTRECORDS01_LOAD",
	},
	{
		DataSource: "CUSTOMERS",
		Id:         "1044",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1044", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Smith", "PRIMARY_NAME_FIRST": "Patricia", "DATE_OF_BIRTH": "3/15/90", "PASSPORT_NUMBER": "10252222", "PASSPORT_COUNTRY": "US", "ADDR_TYPE": "HOME", "ADDR_LINE1": "3212 W. 32nd St Palm Harbor, FL 60527", "DATE": "1/31/18", "STATUS": "Active", "AMOUNT": "300"}`,
		LoadId:     "TESTRECORDS01_LOAD",
	},
}

var TestRecordsWithoutRecordId = []struct {
	DataSource string
	Data       string
	LoadId     string
}{
	{
		DataSource: "CUSTOMERS",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Kellar", "PRIMARY_NAME_FIRST": "Candace", "ADDR_LINE1": "1824 AspenOak Way", "ADDR_CITY": "Elmwood Park", "ADDR_STATE": "CA", "ADDR_POSTAL_CODE": "95865", "EMAIL_ADDRESS": "info@ca-state.gov"}`,
		LoadId:     "TESTRECORDS01_LOAD",
	},
	{
		DataSource: "CUSTOMERS",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Sanders", "PRIMARY_NAME_FIRST": "Sandy", "ADDR_LINE1": "1376 BlueBell Rd", "ADDR_CITY": "Sacramento", "ADDR_STATE": "CA", "ADDR_POSTAL_CODE": "95823", "EMAIL_ADDRESS": "info@ca-state.gov"}`,
		LoadId:     "TESTRECORDS01_LOAD",
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
		LoadId:     "TESTRECORDS01_LOAD",
	},
}
