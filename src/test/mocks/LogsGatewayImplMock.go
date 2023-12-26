package test_mocks

import (
	main_domains "baseapplicationgo/main/domains"
)

type LogsMonitoringGatewayMock struct {
}

func (this *LogsMonitoringGatewayMock) DEBUG(
	span main_domains.SpanLogInfo,
	msg string,
	args ...any,
) {

}

func (this *LogsMonitoringGatewayMock) WARN(
	span main_domains.SpanLogInfo,
	msg string,
	args ...any,
) {

}

func (this *LogsMonitoringGatewayMock) INFO(
	span main_domains.SpanLogInfo,
	msg string,
	args ...any,
) {

}

func (this *LogsMonitoringGatewayMock) ERROR(
	span main_domains.SpanLogInfo,
	msg string,
	args ...any,
) {

}
