package main_gateways_ws_interceptors

import (
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_gateways_ws_commons "baseapplicationgo/main/gateways/ws/commons"
	main_gateways_ws_interceptors_impl "baseapplicationgo/main/gateways/ws/interceptors/impl"
	"net/http"
)

type InterceptorInscription struct {
	appInterceptor AppInterceptor
}

func NewInterceptorInscription(
	fn func(http.ResponseWriter, *http.Request) (main_gateways_ws_commons.ControllerResponse,
		main_domains_exceptions.ApplicationException)) *InterceptorInscription {
	return &InterceptorInscription{
		appInterceptor: main_gateways_ws_interceptors_impl.NewAppCustomErrorHandlerImpl(fn),
	}
}

func (this InterceptorInscription) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	this.appInterceptor.ServeHTTP(w, r)
}
