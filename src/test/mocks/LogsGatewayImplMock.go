package test_mocks

import (
	main_gateways_logs_resources "baseapplicationgo/main/gateways/logs/resources"
)

type LogsMonitoringGateway struct {
}

func (this *LogsMonitoringGateway) DEBUG(
	span main_gateways_logs_resources.SpanLogInfo,
	msg string,
	args ...any,
) {

}

func (this *LogsMonitoringGateway) WARN(
	span main_gateways_logs_resources.SpanLogInfo,
	msg string,
	args ...any,
) {

}

func (this *LogsMonitoringGateway) INFO(
	span main_gateways_logs_resources.SpanLogInfo,
	msg string,
	args ...any,
) {

}

func (this *LogsMonitoringGateway) ERROR(
	span main_gateways_logs_resources.SpanLogInfo,
	msg string,
	args ...any,
) {

}
