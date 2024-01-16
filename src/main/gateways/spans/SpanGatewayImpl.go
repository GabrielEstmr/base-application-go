package main_gateways_spans

import (
	main_domains "baseapplicationgo/main/domains/apm"
	"context"
)

type SpanGatewayImpl struct {
}

func NewSpanGatewayImpl() *SpanGatewayImpl {
	return &SpanGatewayImpl{}
}

func (this *SpanGatewayImpl) Get(ctx context.Context, scope string) main_domains.SpanLogInfo {
	return *main_domains.NewSpanLogInfo(ctx, scope)
}
