package main_usecases_beans_factories

import (
	main_gateways "baseapplicationgo/main/gateways"
	main_gateways_lock "baseapplicationgo/main/gateways/lock"
	main_gateways_logs "baseapplicationgo/main/gateways/logs"
	main_gateways_mongodb "baseapplicationgo/main/gateways/mongodb"
	main_gateways_mongodb_repositories "baseapplicationgo/main/gateways/mongodb/repositories"
	main_gateways_spans "baseapplicationgo/main/gateways/spans"
	main_usecases "baseapplicationgo/main/usecases"
	main_utils_messages "baseapplicationgo/main/utils/messages"
)

type EnableInternalProviderUser struct {
}

func NewEnableInternalProviderUserBean() *EnableInternalProviderUser {
	return &EnableInternalProviderUser{}
}

func (this *EnableInternalProviderUser) Get() *main_usecases.EnableInternalProviderUser {

	validateUserIsInternalAuthProvider := NewValidateUserIsInternalAuthProviderBean().Get()

	userEmailVerificationRepo := *main_gateways_mongodb_repositories.NewUserEmailVerificationRepository()
	var userEmailVerDatabaseGateway main_gateways.UserEmailVerificationDatabaseGateway = main_gateways_mongodb.NewUserEmailVerificationDatabaseGatewayImpl(
		userEmailVerificationRepo,
	)
	var distributedLockGateway main_gateways.DistributedLockGateway = main_gateways_lock.NewDistributedLockGatewayImpl()
	var logsMonitoringGateway main_gateways.LogsMonitoringGateway = main_gateways_logs.NewLogsMonitoringGatewayImpl()
	var spanGatewayImpl main_gateways.SpanGateway = main_gateways_spans.NewSpanGatewayImpl()

	return main_usecases.NewEnableInternalProviderUser(
		validateUserIsInternalAuthProvider,
		NewFindValidUserEmailVerificationBean().Get(),
		userEmailVerDatabaseGateway,
		NewSetUserToEnabledAndEmailVerifiedBean().Get(),
		distributedLockGateway,
		logsMonitoringGateway,
		spanGatewayImpl,
		*main_utils_messages.NewApplicationMessages())
}
