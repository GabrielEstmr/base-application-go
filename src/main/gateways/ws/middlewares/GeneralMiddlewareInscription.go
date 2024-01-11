package main_gateways_ws_middlewares

import (
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_gateways_ws_commons "baseapplicationgo/main/gateways/ws/commons"
	"net/http"
)

type GeneralMiddlewareInscription struct {
	appMiddleware AppMiddleware
}

func NewGeneralMiddlewareInscription(
	fn func(http.ResponseWriter, *http.Request) (main_gateways_ws_commons.ControllerResponse,
		main_domains_exceptions.ApplicationException)) *GeneralMiddlewareInscription {
	return &GeneralMiddlewareInscription{
		appMiddleware: NewLogMiddleware(
			NewCheckTokenMiddleware(
				NewAcceptLanguageMiddleware(
					NewAppCustomErrorHandlerImpl(fn)))),
	}
}

func (this GeneralMiddlewareInscription) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	this.appMiddleware.ServeHTTP(w, r)
}
