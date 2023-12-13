package main_gateways

import main_gateways_logs_resources "baseapplicationgo/main/gateways/logs/resources"

type LogsMonitoringGateway interface {
	DEBUG(
		span main_gateways_logs_resources.SpanLogInfo,
		msg string,
		args ...any,
	)

	WARN(
		span main_gateways_logs_resources.SpanLogInfo,
		msg string,
		args ...any,
	)

	INFO(
		span main_gateways_logs_resources.SpanLogInfo,
		msg string,
		args ...any,
	)

	ERROR(
		span main_gateways_logs_resources.SpanLogInfo,
		msg string,
		args ...any,
	)
}
