package main_gateways_ws_middlewares

import (
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_gateways "baseapplicationgo/main/gateways"
	main_gateways_logs "baseapplicationgo/main/gateways/logs"
	main_gateways_spans "baseapplicationgo/main/gateways/spans"
	main_utils "baseapplicationgo/main/utils"
	main_utils_messages "baseapplicationgo/main/utils/messages"
	"fmt"
	"net/http"
)

const _MSG_KEY_UNAUTHORIZED_ERROR = "exceptions.unauthorized.error"

type CheckTokenBeforeRequestMiddleware struct {
	authHeaderKey         string
	stringUtils           main_utils.StringUtils
	messageUtils          main_utils_messages.ApplicationMessages
	spanGateway           main_gateways.SpanGateway
	logsMonitoringGateway main_gateways.LogsMonitoringGateway
}

func NewCheckTokenBeforeRequestMiddleware() *CheckTokenBeforeRequestMiddleware {
	return &CheckTokenBeforeRequestMiddleware{
		authHeaderKey:         "Authorization",
		stringUtils:           *main_utils.NewStringUtils(),
		messageUtils:          *main_utils_messages.NewApplicationMessages(),
		spanGateway:           main_gateways_spans.NewSpanGatewayImpl(),
		logsMonitoringGateway: main_gateways_logs.NewLogsMonitoringGatewayImpl(),
	}
}

func (this *CheckTokenBeforeRequestMiddleware) ServeHTTP(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		span := this.spanGateway.Get(r.Context(), fmt.Sprintf("CheckTokenBeforeRequestMiddleware %s", r.URL.Path))
		defer span.End()

		languageHeader := r.Header.Get(this.authHeaderKey)
		if this.stringUtils.IsEmpty(languageHeader) {
			errApp := main_domains_exceptions.NewUnauthorizedExceptionSglMsg(this.messageUtils.
				GetDefaultLocale(_MSG_KEY_UNAUTHORIZED_ERROR))
			this.logsMonitoringGateway.ERROR(span, _MSG_KEY_UNAUTHORIZED_ERROR)
			main_utils.ERROR_APP(w, errApp)
			return
		}
		newR := r.WithContext(span.GetCtx())
		h.ServeHTTP(w, newR)
	}
	return http.HandlerFunc(fn)
}
