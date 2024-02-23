package main_usecases_beans_factories

import (
	main_gateways "baseapplicationgo/main/gateways"
	"baseapplicationgo/main/gateways/http/authprovider"
	"baseapplicationgo/main/gateways/http/authprovider/clients"
	main_gateways_lock "baseapplicationgo/main/gateways/lock"
	main_gateways_logs "baseapplicationgo/main/gateways/logs"
	main_gateways_spans "baseapplicationgo/main/gateways/spans"
	main_usecases "baseapplicationgo/main/usecases"
	main_usecases_lockers "baseapplicationgo/main/usecases/lockers"
	main_utils_messages "baseapplicationgo/main/utils/messages"
)

type AtomicLockedCreateExternalAuthProviderUserSessionBeanFactory struct {
}

func NewAtomicLockedCreateExternalAuthProviderUserSessionBeanFactory() *AtomicLockedCreateExternalAuthProviderUserSessionBeanFactory {
	return &AtomicLockedCreateExternalAuthProviderUserSessionBeanFactory{}
}

func (this *AtomicLockedCreateExternalAuthProviderUserSessionBeanFactory) Get() *main_usecases_lockers.AtomicLockedCreateExternalAuthProviderUserSession {

	var authProviderGateway main_gateways.AuthProviderGateway = main_gateways_http_authprovider.NewKeycloakAuthProviderGatewayImpl(
		*main_gateways_http_authprovider_clients.NewKeycloakClient(),
		main_gateways_logs.NewLogsMonitoringGatewayImpl(),
		main_gateways_spans.NewSpanGatewayImpl(),
		*main_utils_messages.NewApplicationMessages(),
	)

	buildTokenClaim := main_usecases.NewBuildTokenClaim(
		main_gateways_logs.NewLogsMonitoringGatewayImpl(),
		main_gateways_spans.NewSpanGatewayImpl(),
		*main_utils_messages.NewApplicationMessages(),
	)

	var distributedLockGateway main_gateways.DistributedLockGateway = main_gateways_lock.NewDistributedLockGatewayImpl()

	return main_usecases_lockers.NewAtomicLockedCreateExternalAuthProviderUserSession(
		NewCreateSessionByIdentityProviderBean().Get(),
		authProviderGateway,
		buildTokenClaim,
		distributedLockGateway,
		main_gateways_logs.NewLogsMonitoringGatewayImpl(),
		main_gateways_spans.NewSpanGatewayImpl(),
		*main_utils_messages.NewApplicationMessages(),
	)
}
