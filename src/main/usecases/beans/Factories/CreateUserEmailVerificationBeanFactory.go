package main_usecases_beans_factories

import (
	main_gateways "baseapplicationgo/main/gateways"
	main_gateways_lock "baseapplicationgo/main/gateways/lock"
	main_gateways_logs "baseapplicationgo/main/gateways/logs"
	main_gateways_spans "baseapplicationgo/main/gateways/spans"
	main_usecases "baseapplicationgo/main/usecases"
	main_usecases_interfaces "baseapplicationgo/main/usecases/interfaces"
	main_utils_messages "baseapplicationgo/main/utils/messages"
)

type CreateUserEmailVerificationBean struct {
}

func NewCreateUserEmailVerificationBean() *CreateUserEmailVerificationBean {
	return &CreateUserEmailVerificationBean{}
}

func (this *CreateUserEmailVerificationBean) Get() *main_usecases.CreateUserEmailVerification {

	var generateEmailVerificationCode main_usecases_interfaces.GenerateEmailVerificationCode = NewGenerateEmailVerificationCodeBean().Get()
	var distributedLockGateway main_gateways.DistributedLockGateway = main_gateways_lock.NewDistributedLockGatewayImpl()
	var logsMonitoringGateway main_gateways.LogsMonitoringGateway = main_gateways_logs.NewLogsMonitoringGatewayImpl()
	var spanGatewayImpl main_gateways.SpanGateway = main_gateways_spans.NewSpanGatewayImpl()

	return main_usecases.NewCreateUserEmailVerification(
		generateEmailVerificationCode,
		distributedLockGateway,
		logsMonitoringGateway,
		spanGatewayImpl,
		*main_utils_messages.NewApplicationMessages())
}
