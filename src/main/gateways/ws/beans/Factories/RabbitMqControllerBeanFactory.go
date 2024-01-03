package main_gateways_ws_beans_factories

import (
	main_gateways "baseapplicationgo/main/gateways"
	main_gateways_logs "baseapplicationgo/main/gateways/logs"
	main_gateways_spans "baseapplicationgo/main/gateways/spans"
	main_gateways_ws_v1 "baseapplicationgo/main/gateways/ws/v1"
	main_usecases_beans "baseapplicationgo/main/usecases/beans"
)

type RabbitMqControllerBean struct {
}

func NewRabbitMqControllerBean() *RabbitMqControllerBean {
	return &RabbitMqControllerBean{}
}

func (this *RabbitMqControllerBean) Get() *main_gateways_ws_v1.RabbitMqController {
	var logsMonitoringGateway main_gateways.LogsMonitoringGateway = main_gateways_logs.NewLogsMonitoringGatewayImpl()
	var spanGatewayImpl main_gateways.SpanGateway = main_gateways_spans.NewSpanGatewayImpl()
	usecaseBeans := main_usecases_beans.GetUsecaseBeans()
	rabbitMqController := *main_gateways_ws_v1.NewRabbitMqController(
		usecaseBeans.CreateTransactionAmqpEvent,
		logsMonitoringGateway,
		spanGatewayImpl,
	)
	return &rabbitMqController
}
