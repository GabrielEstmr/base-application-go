package main_gateways_logs

import (
	main_configs_apm_logs_impl "baseapplicationgo/main/configs/apm/logs/impl"
	main_gateways_logs_resources "baseapplicationgo/main/gateways/logs/resources"
)

type LogsMonitoringGatewayImpl struct {
	logLoki main_configs_apm_logs_impl.LogsMethods
}

func NewLogsMonitoringGatewayImpl(
	logLoki main_configs_apm_logs_impl.LogsMethods) *LogsMonitoringGatewayImpl {
	return &LogsMonitoringGatewayImpl{
		logLoki: logLoki,
	}
}

func (this *LogsMonitoringGatewayImpl) DEBUG(
	span main_gateways_logs_resources.SpanLogInfo,
	msg string,
	args ...any,
) {
	this.logLoki.DEBUG(span.GetSpan(), msg, args)
}

func (this *LogsMonitoringGatewayImpl) WARN(
	span main_gateways_logs_resources.SpanLogInfo,
	msg string,
	args ...any,
) {
	this.logLoki.WARN(span.GetSpan(), msg, args)
}

func (this *LogsMonitoringGatewayImpl) INFO(
	span main_gateways_logs_resources.SpanLogInfo,
	msg string,
	args ...any,
) {
	this.logLoki.INFO(span.GetSpan(), msg, args)
}

func (this *LogsMonitoringGatewayImpl) ERROR(
	span main_gateways_logs_resources.SpanLogInfo,
	msg string,
	args ...any,
) {
	this.logLoki.ERROR(span.GetSpan(), msg, args)
}