package testrecords01

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

// A list of data sources.
var TestDataSources = []struct {
	Data string
}{
	{
		Data: `{"DSRC_CODE": "GO_TEST"}`,
	},
}

// Must match value in sys_cfg.config_data_id.
var TestConfigDataId = 3239924254

// A list of test records.
var TestRecords = []struct {
	DataSource string
	Id         string
	Data       string
	LoadId     string
}{
	{
		DataSource: "GO_TEST",
		Id:         "9001",
		Data:       `{"SOCIAL_HANDLE": "flavorh", "DATE_OF_BIRTH": "4/8/1983", "ADDR_STATE": "LA", "ADDR_POSTAL_CODE": "71232", "SSN_NUMBER": "053-39-3251", "ENTITY_TYPE": "GO_TEST", "GENDER": "F", "srccode": "MDMPER", "CC_ACCOUNT_NUMBER": "5534202208773608", "RECORD_ID": "9001", "DSRC_ACTION": "A", "ADDR_CITY": "Delhi", "DRIVERS_LICENSE_STATE": "DE", "PHONE_NUMBER": "225-671-0796", "NAME_LAST": "JOHNSON", "entityid": "284430058", "ADDR_LINE1": "772 Armstrong RD"}`,
		LoadId:     "GO_TEST",
	},
	{
		DataSource: "GO_TEST",
		Id:         "9002",
		Data:       `{"SOCIAL_HANDLE": "flavorh", "DATE_OF_BIRTH": "4/8/1983", "ADDR_STATE": "LA", "ADDR_POSTAL_CODE": "71232", "SSN_NUMBER": "053-39-3251", "ENTITY_TYPE": "GO_TEST", "GENDER": "F", "srccode": "MDMPER", "CC_ACCOUNT_NUMBER": "5534202208773608", "RECORD_ID": "9002", "DSRC_ACTION": "A", "ADDR_CITY": "Delhi", "DRIVERS_LICENSE_STATE": "DE", "PHONE_NUMBER": "225-671-0796", "NAME_LAST": "JOHNSON", "entityid": "284430058", "ADDR_LINE1": "772 Armstrong RD"}`,
		LoadId:     "GO_TEST",
	},
	{
		DataSource: "GO_TEST",
		Id:         "9003",
		Data:       `{"SOCIAL_HANDLE": "flavorh", "DATE_OF_BIRTH": "4/8/1983", "ADDR_STATE": "LA", "ADDR_POSTAL_CODE": "71232", "SSN_NUMBER": "053-39-3251", "ENTITY_TYPE": "GO_TEST", "GENDER": "F", "srccode": "MDMPER", "CC_ACCOUNT_NUMBER": "5534202208773608", "RECORD_ID": "9003", "DSRC_ACTION": "A", "ADDR_CITY": "Delhi", "DRIVERS_LICENSE_STATE": "DE", "PHONE_NUMBER": "225-671-0796", "NAME_LAST": "JOHNSON", "entityid": "284430058", "ADDR_LINE1": "772 Armstrong RD"}`,
		LoadId:     "GO_TEST",
	},
	{
		DataSource: "GO_TEST",
		Id:         "9004",
		Data:       `{"ADDR_STATE": "LA", "ADDR_POSTAL_CODE": "71232", "ENTITY_TYPE": "GO_TEST", "GENDER": "M", "srccode": "MDMPER", "RECORD_ID": "9004", "DSRC_ACTION": "A", "ADDR_CITY": "Delhi", "PHONE_NUMBER": "225-671-0796", "NAME_LAST": "Smith", "entityid": "284430058", "ADDR_LINE1": "772 Armstrong RD"}`,
		LoadId:     "GO_TEST",
	},
}

var TestRecordsWithoutRecordId = []struct {
	DataSource string
	Data       string
	LoadId     string
}{
	{
		DataSource: "GO_TEST",
		Data:       `{"SOCIAL_HANDLE": "flavorh", "DATE_OF_BIRTH": "4/8/1983", "ADDR_STATE": "LA", "ADDR_POSTAL_CODE": "71232", "SSN_NUMBER": "053-39-3251", "ENTITY_TYPE": "GO_TEST", "GENDER": "F", "srccode": "MDMPER", "CC_ACCOUNT_NUMBER": "5534202208773608", "DSRC_ACTION": "A", "ADDR_CITY": "Delhi", "DRIVERS_LICENSE_STATE": "DE", "PHONE_NUMBER": "225-671-0796", "NAME_LAST": "JOHNSON", "entityid": "284430058", "ADDR_LINE1": "772 Armstrong RD"}`,
		LoadId:     "GO_TEST",
	},
	{
		DataSource: "GO_TEST",
		Data:       `{"SOCIAL_HANDLE": "bobby", "DATE_OF_BIRTH": "1/2/1983", "ADDR_STATE": "WI", "ADDR_POSTAL_CODE": "54434", "SSN_NUMBER": "987-65-4321", "ENTITY_TYPE": "GO_TEST", "GENDER": "F", "srccode": "MDMPER", "CC_ACCOUNT_NUMBER": "5534202208773608", "DSRC_ACTION": "A", "ADDR_CITY": "Delhi", "DRIVERS_LICENSE_STATE": "DE", "PHONE_NUMBER": "225-671-0796", "NAME_LAST": "Smith", "entityid": "284430058", "ADDR_LINE1": "772 Armstrong RD"}`,
		LoadId:     "GO_TEST",
	},
}

var TestRecordsForReplacement = []struct {
	DataSource string
	Id         string
	Data       string
	LoadId     string
}{
	{
		DataSource: "GO_TEST",
		Id:         "9001",
		Data:       `{"SOCIAL_HANDLE": "flavorh", "DATE_OF_BIRTH": "4/8/1984", "ADDR_STATE": "LA", "ADDR_POSTAL_CODE": "71232", "SSN_NUMBER": "053-39-3251", "ENTITY_TYPE": "GO_TEST", "GENDER": "F", "srccode": "MDMPER", "CC_ACCOUNT_NUMBER": "5534202208773608", "RECORD_ID": "9001", "DSRC_ACTION": "A", "ADDR_CITY": "Delhi", "DRIVERS_LICENSE_STATE": "DE", "PHONE_NUMBER": "225-671-0796", "NAME_LAST": "JOHNSON", "entityid": "284430058", "ADDR_LINE1": "772 Armstrong RD"}`,
		LoadId:     "GO_TEST",
	},
	{
		DataSource: "GO_TEST",
		Id:         "9002",
		Data:       `{"SOCIAL_HANDLE": "flavorh", "DATE_OF_BIRTH": "4/8/1984", "ADDR_STATE": "LA", "ADDR_POSTAL_CODE": "71232", "SSN_NUMBER": "053-39-3251", "ENTITY_TYPE": "GO_TEST", "GENDER": "F", "srccode": "MDMPER", "CC_ACCOUNT_NUMBER": "5534202208773608", "RECORD_ID": "9002", "DSRC_ACTION": "A", "ADDR_CITY": "Delhi", "DRIVERS_LICENSE_STATE": "DE", "PHONE_NUMBER": "225-671-0796", "NAME_LAST": "JOHNSON", "entityid": "284430058", "ADDR_LINE1": "772 Armstrong RD"}`,
		LoadId:     "GO_TEST",
	},
	{
		DataSource: "GO_TEST",
		Id:         "9003",
		Data:       `{"SOCIAL_HANDLE": "flavorh", "DATE_OF_BIRTH": "4/8/1984", "ADDR_STATE": "LA", "ADDR_POSTAL_CODE": "71232", "SSN_NUMBER": "053-39-3251", "ENTITY_TYPE": "GO_TEST", "GENDER": "F", "srccode": "MDMPER", "CC_ACCOUNT_NUMBER": "5534202208773608", "RECORD_ID": "9003", "DSRC_ACTION": "A", "ADDR_CITY": "Delhi", "DRIVERS_LICENSE_STATE": "DE", "PHONE_NUMBER": "225-671-0796", "NAME_LAST": "JOHNSON", "entityid": "284430058", "ADDR_LINE1": "772 Armstrong RD"}`,
		LoadId:     "GO_TEST",
	},
	{
		DataSource: "GO_TEST",
		Id:         "9004",
		Data:       `{"ADDR_STATE": "LA", "ADDR_POSTAL_CODE": "71232", "ENTITY_TYPE": "GO_TEST", "GENDER": "M", "srccode": "MDMPER", "RECORD_ID": "9004", "DSRC_ACTION": "A", "ADDR_CITY": "Delhi", "PHONE_NUMBER": "225-671-0796", "NAME_LAST": "Smith", "entityid": "284430058", "ADDR_LINE1": "772 Armstrong RD"}`,
		LoadId:     "GO_TEST",
	},
}
