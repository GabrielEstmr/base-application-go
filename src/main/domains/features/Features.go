package main_domains_features

type Features map[string]FeaturesData

const (
	ENABLE_FIND_BY_ID_ENDPOINT  = "ENABLE_FIND_BY_ID_ENDPOINT"
	RABBITMQ_TES_LISTENER_RETRY = "RABBITMQ_TES_LISTENER_RETRY"
)

const (
	GROUP_ID_ENDPOINTS_MANAGEMENT    = "endpoints-management"
	GROUP_ID_RABBITMQ_LISTENER_RETRY = "rabbitmq-listener-retry"
)

var FEATURES = Features{
	ENABLE_FIND_BY_ID_ENDPOINT: *NewFeaturesDataAllArgs(
		ENABLE_FIND_BY_ID_ENDPOINT,
		GROUP_ID_ENDPOINTS_MANAGEMENT,
		"feature to testing ff lib",
		false),
	RABBITMQ_TES_LISTENER_RETRY: *NewFeaturesDataAllArgs(
		RABBITMQ_TES_LISTENER_RETRY,
		GROUP_ID_RABBITMQ_LISTENER_RETRY,
		"feature to middlewares rabbitmq behaviour",
		false),
}
