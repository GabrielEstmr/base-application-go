package main_gateways_ws_middlewares_impl

import (
	main_gateways_ws_commons "baseapplicationgo/main/gateways/ws/commons"
	"baseapplicationgo/main/gateways/ws/middlewares"
	main_utils "baseapplicationgo/main/utils"
	"net/http"
)

type AcceptLanguageMiddleware struct {
	acceptLanguageHeaderKey string
	stringUtils             main_utils.StringUtils
	appMiddleware           main_gateways_ws_middlewares.AppMiddleware
}

func NewAcceptLanguageMiddleware(appMiddleware main_gateways_ws_middlewares.AppMiddleware) *AcceptLanguageMiddleware {
	return &AcceptLanguageMiddleware{
		acceptLanguageHeaderKey: "Accept-Language",
		stringUtils:             *main_utils.NewStringUtils(),
		appMiddleware:           appMiddleware,
	}
}

func (this *AcceptLanguageMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	languageHeader := r.Header.Get(this.acceptLanguageHeaderKey)
	if this.stringUtils.IsEmpty(languageHeader) {
		r.Header.Set(this.acceptLanguageHeaderKey, main_gateways_ws_commons.EN_US)
	}
	this.appMiddleware.ServeHTTP(w, r)
}
