# go-common errors

## Error strings

### CONFIGPATH: Could not find /path/to/file

The `PIPELINE.CONFIGPATH` is incorrectly specified in the Senzing engine configuration JSON.

If it is correctly specified, then a file that is expected is not in the directory.

### RESOURCEPATH: Could not find /path/to/file

The `PIPELINE.RESOURCEPATH` is incorrectly specified in the Senzing engine configuration JSON.

If it is correctly specified, then a file that is expected is not in the directory.

### SQL.CONNECTION empty in Senzing engine configuration JSON

No database was specified in the Senzing engine configuration JSON.

In `senzing-tools`:

1. If `SENZING_TOOLS_ENGINE_CONFIGURATION_JSON` is used,
   make sure the database are correctly specified.
1. If `SENZING_TOOLS_ENGINE_CONFIGURATION_JSON` is **not** used,
   `SENZING_TOOLS_DATABASE_URL` needs to be set.

References:

1. [SENZING_TOOLS_DATABASE_URL](https://github.com/Senzing/knowledge-base/blob/main/lists/environment-variables.md#senzing_tools_database_url)
1. [SENZING_TOOLS_ENGINE_CONFIGURATION_JSON](https://github.com/Senzing/knowledge-base/blob/main/lists/environment-variables.md#senzing_tools_engine_configuration_json)

### SUPPORTPATH: Could not find /path/to/file

The `PIPELINE.SUPPORTPATH` is incorrectly specified in the Senzing engine configuration JSON.

If it is correctly specified, then a file that is expected is not in the directory.
