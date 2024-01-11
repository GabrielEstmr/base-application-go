package main_gateways_ws_middlewares

import (
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	"net/http"
)

type AppHandlerDecoratorBaseImpl struct {
	fn func(http.ResponseWriter, *http.Request) main_domains_exceptions.ApplicationException
}

func NewAppHandlerDecoratorBaseImpl(fn func(http.ResponseWriter, *http.Request) main_domains_exceptions.ApplicationException) *AppHandlerDecoratorBaseImpl {
	return &AppHandlerDecoratorBaseImpl{fn: fn}
}

func (this AppHandlerDecoratorBaseImpl) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	this.fn(w, r)
}
