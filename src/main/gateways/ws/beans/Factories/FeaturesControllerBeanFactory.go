package main_gateways_ws_beans_factories

import (
	main_configs_apm_logs_impl "baseapplicationgo/main/configs/apm/logs/impl"
	main_gateways "baseapplicationgo/main/gateways"
	main_gateways_logs "baseapplicationgo/main/gateways/logs"
	main_gateways_ws_v1 "baseapplicationgo/main/gateways/ws/v1"
)

type FeatureControllerBean struct {
}

func NewFeatureControllerBean() *FeatureControllerBean {
	return &FeatureControllerBean{}
}

func (this *FeatureControllerBean) Get() *main_gateways_ws_v1.FeaturesController {
	var logsMonitoringGateway main_gateways.LogsMonitoringGateway = main_gateways_logs.NewLogsMonitoringGatewayImpl(
		main_configs_apm_logs_impl.NewLogsGatewayImpl())
	featureController := *main_gateways_ws_v1.NewFeaturesController(
		logsMonitoringGateway,
	)
	return &featureController
}
