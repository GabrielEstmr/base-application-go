package main_gateways_ws_commons

import "net/http"

type Middlewares struct {
	funcs []func(h http.Handler) http.Handler
}

func (this *Middlewares) GetFuncs() []func(h http.Handler) http.Handler {
	return this.funcs
}

func NewMiddlewares(funcs ...func(h http.Handler) http.Handler) *Middlewares {
	return &Middlewares{funcs: funcs}
}
