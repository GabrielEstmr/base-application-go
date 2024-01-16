package main_gateways_ws_commons

const (
	CONTENT_TYPE_JSON = "application/json"
)

func GetAllAvailableContentTypes() map[string]string {
	result := make(map[string]string)
	result[CONTENT_TYPE_JSON] = CONTENT_TYPE_JSON
	return result
}
