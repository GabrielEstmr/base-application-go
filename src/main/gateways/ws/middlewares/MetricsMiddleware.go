package main_gateways_ws_middlewares

import (
	main_configs_apm_metrics "baseapplicationgo/main/configs/apm/metrics"
	"baseapplicationgo/main/domains/apm"
	main_gateways "baseapplicationgo/main/gateways"
	main_gateways_logs "baseapplicationgo/main/gateways/logs"
	main_gateways_spans "baseapplicationgo/main/gateways/spans"
	main_gateways_ws_commons "baseapplicationgo/main/gateways/ws/commons"
	main_utils "baseapplicationgo/main/utils"
	main_utils_messages "baseapplicationgo/main/utils/messages"
	"context"
	"fmt"
	"go.opentelemetry.io/otel/metric"
	"net/http"
	"path"
)

type MetricsMiddleware struct {
	stringUtils           main_utils.StringUtils
	messageUtils          main_utils_messages.ApplicationMessages
	spanGateway           main_gateways.SpanGateway
	logsMonitoringGateway main_gateways.LogsMonitoringGateway
}

func NewMetricsMiddleware() *MetricsMiddleware {
	return &MetricsMiddleware{
		stringUtils:           *main_utils.NewStringUtils(),
		messageUtils:          *main_utils_messages.NewApplicationMessages(),
		spanGateway:           main_gateways_spans.NewSpanGatewayImpl(),
		logsMonitoringGateway: main_gateways_logs.NewLogsMonitoringGatewayImpl(),
	}
}

func (this *MetricsMiddleware) ServeHTTP(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		span := this.spanGateway.Get(r.Context(), fmt.Sprintf("MetricsMiddleware %s", r.URL.Path))
		defer span.End()
		go buildHttpNumberOfHits(r, span)
		newR := r.WithContext(span.GetCtx())

		h.ServeHTTP(w, newR)
	}
	return http.HandlerFunc(fn)
}

func buildHttpNumberOfHits(r *http.Request, _ apm.SpanLogInfo) {

	baseRoute := path.Dir(path.Clean(r.URL.Path))
	routeBean := main_gateways_ws_commons.RoutesURI

	_, ok1 := routeBean[baseRoute]
	if ok1 {
		buildHttpNumberOfHitsPerRouteAndMethodMetric(r.URL.Path, r)
		return
	}

	buildHttpNumberOfHitsPerRouteAndMethodMetric(r.URL.Path, r)
}

func buildHttpNumberOfHitsPerRouteAndMethodMetric(route string, r *http.Request) {
	name := fmt.Sprintf("http_total_requests_route_%s_method_%s", route, r.Method)
	GlobalMeter := *main_configs_apm_metrics.GetGlobalMeterBean()
	if _, err := GlobalMeter.Float64ObservableCounter(
		name,
		metric.WithDescription("Number os hits per route"),
		metric.WithUnit("un"),
		metric.WithFloat64Callback(func(ctx context.Context, o metric.Float64Observer) error {
			o.Observe(1)
			return nil
		}),
	); err != nil {
		panic(err)
	}
}
