package test_mocks

import (
	main_gateways_logs_resources "baseapplicationgo/main/gateways/logs/resources"
)

type LogsMonitoringGatewayMock struct {
}

func (this *LogsMonitoringGatewayMock) DEBUG(
	span main_gateways_logs_resources.SpanLogInfo,
	msg string,
	args ...any,
) {

}

func (this *LogsMonitoringGatewayMock) WARN(
	span main_gateways_logs_resources.SpanLogInfo,
	msg string,
	args ...any,
) {

}

func (this *LogsMonitoringGatewayMock) INFO(
	span main_gateways_logs_resources.SpanLogInfo,
	msg string,
	args ...any,
) {

}

func (this *LogsMonitoringGatewayMock) ERROR(
	span main_gateways_logs_resources.SpanLogInfo,
	msg string,
	args ...any,
) {

}
