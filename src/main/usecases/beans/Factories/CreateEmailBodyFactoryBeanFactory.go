package main_usecases_beans_factories

import (
	main_gateways_logs "baseapplicationgo/main/gateways/logs"
	main_gateways_spans "baseapplicationgo/main/gateways/spans"
	main_usecases_factories "baseapplicationgo/main/usecases/factories"
	main_usecases_interfaces "baseapplicationgo/main/usecases/interfaces"
)

type CreateEmailBodyFactoryBean struct {
}

func NewCreateEmailBodyFactoryBean() *CreateEmailBodyFactoryBean {
	return &CreateEmailBodyFactoryBean{}
}

func (this *CreateEmailBodyFactoryBean) Get() *main_usecases_factories.CreateEmailBodyFactory {
	var createWelcomeEmailTemplateBody main_usecases_interfaces.CreateEmailBody = NewCreateEmailBodyForWelcomeEmailBean().Get()
	logsMonitoringGateway := main_gateways_logs.NewLogsMonitoringGatewayImpl()
	spanGateway := main_gateways_spans.NewSpanGatewayImpl()

	createEmailBodyFactory := *main_usecases_factories.NewCreateEmailBodyFactoryAllArgs(
		createWelcomeEmailTemplateBody, logsMonitoringGateway, spanGateway)
	return &createEmailBodyFactory
}
