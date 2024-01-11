package main_gateways_ws_middlewares

import (
	"log"
	"net/http"
)

type AppHandlerDecoratorImpl struct {
	decorator AppMiddleware
}

func NewAppHandlerDecoratorImpl(decorator AppMiddleware) *AppHandlerDecoratorImpl {
	return &AppHandlerDecoratorImpl{decorator: decorator}
}

func (this AppHandlerDecoratorImpl) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("DECORATOR1")
	this.decorator.ServeHTTP(w, r)
}
