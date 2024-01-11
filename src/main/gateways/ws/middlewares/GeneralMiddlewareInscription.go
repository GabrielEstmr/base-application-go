package main_gateways_ws_middlewares

import (
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_gateways_ws_commons "baseapplicationgo/main/gateways/ws/commons"
	main_gateways_ws_interceptors "baseapplicationgo/main/gateways/ws/interceptors"
	main_gateways_ws_middlewares_impl "baseapplicationgo/main/gateways/ws/middlewares/impl"
	"net/http"
)

type GeneralMiddlewareInscription struct {
	appMiddleware AppMiddleware
}

func NewGeneralMiddlewareInscription(
	fn func(http.ResponseWriter, *http.Request) (main_gateways_ws_commons.ControllerResponse,
		main_domains_exceptions.ApplicationException)) *GeneralMiddlewareInscription {
	return &GeneralMiddlewareInscription{
		appMiddleware: main_gateways_ws_middlewares_impl.NewLogMiddleware(
			main_gateways_ws_middlewares_impl.NewCheckTokenMiddleware(
				main_gateways_ws_middlewares_impl.NewAcceptLanguageMiddleware(
					main_gateways_ws_middlewares_impl.NewAppHandlerDecoratorBaseImpl(
						*main_gateways_ws_interceptors.NewInterceptorInscription(fn))))),
	}
}

func (this GeneralMiddlewareInscription) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	this.appMiddleware.ServeHTTP(w, r)
}
