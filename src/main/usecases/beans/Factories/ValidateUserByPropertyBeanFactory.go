package main_usecases_beans_factories

import (
	main_gateways "baseapplicationgo/main/gateways"
	main_gateways_logs "baseapplicationgo/main/gateways/logs"
	main_gateways_spans "baseapplicationgo/main/gateways/spans"
	main_usecases "baseapplicationgo/main/usecases"
	main_usecases_interfaces "baseapplicationgo/main/usecases/interfaces"
	main_utils_messages "baseapplicationgo/main/utils/messages"
)

type ValidateUserByPropertyBean struct {
}

func NewValidateUserByPropertyBean() *ValidateUserByPropertyBean {
	return &ValidateUserByPropertyBean{}
}

func (this *ValidateUserByPropertyBean) Get() *main_usecases.ValidateHasUserByProperties {

	validateUserByPropertyImpls := make([]main_usecases_interfaces.ValidateUserByProperty, 0)
	validateUserByPropertyImpls = append(validateUserByPropertyImpls, NewValidateHasUserByEmail().Get())
	validateUserByPropertyImpls = append(validateUserByPropertyImpls, NewValidateHasUserByUserNameBean().Get())
	validateUserByPropertyImpls = append(validateUserByPropertyImpls, NewValidateHasUserByDocumentIdBean().Get())

	var logsMonitoringGateway main_gateways.LogsMonitoringGateway = main_gateways_logs.NewLogsMonitoringGatewayImpl()
	var spanGatewayImpl main_gateways.SpanGateway = main_gateways_spans.NewSpanGatewayImpl()

	return main_usecases.NewValidateHasUserByProperties(
		validateUserByPropertyImpls,
		logsMonitoringGateway,
		spanGatewayImpl,
		*main_utils_messages.NewApplicationMessages())
}
