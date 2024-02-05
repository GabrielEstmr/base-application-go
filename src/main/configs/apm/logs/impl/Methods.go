package main_configs_apm_logs_impl

type LogLevelMethod string

const (
	METHOD_DEBUG = "DEBUG"
	METHOD_WARN  = "WARN"
	METHOD_INFO  = "INFO"
	METHOD_ERROR = "ERROR"
)

var logProfileEnum = map[LogLevelMethod]LogLevelMethod{
	METHOD_DEBUG: METHOD_DEBUG,
	METHOD_WARN:  METHOD_WARN,
	METHOD_INFO:  METHOD_INFO,
	METHOD_ERROR: METHOD_ERROR,
}

func GetLogProfileDescription(value LogLevelMethod) string {
	valueMap, exists := logProfileEnum[value]
	if exists {
		return string(valueMap)
	}
	return ""
}

func LogProfileFromDescription(description string) LogLevelMethod {
	valueMap, exists := logProfileEnum[LogLevelMethod(description)]
	if exists {
		return valueMap
	}
	return ""
}
