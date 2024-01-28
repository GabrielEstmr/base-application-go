package main_usecases_beans_factories

import (
	main_gateways_logs "baseapplicationgo/main/gateways/logs"
	main_gateways_spans "baseapplicationgo/main/gateways/spans"
	main_usecases "baseapplicationgo/main/usecases"
	main_utils_messages "baseapplicationgo/main/utils/messages"
)

type CreateEmailBodyForWelcomeEmailBean struct {
}

func NewCreateEmailBodyForWelcomeEmailBean() *CreateEmailBodyForWelcomeEmailBean {
	return &CreateEmailBodyForWelcomeEmailBean{}
}

func (this *CreateEmailBodyForWelcomeEmailBean) Get() *main_usecases.CreateEmailBodyForWelcomeEmail {
	return main_usecases.NewCreateEmailBodyForWelcomeEmail(
		main_gateways_logs.NewLogsMonitoringGatewayImpl(),
		main_gateways_spans.NewSpanGatewayImpl(),
		*main_utils_messages.NewApplicationMessages(),
	)
}
