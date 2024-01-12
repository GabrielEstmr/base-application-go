package main_gateways_ws_commons

const (
	PT_BR = "pt-BR"
	EN_US = "en-US"
)

func GetAllAvailableLanguages() map[string]string {
	result := make(map[string]string)
	result[PT_BR] = PT_BR
	result[EN_US] = EN_US
	return result
}
