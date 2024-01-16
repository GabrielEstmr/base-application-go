package main_gateways_ws

import (
	"baseapplicationgo/main/gateways/ws/beans"
	"github.com/gorilla/mux"
	"net/http"
)

// TODO: ver como colocar swagger
func ConfigRoutes(r *mux.Router, routes map[string]main_gateways_ws_beans.Route) *mux.Router {
	for _, route := range routes {
		subRouter := r.PathPrefix(route.URI).Subrouter()

		for _, v := range NewGeneralHandlersInscription().Build() {
			subRouter.Use(v)
		}

		for _, v := range route.Handler.GetFuncs() {
			subRouter.Use(v)
		}
		subRouter.HandleFunc("", route.ControllerParams.GetHttpFunc()).Methods(route.Method)
	}

	sh := http.StripPrefix("/swagger-ui/", http.FileServer(http.Dir("./main/configs/doc/dist/")))
	r.PathPrefix("/swagger-ui/").Handler(sh)
	return r
}
