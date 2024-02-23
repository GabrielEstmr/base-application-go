package main_usecases_beans_factories

import (
	main_gateways "baseapplicationgo/main/gateways"
	main_gateways_http_authprovider "baseapplicationgo/main/gateways/http/authprovider"
	main_gateways_http_authprovider_clients "baseapplicationgo/main/gateways/http/authprovider/clients"
	main_gateways_logs "baseapplicationgo/main/gateways/logs"
	main_gateways_spans "baseapplicationgo/main/gateways/spans"
	main_usecases "baseapplicationgo/main/usecases"
	main_utils_messages "baseapplicationgo/main/utils/messages"
)

type CreateOrRetrieveUserFromAuthProviderBean struct {
}

func NewCreateOrRetrieveUserFromAuthProviderBean() *CreateOrRetrieveUserFromAuthProviderBean {
	return &CreateOrRetrieveUserFromAuthProviderBean{}
}

func (this *CreateOrRetrieveUserFromAuthProviderBean) Get() *main_usecases.CreateOrRetrieveUserFromAuthProvider {

	var logsMonitoringGateway main_gateways.LogsMonitoringGateway = main_gateways_logs.NewLogsMonitoringGatewayImpl()
	var spanGatewayImpl main_gateways.SpanGateway = main_gateways_spans.NewSpanGatewayImpl()

	var authProviderGateway main_gateways.AuthProviderGateway = main_gateways_http_authprovider.NewKeycloakAuthProviderGatewayImpl(
		*main_gateways_http_authprovider_clients.NewKeycloakClient(),
		main_gateways_logs.NewLogsMonitoringGatewayImpl(),
		main_gateways_spans.NewSpanGatewayImpl(),
		*main_utils_messages.NewApplicationMessages(),
	)

	return main_usecases.NewCreateOrRetrieveUserFromAuthProvider(
		authProviderGateway,
		logsMonitoringGateway,
		spanGatewayImpl,
		*main_utils_messages.NewApplicationMessages())
}
