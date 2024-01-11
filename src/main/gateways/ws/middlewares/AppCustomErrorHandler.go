package main_gateways_ws_middlewares

import (
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_gateways_ws_commons "baseapplicationgo/main/gateways/ws/commons"
	main_utils "baseapplicationgo/main/utils"
	"net/http"
)

type AppCustomErrorHandlerImpl struct {
	fn func(http.ResponseWriter, *http.Request) (main_gateways_ws_commons.ControllerResponse,
		main_domains_exceptions.ApplicationException)
}

func NewAppCustomErrorHandlerImpl(fn func(http.ResponseWriter, *http.Request) (main_gateways_ws_commons.ControllerResponse,
	main_domains_exceptions.ApplicationException)) *AppCustomErrorHandlerImpl {
	return &AppCustomErrorHandlerImpl{
		fn: fn,
	}
}

func (this AppCustomErrorHandlerImpl) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	v, e := this.fn(w, r)
	if e != nil {
		main_utils.ERROR_APP(w, e)
		return
	}
	main_utils.JSON(w, v.GetStatusCode(), v.GetData())
}
