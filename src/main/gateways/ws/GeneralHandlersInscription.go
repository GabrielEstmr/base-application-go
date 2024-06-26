package main_gateways_ws

import (
	main_gateways_ws_commons "baseapplicationgo/main/gateways/ws/commons"
	main_gateways_ws_interceptors "baseapplicationgo/main/gateways/ws/interceptors"
	main_gateways_ws_middlewares "baseapplicationgo/main/gateways/ws/middlewares"
	"net/http"
)

type GeneralHandlersInscription struct {
	middlewares main_gateways_ws_commons.Middlewares
}

func NewGeneralHandlersInscription() *GeneralHandlersInscription {
	request := main_gateways_ws_interceptors.NewMetricsAfterRequest()
	return &GeneralHandlersInscription{
		middlewares: *main_gateways_ws_commons.NewMiddlewares(
			main_gateways_ws_middlewares.NewRouteLoggingBeforeRequestMiddleware().ServeHTTP,
			request.ServeHTTP,
			main_gateways_ws_middlewares.NewCheckTokenBeforeRequestMiddleware().ServeHTTP,
			main_gateways_ws_middlewares.NewAcceptLanguageValidationMiddleware().ServeHTTP,
		)}
}

func (this *GeneralHandlersInscription) Build() []func(h http.Handler) http.Handler {
	return this.middlewares.GetFuncs()
}
