package main_gateways_ws_middlewares

import (
	main_gateways_ws_commons "baseapplicationgo/main/gateways/ws/commons"
	main_utils "baseapplicationgo/main/utils"
	"net/http"
)

type AcceptJsonContentTypeHeaderMiddleware struct {
	acceptLanguageHeaderKey string
	stringUtils             main_utils.StringUtils
}

func NewAcceptJsonContentTypeHeaderMiddleware() *AcceptJsonContentTypeHeaderMiddleware {
	return &AcceptJsonContentTypeHeaderMiddleware{
		acceptLanguageHeaderKey: "Accept-Language",
		stringUtils:             *main_utils.NewStringUtils(),
	}
}

func (this *AcceptJsonContentTypeHeaderMiddleware) ServeHTTP(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		languageHeader := r.Header.Get(this.acceptLanguageHeaderKey)
		if this.stringUtils.IsEmpty(languageHeader) {
			r.Header.Set(this.acceptLanguageHeaderKey, main_gateways_ws_commons.EN_US)
		}
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
