package main_gateways_ws_commons

import (
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	"baseapplicationgo/main/gateways/ws/commonsresources"
	main_gateways_ws_errorhandler "baseapplicationgo/main/gateways/ws/errorhandler"
	"net/http"
)

type ControllerParams struct {
	httpFunc func(http.ResponseWriter, *http.Request)
}

type ControllerFunc func(http.ResponseWriter, *http.Request) (main_gateways_ws_commonsresources.ControllerResponse, main_domains_exceptions.ApplicationException)

func NewControllerParams(controllerFunc ControllerFunc) *ControllerParams {
	return &ControllerParams{
		httpFunc: main_gateways_ws_errorhandler.NewAppCustomErrorHandlerImpl(controllerFunc).ServeHTTP,
	}
}

func (this *ControllerParams) GetHttpFunc() func(http.ResponseWriter, *http.Request) {
	return this.httpFunc
}
