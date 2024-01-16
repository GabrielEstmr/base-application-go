package main_gateways_ws_interceptors

import "net/http"

type RunAfterRequest interface {
	ServeHTTP(h http.Handler) http.Handler
}
