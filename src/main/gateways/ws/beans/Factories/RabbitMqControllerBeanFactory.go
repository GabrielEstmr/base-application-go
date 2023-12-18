package main_gateways_ws_beans_factories

import (
	main_configs_apm_logs_impl "baseapplicationgo/main/configs/apm/logs/impl"
	main_gateways "baseapplicationgo/main/gateways"
	main_gateways_logs "baseapplicationgo/main/gateways/logs"
	main_gateways_ws_v1 "baseapplicationgo/main/gateways/ws/v1"
	main_usecases_beans "baseapplicationgo/main/usecases/beans"
)

type RabbitMqControllerBean struct {
}

func NewRabbitMqControllerBean() *RabbitMqControllerBean {
	return &RabbitMqControllerBean{}
}

func (this *RabbitMqControllerBean) Get() *main_gateways_ws_v1.RabbitMqController {
	var logsMonitoringGateway main_gateways.LogsMonitoringGateway = main_gateways_logs.NewLogsMonitoringGatewayImpl(
		main_configs_apm_logs_impl.NewLogsGatewayImpl())
	usecaseBeans := main_usecases_beans.GetUsecaseBeans()
	rabbitMqController := *main_gateways_ws_v1.NewRabbitMqController(
		usecaseBeans.CreateTransactionAmqpEvent,
		logsMonitoringGateway,
	)
	return &rabbitMqController
}
