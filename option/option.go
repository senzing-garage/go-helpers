package option

import (
	"fmt"
	"time"

	"github.com/senzing/go-common/option/optiontype"
)

func (v ContextVariable) SetDefault(newDefault any) ContextVariable {
	v.Default = newDefault
	return v
}

var ConfigPath = ContextVariable{
	Default: OsLookupEnvString("SENZING_TOOLS_CONFIG_PATH", ""),
	Envar:   "SENZING_TOOLS_CONFIG_PATH",
	Help:    "Path to SenzingAPI's configuration directory [%s]",
	Arg:     "config-path",
	Type:    optiontype.String,
}

var Configuration = ContextVariable{
	Default: OsLookupEnvString("SENZING_TOOLS_CONFIGURATION", ""),
	Envar:   "SENZING_TOOLS_CONFIGURATION",
	Help:    "Path to configuration file [%s]",
	Arg:     "configuration",
	Type:    optiontype.String,
}

var DatabaseUrl = ContextVariable{
	Default: OsLookupEnvString("SENZING_TOOLS_DATABASE_URL", ""),
	Envar:   "SENZING_TOOLS_DATABASE_URL",
	Help:    "URL of database to initialize [%s]",
	Arg:     "database-url",
	Type:    optiontype.String,
}

var Datasources = ContextVariable{
	Default: []string{},
	Envar:   "SENZING_TOOLS_DATASOURCES",
	Help:    "Datasources to be added to initial Senzing configuration [%s]",
	Arg:     "datasources",
	Type:    optiontype.StringSlice,
}

var DelayInSeconds = ContextVariable{
	Default: OsLookupEnvInt("SENZING_TOOLS_DELAY_IN_SECONDS", 0),
	Envar:   "SENZING_TOOLS_DELAY_IN_SECONDS",
	Help:    "Number of seconds to wait before starting process [%s]",
	Arg:     "delay-in-seconds",
	Type:    optiontype.Int,
}

var EnableAll = ContextVariable{
	Default: OsLookupEnvBool("SENZING_TOOLS_ENABLE_ALL", false),
	Envar:   "SENZING_TOOLS_ENABLE_ALL",
	Help:    "Enable all services [%s]",
	Arg:     "enable-all",
	Type:    optiontype.Bool,
}

var EnableG2config = ContextVariable{
	Default: OsLookupEnvBool("SENZING_TOOLS_ENABLE_G2CONFIG", false),
	Envar:   "SENZING_TOOLS_ENABLE_G2CONFIG",
	Help:    "Enable G2Config service [%s]",
	Arg:     "enable-g2config",
	Type:    optiontype.Bool,
}

var EnableG2configmgr = ContextVariable{
	Default: OsLookupEnvBool("SENZING_TOOLS_ENABLE_G2CONFIGMGR", false),
	Envar:   "SENZING_TOOLS_ENABLE_G2CONFIGMGR",
	Help:    "Enable G2ConfigMgr service [%s]",
	Arg:     "enable-g2configmgr",
	Type:    optiontype.Bool,
}

var EnableG2diagnostic = ContextVariable{
	Default: OsLookupEnvBool("SENZING_TOOLS_ENABLE_G2DIAGNOSTIC", false),
	Envar:   "SENZING_TOOLS_ENABLE_G2DIAGNOSTIC",
	Help:    "Enable G2Diagnostic service [%s]",
	Arg:     "enable-g2diagnostic",
	Type:    optiontype.Bool,
}

var EnableG2engine = ContextVariable{
	Default: OsLookupEnvBool("SENZING_TOOLS_ENABLE_G2ENGINE", false),
	Envar:   "SENZING_TOOLS_ENABLE_G2ENGINE",
	Help:    "Enable G2Config service [%s]",
	Arg:     "enable-g2engine",
	Type:    optiontype.Bool,
}

var EnableG2product = ContextVariable{
	Default: OsLookupEnvBool("SENZING_TOOLS_ENABLE_G2PRODUCT", false),
	Envar:   "SENZING_TOOLS_ENABLE_G2PRODUCT",
	Help:    "Enable G2Config service [%s]",
	Arg:     "enable-g2product",
	Type:    optiontype.Bool,
}

var EnableSenzingChatApi = ContextVariable{
	Default: OsLookupEnvBool("SENZING_TOOLS_ENABLE_SENZING_CHAT_API", false),
	Envar:   "SENZING_TOOLS_ENABLE_SENZING_CHAT_API",
	Help:    "Enable the Senzing REST Chat service [%s]",
	Arg:     "enable-senzing-chat-api",
	Type:    optiontype.Bool,
}

var EnableSenzingRestApi = ContextVariable{
	Default: OsLookupEnvBool("SENZING_TOOLS_ENABLE_SENZING_REST_API", false),
	Envar:   "SENZING_TOOLS_ENABLE_SENZING_REST_API",
	Help:    "Enable the Senzing REST API service [%s]",
	Arg:     "enable-senzing-rest-api",
	Type:    optiontype.Bool,
}

var EnableSwaggerUi = ContextVariable{
	Default: OsLookupEnvBool("SENZING_TOOLS_ENABLE_SWAGGER_UI", false),
	Envar:   "SENZING_TOOLS_ENABLE_SWAGGER_UI",
	Help:    "Enable the Swagger UI service [%s]",
	Arg:     "enable-swagger-ui",
	Type:    optiontype.Bool,
}

var EnableXterm = ContextVariable{
	Default: OsLookupEnvBool("SENZING_TOOLS_ENABLE_XTERM", false),
	Envar:   "SENZING_TOOLS_ENABLE_XTERM",
	Help:    "Enable the XTerm service [%s]",
	Arg:     "enable-xterm",
	Type:    optiontype.Bool,
}

var EngineConfigurationJson = ContextVariable{
	Default: OsLookupEnvString("SENZING_TOOLS_ENGINE_CONFIGURATION_JSON", ""),
	Envar:   "SENZING_TOOLS_ENGINE_CONFIGURATION_JSON",
	Help:    "JSON string sent to Senzing's init() function [%s]",
	Arg:     "engine-configuration-json",
	Type:    optiontype.String,
}

var EngineLogLevel = ContextVariable{
	Default: OsLookupEnvInt("SENZING_TOOLS_ENGINE_LOG_LEVEL", 0),
	Envar:   "SENZING_TOOLS_ENGINE_LOG_LEVEL",
	Help:    "Log level for Senzing Engine [%s]",
	Arg:     "engine-log-level",
	Type:    optiontype.Int,
}

var EngineModuleName = ContextVariable{
	Default: fmt.Sprintf("senzing-tools-%d", time.Now().Unix()),
	Envar:   "SENZING_TOOLS_ENGINE_MODULE_NAME",
	Help:    "Identifier given to the Senzing engine [%s]",
	Arg:     "engine-module-name",
	Type:    optiontype.String,
}

var GrpcPort = ContextVariable{
	Default: OsLookupEnvInt("SENZING_TOOLS_GRPC_PORT", 8260),
	Envar:   "SENZING_TOOLS_GRPC_PORT",
	Help:    "Port used to serve gRPC [%s]",
	Arg:     "grpc-port",
	Type:    optiontype.Int,
}

var GrpcUrl = ContextVariable{
	Default: OsLookupEnvString("SENZING_TOOLS_GRPC_URL", ""),
	Envar:   "SENZING_TOOLS_GRPC_URL",
	Help:    "URL of Senzing gRPC service [%s]",
	Arg:     "grpc-url",
	Type:    optiontype.String,
}

var HttpPort = ContextVariable{
	Default: OsLookupEnvInt("SENZING_TOOLS_HTTP_PORT", 8261),
	Envar:   "SENZING_TOOLS_HTTP_PORT",
	Help:    "Port to serve HTTP [%s]",
	Arg:     "http-port",
	Type:    optiontype.Int,
}

var InputFileType = ContextVariable{
	Default: OsLookupEnvString("SENZING_TOOLS_INPUT_FILE_TYPE", ""),
	Envar:   "SENZING_TOOLS_INPUT_FILE_TYPE",
	Help:    "Input file type to override auto-detect based on file name [%s]",
	Arg:     "input-file-type",
	Type:    optiontype.String,
}

var InputURL = ContextVariable{
	Default: OsLookupEnvString("SENZING_TOOLS_INPUT_URL", ""),
	Envar:   "SENZING_TOOLS_INPUT_URL",
	Help:    "Input URL used for processing [%s]",
	Arg:     "input-url",
	Type:    optiontype.String,
}

var LicenseStringBase64 = ContextVariable{
	Default: OsLookupEnvString("SENZING_TOOLS_LICENSE_STRING_BASE64", ""),
	Envar:   "SENZING_TOOLS_LICENSE_STRING_BASE64",
	Help:    "Base64 representation of a Senzing license [%s]",
	Arg:     "license-string-base64",
	Type:    optiontype.String,
}

var LogLevel = ContextVariable{
	Default: OsLookupEnvString("SENZING_TOOLS_LOG_LEVEL", "INFO"),
	Envar:   "SENZING_TOOLS_LOG_LEVEL",
	Help:    "Log level of TRACE, DEBUG, INFO, WARN, ERROR, FATAL, or PANIC [%s]",
	Arg:     "log-level",
	Type:    optiontype.String,
}

var MonitoringPeriodInSeconds = ContextVariable{
	Default: OsLookupEnvInt("SENZING_TOOLS_MONITORING_PERIOD_IN_SECONDS", 60),
	Envar:   "SENZING_TOOLS_MONITORING_PERIOD_IN_SECONDS",
	Help:    "Print monitoring log messages with the period given in seconds [%s]",
	Arg:     "monitoring-period-in-seconds",
	Type:    optiontype.Int,
}

var NumberOfWorkers = ContextVariable{
	Default: OsLookupEnvInt("SENZING_TOOLS_NUMBER_OF_WORKERS", 0),
	Envar:   "SENZING_TOOLS_NUMBER_OF_WORKERS",
	Help:    "Override the default number of worker routines. Default is GOMAXPROCS [%s]",
	Arg:     "number-of-workers",
	Type:    optiontype.Int,
}

var ObserverOrigin = ContextVariable{
	Default: OsLookupEnvString("SENZING_TOOLS_OBSERVER_ORIGIN", ""),
	Envar:   "SENZING_TOOLS_OBSERVER_ORIGIN",
	Help:    "Identify this instance to the Observer [%s]",
	Arg:     "observer-origin",
	Type:    optiontype.String,
}

var ObserverUrl = ContextVariable{
	Default: OsLookupEnvString("SENZING_TOOLS_OBSERVER_URL", ""),
	Envar:   "SENZING_TOOLS_OBSERVER_URL",
	Help:    "URL of Observer [%s]",
	Arg:     "observer-url",
	Type:    optiontype.String,
}

var OutputURL = ContextVariable{
	Default: OsLookupEnvString("SENZING_TOOLS_OUTPUT_URL", ""),
	Envar:   "SENZING_TOOLS_OUTPUT_URL",
	Help:    "Output URL used for processing [%s]",
	Arg:     "output-url",
	Type:    optiontype.String,
}

var RecordMax = ContextVariable{
	Default: OsLookupEnvInt("SENZING_TOOLS_RECORD_MAX", 0),
	Envar:   "SENZING_TOOLS_RECORD_MAX",
	Help:    "Process a maximum number of records equal to this number [%s]",
	Arg:     "record-max",
	Type:    optiontype.Int,
}

var RecordMin = ContextVariable{
	Default: OsLookupEnvInt("SENZING_TOOLS_RECORD_MIN", 0),
	Envar:   "SENZING_TOOLS_RECORD_MIN",
	Help:    "Process records starting at this record number, discarding all before [%s]",
	Arg:     "record-min",
	Type:    optiontype.Int,
}

var RecordMonitor = ContextVariable{
	Default: OsLookupEnvInt("SENZING_TOOLS_RECORD_MONITOR", 100000),
	Envar:   "SENZING_TOOLS_RECORD_MONITOR",
	Help:    "Log a monitor message after this number of records have been processed [%s]",
	Arg:     "record-monitor",
	Type:    optiontype.Int,
}

var ResourcePath = ContextVariable{
	Default: OsLookupEnvString("SENZING_TOOLS_RESOURCE_PATH", ""),
	Envar:   "SENZING_TOOLS_RESOURCE_PATH",
	Help:    "Path to SenzingAPI's config, schema, and templates directory [%s]",
	Arg:     "resource-path",
	Type:    optiontype.String,
}

var SenzingDirectory = ContextVariable{
	Default: OsLookupEnvString("SENZING_TOOLS_SENZING_DIRECTORY", ""),
	Envar:   "SENZING_TOOLS_SENZING_DIRECTORY",
	Help:    "Path to the SenzingAPI installation directory [%s]",
	Arg:     "senzing-directory",
	Type:    optiontype.String,
}

var ServerAddress = ContextVariable{
	Default: OsLookupEnvString("SENZING_TOOLS_SERVER_ADDRESS", "0.0.0.0"),
	Envar:   "SENZING_TOOLS_SERVER_ADDRESS",
	Help:    "IP interface server listens on [%s]",
	Arg:     "server-address",
	Type:    optiontype.String,
}

var SupportPath = ContextVariable{
	Default: OsLookupEnvString("SENZING_TOOLS_SUPPORT_PATH", "8261"),
	Envar:   "SENZING_TOOLS_SUPPORT_PATH",
	Help:    "Path to SenzingAPI's data directory [%s]",
	Arg:     "support-path",
	Type:    optiontype.String,
}

var VisibilityPeriodInSeconds = ContextVariable{
	Default: OsLookupEnvInt("SENZING_TOOLS_VISIBILITY_PERIOD_IN_SECONDS", 60),
	Envar:   "SENZING_TOOLS_VISIBILITY_PERIOD_IN_SECONDS",
	Help:    "Number of seconds a record held for processing.  This is renewed if processing takes longer [%s]",
	Arg:     "visibility-period-in-seconds",
	Type:    optiontype.Int,
}

var XtermAllowedHostnames = ContextVariable{
	Default: []string{"localhost"},
	Envar:   "SENZING_TOOLS_XTERM_ALLOWED_HOSTNAMES",
	Help:    "Comma-delimited list of hostnames permitted to connect to the websocket [%s]",
	Arg:     "xterm-allowed-hostnames",
	Type:    optiontype.StringSlice,
}

var XtermArguments = ContextVariable{
	Default: []string{},
	Envar:   "SENZING_TOOLS_XTERM_ARGUMENTS",
	Help:    "Comma-delimited list of arguments passed to the terminal command prompt [%s]",
	Arg:     "xterm-arguments",
	Type:    optiontype.StringSlice,
}

var XtermCommand = ContextVariable{
	Default: OsLookupEnvString("SENZING_TOOLS_XTERM_COMMAND", "/bin/bash"),
	Envar:   "SENZING_TOOLS_XTERM_COMMAND",
	Help:    "Path of shell command [%s]",
	Arg:     "xterm-command",
	Type:    optiontype.String,
}

var XtermConnectionErrorLimit = ContextVariable{
	Default: OsLookupEnvInt("SENZING_TOOLS_XTERM_CONNECTION_ERROR_LIMIT", 10),
	Envar:   "SENZING_TOOLS_XTERM_CONNECTION_ERROR_LIMIT",
	Help:    "Connection re-attempts before terminating [%s]",
	Arg:     "xterm-connection-error-limit",
	Type:    optiontype.Int,
}

var XtermKeepalivePingTimeout = ContextVariable{
	Default: OsLookupEnvInt("SENZING_TOOLS_XTERM_KEEPALIVE_PING_TIMEOUT", 20),
	Envar:   "SENZING_TOOLS_XTERM_KEEPALIVE_PING_TIMEOUT",
	Help:    "Maximum allowable seconds between a ping message and its response [%s]",
	Arg:     "xterm-keepalive-ping-timeout",
	Type:    optiontype.Int,
}

var XtermMaxBufferSizeBytes = ContextVariable{
	Default: OsLookupEnvInt("SENZING_TOOLS_XTERM_MAX_BUFFER_SIZE_BYTES", 512),
	Envar:   "SENZING_TOOLS_XTERM_MAX_BUFFER_SIZE_BYTES",
	Help:    "Maximum length of terminal input [%s]",
	Arg:     "xterm-max-buffer-size-bytes",
	Type:    optiontype.Int,
}
