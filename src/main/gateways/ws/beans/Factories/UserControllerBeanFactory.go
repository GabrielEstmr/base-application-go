package main_gateways_ws_beans_factories

import (
	main_configs_apm_logs_impl "baseapplicationgo/main/configs/apm/logs/impl"
	main_gateways "baseapplicationgo/main/gateways"
	main_gateways_logs "baseapplicationgo/main/gateways/logs"
	main_gateways_spans "baseapplicationgo/main/gateways/spans"
	main_gateways_ws_v1 "baseapplicationgo/main/gateways/ws/v1"
	main_usecases_beans "baseapplicationgo/main/usecases/beans"
)

type UserControllerBean struct {
}

func NewUserControllerBean() *UserControllerBean {
	return &UserControllerBean{}
}

func (this *UserControllerBean) Get() *main_gateways_ws_v1.UserController {
	var logsMonitoringGateway main_gateways.LogsMonitoringGateway = main_gateways_logs.NewLogsMonitoringGatewayImpl(
		main_configs_apm_logs_impl.NewLogsGatewayImpl())
	var spanGatewayImpl main_gateways.SpanGateway = main_gateways_spans.NewSpanGatewayImpl()
	usecaseBeans := main_usecases_beans.GetUsecaseBeans()
	userController := *main_gateways_ws_v1.NewUserController(
		usecaseBeans.CreateNewUser,
		usecaseBeans.FindUserById,
		usecaseBeans.FindUsersByFilter,
		logsMonitoringGateway,
		spanGatewayImpl,
	)
	return &userController
}
