package main_gateways

import (
	main_domains "baseapplicationgo/main/domains"
)

type LogsMonitoringGateway interface {
	DEBUG(
		span main_domains.SpanLogInfo,
		msg string,
		args ...any,
	)

	WARN(
		span main_domains.SpanLogInfo,
		msg string,
		args ...any,
	)

	INFO(
		span main_domains.SpanLogInfo,
		msg string,
		args ...any,
	)

	ERROR(
		span main_domains.SpanLogInfo,
		msg string,
		args ...any,
	)
}
