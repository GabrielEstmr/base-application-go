package main_gateways_ws_middlewares

import (
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_gateways "baseapplicationgo/main/gateways"
	main_gateways_logs "baseapplicationgo/main/gateways/logs"
	main_gateways_spans "baseapplicationgo/main/gateways/spans"
	main_gateways_ws_commons "baseapplicationgo/main/gateways/ws/commons"
	main_utils "baseapplicationgo/main/utils"
	main_utils_messages "baseapplicationgo/main/utils/messages"
	"fmt"
	"net/http"
)

const _MSG_KEY_CONTENT_TYPE_NOT_FOUND = "middlewares.headers.content.type.not.found.error"

type AcceptJsonContentTypeHeaderMiddleware struct {
	acceptLanguageHeaderKey string
	stringUtils             main_utils.StringUtils
	messageUtils            main_utils_messages.ApplicationMessages
	spanGateway             main_gateways.SpanGateway
	logsMonitoringGateway   main_gateways.LogsMonitoringGateway
}

func NewAcceptJsonContentTypeHeaderMiddleware() *AcceptJsonContentTypeHeaderMiddleware {
	return &AcceptJsonContentTypeHeaderMiddleware{
		acceptLanguageHeaderKey: "Content-Type",
		stringUtils:             *main_utils.NewStringUtils(),
		messageUtils:            *main_utils_messages.NewApplicationMessages(),
		spanGateway:             main_gateways_spans.NewSpanGatewayImpl(),
		logsMonitoringGateway:   main_gateways_logs.NewLogsMonitoringGatewayImpl(),
	}
}

func (this *AcceptJsonContentTypeHeaderMiddleware) ServeHTTP(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		span := this.spanGateway.Get(r.Context(), fmt.Sprintf("AcceptJsonContentTypeHeaderMiddleware %s", r.URL.Path))
		defer span.End()

		contentTypeHeader := r.Header.Get(this.acceptLanguageHeaderKey)

		if !this.stringUtils.IsEmpty(contentTypeHeader) {
			if main_gateways_ws_commons.CONTENT_TYPE_JSON != contentTypeHeader {
				message := this.messageUtils.
					GetDefaultLocale(_MSG_KEY_CONTENT_TYPE_NOT_FOUND)
				errApp := main_domains_exceptions.NewResourceNotFoundExceptionSglMsg(message)
				this.logsMonitoringGateway.ERROR(span, message)
				main_utils.ERROR_APP(w, errApp)
				return
			}
		}
		newR := r.WithContext(span.GetCtx())
		h.ServeHTTP(w, newR)
	}
	return http.HandlerFunc(fn)
}
