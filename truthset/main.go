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

// A list of test records.
var TruthsetCustomerRecords = map[string]struct {
	DataSource string
	Id         string
	Data       string
	LoadId     string
}{
	"1001": {
		DataSource: "CUSTOMERS",
		Id:         "1001",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1001", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Smith", "PRIMARY_NAME_FIRST": "Robert", "DATE_OF_BIRTH": "12/11/1978", "ADDR_TYPE": "MAILING", "ADDR_LINE1": "123 Main Street, Las Vegas NV 89132", "PHONE_TYPE": "HOME", "PHONE_NUMBER": "702-919-1300", "EMAIL_ADDRESS": "bsmith@work.com", "DATE": "1/2/18", "STATUS": "Active", "AMOUNT": "100"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1002": {
		DataSource: "CUSTOMERS",
		Id:         "1002",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1002", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Smith", "PRIMARY_NAME_FIRST": "Bob", "DATE_OF_BIRTH": "11/12/1978", "ADDR_TYPE": "HOME", "ADDR_LINE1": "1515 Adela Lane", "ADDR_CITY": "Las Vegas", "ADDR_STATE": "NV", "ADDR_POSTAL_CODE": "89111", "PHONE_TYPE": "MOBILE", "PHONE_NUMBER": "702-919-1300", "DATE": "3/10/17", "STATUS": "Inactive", "AMOUNT": "200"}`,
		LoadId:     "TESTRECORDS01_LOAD",
	},
	"1003": {
		DataSource: "CUSTOMERS",
		Id:         "1003",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1003", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Smith", "PRIMARY_NAME_FIRST": "Bob", "PRIMARY_NAME_MIDDLE": "J", "DATE_OF_BIRTH": "12/11/1978", "EMAIL_ADDRESS": "bsmith@work.com", "DATE": "4/9/16", "STATUS": "Inactive", "AMOUNT": "300"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1004": {
		DataSource: "CUSTOMERS",
		Id:         "1004",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1004", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Smith", "PRIMARY_NAME_FIRST": "B", "DATE_OF_BIRTH": "11/12/1979", "ADDR_TYPE": "HOME", "ADDR_LINE1": "1515 Adela Ln", "ADDR_CITY": "Las Vegas", "ADDR_STATE": "NV", "ADDR_POSTAL_CODE": "89132", "EMAIL_ADDRESS": "bsmith@work.com", "DATE": "1/5/15", "STATUS": "Inactive", "AMOUNT": "400"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1005": {
		DataSource: "CUSTOMERS",
		Id:         "1005",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1005", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Smith", "PRIMARY_NAME_FIRST": "Robbie", "DRIVERS_LICENSE_NUMBER": "112233", "DRIVERS_LICENSE_STATE": "NV", "ADDR_TYPE": "MAILING", "ADDR_LINE1": "123 E Main St", "ADDR_CITY": "Henderson", "ADDR_STATE": "NV", "ADDR_POSTAL_CODE": "89132", "DATE": "7/16/19", "STATUS": "Active", "AMOUNT": "500"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1009": {
		DataSource: "CUSTOMERS",
		Id:         "1009",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1009", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Kusha", "PRIMARY_NAME_FIRST": "Edward", "DATE_OF_BIRTH": "3/1/1970", "SSN_NUMBER": "294-66-9999", "ADDR_TYPE": "HOME", "ADDR_LINE1": "1304 Poppy Hills Dr", "ADDR_CITY": "Blacklick", "ADDR_STATE": "OH", "ADDR_POSTAL_CODE": "43004", "EMAIL_ADDRESS": "Kusha123@hmail.com", "DATE": "1/7/18", "STATUS": "Active", "AMOUNT": "600"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1010": {
		DataSource: "CUSTOMERS",
		Id:         "1010",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1010", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Kusha", "PRIMARY_NAME_FIRST": "Eddie", "DATE_OF_BIRTH": "Mar 1 1970", "ADDR_TYPE": "HOME", "ADDR_LINE1": "1304 Poppy Hills Dr", "ADDR_CITY": "Blacklick", "ADDR_STATE": "OHIO", "DATE": "1/8/16", "STATUS": "Inactive", "AMOUNT": "700"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1011": {
		DataSource: "CUSTOMERS",
		Id:         "1011",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1011", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Knight", "PRIMARY_NAME_FIRST": "Ed", "DATE_OF_BIRTH": "3/1/70", "ADDR_TYPE": "HOME", "ADDR_LINE1": "1602 Brenville Pl", "ADDR_CITY": "San Francisco", "ADDR_STATE": "CA", "ADDR_POSTAL_CODE": "94105", "DATE": "10/9/15", "STATUS": "Terminated", "AMOUNT": "800"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1015": {
		DataSource: "CUSTOMERS",
		Id:         "1015",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1015", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Kusha", "PRIMARY_NAME_FIRST": "Mary ", "DATE_OF_BIRTH": "10/27/76", "SSN_NUMBER": "293-90-9090", "ADDR_TYPE": "HOME", "ADDR_LINE1": "1304 Poppy Hills Dr", "ADDR_CITY": "Blacklick", "ADDR_STATE": "OH", "ADDR_POSTAL_CODE": "43004", "PHONE_TYPE": "HOME", "PHONE_NUMBER": "512-353-8633", "EMAIL_ADDRESS": "Kusha123@hmail.com", "DATE": "1/10/18", "STATUS": "Active", "AMOUNT": "900"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1016": {
		DataSource: "CUSTOMERS",
		Id:         "1016",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1016", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Kusha", "PRIMARY_NAME_FIRST": "Marie", "DATE_OF_BIRTH": "10/27/76", "ADDR_TYPE": "HOME", "ADDR_LINE1": "1304 Poppy Hills Dr", "ADDR_POSTAL_CODE": "43004", "DATE": "1/11/18", "STATUS": "Active", "AMOUNT": "100"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1017": {
		DataSource: "CUSTOMERS",
		Id:         "1017",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1017", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Kusha", "PRIMARY_NAME_FIRST": "Mary ", "SSN_NUMBER": "293-90-9090", "DATE": "1/12/18", "STATUS": "Active", "AMOUNT": "200"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1018": {
		DataSource: "CUSTOMERS",
		Id:         "1018",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1018", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Kusha", "PRIMARY_NAME_FIRST": "Marie", "DATE_OF_BIRTH": "10/28/76", "PHONE_TYPE": "HOME", "PHONE_NUMBER": "512-353-8633", "DATE": "1/13/18", "STATUS": "Active", "AMOUNT": "300"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1019": {
		DataSource: "CUSTOMERS",
		Id:         "1019",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1019", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Kusha", "PRIMARY_NAME_FIRST": "Mark", "DATE_OF_BIRTH": "9/28/97", "ADDR_TYPE": "HOME", "ADDR_LINE1": "1304 Poppy Hills Dr", "ADDR_CITY": "Blacklick", "ADDR_STATE": "OH", "ADDR_POSTAL_CODE": "43004", "EMAIL_ADDRESS": "Kusha123@hmail.com", "DATE": "1/14/18", "STATUS": "Active", "AMOUNT": "400"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1020": {
		DataSource: "CUSTOMERS",
		Id:         "1020",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1020", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Kusha", "PRIMARY_NAME_FIRST": "Marsha", "DATE_OF_BIRTH": "9/28/97", "SSN_NUMBER": "201-77-7719", "ADDR_TYPE": "HOME", "ADDR_LINE1": "1304 Poppy Hills Dr", "ADDR_CITY": "Blacklick", "ADDR_STATE": "OH", "ADDR_POSTAL_CODE": "43004", "EMAIL_ADDRESS": "Kusha123@hmail.com", "DATE": "1/15/18", "STATUS": "Active", "AMOUNT": "500"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1022": {
		DataSource: "CUSTOMERS",
		Id:         "1022",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1022", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Antoun", "PRIMARY_NAME_FIRST": "Mohamed", "DATE_OF_BIRTH": "1/7/80", "DATE": "1/16/18", "STATUS": "Active", "AMOUNT": "600"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1023": {
		DataSource: "CUSTOMERS",
		Id:         "1023",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1023", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Antoun", "PRIMARY_NAME_FIRST": "Muhammed", "DATE_OF_BIRTH": "1/7/80", "DATE": "1/17/18", "STATUS": "Active", "AMOUNT": "700"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1025": {
		DataSource: "CUSTOMERS",
		Id:         "1025",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1025", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Anderson", "PRIMARY_NAME_FIRST": "Darla", "DATE_OF_BIRTH": "1/7/80", "DATE": "1/18/18", "STATUS": "Active", "AMOUNT": "800"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1026": {
		DataSource: "CUSTOMERS",
		Id:         "1026",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1026", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Anderson", "PRIMARY_NAME_FIRST": "Darlene", "DATE_OF_BIRTH": "1/7/80", "DATE": "1/19/18", "STATUS": "Active", "AMOUNT": "900"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1028": {
		DataSource: "CUSTOMERS",
		Id:         "1028",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1028", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Dobbins Jr", "PRIMARY_NAME_FIRST": "David", "ADDR_TYPE": "MAILING", "ADDR_LINE1": "1450 N City Rd Suite 900", "ADDR_CITY": "Arlington", "ADDR_STATE": "VA", "ADDR_POSTAL_CODE": "23208", "DATE": "1/20/18", "STATUS": "Active", "AMOUNT": "100"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1030": {
		DataSource: "CUSTOMERS",
		Id:         "1030",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1030", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Garski", "PRIMARY_NAME_FIRST": "Luis", "DATE_OF_BIRTH": "3/25/89", "ADDR_TYPE": "MAILING", "ADDR_LINE1": "445 Overpass Rd ", "ADDR_CITY": "San Ramon ", "ADDR_STATE": "CA ", "ADDR_POSTAL_CODE": "927230000", "DATE": "1/21/18", "STATUS": "Active", "AMOUNT": "200"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1031": {
		DataSource: "CUSTOMERS",
		Id:         "1031",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1031", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Garsky", "PRIMARY_NAME_FIRST": "Louis", "DATE_OF_BIRTH": "3/25/89", "ADDR_TYPE": "HOME", "ADDR_LINE1": "445 Overpass Rd San Ramon", "DATE": "1/22/18", "STATUS": "Active", "AMOUNT": "300"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1032": {
		DataSource: "CUSTOMERS",
		Id:         "1032",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1032", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Shaw", "PRIMARY_NAME_FIRST": "Daniella", "DATE_OF_BIRTH": "20/8/1991", "PHONE_TYPE": "HOME", "PHONE_NUMBER": "202-321-3212", "DATE": "1/23/18", "STATUS": "Active", "AMOUNT": "400"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1033": {
		DataSource: "CUSTOMERS",
		Id:         "1033",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1033", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Daniella", "PRIMARY_NAME_FIRST": "Shaw", "DATE_OF_BIRTH": "8/20/91", "ADDR_TYPE": "HOME", "ADDR_LINE1": "80 Delaware Ave SE Washington DC 40040", "PHONE_TYPE": "HOME", "PHONE_NUMBER": "321-3212", "DATE": "1/24/18", "STATUS": "Active", "AMOUNT": "500"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1034": {
		DataSource: "CUSTOMERS",
		Id:         "1034",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1034", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Medina Sentosa", "PRIMARY_NAME_FIRST": "Maria Luis", "DATE_OF_BIRTH": "11/21/73", "ADDR_TYPE": "HOME", "ADDR_LINE1": "9304 W. 15th St La Blanca, FL 60527", "EMAIL_ADDRESS": "Maria Sentosa<msentosa@fmail.com>", "DATE": "1/25/18", "STATUS": "Active", "AMOUNT": "600"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1035": {
		DataSource: "CUSTOMERS",
		Id:         "1035",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1035", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "M Sentosa", "PRIMARY_NAME_FIRST": "Maria Luis", "DATE_OF_BIRTH": "11/12/73", "ADDR_TYPE": "HOME", "ADDR_LINE1": "9304 W. 15th St La Blanca, FL 60527", "DATE": "1/26/18", "STATUS": "Active", "AMOUNT": "700"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1036": {
		DataSource: "CUSTOMERS",
		Id:         "1036",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1036", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Sentosa", "PRIMARY_NAME_FIRST": "Maria Luis", "DATE_OF_BIRTH": "11/12/73", "ADDR_TYPE": "HOME", "ADDR_LINE1": "9304 W. 15th St La Blanca, FL 60527", "DATE": "1/27/18", "STATUS": "Active", "AMOUNT": "800"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1039": {
		DataSource: "CUSTOMERS",
		Id:         "1039",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1039", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Smith", "PRIMARY_NAME_FIRST": "John", "GENDER": "M", "DATE_OF_BIRTH": "10/10/70", "ADDR_TYPE": "HOME", "ADDR_LINE1": "3212 W. 32nd St Palm Harbor, FL 60527", "DATE": "1/28/18", "STATUS": "Active", "AMOUNT": "900"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1040": {
		DataSource: "CUSTOMERS",
		Id:         "1040",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1040", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Smith", "PRIMARY_NAME_FIRST": "John", "DATE_OF_BIRTH": "3/15/90", "ADDR_TYPE": "HOME", "ADDR_LINE1": "3212 W. 32nd St Palm Harbor, FL 60527", "DATE": "1/29/18", "STATUS": "Active", "AMOUNT": "100"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1043": {
		DataSource: "CUSTOMERS",
		Id:         "1043",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1043", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Smith", "PRIMARY_NAME_FIRST": "Patrick", "DATE_OF_BIRTH": "10/10/70", "PASSPORT_NUMBER": "10251111", "PASSPORT_COUNTRY": "US", "ADDR_TYPE": "HOME", "ADDR_LINE1": "3212 W. 32nd St Palm Harbor, FL 60527", "DATE": "1/30/18", "STATUS": "Active", "AMOUNT": "200"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1044": {
		DataSource: "CUSTOMERS",
		Id:         "1044",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1044", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Smith", "PRIMARY_NAME_FIRST": "Patricia", "DATE_OF_BIRTH": "3/15/90", "PASSPORT_NUMBER": "10252222", "PASSPORT_COUNTRY": "US", "ADDR_TYPE": "HOME", "ADDR_LINE1": "3212 W. 32nd St Palm Harbor, FL 60527", "DATE": "1/31/18", "STATUS": "Active", "AMOUNT": "300"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1045": {
		DataSource: "CUSTOMERS",
		Id:         "1045",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1045", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Smith", "PRIMARY_NAME_FIRST": "Pat", "PASSPORT_NUMBER": "10251111", "ADDR_TYPE": "HOME", "ADDR_LINE1": "3212 W. 32nd St Palm Harbor, FL 60527", "DATE": "1/2/18", "STATUS": "Active", "AMOUNT": "400"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1046": {
		DataSource: "CUSTOMERS",
		Id:         "1046",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1046", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Smith", "PRIMARY_NAME_FIRST": "Pat", "PASSPORT_NUMBER": "10252222", "PASSPORT_COUNTRY": "USA", "ADDR_TYPE": "HOME", "ADDR_LINE1": "3212 W. 32nd St Palm Harbor, FL 60527", "DATE": "1/3/18", "STATUS": "Active", "AMOUNT": "500"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1047": {
		DataSource: "CUSTOMERS",
		Id:         "1047",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1047", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Thompson", "PRIMARY_NAME_FIRST": "Zara", "EMAIL_ADDRESS": "sthomp45@fmail.com", "DATE": "1/4/18", "STATUS": "Active", "AMOUNT": "600"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1048": {
		DataSource: "CUSTOMERS",
		Id:         "1048",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1048", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Tompson", "PRIMARY_NAME_FIRST": "Sarah", "EMAIL_ADDRESS": "sthomp45@fmail.com", "DATE": "1/5/18", "STATUS": "Active", "AMOUNT": "700"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1049": {
		DataSource: "CUSTOMERS",
		Id:         "1049",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1049", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Thompson", "PRIMARY_NAME_FIRST": "Sahra", "EMAIL_ADDRESS": "sthomp45@fmail.com", "DATE": "1/6/18", "STATUS": "Active", "AMOUNT": "800"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1050": {
		DataSource: "CUSTOMERS",
		Id:         "1050",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1050", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Toulouse", "PRIMARY_NAME_FIRST": "Lee", "DATE_OF_BIRTH": "2/1/85", "PASSPORT_NUMBER": "483290175", "PASSPORT_COUNTRY": "USA", "DATE": "1/7/18", "STATUS": "Active", "AMOUNT": "900"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1051": {
		DataSource: "CUSTOMERS",
		Id:         "1051",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1051", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Toulouse", "PRIMARY_NAME_FIRST": "Leigh", "DATE_OF_BIRTH": "1/2/85", "PASSPORT_NUMBER": "483290175", "PASSPORT_COUNTRY": "US", "DATE": "1/8/18", "STATUS": "Active", "AMOUNT": "100"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1052": {
		DataSource: "CUSTOMERS",
		Id:         "1052",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1052", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Toulouse", "PRIMARY_NAME_FIRST": "Lea", "PASSPORT_NUMBER": "483290175", "PASSPORT_COUNTRY": "US", "DATE": "1/9/18", "STATUS": "Active", "AMOUNT": "200"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1053": {
		DataSource: "CUSTOMERS",
		Id:         "1053",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1053", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Smith", "PRIMARY_NAME_FIRST": "Beau", "PASSPORT_NUMBER": "72129291", "PASSPORT_COUNTRY": "CA", "ADDR_TYPE": "HOME", "ADDR_LINE1": "6371 E Foothill Dr, Orroville, CA ", "DATE": "1/10/18", "STATUS": "Active", "AMOUNT": "300"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1054": {
		DataSource: "CUSTOMERS",
		Id:         "1054",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1054", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Smith", "PRIMARY_NAME_FIRST": "Magdalena", "DATE_OF_BIRTH": "24-May-11", "DRIVERS_LICENSE_NUMBER": "93939211", "DRIVERS_LICENSE_STATE": "CA", "ADDR_TYPE": "HOME", "ADDR_LINE1": "6371 E Foothill Dr, Orroville, CA 95915", "DATE": "1/11/18", "STATUS": "Active", "AMOUNT": "400"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1055": {
		DataSource: "CUSTOMERS",
		Id:         "1055",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1055", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_FIRST": "Beau", "PASSPORT_NUMBER": "72129291", "PASSPORT_COUNTRY": "CAN", "ADDR_TYPE": "HOME", "ADDR_LINE1": "6371 E Foothill Dr, 95915", "DATE": "1/12/18", "STATUS": "Active", "AMOUNT": "500"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1056": {
		DataSource: "CUSTOMERS",
		Id:         "1056",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1056", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Jones", "PRIMARY_NAME_FIRST": "Magdalena", "DATE_OF_BIRTH": "5/24/11", "DRIVERS_LICENSE_NUMBER": "93939211", "DRIVERS_LICENSE_STATE": "CA", "ADDR_TYPE": "HOME", "ADDR_LINE1": "6371 E Foothill Dr, Orroville, CA ", "DATE": "1/13/18", "STATUS": "Active", "AMOUNT": "600"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1057": {
		DataSource: "CUSTOMERS",
		Id:         "1057",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1057", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Jones", "PRIMARY_NAME_FIRST": "Jay", "EMAIL_ADDRESS": "jjones@jones.com", "DATE": "1/14/18", "STATUS": "Active", "AMOUNT": "700"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1058": {
		DataSource: "CUSTOMERS",
		Id:         "1058",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1058", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Jay", "PRIMARY_NAME_FIRST": "Jones", "EMAIL_ADDRESS": "\"Jay Jones\" <jjones@fmail.com>", "DATE": "1/15/18", "STATUS": "Active", "AMOUNT": "800"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1059": {
		DataSource: "CUSTOMERS",
		Id:         "1059",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1059", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Roderick", "PRIMARY_NAME_FIRST": "Ray", "PHONE_TYPE": "HOME", "PHONE_NUMBER": "971-421-8250", "DATE": "1/16/18", "STATUS": "Active", "AMOUNT": "900"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1060": {
		DataSource: "CUSTOMERS",
		Id:         "1060",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1060", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Roderick", "PRIMARY_NAME_FIRST": "R", "PHONE_TYPE": "HOME", "PHONE_NUMBER": "9714218250", "DATE": "1/17/18", "STATUS": "Active", "AMOUNT": "100"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1061": {
		DataSource: "CUSTOMERS",
		Id:         "1061",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1061", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Andreason", "PHONE_TYPE": "MOBILE", "PHONE_NUMBER": "(807) 422-9031", "DATE": "1/18/18", "STATUS": "Active", "AMOUNT": "200"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1062": {
		DataSource: "CUSTOMERS",
		Id:         "1062",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1062", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Andreason", "PHONE_NUMBER": "807-422-9031", "DATE": "1/19/18", "STATUS": "Active", "AMOUNT": "300"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1063": {
		DataSource: "CUSTOMERS",
		Id:         "1063",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1063", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Mooney", "PRIMARY_NAME_FIRST": "Susan", "DATE_OF_BIRTH": "6/15/98", "DATE": "1/20/18", "STATUS": "Active", "AMOUNT": "400"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1064": {
		DataSource: "CUSTOMERS",
		Id:         "1064",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1064", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Mooney", "PRIMARY_NAME_FIRST": "Susanne", "DATE_OF_BIRTH": "6/15/98", "PASSPORT_NUMBER": "1231345345", "PASSPORT_COUNTRY": "US", "DATE": "1/21/18", "STATUS": "Active", "AMOUNT": "500"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1065": {
		DataSource: "CUSTOMERS",
		Id:         "1065",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1065", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Mooney", "PRIMARY_NAME_FIRST": "Susan", "PASSPORT_NUMBER": "1231345345", "PASSPORT_COUNTRY": "US", "DRIVERS_LICENSE_NUMBER": "8923322", "DRIVERS_LICENSE_STATE": "OR", "DATE": "1/22/18", "STATUS": "Active", "AMOUNT": "600"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1066": {
		DataSource: "CUSTOMERS",
		Id:         "1066",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1066", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Mooney", "PRIMARY_NAME_FIRST": "Susan", "DRIVERS_LICENSE_NUMBER": "8923322", "DRIVERS_LICENSE_STATE": "OR", "SSN_NUMBER": "521-21-2123", "DATE": "1/23/18", "STATUS": "Active", "AMOUNT": "700"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1067": {
		DataSource: "CUSTOMERS",
		Id:         "1067",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1067", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Moonie", "PRIMARY_NAME_FIRST": "Susan", "SSN_NUMBER": "521212123", "ADDR_TYPE": "HOME", "ADDR_LINE1": "638 Downey St, Salem, OR", "DATE": "1/24/18", "STATUS": "Active", "AMOUNT": "800"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1068": {
		DataSource: "CUSTOMERS",
		Id:         "1068",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1068", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Moony", "PRIMARY_NAME_FIRST": "Susan", "ADDR_TYPE": "MAILING", "ADDR_LINE1": "Adventura Aparments 638 Downey St, Salem, OR", "DATE": "1/25/18", "STATUS": "Active", "AMOUNT": "900"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1069": {
		DataSource: "CUSTOMERS",
		Id:         "1069",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1069", "RECORD_TYPE": "PERSON", "NATIVE_NAME_FULL": "\u738b\u6770", "GENDER": "M", "DATE_OF_BIRTH": "9/14/93", "NATIONAL_ID_NUMBER": "832721", "ADDR_TYPE": "HOME", "ADDR_LINE1": "12 Constitution Street ", "DATE": "1/26/18", "STATUS": "Active", "AMOUNT": "100"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1070": {
		DataSource: "CUSTOMERS",
		Id:         "1070",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1070", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Wang", "PRIMARY_NAME_FIRST": "Jie", "GENDER": "Male", "DATE_OF_BIRTH": "9/14/93", "NATIONAL_ID_NUMBER": "832721", "NATIONAL_ID_COUNTRY": "Hong Kong", "ADDR_TYPE": "HOME", "ADDR_LINE1": "12 Constitution Street ", "DATE": "1/27/18", "STATUS": "Active", "AMOUNT": "200"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1071": {
		DataSource: "CUSTOMERS",
		Id:         "1071",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1071", "RECORD_TYPE": "PERSON", "NATIVE_NAME_FULL": "\u738b\u4f1f", "GENDER": "F", "DATE_OF_BIRTH": "9/14/97", "NATIONAL_ID_NUMBER": "7123833", "NATIONAL_ID_COUNTRY": "China", "ADDR_TYPE": "HOME", "ADDR_LINE1": "169 3rd Ave. Camden, NJ 08030", "DATE": "1/28/18", "STATUS": "Active", "AMOUNT": "300"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1072": {
		DataSource: "CUSTOMERS",
		Id:         "1072",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1072", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Wang", "PRIMARY_NAME_FIRST": "Wei", "GENDER": "Female", "DATE_OF_BIRTH": "9/14/97", "NATIONAL_ID_NUMBER": "7123833", "NATIONAL_ID_COUNTRY": "China", "ADDR_TYPE": "HOME", "ADDR_LINE1": "169 3rd Ave. Camden, NJ 08030", "DATE": "1/29/18", "STATUS": "Active", "AMOUNT": "400"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1073": {
		DataSource: "CUSTOMERS",
		Id:         "1073",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1073", "RECORD_TYPE": "PERSON", "NATIVE_NAME_FULL": "\u5f20\u4f1f", "GENDER": "M", "DATE_OF_BIRTH": "8/2/06", "ADDR_TYPE": "HOME", "ADDR_LINE1": "173 John Lane, Camden, NJ 08030", "DATE": "1/30/18", "STATUS": "Active", "AMOUNT": "500"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1074": {
		DataSource: "CUSTOMERS",
		Id:         "1074",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1074", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Zhang", "PRIMARY_NAME_FIRST": "Wei", "GENDER": "Male", "DATE_OF_BIRTH": "2/8/06", "ADDR_TYPE": "HOME", "ADDR_LINE1": "173 John Lane, 08030", "DATE": "1/31/18", "STATUS": "Active", "AMOUNT": "600"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1075": {
		DataSource: "CUSTOMERS",
		Id:         "1075",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1075", "RECORD_TYPE": "PERSON", "NATIVE_NAME_FULL": "\u5f20\u79c0\u82f1", "GENDER": "F", "DATE_OF_BIRTH": "2/4/31", "ADDR_TYPE": "HOME", "ADDR_LINE1": "329 Leatherwood Street, Las Vegas, 89117", "DATE": "1/2/18", "STATUS": "Active", "AMOUNT": "700"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1076": {
		DataSource: "CUSTOMERS",
		Id:         "1076",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1076", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Zhang", "PRIMARY_NAME_FIRST": "Xiu Ying", "GENDER": "Female", "DATE_OF_BIRTH": "4/2/31", "ADDR_TYPE": "HOME", "ADDR_LINE1": "329 Leatherwood Street, Las Vegas, NV", "DATE": "1/3/18", "STATUS": "Active", "AMOUNT": "800"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1077": {
		DataSource: "CUSTOMERS",
		Id:         "1077",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1077", "RECORD_TYPE": "PERSON", "NATIVE_NAME_FULL": "\u5218\u6770", "GENDER": "F", "DATE_OF_BIRTH": "6/25/08", "ADDR_TYPE": "HOME", "ADDR_LINE1": "37 Campfire St. ", "DATE": "1/4/18", "STATUS": "Active", "AMOUNT": "900"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1078": {
		DataSource: "CUSTOMERS",
		Id:         "1078",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1078", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Liu", "PRIMARY_NAME_FIRST": "Jie", "GENDER": "Unknown", "DATE_OF_BIRTH": "25-Jun-08", "ADDR_TYPE": "HOME", "ADDR_LINE1": "37 Campfire St. ", "DATE": "1/5/18", "STATUS": "Active", "AMOUNT": "100"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1079": {
		DataSource: "CUSTOMERS",
		Id:         "1079",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1079", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Brown", "PRIMARY_NAME_FIRST": "Jeffrey", "GENDER": "U", "DATE_OF_BIRTH": "6/21/82", "SSN_NUMBER": "3241", "DATE": "1/6/18", "STATUS": "Active", "AMOUNT": "200"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1080": {
		DataSource: "CUSTOMERS",
		Id:         "1080",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1080", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Brown Jr", "PRIMARY_NAME_FIRST": "Geoffrey", "GENDER": "M", "DATE_OF_BIRTH": "6/21/82", "SSN_NUMBER": "3241", "DATE": "1/7/18", "STATUS": "Active", "AMOUNT": "300"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1081": {
		DataSource: "CUSTOMERS",
		Id:         "1081",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1081", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Frankens", "PRIMARY_NAME_FIRST": "George", "DATE_OF_BIRTH": "15-Mar-92", "PASSPORT_NUMBER": "234456456", "PASSPORT_COUNTRY": "DE", "ADDR_TYPE": "HOME", "ADDR_LINE1": "Ansbacher Strasse 23, 56422 Dusseldorf", "ADDR_POSTAL_CODE": "56244", "DATE": "1/8/18", "STATUS": "Active", "AMOUNT": "400"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1082": {
		DataSource: "CUSTOMERS",
		Id:         "1082",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1082", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Frankens", "PRIMARY_NAME_FIRST": "Georg", "DATE_OF_BIRTH": "15-Mar-92", "PASSPORT_NUMBER": "234456456", "PASSPORT_COUNTRY": "Germany", "ADDR_TYPE": "MAILING", "ADDR_LINE1": "23 Ansbacher Street", "ADDR_CITY": "Dusseldorf", "ADDR_POSTAL_CODE": "56244", "ADDR_COUNTRY": "Germany", "DATE": "1/9/18", "STATUS": "Active", "AMOUNT": "500"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1083": {
		DataSource: "CUSTOMERS",
		Id:         "1083",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1083", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Klempski", "PRIMARY_NAME_FIRST": "Morris", "DATE_OF_BIRTH": "17-May-90", "PASSPORT_NUMBER": "34543555", "PASSPORT_COUNTRY": "CA", "ADDR_TYPE": "HOME", "ADDR_LINE1": "Skyline Apartments, 705 Sheppard Ave", "ADDR_CITY": "Toronto", "ADDR_POSTAL_CODE": "M1S 1T4", "DATE": "1/10/18", "STATUS": "Active", "AMOUNT": "600"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1084": {
		DataSource: "CUSTOMERS",
		Id:         "1084",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1084", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Klempsky", "PRIMARY_NAME_FIRST": "Morrie", "DATE_OF_BIRTH": "17-May-90", "PASSPORT_NUMBER": "34543555", "PASSPORT_COUNTRY": "Canada", "ADDR_TYPE": "MAILING", "ADDR_LINE1": "705 Sheppard Ave", "ADDR_CITY": "Toronto", "ADDR_POSTAL_CODE": "M1S 1T4", "ADDR_COUNTRY": "CAN", "DATE": "1/11/18", "STATUS": "Active", "AMOUNT": "700"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1085": {
		DataSource: "CUSTOMERS",
		Id:         "1085",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1085", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "D'esquire", "PRIMARY_NAME_FIRST": "Ellie", "DATE_OF_BIRTH": "19-Feb-91", "PHONE_TYPE": "HOME", "PHONE_NUMBER": "0352 6553537", "EMAIL_ADDRESS": "dellie@fmail.com", "DATE": "1/12/18", "STATUS": "Active", "AMOUNT": "800"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1086": {
		DataSource: "CUSTOMERS",
		Id:         "1086",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1086", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Desqueir", "PRIMARY_NAME_FIRST": "Ellie", "DATE_OF_BIRTH": "19-Feb-91", "PHONE_TYPE": "HOME", "PHONE_NUMBER": "+39 0352 6553537", "EMAIL_ADDRESS": "dellie@fmail.com", "DATE": "1/13/18", "STATUS": "Active", "AMOUNT": "900"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1087": {
		DataSource: "CUSTOMERS",
		Id:         "1087",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1087", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Wiest", "PRIMARY_NAME_FIRST": "George", "GENDER": "M", "DATE_OF_BIRTH": "3/12/87", "PHONE_TYPE": "HOME", "PHONE_NUMBER": "702-221-2412", "EMAIL_ADDRESS": "pfranks@ishmail.com", "DATE": "1/14/18", "STATUS": "Active", "AMOUNT": "100"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1088": {
		DataSource: "CUSTOMERS",
		Id:         "1088",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1088", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Weest", "PRIMARY_NAME_FIRST": "George", "GENDER": "F", "DATE_OF_BIRTH": "3/12/87", "PHONE_TYPE": "HOME", "PHONE_NUMBER": "221-2412", "EMAIL_ADDRESS": "pfranks@ishmail.com", "DATE": "1/15/18", "STATUS": "Active", "AMOUNT": "200"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1089": {
		DataSource: "CUSTOMERS",
		Id:         "1089",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1089", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Klein", "PRIMARY_NAME_FIRST": "Morris I", "DATE_OF_BIRTH": "4/12/82", "DATE": "1/16/18", "STATUS": "Active", "AMOUNT": "300"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1090": {
		DataSource: "CUSTOMERS",
		Id:         "1090",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1090", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Klein", "PRIMARY_NAME_FIRST": "Morris II", "DATE_OF_BIRTH": "4/12/82", "DATE": "1/17/18", "STATUS": "Active", "AMOUNT": "400"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1091": {
		DataSource: "CUSTOMERS",
		Id:         "1091",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1091", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Ohare", "PRIMARY_NAME_FIRST": "Ellie", "DATE_OF_BIRTH": "8/15/67", "PHONE_TYPE": "HOME", "PHONE_NUMBER": "0352 6553537", "EMAIL_ADDRESS": "ellie.ohare@fmail.com", "DATE": "1/18/18", "STATUS": "Active", "AMOUNT": "500"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1092": {
		DataSource: "CUSTOMERS",
		Id:         "1092",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1092", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "O'hare", "PRIMARY_NAME_FIRST": "Ellie", "DATE_OF_BIRTH": "8/15/67", "PHONE_TYPE": "HOME", "PHONE_NUMBER": "+39 0352 6553537", "EMAIL_ADDRESS": "ellie.ohare@fmail.com", "DATE": "1/19/18", "STATUS": "Active", "AMOUNT": "600"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1093": {
		DataSource: "CUSTOMERS",
		Id:         "1093",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1093", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Anderson", "PRIMARY_NAME_FIRST": "Amanda", "DATE_OF_BIRTH": "3/12/87", "DRIVERS_LICENSE_NUMBER": "73423499", "DRIVERS_LICENSE_STATE": "MN", "DATE": "1/20/18", "STATUS": "Active", "AMOUNT": "700"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1094": {
		DataSource: "CUSTOMERS",
		Id:         "1094",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1094", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Jones", "PRIMARY_NAME_FIRST": "Amanda", "DATE_OF_BIRTH": "3/12/87", "DRIVERS_LICENSE_NUMBER": "73423499", "DRIVERS_LICENSE_STATE": "MN", "DATE": "1/21/18", "STATUS": "Active", "AMOUNT": "800"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1095": {
		DataSource: "CUSTOMERS",
		Id:         "1095",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1095", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Aguilar", "PRIMARY_NAME_FIRST": "Juan", "GENDER": "Male", "DATE_OF_BIRTH": "4/12/82", "DRIVERS_LICENSE_NUMBER": "234234455", "DRIVERS_LICENSE_STATE": "MN", "DATE": "1/22/18", "STATUS": "Active", "AMOUNT": "900"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1096": {
		DataSource: "CUSTOMERS",
		Id:         "1096",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1096", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Aguilar", "PRIMARY_NAME_FIRST": "Juann", "DATE_OF_BIRTH": "4/12/82", "DRIVERS_LICENSE_NUMBER": "234234455", "DRIVERS_LICENSE_STATE": "MN", "DATE": "1/23/18", "STATUS": "Active", "AMOUNT": "100"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1097": {
		DataSource: "CUSTOMERS",
		Id:         "1097",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1097", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Sanchez", "PRIMARY_NAME_FIRST": "Marie", "ADDR_TYPE": "MAILING", "ADDR_LINE1": "P.O. Box 12987", "ADDR_CITY": "Andersonville", "ADDR_STATE": "IL", "ADDR_POSTAL_CODE": "60611", "PHONE_TYPE": "MOBILE", "EMAIL_ADDRESS": "mickey@mmail.com", "DATE": "1/24/18", "STATUS": "Active", "AMOUNT": "200"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1098": {
		DataSource: "CUSTOMERS",
		Id:         "1098",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1098", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Sanchez Mendoza", "PRIMARY_NAME_FIRST": "Marie", "ADDR_TYPE": "MAILING", "ADDR_LINE1": "PO BOX 12987", "ADDR_CITY": "Chicago", "ADDR_STATE": "IL", "ADDR_POSTAL_CODE": "60611", "PHONE_TYPE": "MOBILE", "EMAIL_ADDRESS": "mickey@mmail.com", "DATE": "1/25/18", "STATUS": "Active", "AMOUNT": "300"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1099": {
		DataSource: "CUSTOMERS",
		Id:         "1099",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1099", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Aguilar", "PRIMARY_NAME_FIRST": "Anna Maria", "GENDER": "Female", "ADDR_TYPE": "HOME", "ADDR_LINE1": "1812 Overture way", "ADDR_CITY": "Chicago", "ADDR_STATE": "IL", "PHONE_TYPE": "MOBILE", "EMAIL_ADDRESS": "mouse@mmail.com", "DATE": "1/26/18", "STATUS": "Active", "AMOUNT": "400"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1100": {
		DataSource: "CUSTOMERS",
		Id:         "1100",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1100", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Aguilar", "PRIMARY_NAME_FIRST": "Anna", "PRIMARY_NAME_MIDDLE": "Marie", "GENDER": "Unknown", "ADDR_TYPE": "HOME", "ADDR_LINE1": "9881 Freedom way", "ADDR_CITY": "Chicago", "ADDR_STATE": "IL", "PHONE_TYPE": "MOBILE", "EMAIL_ADDRESS": "mouse@mmail.com", "DATE": "1/27/18", "STATUS": "Active", "AMOUNT": "500"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1101": {
		DataSource: "CUSTOMERS",
		Id:         "1101",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1101", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Miller", "PRIMARY_NAME_FIRST": "Mark", "EMAIL_ADDRESS": "mark@marksfoods.com", "DATE": "1/28/18", "STATUS": "Active", "AMOUNT": "600"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1102": {
		DataSource: "CUSTOMERS",
		Id:         "1102",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1102", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Milner", "PRIMARY_NAME_FIRST": "Mark", "EMAIL_ADDRESS": "mark@marksfoods.com", "DATE": "1/29/18", "STATUS": "Active", "AMOUNT": "700"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1103": {
		DataSource: "CUSTOMERS",
		Id:         "1103",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1103", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Becker", "PRIMARY_NAME_FIRST": "Anabella", "GENDER": "U", "DRIVERS_LICENSE_NUMBER": "823123", "DRIVERS_LICENSE_STATE": "TX", "DATE": "1/30/18", "STATUS": "Active", "AMOUNT": "800"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"1104": {
		DataSource: "CUSTOMERS",
		Id:         "1104",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "1104", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Becker", "PRIMARY_NAME_FIRST": "Annabelle", "GENDER": "F", "DRIVERS_LICENSE_NUMBER": "823123", "DRIVERS_LICENSE_STATE": "Texas", "DATE": "1/31/18", "STATUS": "Active", "AMOUNT": "900"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"2011": {
		DataSource: "CUSTOMERS",
		Id:         "2011",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "2011", "RECORD_TYPE": "ORGANIZATION", "PRIMARY_NAME_ORG": "Hajah Mamunah (Jln Pisang)", "ADDR_TYPE": "BUSINESS", "ADDR_FULL": "#01-11, HillV2, 4 Hillview Rise, 667979", "ADDR_COUNTRY": "Singapore", "DATE": "1/31/18", "STATUS": "Inactive", "CATEGORY": "Platinum"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"2031": {
		DataSource: "CUSTOMERS",
		Id:         "2031",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "2031", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "\u1782\u17b9\u1798", "PRIMARY_NAME_FIRST": "\u178f\u17b6\u179a\u17b6", "ADDR_TYPE": "PRIMARY", "ADDR_FULL": "Street 128 Phnom Penh Cambodia", "DATE": "3/15/1992", "STATUS": "Active", "CATEGORY": "Gold"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"2032": {
		DataSource: "CUSTOMERS",
		Id:         "2032",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "2032", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Kim", "PRIMARY_NAME_FIRST": "Dara", "ADDR_TYPE": "PRIMARY", "ADDR_LINE1": "Street 128 ", "ADDR_CITY": "Phnom Penh", "ADDR_COUNTRY": "Cambodia", "DATE": "3/12/1998", "STATUS": "Active", "CATEGORY": "Silver"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"2042": {
		DataSource: "CUSTOMERS",
		Id:         "2042",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "2042", "RECORD_TYPE": "ORGANIZATION", "PRIMARY_NAME_ORG": "Mullenkrants ", "SECONDARY_NAME_ORG": "Autoworks", "ADDR_TYPE": "PRIMARY", "ADDR_LINE1": "Hardenbergstrasse 87", "ADDR_POSTAL_CODE": "66879", "ADDR_COUNTRY": "Germany", "DATE": "3/15/2019", "STATUS": "Terminated", "CATEGORY": "Platinum"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"2063": {
		DataSource: "CUSTOMERS",
		Id:         "2063",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "2063", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_FULL": "Alexander Pavlovich Vasiliev", "PHONE_NUMBER": "481-285-6234", "DATE": "1/15/2000", "STATUS": "Active", "CATEGORY": "Platinum"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"2072": {
		DataSource: "CUSTOMERS",
		Id:         "2072",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "2072", "RECORD_TYPE": "ORGANIZATION", "PRIMARY_NAME_ORG": "Univrsl Export Inc", "ADDR_TYPE": "BUSINESS", "ADDR_LINE1": "100 Howard Hughs Plaza", "ADDR_CITY": "Las Vegas", "ADDR_STATE": "NV", "ADDR_POSTAL_CODE": "89111", "PHONE_NUMBER": "800-111-1234", "DATE": "6/15/2005", "STATUS": "Active", "CATEGORY": "Silver"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"2073": {
		DataSource: "CUSTOMERS",
		Id:         "2073",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "2073", "RECORD_TYPE": "ORGANIZATION", "PRIMARY_NAME_ORG": "Worldwide Exports ", "ADDR_TYPE": "REGISTERED", "ADDR_LINE1": "Chrysler Building, 405 Lexington Avenue", "ADDR_CITY": "New York", "ADDR_STATE": "NY", "ADDR_POSTAL_CODE": "10174", "DATE": "12/10/2020", "STATUS": "Active", "CATEGORY": "Platinum"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"2142": {
		DataSource: "CUSTOMERS",
		Id:         "2142",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "2142", "RECORD_TYPE": "ORGANIZATION", "PRIMARY_NAME_ORG": "Singapore exports", "ADDR_TYPE": "PRIMARY", "ADDR_FULL": "133 New Bridge Road, Chinatown Point, Singapore 059413", "ADDR_COUNTRY": "Singapore", "DATE": "2/4/2012", "STATUS": "Active", "CATEGORY": "Silver"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"2152": {
		DataSource: "CUSTOMERS",
		Id:         "2152",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "2152", "RECORD_TYPE": "ORGANIZATION", "PRIMARY_NAME_ORG": "India Exports", "ADDR_TYPE": "PRIMARY", "ADDR_FULL": "Mullanpara Road, Old Vythiri, Vythiri, Wayanad, 673576, India", "DATE": "3/1/2010", "STATUS": "Active", "CATEGORY": "Gold"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"2171": {
		DataSource: "CUSTOMERS",
		Id:         "2171",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "2171", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Anderson", "PRIMARY_NAME_FIRST": "Andrew", "EMAIL_ADDRESS": "info@ca-state.gov"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"2172": {
		DataSource: "CUSTOMERS",
		Id:         "2172",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "2172", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Anderson", "PRIMARY_NAME_FIRST": "Andy ", "EMAIL_ADDRESS": "info@ca-state.gov"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"2181": {
		DataSource: "CUSTOMERS",
		Id:         "2181",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "2181", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Anderson", "PRIMARY_NAME_FIRST": "Anna", "PHONE_NUMBER": "702-221-2211", "EMAIL_ADDRESS": "info@ca-state.gov"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"2182": {
		DataSource: "CUSTOMERS",
		Id:         "2182",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "2182", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Anderson", "PRIMARY_NAME_FIRST": "Annabelle", "PHONE_NUMBER": "702-221-2211", "EMAIL_ADDRESS": "info@ca-state.gov"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"2191": {
		DataSource: "CUSTOMERS",
		Id:         "2191",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "2191", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Muir", "PRIMARY_NAME_FIRST": "Jim", "DATE_OF_BIRTH": "1997-11-12", "ADDR_LINE1": "12396 Austin Rd", "ADDR_CITY": "Sacramento", "ADDR_STATE": "CA", "ADDR_POSTAL_CODE": "95823", "EMAIL_ADDRESS": "info@ca-state.gov"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"2192": {
		DataSource: "CUSTOMERS",
		Id:         "2192",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "2192", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Muir", "PRIMARY_NAME_FIRST": "Jane", "DATE_OF_BIRTH": "1999-12-10", "ADDR_LINE1": "12396 Austin Rd", "ADDR_CITY": "Sacramento", "ADDR_STATE": "CA", "ADDR_POSTAL_CODE": "95823", "EMAIL_ADDRESS": "info@ca-state.gov"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"2193": {
		DataSource: "CUSTOMERS",
		Id:         "2193",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "2193", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Muir", "PRIMARY_NAME_FIRST": "J", "DATE_OF_BIRTH": "1999-12-10", "ADDR_LINE1": "12396 Austin Rd", "ADDR_CITY": "Sacramento", "ADDR_STATE": "CA", "ADDR_POSTAL_CODE": "95823", "EMAIL_ADDRESS": "info@ca-state.gov"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"2201": {
		DataSource: "CUSTOMERS",
		Id:         "2201",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "2201", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Swarm", "PRIMARY_NAME_FIRST": "Jorg", "ADDR_LINE1": "127 14th Ave", "ADDR_CITY": "Elmwood Park", "ADDR_STATE": "CA", "ADDR_POSTAL_CODE": "95865", "EMAIL_ADDRESS": "info@ca-state.gov"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"2202": {
		DataSource: "CUSTOMERS",
		Id:         "2202",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "2202", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Barge", "PRIMARY_NAME_FIRST": "Jorge", "ADDR_LINE1": "4362 Belmont Lane", "ADDR_CITY": "Elmwood Park", "ADDR_STATE": "CA", "ADDR_POSTAL_CODE": "95865", "EMAIL_ADDRESS": "info@ca-state.gov"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"2203": {
		DataSource: "CUSTOMERS",
		Id:         "2203",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "2203", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Gray", "PRIMARY_NAME_FIRST": "Gaston", "ADDR_LINE1": "1376 BlueBell Road", "ADDR_CITY": "Sacramento", "ADDR_STATE": "CA", "ADDR_POSTAL_CODE": "95823", "EMAIL_ADDRESS": "info@ca-state.gov"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"2204": {
		DataSource: "CUSTOMERS",
		Id:         "2204",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "2204", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Howard", "PRIMARY_NAME_FIRST": "Henry", "ADDR_LINE1": "538 Blanco St", "ADDR_CITY": "Sacramento", "ADDR_STATE": "CA", "ADDR_POSTAL_CODE": "95823", "EMAIL_ADDRESS": "info@ca-state.gov"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"2205": {
		DataSource: "CUSTOMERS",
		Id:         "2205",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "2205", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Inverness", "PRIMARY_NAME_FIRST": "Inez", "ADDR_LINE1": "2516 BentTree Ln", "ADDR_CITY": "Sacramento", "ADDR_STATE": "CA", "ADDR_POSTAL_CODE": "95823", "EMAIL_ADDRESS": "info@ca-state.gov"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"2206": {
		DataSource: "CUSTOMERS",
		Id:         "2206",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "2206", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Jackson", "PRIMARY_NAME_FIRST": "Julia", "ADDR_LINE1": "319 Cody Road", "ADDR_CITY": "Elmwood Park", "ADDR_STATE": "CA", "ADDR_POSTAL_CODE": "95865", "EMAIL_ADDRESS": "info@ca-state.gov"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"2207": {
		DataSource: "CUSTOMERS",
		Id:         "2207",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "2207", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Kellar", "PRIMARY_NAME_FIRST": "Kandace", "ADDR_LINE1": "1824 AspenOak Way", "ADDR_CITY": "Elmwood Park", "ADDR_STATE": "CA", "ADDR_POSTAL_CODE": "95865", "EMAIL_ADDRESS": "info@ca-state.gov"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"2208": {
		DataSource: "CUSTOMERS",
		Id:         "2208",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "2208", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Leonard", "PRIMARY_NAME_FIRST": "Leslie", "ADDR_LINE1": "4362 Belmont Lane", "ADDR_CITY": "Elmwood Park", "ADDR_STATE": "CA", "ADDR_POSTAL_CODE": "95865", "EMAIL_ADDRESS": "info@ca-state.gov"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"2209": {
		DataSource: "CUSTOMERS",
		Id:         "2209",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "2209", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Miller", "PRIMARY_NAME_FIRST": "Millie", "ADDR_LINE1": "1376 BlueBell Road", "ADDR_CITY": "Sacramento", "ADDR_STATE": "CA", "ADDR_POSTAL_CODE": "95823", "EMAIL_ADDRESS": "info@ca-state.gov"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"2210": {
		DataSource: "CUSTOMERS",
		Id:         "2210",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "2210", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Nice", "PRIMARY_NAME_FIRST": "Nelson", "ADDR_LINE1": "319 Cody Road", "ADDR_CITY": "Elmwood Park", "ADDR_STATE": "CA", "ADDR_POSTAL_CODE": "95865", "EMAIL_ADDRESS": "info@ca-state.gov"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"2211": {
		DataSource: "CUSTOMERS",
		Id:         "2211",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "2211", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Otter", "PRIMARY_NAME_FIRST": "Otto", "ADDR_LINE1": "1824 AspenOak Way", "ADDR_CITY": "Elmwood Park", "ADDR_STATE": "CA", "ADDR_POSTAL_CODE": "95865", "EMAIL_ADDRESS": "info@ca-state.gov"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"2212": {
		DataSource: "CUSTOMERS",
		Id:         "2212",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "2212", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Pemberton", "PRIMARY_NAME_FIRST": "Penny", "ADDR_LINE1": "1824 AspenOak Way", "ADDR_CITY": "Elmwood Park", "ADDR_STATE": "CA", "ADDR_POSTAL_CODE": "95823", "EMAIL_ADDRESS": "info@ca-state.gov"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"2213": {
		DataSource: "CUSTOMERS",
		Id:         "2213",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "2213", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Kellar", "PRIMARY_NAME_FIRST": "Candace", "ADDR_LINE1": "1824 AspenOak Way", "ADDR_CITY": "Elmwood Park", "ADDR_STATE": "CA", "ADDR_POSTAL_CODE": "95865", "EMAIL_ADDRESS": "info@ca-state.gov"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
	},
	"2214": {
		DataSource: "CUSTOMERS",
		Id:         "2214",
		Data:       `{"DATA_SOURCE": "CUSTOMERS", "RECORD_ID": "2214", "RECORD_TYPE": "PERSON", "PRIMARY_NAME_LAST": "Sanders", "PRIMARY_NAME_FIRST": "Sandy", "ADDR_LINE1": "1376 BlueBell Rd", "ADDR_CITY": "Sacramento", "ADDR_STATE": "CA", "ADDR_POSTAL_CODE": "95823", "EMAIL_ADDRESS": "info@ca-state.gov"}`,
		LoadId:     "TRUTHSET_CUSTOMER_LOAD",
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
