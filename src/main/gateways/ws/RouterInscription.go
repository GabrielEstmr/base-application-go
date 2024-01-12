package main_gateways_ws

import (
	mainGatewaysWsBeans "baseapplicationgo/main/gateways/ws/beans"
	main_gateways_ws_commons "baseapplicationgo/main/gateways/ws/commons"
	main_gateways_ws_interceptors "baseapplicationgo/main/gateways/ws/interceptors"
	main_gateways_ws_middlewares "baseapplicationgo/main/gateways/ws/middlewares"
	"github.com/gorilla/mux"
	"net/http"
	"sync"
)

const API_V1_PREFIX = "/api/v1"

type Route struct {
	URI              string
	Method           string
	ControllerParams main_gateways_ws_commons.ControllerParams
	AuthRequired     bool
	Handler          main_gateways_ws_commons.Middlewares
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
			URI:    API_V1_PREFIX + "/users",
			Method: http.MethodPost,
			ControllerParams: *main_gateways_ws_commons.NewControllerParams(
				beans.UserControllerV1Bean.CreateUser),
			AuthRequired: false,
			Handler: *main_gateways_ws_commons.NewMiddlewares(
				main_gateways_ws_middlewares.NewCheckTokenMiddleware().ServeHTTP,
				main_gateways_ws_interceptors.NewRunAfterTestImpl().ServeHTTP,
			),
		},
		{
			URI:    API_V1_PREFIX + "/users/{id}",
			Method: http.MethodGet,
			ControllerParams: *main_gateways_ws_commons.NewControllerParams(
				beans.UserControllerV1Bean.FindUserById),
			AuthRequired: false,
			Handler: *main_gateways_ws_commons.NewMiddlewares(
				main_gateways_ws_middlewares.NewCheckTokenMiddleware().ServeHTTP,
				main_gateways_ws_interceptors.NewRunAfterTestImpl().ServeHTTP,
			),
		},
		{
			URI:    API_V1_PREFIX + "/users",
			Method: http.MethodGet,
			ControllerParams: *main_gateways_ws_commons.NewControllerParams(
				beans.UserControllerV1Bean.FindUser),
			AuthRequired: false,
			Handler: *main_gateways_ws_commons.NewMiddlewares(
				main_gateways_ws_middlewares.NewCheckTokenMiddleware().ServeHTTP,
				main_gateways_ws_interceptors.NewRunAfterTestImpl().ServeHTTP,
			),
		},
		{
			URI:    API_V1_PREFIX + "/features/{key}/enable",
			Method: http.MethodPost,
			ControllerParams: *main_gateways_ws_commons.NewControllerParams(
				beans.FeatureControllerV1Bean.EnableFeatureByKey),
			AuthRequired: false,
			Handler: *main_gateways_ws_commons.NewMiddlewares(
				main_gateways_ws_middlewares.NewCheckTokenMiddleware().ServeHTTP,
				main_gateways_ws_interceptors.NewRunAfterTestImpl().ServeHTTP,
			),
		},
		{
			URI:    API_V1_PREFIX + "/features/{key}/disable",
			Method: http.MethodPost,
			ControllerParams: *main_gateways_ws_commons.NewControllerParams(
				beans.FeatureControllerV1Bean.DisableFeatureByKey),
			AuthRequired: false,
			Handler: *main_gateways_ws_commons.NewMiddlewares(
				main_gateways_ws_middlewares.NewCheckTokenMiddleware().ServeHTTP,
				main_gateways_ws_interceptors.NewRunAfterTestImpl().ServeHTTP,
			),
		},
		{
			URI:    API_V1_PREFIX + "/rabbitmq/send-event",
			Method: http.MethodPost,
			ControllerParams: *main_gateways_ws_commons.NewControllerParams(
				beans.RabbitMqControllerV1Bean.CreateRabbitMqTransactionEvent),
			AuthRequired: false,
			Handler: *main_gateways_ws_commons.NewMiddlewares(
				main_gateways_ws_middlewares.NewCheckTokenMiddleware().ServeHTTP,
				main_gateways_ws_interceptors.NewRunAfterTestImpl().ServeHTTP,
			),
		},
		{
			URI:    API_V1_PREFIX + "/transactions",
			Method: http.MethodPost,
			ControllerParams: *main_gateways_ws_commons.NewControllerParams(
				beans.TransactionControllerV1Bean.CreateTransaction),
			AuthRequired: false,
			Handler: *main_gateways_ws_commons.NewMiddlewares(
				main_gateways_ws_middlewares.NewCheckTokenMiddleware().ServeHTTP,
				main_gateways_ws_interceptors.NewRunAfterTestImpl().ServeHTTP,
			),
		},
	}
	return RoutesConfig
}

// TODO: ver como colocar swagger
func ConfigRoutes(r *mux.Router, routes []Route) *mux.Router {
	for _, route := range routes {
		subRouter := r.PathPrefix(route.URI).Subrouter()
		for _, v := range route.Handler.GetFuncs() {
			subRouter.Use(v)
		}
		subRouter.HandleFunc("", route.ControllerParams.GetHttpFunc()).Methods(route.Method)
	}

	sh := http.StripPrefix("/swagger-ui/", http.FileServer(http.Dir("./main/configs/doc/dist/")))
	r.PathPrefix("/swagger-ui/").Handler(sh)
	return r
}
