package test_mocks

import (
	main_configs_apm_logs_resources "baseapplicationgo/main/configs/apm/logs/resources"
	"go.opentelemetry.io/otel/trace"
)

type LogsGatewayImplMock struct {
	logConfig main_configs_apm_logs_resources.LogProviderConfig
}

func (this LogsGatewayImplMock) DEBUG(
	span trace.Span,
	msg string,
	args ...any,
) {

}

func (this LogsGatewayImplMock) WARN(
	span trace.Span,
	msg string,
	args ...any,
) {

}

func (this LogsGatewayImplMock) INFO(
	span trace.Span,
	msg string,
	args ...any,
) {

}

func (this LogsGatewayImplMock) ERROR(
	span trace.Span,
	msg string,
	args ...any,
) {

}
