package main_usecases_beans_factories

import (
	main_gateways "baseapplicationgo/main/gateways"
	main_gateways_lock "baseapplicationgo/main/gateways/lock"
	main_gateways_logs "baseapplicationgo/main/gateways/logs"
	main_gateways_mongodb "baseapplicationgo/main/gateways/mongodb"
	main_gateways_mongodb_repositories "baseapplicationgo/main/gateways/mongodb/repositories"
	main_gateways_spans "baseapplicationgo/main/gateways/spans"
	main_usecases "baseapplicationgo/main/usecases"
	main_usecases_interfaces "baseapplicationgo/main/usecases/interfaces"
	main_utils_messages "baseapplicationgo/main/utils/messages"
)

type CreateEmailBean struct {
}

func NewCreateEmailBean() *CreateEmailBean {
	return &CreateEmailBean{}
}

func (this *CreateEmailBean) Get() *main_usecases.CreateEmail {

	emailRepo := main_gateways_mongodb_repositories.NewEmailRepository()
	var emailDatabaseGateway main_gateways.EmailDatabaseGateway = main_gateways_mongodb.NewEmailDatabaseGatewayImpl(*emailRepo)

	var sendEmailGatewayFactory main_usecases_interfaces.CreateEmailBodySendAndPersistAsSent = NewCreateEmailBodySendAndPersistAsSentBeanFactory().Get()

	logsMonitoringGateway := main_gateways_logs.NewLogsMonitoringGatewayImpl()
	spanGateway := main_gateways_spans.NewSpanGatewayImpl()
	messageUtils := *main_utils_messages.NewApplicationMessages()

	var distributedLockGateway main_gateways.DistributedLockGateway = main_gateways_lock.NewDistributedLockGatewayImpl()

	createEmail := main_usecases.NewCreateEmailAllArgs(
		emailDatabaseGateway,
		sendEmailGatewayFactory,
		distributedLockGateway,
		logsMonitoringGateway,
		spanGateway,
		messageUtils,
	)

	return createEmail
}
