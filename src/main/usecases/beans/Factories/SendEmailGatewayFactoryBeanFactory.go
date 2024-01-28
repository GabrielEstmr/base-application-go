package main_usecases_beans_factories

import (
	main_gateways "baseapplicationgo/main/gateways"
	main_gateways_email "baseapplicationgo/main/gateways/email"
	main_gateways_logs "baseapplicationgo/main/gateways/logs"
	main_gateways_spans "baseapplicationgo/main/gateways/spans"
	main_usecases_factories "baseapplicationgo/main/usecases/factories"
)

type SendEmailGatewayFactoryBean struct {
}

func NewSendEmailGatewayFactoryBean() *SendEmailGatewayFactoryBean {
	return &SendEmailGatewayFactoryBean{}
}

func (this *SendEmailGatewayFactoryBean) Get() *main_usecases_factories.SendEmailGatewayFactory {
	logsMonitoringGateway := main_gateways_logs.NewLogsMonitoringGatewayImpl()
	spanGateway := main_gateways_spans.NewSpanGatewayImpl()
	var gmailEmailGatewayImpl main_gateways.EmailGateway = main_gateways_email.NewGmailEmailGatewayImpl()
	return main_usecases_factories.NewSendEmailGatewayFactoryAllArgs(gmailEmailGatewayImpl, logsMonitoringGateway, spanGateway)
}
