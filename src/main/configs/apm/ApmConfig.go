package main_configs_apm

import (
	main_configs_apm_metrics "baseapplicationgo/main/configs/apm/metrics"
	main_configs_apm_tracer "baseapplicationgo/main/configs/apm/tracer"
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/trace"
)

func InitiateApmConfig(mainCtx *context.Context) (*trace.TracerProvider, *metric.MeterProvider) {
	tracerProvider := main_configs_apm_tracer.GetTracerProviderBean(mainCtx)
	meterProvider := main_configs_apm_metrics.GetMetricProviderBean(mainCtx)
	otel.SetTracerProvider(tracerProvider)
	otel.SetMeterProvider(meterProvider)
	return tracerProvider, meterProvider
}
