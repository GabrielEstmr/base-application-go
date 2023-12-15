package main_configs_ff

import main_configs_ff_lib_resources "baseapplicationgo/main/configs/ff/lib/resources"

const (
	ENABLE_FIND_BY_ID_ENDPOINT  = "ENABLE_FIND_BY_ID_ENDPOINT"
	RABBITMQ_TES_LISTENER_RETRY = "RABBITMQ_TES_LISTENER_RETRY"
)

const (
	GROUP_ID_ENDPOINTS_MANAGEMENT    = "endpoints-management"
	GROUP_ID_RABBITMQ_LISTENER_RETRY = "rabbitmq-listener-retry"
)

var FEATURES = map[string]main_configs_ff_lib_resources.FeaturesData{
	ENABLE_FIND_BY_ID_ENDPOINT: *main_configs_ff_lib_resources.NewFeaturesData(
		ENABLE_FIND_BY_ID_ENDPOINT,
		GROUP_ID_ENDPOINTS_MANAGEMENT,
		"feature to testing ff lib",
		false),
	RABBITMQ_TES_LISTENER_RETRY: *main_configs_ff_lib_resources.NewFeaturesData(
		RABBITMQ_TES_LISTENER_RETRY,
		GROUP_ID_RABBITMQ_LISTENER_RETRY,
		"feature to test rabbitmq behaviour",
		false),
}
