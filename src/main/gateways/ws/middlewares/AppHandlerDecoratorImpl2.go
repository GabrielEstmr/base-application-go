package main_gateways_ws_middlewares

import (
	"log"
	"net/http"
)

type AppHandlerDecoratorImpl2 struct {
	decorator AppMiddleware
}

func NewAppHandlerDecoratorImpl2(decorator AppMiddleware) *AppHandlerDecoratorImpl2 {
	return &AppHandlerDecoratorImpl2{decorator: decorator}
}

func (this AppHandlerDecoratorImpl2) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("DECORATOR2")
	this.decorator.ServeHTTP(w, r)
}
