package test_mocks

import (
	main_domains "baseapplicationgo/main/domains"
	main_gateways_logs_resources "baseapplicationgo/main/gateways/logs/resources"
	"context"
	"go.opentelemetry.io/otel/trace"
)

type SpanGatewayMockImpl struct {
}

func NewSpanGatewayMockImpl() *SpanGatewayMockImpl {
	return &SpanGatewayMockImpl{}
}

func (this *SpanGatewayMockImpl) Get(ctx context.Context, scope string) main_domains.SpanLogInfo {
	fromContext := trace.SpanFromContext(ctx)
	args := *main_gateways_logs_resources.NewSpanLogInfoResourceAllArgs(fromContext, ctx)
	return *main_domains.NewSpanLogInfoAllArgs(args)
}
