package main_gateways_ws

import (
	mainGatewaysWsBeans "baseapplicationgo/main/gateways/ws/beans"
	"github.com/gorilla/mux"
	"net/http"
	"sync"
)

type Routes struct {
	URI          string
	Method       string
	Function     func(http.ResponseWriter, *http.Request)
	AuthRequired bool
}

var once sync.Once

var RoutesBean *[]Routes = nil

func GetRoutesBean() *[]Routes {
	once.Do(func() {

		if RoutesBean == nil {
			RoutesBean = getFunctionBeans()
		}

	})
	return RoutesBean
}

func getFunctionBeans() *[]Routes {

	beans := mainGatewaysWsBeans.GetControllerBeans()

	var RoutesConfig = []Routes{
		{
			URI:          "/accounts",
			Method:       http.MethodGet,
			Function:     beans.AccountControllerBean.FindAccount,
			AuthRequired: false,
		},
	}

	return &RoutesConfig
}

func ConfigRoutes(r *mux.Router, routes []Routes) *mux.Router {
	for _, route := range routes {
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}
	sh := http.StripPrefix("/swagger-ui/", http.FileServer(http.Dir("./main/configurations/doc/dist/")))
	r.PathPrefix("/swagger-ui/").Handler(sh)
	return r
}
