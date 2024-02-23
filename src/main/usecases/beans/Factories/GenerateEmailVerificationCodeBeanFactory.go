package main_usecases_beans_factories

import (
	main_gateways_logs "baseapplicationgo/main/gateways/logs"
	main_gateways_spans "baseapplicationgo/main/gateways/spans"
	main_usecases "baseapplicationgo/main/usecases"
	main_utils_messages "baseapplicationgo/main/utils/messages"
)

type GenerateEmailVerificationCodeBean struct {
}

func NewGenerateEmailVerificationCodeBean() *GenerateEmailVerificationCodeBean {
	return &GenerateEmailVerificationCodeBean{}
}

func (this *GenerateEmailVerificationCodeBean) Get() *main_usecases.GenerateEmailVerificationCode {
	return main_usecases.NewGenerateEmailVerificationCode(
		main_gateways_logs.NewLogsMonitoringGatewayImpl(),
		main_gateways_spans.NewSpanGatewayImpl(),
		*main_utils_messages.NewApplicationMessages(),
	)
}
