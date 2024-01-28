package main_usecases_beans_factories

import (
	main_gateways "baseapplicationgo/main/gateways"
	main_gateways_logs "baseapplicationgo/main/gateways/logs"
	main_gateways_mongodb "baseapplicationgo/main/gateways/mongodb"
	main_gateways_mongodb_repositories "baseapplicationgo/main/gateways/mongodb/repositories"
	main_gateways_spans "baseapplicationgo/main/gateways/spans"
	main_usecases "baseapplicationgo/main/usecases"
	main_usecases_factories "baseapplicationgo/main/usecases/factories"
	main_usecases_interfaces "baseapplicationgo/main/usecases/interfaces"
)

type CreateEmailBodySendAndPersistAsSentBeanFactory struct {
}

func NewCreateEmailBodySendAndPersistAsSentBeanFactory() *CreateEmailBodySendAndPersistAsSentBeanFactory {
	return &CreateEmailBodySendAndPersistAsSentBeanFactory{}
}

func (this *CreateEmailBodySendAndPersistAsSentBeanFactory) Get() *main_usecases.CreateEmailBodySendAndPersistAsSent {

	logsMonitoringGateway := main_gateways_logs.NewLogsMonitoringGatewayImpl()
	spanGateway := main_gateways_spans.NewSpanGatewayImpl()

	var createWelcomeEmailTemplateBody main_usecases_interfaces.CreateEmailBody = NewCreateEmailBodyForWelcomeEmailBean().Get()
	createEmailBodyFactory := main_usecases_factories.NewCreateEmailBodyFactoryAllArgs(
		createWelcomeEmailTemplateBody, logsMonitoringGateway, spanGateway)

	emailRepo := main_gateways_mongodb_repositories.NewEmailRepository()
	var emailDatabaseGateway main_gateways.EmailDatabaseGateway = main_gateways_mongodb.NewEmailDatabaseGatewayImpl(*emailRepo)

	sendEmailGatewayFactory := NewSendEmailGatewayFactoryBean().Get()

	createEmail := main_usecases.NewCreateEmailBodySendAndPersistAsSentAllArgs(
		emailDatabaseGateway,
		sendEmailGatewayFactory,
		createEmailBodyFactory,
		logsMonitoringGateway,
		spanGateway)

	return createEmail
}
