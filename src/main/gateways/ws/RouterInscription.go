package main_gateways_ws

import (
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	mainGatewaysWsBeans "baseapplicationgo/main/gateways/ws/beans"
	main_gateways_ws_commons "baseapplicationgo/main/gateways/ws/commons"
	main_gateways_ws_middlewares "baseapplicationgo/main/gateways/ws/middlewares"
	"github.com/gorilla/mux"
	"net/http"
	"sync"
)

const API_V1_PREFIX = "/api/v1"

type Route struct {
	URI      string
	Method   string
	Function func(http.ResponseWriter, *http.Request) (
		main_gateways_ws_commons.ControllerResponse,
		main_domains_exceptions.ApplicationException)
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
			URI:          API_V1_PREFIX + "/users",
			Method:       http.MethodPost,
			Function:     beans.UserControllerV1Bean.CreateUser,
			AuthRequired: false,
		},
		{
			URI:          API_V1_PREFIX + "/users/{id}",
			Method:       http.MethodGet,
			Function:     beans.UserControllerV1Bean.FindUserById,
			AuthRequired: false,
		},
		{
			URI:          API_V1_PREFIX + "/users",
			Method:       http.MethodGet,
			Function:     beans.UserControllerV1Bean.FindUser,
			AuthRequired: false,
		},
		{
			URI:          API_V1_PREFIX + "/features/{key}/enable",
			Method:       http.MethodPost,
			Function:     beans.FeatureControllerV1Bean.EnableFeatureByKey,
			AuthRequired: false,
		},
		{
			URI:          API_V1_PREFIX + "/features/{key}/disable",
			Method:       http.MethodPost,
			Function:     beans.FeatureControllerV1Bean.DisableFeatureByKey,
			AuthRequired: false,
		},
		{
			URI:          API_V1_PREFIX + "/rabbitmq/send-event",
			Method:       http.MethodPost,
			Function:     beans.RabbitMqControllerV1Bean.CreateRabbitMqTransactionEvent,
			AuthRequired: false,
		},
		{
			URI:          API_V1_PREFIX + "/transactions",
			Method:       http.MethodPost,
			Function:     beans.TransactionControllerV1Bean.CreateTransaction,
			AuthRequired: false,
		},
	}
	return RoutesConfig
}

// TODO: ver como colocar swagger
func ConfigRoutes(r *mux.Router, routes []Route) *mux.Router {

	for _, route := range routes {
		middleware := main_gateways_ws_middlewares.NewGeneralMiddlewareInscription(route.Function)
		r.HandleFunc(route.URI, middleware.ServeHTTP).Methods(route.Method)
	}

	sh := http.StripPrefix("/swagger-ui/", http.FileServer(http.Dir("./main/configs/doc/dist/")))
	r.PathPrefix("/swagger-ui/").Handler(sh)
	return r
}
