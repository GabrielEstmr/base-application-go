package main_gateways_logs

import (
	main_configs_apm_logs_impl "baseapplicationgo/main/configs/apm/logs/impl"
	main_domains "baseapplicationgo/main/domains/apm"
)

type LogsMonitoringGatewayImpl struct {
	logLoki main_configs_apm_logs_impl.LogsMethods
}

func NewLogsMonitoringGatewayImplAllArgs(
	logLoki main_configs_apm_logs_impl.LogsMethods) *LogsMonitoringGatewayImpl {
	return &LogsMonitoringGatewayImpl{
		logLoki: logLoki,
	}
}

func NewLogsMonitoringGatewayImpl() *LogsMonitoringGatewayImpl {
	return &LogsMonitoringGatewayImpl{
		logLoki: main_configs_apm_logs_impl.NewLogsGatewayImpl(),
	}
}

func (this *LogsMonitoringGatewayImpl) DEBUG(
	span main_domains.SpanLogInfo,
	msg string,
	args ...any,
) {
	this.logLoki.DEBUG(span.GetSpan(), msg, args)
}

func (this *LogsMonitoringGatewayImpl) WARN(
	span main_domains.SpanLogInfo,
	msg string,
	args ...any,
) {
	this.logLoki.WARN(span.GetSpan(), msg, args)
}

func (this *LogsMonitoringGatewayImpl) INFO(
	span main_domains.SpanLogInfo,
	msg string,
	args ...any,
) {
	this.logLoki.INFO(span.GetSpan(), msg, args)
}

func (this *LogsMonitoringGatewayImpl) ERROR(
	span main_domains.SpanLogInfo,
	msg string,
	args ...any,
) {
	this.logLoki.ERROR(span.GetSpan(), msg, args)
}
