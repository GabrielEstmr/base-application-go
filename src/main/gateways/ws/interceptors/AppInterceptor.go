package main_gateways_ws_interceptors

import (
	"net/http"
)

type AppInterceptor interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}
