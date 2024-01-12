package main_gateways_ws_middlewares

import "net/http"

type RunBeforeRequest interface {
	ServeHTTP(h http.Handler) http.Handler
}
