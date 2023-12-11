package main_configs_apm_metrics

import (
	main_configs_error "baseapplicationgo/main/configs/error"
	main_configs_yml "baseapplicationgo/main/configs/yml"
	"context"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/sdk/metric"
	"sync"
)

const _MSG_ERROR_METRIC_EXPORTER = "Error to instantiate metric exporter"
const _MSG_ERROR_SHUTDOWN_METRIC_PROVIDER = "Error to shutdown metric provider"

const _METRIC_APM_CLUSTER_OTLP_ENDPOINT_YML = "Apm.server.otlp.collector.grpc.host"

var onceMetric sync.Once
var metricProviderBean *metric.MeterProvider

func GetMetricProviderBean(mainCtx *context.Context) *metric.MeterProvider {
	onceMetric.Do(func() {
		if metricProviderBean == nil {
			metricProviderBean = getMetricProvider(mainCtx)
		}
	})
	return metricProviderBean
}

func getMetricProvider(mainCtx *context.Context) *metric.MeterProvider {

	ctx := *mainCtx

	otlpEndpoint := main_configs_yml.GetYmlValueByName(_METRIC_APM_CLUSTER_OTLP_ENDPOINT_YML)
	insecureMetricOpt := otlpmetricgrpc.WithInsecure()
	endpointMetricOpt := otlpmetricgrpc.WithEndpoint(otlpEndpoint)

	metricsExporter, err := otlpmetricgrpc.New(ctx, insecureMetricOpt, endpointMetricOpt)
	main_configs_error.FailOnError(err, _MSG_ERROR_METRIC_EXPORTER)

	meterProvider := metric.NewMeterProvider(
		metric.WithReader(
			metric.NewPeriodicReader(metricsExporter)))
	return meterProvider
}

func ShutdownMetricProvider(mainCtx *context.Context) {
	if err := GetMetricProviderBean(mainCtx).Shutdown(*mainCtx); err != nil {
		main_configs_error.FailOnError(err, _MSG_ERROR_SHUTDOWN_METRIC_PROVIDER)
	}
}
