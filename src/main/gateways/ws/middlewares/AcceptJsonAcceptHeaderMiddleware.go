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

const _MSG_KEY_ACCEPT_TYPE_NOT_FOUND = "middlewares.headers.accept.not.found.error"

type AcceptJsonAcceptHeaderMiddleware struct {
	acceptLanguageHeaderKey string
	stringUtils             main_utils.StringUtils
	messageUtils            main_utils_messages.ApplicationMessages
	spanGateway             main_gateways.SpanGateway
	logsMonitoringGateway   main_gateways.LogsMonitoringGateway
}

func NewAcceptJsonAcceptHeaderMiddleware() *AcceptJsonAcceptHeaderMiddleware {
	return &AcceptJsonAcceptHeaderMiddleware{
		acceptLanguageHeaderKey: "accept",
		stringUtils:             *main_utils.NewStringUtils(),
		messageUtils:            *main_utils_messages.NewApplicationMessages(),
		spanGateway:             main_gateways_spans.NewSpanGatewayImpl(),
		logsMonitoringGateway:   main_gateways_logs.NewLogsMonitoringGatewayImpl(),
	}
}

func (this *AcceptJsonAcceptHeaderMiddleware) ServeHTTP(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		span := this.spanGateway.Get(r.Context(), fmt.Sprintf("AcceptJsonAcceptHeaderMiddleware %s", r.URL.Path))
		defer span.End()

		contentTypeHeader := r.Header.Get(this.acceptLanguageHeaderKey)

		if !this.stringUtils.IsEmpty(contentTypeHeader) {
			if isNotJsonOrAllMediaTypeHeader(contentTypeHeader) {
				message := this.messageUtils.
					GetDefaultLocale(_MSG_KEY_ACCEPT_TYPE_NOT_FOUND)
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

func isNotJsonOrAllMediaTypeHeader(contentTypeHeader string) bool {
	return main_gateways_ws_commons.ACCEPT_TYPE_JSON != contentTypeHeader &&
		main_gateways_ws_commons.ACCEPT_TYPE_ALL != contentTypeHeader
}
