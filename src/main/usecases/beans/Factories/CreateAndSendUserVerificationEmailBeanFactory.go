package main_usecases_beans_factories

import (
	main_gateways "baseapplicationgo/main/gateways"
	main_gateways_lock "baseapplicationgo/main/gateways/lock"
	main_gateways_logs "baseapplicationgo/main/gateways/logs"
	main_gateways_mongodb "baseapplicationgo/main/gateways/mongodb"
	main_gateways_mongodb_repositories "baseapplicationgo/main/gateways/mongodb/repositories"
	main_gateways_rabbitmq "baseapplicationgo/main/gateways/rabbitmq"
	main_gateways_rabbitmq_producers "baseapplicationgo/main/gateways/rabbitmq/producers"
	main_gateways_spans "baseapplicationgo/main/gateways/spans"
	main_usecases "baseapplicationgo/main/usecases"
	main_utils_messages "baseapplicationgo/main/utils/messages"
)

type CreateAndSendUserVerificationEmailBean struct {
}

func NewCreateAndSendUserVerificationEmailBean() *CreateAndSendUserVerificationEmailBean {
	return &CreateAndSendUserVerificationEmailBean{}
}

func (this *CreateAndSendUserVerificationEmailBean) Get() *main_usecases.CreateAndSendUserVerificationEmail {

	userEmailVerificationRepo := *main_gateways_mongodb_repositories.NewUserEmailVerificationRepository()
	var userEmailVerDatabaseGateway main_gateways.UserEmailVerificationDatabaseGateway = main_gateways_mongodb.NewUserEmailVerificationDatabaseGatewayImpl(
		userEmailVerificationRepo,
	)
	var emailInternalProviderGateway main_gateways.EmailInternalProviderGateway = main_gateways_rabbitmq.NewEmailInternalProviderGateway(
		*main_gateways_rabbitmq_producers.NewEmailEventProducer(),
	)
	var distributedLockGateway main_gateways.DistributedLockGateway = main_gateways_lock.NewDistributedLockGatewayImpl()
	var logsMonitoringGateway main_gateways.LogsMonitoringGateway = main_gateways_logs.NewLogsMonitoringGatewayImpl()
	var spanGatewayImpl main_gateways.SpanGateway = main_gateways_spans.NewSpanGatewayImpl()

	return main_usecases.NewCreateAndSendUserVerificationEmail(
		userEmailVerDatabaseGateway,
		emailInternalProviderGateway,
		NewCreateUserEmailVerificationBean().Get(),
		distributedLockGateway,
		logsMonitoringGateway,
		spanGatewayImpl,
		*main_utils_messages.NewApplicationMessages())
}
