package main_gateways_ws_middlewares

import (
	"net/http"
)

type AppMiddleware interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}
