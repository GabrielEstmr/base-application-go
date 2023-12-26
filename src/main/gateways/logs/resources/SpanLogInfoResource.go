package main_gateways_logs_resources

import (
	main_configs_apm_tracer "baseapplicationgo/main/configs/apm/tracer"
	"context"
	"go.opentelemetry.io/otel/trace"
)

type SpanLogInfoResource struct {
	span trace.Span
	ctx  context.Context
}

func NewSpanLogInfoResourceAllArgs(
	span trace.Span,
	ctx context.Context) *SpanLogInfoResource {
	return &SpanLogInfoResource{
		span: span,
		ctx:  ctx}
}

func NewSpanLogInfoResource(ctx context.Context, scope string) *SpanLogInfoResource {
	tracer := main_configs_apm_tracer.GetTracerProviderBean(&ctx).Tracer(scope)
	ctxSpan, spanN := tracer.Start(ctx, scope)
	return &SpanLogInfoResource{span: spanN, ctx: ctxSpan}
}

func (this *SpanLogInfoResource) GetSpan() trace.Span {
	return this.span
}

func (this *SpanLogInfoResource) GetCtx() context.Context {
	return this.ctx
}

func (this *SpanLogInfoResource) End() {
	this.span.End()
}
