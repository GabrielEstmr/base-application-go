package main_gateways_logs_resources

import (
	"context"
	"go.opentelemetry.io/otel/trace"
)

type SpanLogInfo struct {
	span trace.Span
}

func NewSpanLogInfo(ctx context.Context) *SpanLogInfo {
	return &SpanLogInfo{span: trace.SpanFromContext(ctx)}
}

func (this *SpanLogInfo) GetSpan() trace.Span {
	return this.span
}

func (this *SpanLogInfo) End() {
	this.span.End()
}
