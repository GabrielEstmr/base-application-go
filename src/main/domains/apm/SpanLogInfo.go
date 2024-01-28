package main_domains_apm

import (
	main_gateways_logs_resources "baseapplicationgo/main/gateways/logs/resources"
	"context"
	"go.opentelemetry.io/otel/trace"
)

type SpanLogInfo struct {
	span main_gateways_logs_resources.SpanLogInfoResource
}

func NewSpanLogInfoAllArgs(
	span main_gateways_logs_resources.SpanLogInfoResource,
) *SpanLogInfo {
	return &SpanLogInfo{span: span}
}

func NewSpanLogInfo(
	ctx context.Context,
	scope string,
) *SpanLogInfo {
	return &SpanLogInfo{
		span: *main_gateways_logs_resources.NewSpanLogInfoResource(ctx, scope),
	}
}

func (this *SpanLogInfo) GetSpan() trace.Span {
	return this.span.GetSpan()
}

func (this *SpanLogInfo) GetCtx() context.Context {
	return this.span.GetCtx()
}

func (this *SpanLogInfo) End() {
	this.span.End()
}
