package main_gateways_ws_middlewares_impl

import (
	main_gateways_ws_interceptors "baseapplicationgo/main/gateways/ws/interceptors"
	"net/http"
)

type AppHandlerDecoratorBaseImpl struct {
	interceptorInscription main_gateways_ws_interceptors.InterceptorInscription
}

func NewAppHandlerDecoratorBaseImpl(interceptorInscription main_gateways_ws_interceptors.InterceptorInscription) *AppHandlerDecoratorBaseImpl {
	return &AppHandlerDecoratorBaseImpl{interceptorInscription}
}

func (this AppHandlerDecoratorBaseImpl) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	this.interceptorInscription.ServeHTTP(w, r)
}
