package main_configs_authprovider

import (
	main_configs_authprovider_resource "baseapplicationgo/main/configs/authprovider/resource"
	main_error "baseapplicationgo/main/configs/error"
	main_configs_yml "baseapplicationgo/main/configs/yml"
	"log"
	"strconv"
	"sync"
)

const _MSG_INIT_AUTH_EXPORTER = "Initializing Auth exporter"
const _MSG_FINAL_AUTH_PROVIDER = "Auth provider has been instantiated"
const _MSG_ERROR_AUTH_EXPORTER_TIMEOUT_CONFIG = "Error to instantiate Auth exporter."

var onceLogs sync.Once
var authProviderBean *main_configs_authprovider_resource.AuthProviderConfig

func GetAuthProviderBean() *main_configs_authprovider_resource.AuthProviderConfig {
	onceLogs.Do(func() {
		if authProviderBean == nil {
			authProviderBean = getAuthProvider()
		}
	})
	return authProviderBean
}

func getAuthProvider() *main_configs_authprovider_resource.AuthProviderConfig {
	log.Println(_MSG_INIT_AUTH_EXPORTER)

	authProviderTimeout, err := strconv.ParseInt(
		main_configs_yml.GetYmlValueByName(main_configs_yml.IntegrationAuthProviderTimeoutInMilliseconds), 10, 64)
	main_error.FailOnError(err, _MSG_ERROR_AUTH_EXPORTER_TIMEOUT_CONFIG)

	authProviderURL := main_configs_yml.GetYmlValueByName(main_configs_yml.IntegrationAuthProviderUrl)

	clientId := main_configs_yml.GetYmlValueByName(main_configs_yml.IntegrationAuthProviderClientId)
	clientSecret := main_configs_yml.GetYmlValueByName(main_configs_yml.IntegrationAuthProviderClientSecret)
	clientRealm := main_configs_yml.GetYmlValueByName(main_configs_yml.IntegrationAuthProviderRealm)
	tokenExchangeClientId := main_configs_yml.GetYmlValueByName(main_configs_yml.IntegrationAuthProviderClientIdTokenExchange)

	log.Println(_MSG_FINAL_AUTH_PROVIDER)
	return main_configs_authprovider_resource.NewAuthProviderConfig(
		authProviderURL,
		authProviderTimeout,
		clientId,
		clientSecret,
		clientRealm,
		tokenExchangeClientId)
}
