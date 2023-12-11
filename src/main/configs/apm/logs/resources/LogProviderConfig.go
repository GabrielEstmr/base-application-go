package main_configs_apm_logs_resources

type LogProviderConfig struct {
	appName             string
	host                string
	timeoutMilliseconds int
}

func NewLogProviderConfig(
	appName string,
	host string,
	timeoutMilliseconds int) *LogProviderConfig {
	return &LogProviderConfig{
		appName:             appName,
		host:                host,
		timeoutMilliseconds: timeoutMilliseconds}
}

func (this LogProviderConfig) GetAppName() string {
	return this.appName
}

func (this LogProviderConfig) GetHost() string {
	return this.host
}

func (this LogProviderConfig) GetTimeoutMilliseconds() int {
	return this.timeoutMilliseconds
}
