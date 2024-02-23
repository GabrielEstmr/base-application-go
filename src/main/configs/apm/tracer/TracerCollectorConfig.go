package main_configs_apm_tracer

import (
	main_configs_error "baseapplicationgo/main/configs/error"
	main_configs_yml "baseapplicationgo/main/configs/yml"
	"context"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
	"log"
	"sync"
)

const _MSG_ERROR_TRACER_EXPORTER = "Error to instantiate tracer exporter"
const _MSG_ERROR_TRACER_RESOURCE = "Error to instantiate tracer resource"
const _MSG_ERROR_SHUTDOWN_TRACER_PROVIDER = "Error to shutdown tracer provider"

var onceTracer sync.Once
var tracerProviderBean *trace.TracerProvider

func GetTracerProviderBean(mainCtx *context.Context) *trace.TracerProvider {
	onceTracer.Do(func() {
		if tracerProviderBean == nil {
			tracerProviderBean = getTracerProvider(mainCtx)
		}
	})
	return tracerProviderBean
}

func getTracerProvider(mainCtx *context.Context) *trace.TracerProvider {

	ctx := *mainCtx

	otlpEndpoint := main_configs_yml.GetYmlValueByName(main_configs_yml.ApmServerOtlpCollectorGrpcHost)
	log.Println(otlpEndpoint)

	insecureOpt := otlptracegrpc.WithInsecure()
	endpointOpt := otlptracegrpc.WithEndpoint(otlpEndpoint)

	otlptracegrpc.NewClient(endpointOpt)
	tracerExporter, err := otlptracegrpc.New(ctx, endpointOpt, insecureOpt)
	main_configs_error.FailOnError(err, _MSG_ERROR_TRACER_EXPORTER)

	serverName := main_configs_yml.GetYmlValueByName(main_configs_yml.ApmServerName)
	log.Println(serverName)

	tracerResource, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(serverName),
		),
	)
	main_configs_error.FailOnError(err, _MSG_ERROR_TRACER_RESOURCE)

	tracerProvider := trace.NewTracerProvider(
		trace.WithBatcher(tracerExporter),
		trace.WithResource(tracerResource),
	)

	return tracerProvider
}

func ShutdownTracerProvider(mainCtx *context.Context) {
	if err := GetTracerProviderBean(mainCtx).Shutdown(*mainCtx); err != nil {
		main_configs_error.FailOnError(err, _MSG_ERROR_SHUTDOWN_TRACER_PROVIDER)
	}
}
