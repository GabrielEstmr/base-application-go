package main_usecases_beans_factories

import (
	main_gateways "baseapplicationgo/main/gateways"
	"baseapplicationgo/main/gateways/http/authprovider"
	"baseapplicationgo/main/gateways/http/authprovider/clients"
	main_gateways_logs "baseapplicationgo/main/gateways/logs"
	main_gateways_mongodb "baseapplicationgo/main/gateways/mongodb"
	main_gateways_mongodb_repositories "baseapplicationgo/main/gateways/mongodb/repositories"
	main_gateways_redis "baseapplicationgo/main/gateways/redis"
	main_gateways_redis_repositories "baseapplicationgo/main/gateways/redis/repositories"
	main_gateways_spans "baseapplicationgo/main/gateways/spans"
	main_usecases "baseapplicationgo/main/usecases"
	main_usecases_interfaces "baseapplicationgo/main/usecases/interfaces"
	main_utils_messages "baseapplicationgo/main/utils/messages"
)

type ChangeInternalProviderUserPasswordBean struct {
}

func NewChangeInternalProviderUserPasswordBean() *ChangeInternalProviderUserPasswordBean {
	return &ChangeInternalProviderUserPasswordBean{}
}

func (this *ChangeInternalProviderUserPasswordBean) Get() *main_usecases.ChangeInternalProviderUserPassword {

	userEmailVerificationRepo := *main_gateways_mongodb_repositories.NewUserEmailVerificationRepository()
	var userEmailVerDatabaseGateway main_gateways.UserEmailVerificationDatabaseGateway = main_gateways_mongodb.NewUserEmailVerificationDatabaseGatewayImpl(
		userEmailVerificationRepo,
	)

	userRepository := main_gateways_mongodb_repositories.NewUserRepository()
	var userDatabaseGateway main_gateways.UserDatabaseGateway = main_gateways_mongodb.NewUserDatabaseGatewayImpl(*userRepository)
	userRedisRepository := main_gateways_redis_repositories.NewUserRedisRepository()
	var userDatabaseCacheGateway main_gateways.UserDatabaseCacheGateway = main_gateways_redis.NewUserDatabaseCacheGatewayImpl(*userRedisRepository)
	var cachedUserDatabaseGateway main_gateways.UserDatabaseGateway = main_gateways_mongodb.NewCachedUserDatabaseGatewayImpl(userDatabaseGateway, userDatabaseCacheGateway)

	var logsMonitoringGateway main_gateways.LogsMonitoringGateway = main_gateways_logs.NewLogsMonitoringGatewayImpl()
	var spanGatewayImpl main_gateways.SpanGateway = main_gateways_spans.NewSpanGatewayImpl()

	var authProviderGateway main_gateways.AuthProviderGateway = main_gateways_http_authprovider.NewKeycloakAuthProviderGatewayImpl(
		*main_gateways_http_authprovider_clients.NewKeycloakClient(),
		main_gateways_logs.NewLogsMonitoringGatewayImpl(),
		main_gateways_spans.NewSpanGatewayImpl(),
		*main_utils_messages.NewApplicationMessages(),
	)

	var validateUserIsInternalAuthProvider main_usecases_interfaces.ValidateUserAuthProviderOrigin = NewValidateUserIsInternalAuthProviderBean().Get()

	return main_usecases.NewChangeInternalProviderUserPassword(
		NewValidatePasswordFormatBean().Get(),
		validateUserIsInternalAuthProvider,
		NewFindValidUserEmailVerificationBean().Get(),
		userEmailVerDatabaseGateway,
		authProviderGateway,
		cachedUserDatabaseGateway,
		logsMonitoringGateway,
		spanGatewayImpl,
		*main_utils_messages.NewApplicationMessages())
}