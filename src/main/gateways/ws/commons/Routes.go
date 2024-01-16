package main_gateways_ws_commons

const API_V1_PREFIX = "/api/v1"

const (
	API_V1_USERS                 = API_V1_PREFIX + "/users"
	API_V1_USERS_ID              = API_V1_PREFIX + "/users/{id}"
	API_V1_FEATURES_KEY_ENABLED  = API_V1_PREFIX + "/features/{key}/enable"
	API_V1_FEATURES_KEY_DISABLED = API_V1_PREFIX + "/features/{key}/disabled"
	API_V1_RABBITMQ_SEND_EVENT   = API_V1_PREFIX + "/rabbitmq/send-event"
	API_V1_TRANSACTIONS          = API_V1_PREFIX + "/transactions"
)

var RoutesURI = map[string]string{
	API_V1_USERS:                 API_V1_USERS,
	API_V1_USERS_ID:              API_V1_USERS_ID,
	API_V1_FEATURES_KEY_ENABLED:  API_V1_FEATURES_KEY_ENABLED,
	API_V1_FEATURES_KEY_DISABLED: API_V1_FEATURES_KEY_DISABLED,
	API_V1_RABBITMQ_SEND_EVENT:   API_V1_RABBITMQ_SEND_EVENT,
	API_V1_TRANSACTIONS:          API_V1_TRANSACTIONS,
}
