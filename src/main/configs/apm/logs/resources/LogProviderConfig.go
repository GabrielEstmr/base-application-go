package main_configs_apm_logs_resources

import (
	"net/http"
	"time"
)

type LogProviderConfig struct {
	appName             string
	host                string
	logProfile          LogProfile
	timeoutMilliseconds int
	client              http.Client
	baseUrl             string
}

func NewLogProviderConfig(
	appName string,
	host string,
	logProfile LogProfile,
	timeoutMilliseconds int) *LogProviderConfig {
	return &LogProviderConfig{
		appName:             appName,
		host:                host,
		logProfile:          logProfile,
		timeoutMilliseconds: timeoutMilliseconds,
		client: http.Client{
			Timeout: time.Duration(timeoutMilliseconds) * time.Millisecond,
		},
		baseUrl: "http://" + host + "/loki/api/v1/push",
	}
}

func (this LogProviderConfig) GetAppName() string {
	return this.appName
}

func (this LogProviderConfig) GetHost() string {
	return this.host
}

func (this LogProviderConfig) GetLogProfile() LogProfile {
	return this.logProfile
}

func (this LogProviderConfig) GetTimeoutMilliseconds() int {
	return this.timeoutMilliseconds
}

func (this LogProviderConfig) GetClient() *http.Client {
	return &this.client
}

func (this LogProviderConfig) GetBaseUrl() string {
	return this.baseUrl
}
