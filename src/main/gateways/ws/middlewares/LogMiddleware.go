package main_gateways_ws_middlewares

import (
	main_gateways "baseapplicationgo/main/gateways"
	main_gateways_logs "baseapplicationgo/main/gateways/logs"
	main_gateways_spans "baseapplicationgo/main/gateways/spans"
	"fmt"
	"log/slog"
	"net/http"
)

type LogMiddleware struct {
	spanGateway           main_gateways.SpanGateway
	logsMonitoringGateway main_gateways.LogsMonitoringGateway
	appMiddleware         AppMiddleware
}

func NewLogMiddleware(appMiddleware AppMiddleware) *LogMiddleware {
	return &LogMiddleware{
		main_gateways_spans.NewSpanGatewayImpl(),
		main_gateways_logs.NewLogsMonitoringGatewayImpl(),
		appMiddleware,
	}
}

func (this *LogMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	span := this.spanGateway.Get(r.Context(), fmt.Sprintf("Middleware %s", r.URL.Path))
	defer span.End()

	this.logsMonitoringGateway.INFO(span, fmt.Sprintf("%s - %s (%s)", r.Method, r.URL.Path, r.RemoteAddr))
	slog.Info(fmt.Sprintf("%s - %s (%s)", r.Method, r.URL.Path, r.RemoteAddr))

	r = r.WithContext(span.GetCtx())
	this.appMiddleware.ServeHTTP(w, r)
}
