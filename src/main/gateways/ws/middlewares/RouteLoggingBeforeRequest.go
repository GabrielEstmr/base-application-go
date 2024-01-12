package main_gateways_ws_middlewares

import (
	main_gateways "baseapplicationgo/main/gateways"
	main_gateways_logs "baseapplicationgo/main/gateways/logs"
	main_gateways_spans "baseapplicationgo/main/gateways/spans"
	"fmt"
	"log/slog"
	"net/http"
)

const _ROUTE_LOGGING_MSG_KEY = "%s - %s (%s)"

type RouteLoggingBeforeRequestMiddleware struct {
	spanGateway           main_gateways.SpanGateway
	logsMonitoringGateway main_gateways.LogsMonitoringGateway
}

func NewRouteLoggingBeforeRequestMiddleware() *RouteLoggingBeforeRequestMiddleware {
	return &RouteLoggingBeforeRequestMiddleware{
		main_gateways_spans.NewSpanGatewayImpl(),
		main_gateways_logs.NewLogsMonitoringGatewayImpl(),
	}
}

func (this *RouteLoggingBeforeRequestMiddleware) ServeHTTP(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		span := this.spanGateway.Get(r.Context(), fmt.Sprintf("Middleware %s", r.URL.Path))
		defer span.End()
		this.logsMonitoringGateway.INFO(span, fmt.Sprintf(_ROUTE_LOGGING_MSG_KEY, r.Method, r.URL.Path, r.RemoteAddr))
		slog.Info(fmt.Sprintf(_ROUTE_LOGGING_MSG_KEY, r.Method, r.URL.Path, r.RemoteAddr))
		r = r.WithContext(span.GetCtx())
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
