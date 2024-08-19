/*
Package settingsparser is used to pull information from the Senzing engine settings string.

Single-database example of a _ENGINE_CONFIGURATION_JSON string:

	{
		"PIPELINE": {
			"CONFIGPATH": "/etc/opt/senzing",
			"LICENSESTRINGBASE64": "${SENZING_LICENSE_BASE64_ENCODED}"
			"RESOURCEPATH": "/opt/senzing/er/resources",
			"SUPPORTPATH": "/opt/senzing/data",
		},
		"SQL": {
			"CONNECTION": "postgresql://username:password@db.example.com:5432:G2"
		}
	}

Multi-database example of a _ENGINE_CONFIGURATION_JSON string:

	{
		"PIPELINE": {
			"CONFIGPATH": "/etc/opt/senzing",
			"LICENSESTRINGBASE64": "${SENZING_LICENSE_BASE64_ENCODED}",
			"RESOURCEPATH": "/opt/senzing/er/resources",
			"SUPPORTPATH": "/opt/senzing/data"
		},
		"SQL": {
			"BACKEND": "HYBRID",
			"CONNECTION": "postgresql://username:password@db-1.example.com:5432:G2"
		},
		"C1": {
			"CLUSTER_SIZE": "1",
			"DB_1": "postgresql://username:password@db-2.example.com:5432:G2"
		},
		"C2": {
			"CLUSTER_SIZE": "1",
			"DB_1": "postgresql://username:password@db-3.example.com:5432:G2"
		},
		"HYBRID": {
			"RES_FEAT": "C1",
			"RES_FEAT_EKEY": "C1",
			"RES_FEAT_LKEY": "C1",
			"RES_FEAT_STAT": "C1",
			"LIB_FEAT": "C2",
			"LIB_FEAT_HKEY": "C2"
		}
	}

If SQL.BACKEND is omitted or is equal to "SQL", then there is only one database specified by SQL.CONNECTION.
For multiple databases, SQL.BACKEND specifies the JSON key for finding the JSON stanza that maps RES_FEAT_* and LIB_FEAT_* variables to databases.
Note that in the example, HYBRID, C1, and C2 variable names are user-specified.
A user could specify alternative names such as MULTI, DB1, and DB2.

# Documentation

- https://senzing.zendesk.com/hc/en-us/articles/360038774134-G2Engine-Configuration-and-the-Senzing-API
- https://senzing.zendesk.com/hc/en-us/articles/360010599254-Scaling-Out-Your-Database-With-Clustering
*/
package settingsparser
