package main_configs_apm_logs

import (
	main_configs_apm_logs_resources "baseapplicationgo/main/configs/apm/logs/resources"
	main_error "baseapplicationgo/main/configs/error"
	main_configs_yml "baseapplicationgo/main/configs/yml"
	"log"
	"strconv"
	"sync"
)

const _MSG_INIT_LOG_EXPORTER = "Initializing logs exporter"
const _MSG_FINAL_LOG_PROVIDER = "Log provider has been instantiated"
const _MSG_ERROR_LOG_EXPORTER_TIMEOUT_CONFIG = "Error to instantiate logs exporter. Timeout config is not a number"

const _METRIC_APM_CLUSTER_LOKI_HOST_YML = "Apm.server.loki.collector.http.host"
const _METRIC_APM_CLUSTER_LOKI_TIMEOUT_YML = "Apm.server.loki.collector.http.timeout.milliseconds"
const _APP_NAME_YML = "Apm.server.name"

var onceLogs sync.Once
var logProviderBean *main_configs_apm_logs_resources.LogProviderConfig

func GetLogProviderBean() *main_configs_apm_logs_resources.LogProviderConfig {
	onceLogs.Do(func() {
		if logProviderBean == nil {
			logProviderBean = getLogProvider()
		}
	})
	return logProviderBean
}

func getLogProvider() *main_configs_apm_logs_resources.LogProviderConfig {
	log.Println(_MSG_INIT_LOG_EXPORTER)

	timeout, err := strconv.Atoi(main_configs_yml.GetYmlValueByName(_METRIC_APM_CLUSTER_LOKI_TIMEOUT_YML))
	main_error.FailOnError(err, _MSG_ERROR_LOG_EXPORTER_TIMEOUT_CONFIG)

	log.Println(_MSG_FINAL_LOG_PROVIDER)
	return main_configs_apm_logs_resources.NewLogProviderConfig(
		main_configs_yml.GetYmlValueByName(_APP_NAME_YML),
		main_configs_yml.GetYmlValueByName(_METRIC_APM_CLUSTER_LOKI_HOST_YML),
		timeout,
	)
}
