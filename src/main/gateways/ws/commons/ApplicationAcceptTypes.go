package main_gateways_ws_commons

const (
	ACCEPT_TYPE_JSON = "application/json"
	ACCEPT_TYPE_ALL  = "*/*"
)

func GetAllAvailableAcceptTypes() map[string]string {
	result := make(map[string]string)
	result[ACCEPT_TYPE_JSON] = ACCEPT_TYPE_JSON
	result[ACCEPT_TYPE_ALL] = ACCEPT_TYPE_ALL
	return result
}
