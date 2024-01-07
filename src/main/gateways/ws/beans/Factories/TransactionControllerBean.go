package main_gateways_ws_beans_factories

import (
	main_gateways "baseapplicationgo/main/gateways"
	main_gateways_logs "baseapplicationgo/main/gateways/logs"
	main_gateways_spans "baseapplicationgo/main/gateways/spans"
	main_gateways_ws_v1 "baseapplicationgo/main/gateways/ws/v1"
	main_usecases_beans "baseapplicationgo/main/usecases/beans"
)

type TransactionControllerBean struct {
}

func NewTransactionControllerBean() *TransactionControllerBean {
	return &TransactionControllerBean{}
}

func (this *TransactionControllerBean) Get() *main_gateways_ws_v1.TransactionController {
	var logsMonitoringGateway main_gateways.LogsMonitoringGateway = main_gateways_logs.NewLogsMonitoringGatewayImpl()
	var spanGatewayImpl main_gateways.SpanGateway = main_gateways_spans.NewSpanGatewayImpl()
	usecaseBeans := main_usecases_beans.GetUsecaseBeans()

	userController := *main_gateways_ws_v1.NewTransactionController(
		usecaseBeans.PersistTransaction,
		logsMonitoringGateway,
		spanGatewayImpl,
	)
	return &userController
}
