package main_configs_apm_logs_resources

type LogProfile string

const (
	DEBUG LogProfile = "DEBUG"
	WARN  LogProfile = "WARN"
	INFO  LogProfile = "INFO"
	ERROR LogProfile = "ERROR"
)

var logProfileEnum = map[LogProfile]LogProfile{
	DEBUG: DEBUG,
	WARN:  WARN,
	INFO:  INFO,
	ERROR: ERROR,
}

var logProfileEnumFromNames = map[string]LogProfile{
	"DEBUG": DEBUG,
	"WARN":  WARN,
	"INFO":  INFO,
	"ERROR": ERROR,
}

var logLevelCompareEnum = map[LogProfile]map[LogProfile]LogProfile{
	DEBUG: debugLogProfileEnum,
	WARN:  warnLogProfileEnum,
	INFO:  infoLogProfileEnum,
	ERROR: errorLogProfileEnum,
}

var debugLogProfileEnum = map[LogProfile]LogProfile{
	DEBUG: DEBUG,
	WARN:  WARN,
	INFO:  INFO,
	ERROR: ERROR,
}

var warnLogProfileEnum = map[LogProfile]LogProfile{
	WARN:  WARN,
	ERROR: ERROR,
}

var infoLogProfileEnum = map[LogProfile]LogProfile{
	WARN:  WARN,
	INFO:  INFO,
	ERROR: ERROR,
}

var errorLogProfileEnum = map[LogProfile]LogProfile{
	ERROR: ERROR,
}

func (this LogProfile) Exists() bool {
	_, exists := logProfileEnum[this]
	return exists
}

func (this LogProfile) FromValue(value string) LogProfile {
	valueMap, exists := logProfileEnumFromNames[value]
	if exists {
		return valueMap
	}
	return ""
}

func (this LogProfile) Name() string {
	valueMap, exists := logProfileEnum[this]
	if exists {
		return string(valueMap)
	}
	return ""
}

func (this LogProfile) IsHigher(value LogProfile) bool {
	valueMap, exists := logLevelCompareEnum[value]
	if exists {
		_, existsLevel := valueMap[this]
		return existsLevel
	}
	return false
}
