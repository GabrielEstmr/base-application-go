package main_usecases_beans_factories

import (
	main_gateways "baseapplicationgo/main/gateways"
	main_gateways_logs "baseapplicationgo/main/gateways/logs"
	main_gateways_mongodb "baseapplicationgo/main/gateways/mongodb"
	main_gateways_mongodb_repositories "baseapplicationgo/main/gateways/mongodb/repositories"
	main_gateways_spans "baseapplicationgo/main/gateways/spans"
	main_usecases "baseapplicationgo/main/usecases"
	main_utils_messages "baseapplicationgo/main/utils/messages"
)

type FindValidUserEmailVerificationBean struct {
}

func NewFindValidUserEmailVerificationBean() *FindValidUserEmailVerificationBean {
	return &FindValidUserEmailVerificationBean{}
}

func (this *FindValidUserEmailVerificationBean) Get() *main_usecases.FindValidUserEmailVerification {

	userEmailVerificationRepo := *main_gateways_mongodb_repositories.NewUserEmailVerificationRepository()
	var userEmailVerDatabaseGateway main_gateways.UserEmailVerificationDatabaseGateway = main_gateways_mongodb.NewUserEmailVerificationDatabaseGatewayImpl(
		userEmailVerificationRepo,
	)
	var logsMonitoringGateway main_gateways.LogsMonitoringGateway = main_gateways_logs.NewLogsMonitoringGatewayImpl()
	var spanGatewayImpl main_gateways.SpanGateway = main_gateways_spans.NewSpanGatewayImpl()

	return main_usecases.NewFindValidUserEmailVerification(
		userEmailVerDatabaseGateway,
		logsMonitoringGateway,
		spanGatewayImpl,
		*main_utils_messages.NewApplicationMessages())
}
