package main_gateways_ws_interceptors

import (
	main_configs_apm_metrics "baseapplicationgo/main/configs/apm/metrics"
	main_gateways "baseapplicationgo/main/gateways"
	main_gateways_logs "baseapplicationgo/main/gateways/logs"
	main_gateways_spans "baseapplicationgo/main/gateways/spans"
	main_gateways_ws_commons "baseapplicationgo/main/gateways/ws/commons"
	main_utils "baseapplicationgo/main/utils"
	main_utils_messages "baseapplicationgo/main/utils/messages"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
	"log"
	"log/slog"
	"net/http"
	"time"
)

type MetricsAfterRequest struct {
	stringUtils              main_utils.StringUtils
	messageUtils             main_utils_messages.ApplicationMessages
	spanGateway              main_gateways.SpanGateway
	logsMonitoringGateway    main_gateways.LogsMonitoringGateway
	requestDurationHistogram metric.Int64Histogram
}

func NewMetricsAfterRequest() MetricsAfterRequest {
	return MetricsAfterRequest{
		stringUtils:              *main_utils.NewStringUtils(),
		messageUtils:             *main_utils_messages.NewApplicationMessages(),
		spanGateway:              main_gateways_spans.NewSpanGatewayImpl(),
		logsMonitoringGateway:    main_gateways_logs.NewLogsMonitoringGatewayImpl(),
		requestDurationHistogram: buildHttpDurationHistogram(),
	}
}

func buildHttpDurationHistogram() metric.Int64Histogram {
	GlobalMeter := *main_configs_apm_metrics.GetGlobalMeterBean()

	histogram, err := GlobalMeter.Int64Histogram("http_request_duration_per_route_and_method_v2",
		metric.WithDescription("Duration per request"),
		metric.WithUnit("ms"),
	)
	if err != nil {
		slog.Error("Error to Instantiate the httpDurationHistogram")
		log.Fatal(err)
		return nil
	}
	return histogram
}

func (this *MetricsAfterRequest) ServeHTTP(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {

		span := this.spanGateway.Get(r.Context(), fmt.Sprintf("MetricsAfterRequestMiddleware %s", r.URL.Path))
		defer span.End()

		appResponseWriter := main_gateways_ws_commons.NewAppResponseWriter(w)
		initialTime := time.Now()

		h.ServeHTTP(appResponseWriter, r)
		duration := time.Since(initialTime)

		go this.buildMetrics(*appResponseWriter, r, duration)
	}
	return http.HandlerFunc(fn)
}

func (this *MetricsAfterRequest) buildMetrics(
	w main_gateways_ws_commons.AppResponseWriter,
	r *http.Request,
	duration time.Duration,
) {
	route := mux.CurrentRoute(r)
	pathTemplate, _ := route.GetPathTemplate()
	buildHttpNumberOfHitsPerRouteMethodAndStatusCodeMetric(w, pathTemplate, r)
	buildHttpNumberOfHitsPerRouteMethod(pathTemplate, r)
	this.addHttpDurationHistogramMetric(w, pathTemplate, r, duration)
}

func (this *MetricsAfterRequest) addHttpDurationHistogramMetric(
	w main_gateways_ws_commons.AppResponseWriter,
	route string,
	r *http.Request,
	duration time.Duration,
) {
	attributes := metric.WithAttributeSet(attribute.NewSet(
		semconv.HTTPMethod(r.Method),
		semconv.HTTPRoute(route),
		semconv.HTTPStatusCode(w.GetStatusCode()),
	))

	this.requestDurationHistogram.Record(r.Context(), duration.Microseconds(), attributes)

}

func buildHttpNumberOfHitsPerRouteMethodAndStatusCodeMetric(w main_gateways_ws_commons.AppResponseWriter, route string, r *http.Request) {
	GlobalMeter := *main_configs_apm_metrics.GetGlobalMeterBean()

	attributes := metric.WithAttributeSet(attribute.NewSet(
		semconv.HTTPMethod(r.Method),
		semconv.HTTPRoute(route),
		semconv.HTTPStatusCode(w.GetStatusCode()),
	))

	if _, err := GlobalMeter.Float64ObservableCounter(
		"http_request_per_route_and_method_and_status",
		metric.WithDescription("Number os hits per route method and status"),
		metric.WithUnit("un"),
		metric.WithFloat64Callback(func(ctx context.Context, o metric.Float64Observer) error {
			o.Observe(1, attributes)
			return nil
		}),
	); err != nil {
		panic(err)
	}
}

func buildHttpNumberOfHitsPerRouteMethod(route string, r *http.Request) {
	GlobalMeter := *main_configs_apm_metrics.GetGlobalMeterBean()

	attributes := metric.WithAttributeSet(attribute.NewSet(
		semconv.HTTPMethod(r.Method),
		semconv.HTTPRoute(route),
	))

	if _, err := GlobalMeter.Float64ObservableCounter(
		"http_request_per_route_and_method",
		metric.WithDescription("Number os hits per route and method"),
		metric.WithUnit("un"),
		metric.WithFloat64Callback(func(ctx context.Context, o metric.Float64Observer) error {
			o.Observe(1, attributes)
			return nil
		}),
	); err != nil {
		panic(err)
	}
}
