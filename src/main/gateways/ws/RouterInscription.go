package main_gateways_ws

import (
	mainGatewaysWsBeans "baseapplicationgo/main/gateways/ws/beans"
	"github.com/gorilla/mux"
	"net/http"
	"sync"
)

type Route struct {
	URI          string
	Method       string
	Function     func(http.ResponseWriter, *http.Request)
	AuthRequired bool
}

var once sync.Once

var routes *[]Route = nil
var routesBean []Route = nil

func GetRoutesBean() *[]Route {
	once.Do(func() {
		if routesBean == nil {
			routesBean = getFunctionBeans()
			routes = &routesBean
		}
	})
	return routes
}

func getFunctionBeans() []Route {
	beans := mainGatewaysWsBeans.GetControllerBeans()
	var RoutesConfig = []Route{
		{
			URI:          "/users",
			Method:       http.MethodPost,
			Function:     beans.UserControllerV1Bean.CreateUser,
			AuthRequired: false,
		},
		{
			URI:          "/users/{id}",
			Method:       http.MethodGet,
			Function:     beans.UserControllerV1Bean.FindUserById,
			AuthRequired: false,
		},
		{
			URI:          "/users",
			Method:       http.MethodGet,
			Function:     beans.UserControllerV1Bean.FindUser,
			AuthRequired: false,
		},
	}
	return RoutesConfig
}

// TODO: ver como colocar swagger
func ConfigRoutes(r *mux.Router, routes []Route) *mux.Router {
	for _, route := range routes {
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	sh := http.StripPrefix("/swagger-ui/", http.FileServer(http.Dir("./main/configs/doc/dist/")))
	r.PathPrefix("/swagger-ui/").Handler(sh)
	return r
}
