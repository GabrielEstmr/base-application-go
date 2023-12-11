package main_configs_apm_logs_resources

import "net/http"

type LogProvider struct {
	client *http.Client
}

func NewLogProvider(client *http.Client) *LogProvider {
	return &LogProvider{client: client}
}

func (this *LogProvider) Client() *http.Client {
	return this.client
}
