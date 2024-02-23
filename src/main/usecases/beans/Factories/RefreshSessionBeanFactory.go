package main_usecases_beans_factories

import (
	main_gateways "baseapplicationgo/main/gateways"
	"baseapplicationgo/main/gateways/http/authprovider"
	"baseapplicationgo/main/gateways/http/authprovider/clients"
	main_gateways_logs "baseapplicationgo/main/gateways/logs"
	main_gateways_spans "baseapplicationgo/main/gateways/spans"
	main_usecases "baseapplicationgo/main/usecases"
	main_utils_messages "baseapplicationgo/main/utils/messages"
)

type RefreshSessionBean struct {
}

func NewRefreshSessionBean() *RefreshSessionBean {
	return &RefreshSessionBean{}
}

func (this *RefreshSessionBean) Get() *main_usecases.RefreshSession {

	var authProviderGateway main_gateways.AuthProviderGateway = main_gateways_http_authprovider.NewKeycloakAuthProviderGatewayImpl(
		*main_gateways_http_authprovider_clients.NewKeycloakClient(),
		main_gateways_logs.NewLogsMonitoringGatewayImpl(),
		main_gateways_spans.NewSpanGatewayImpl(),
		*main_utils_messages.NewApplicationMessages(),
	)

	return main_usecases.NewRefreshSession(
		authProviderGateway,
		main_gateways_logs.NewLogsMonitoringGatewayImpl(),
		main_gateways_spans.NewSpanGatewayImpl(),
		*main_utils_messages.NewApplicationMessages(),
	)
}
