package main_gateways_ws

import (
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	mainGatewaysWsBeans "baseapplicationgo/main/gateways/ws/beans"
	main_gateways_ws_commons "baseapplicationgo/main/gateways/ws/commons"
	main_gateways_ws_errorhandler "baseapplicationgo/main/gateways/ws/errorhandler"
	main_gateways_ws_interceptors "baseapplicationgo/main/gateways/ws/interceptors"
	main_gateways_ws_middlewares "baseapplicationgo/main/gateways/ws/middlewares"
	"github.com/gorilla/mux"
	"net/http"
	"sync"
)

const API_V1_PREFIX = "/api/v1"

type Route struct {
	URI          string
	Method       string
	Function     ControllerParams
	AuthRequired bool
	Handler      Middlewares
}

type ControllerParams struct {
	controllerFunc func(http.ResponseWriter, *http.Request)
}

func NewControllerParams(controllerFunc func(http.ResponseWriter, *http.Request) (main_gateways_ws_commons.ControllerResponse, main_domains_exceptions.ApplicationException)) *ControllerParams {
	return &ControllerParams{controllerFunc: main_gateways_ws_errorhandler.NewAppCustomErrorHandlerImpl(controllerFunc).ServeHTTP}
}

type Middlewares struct {
	funcs []func(h http.Handler) http.Handler
}

func NewMiddlewares(funcs ...func(h http.Handler) http.Handler) *Middlewares {
	return &Middlewares{funcs: funcs}
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
			Function:     *NewControllerParams(beans.UserControllerV1Bean.CreateUser),
			AuthRequired: false,
			Handler: *NewMiddlewares(
				main_gateways_ws_middlewares.NewCheckTokenMiddleware().ServeHTTP,
				main_gateways_ws_interceptors.NewRunAfterTestImpl().ServeHTTP,
			),
		},
		{
			URI:          API_V1_PREFIX + "/users/{id}",
			Method:       http.MethodGet,
			Function:     *NewControllerParams(beans.UserControllerV1Bean.FindUserById),
			AuthRequired: false,
			Handler: *NewMiddlewares(
				main_gateways_ws_middlewares.NewCheckTokenMiddleware().ServeHTTP,
				main_gateways_ws_interceptors.NewRunAfterTestImpl().ServeHTTP,
			),
		},
		{
			URI:          API_V1_PREFIX + "/users",
			Method:       http.MethodGet,
			Function:     *NewControllerParams(beans.UserControllerV1Bean.FindUser),
			AuthRequired: false,
			Handler: *NewMiddlewares(
				main_gateways_ws_middlewares.NewCheckTokenMiddleware().ServeHTTP,
				main_gateways_ws_interceptors.NewRunAfterTestImpl().ServeHTTP,
			),
		},
		{
			URI:          API_V1_PREFIX + "/features/{key}/enable",
			Method:       http.MethodPost,
			Function:     *NewControllerParams(beans.FeatureControllerV1Bean.EnableFeatureByKey),
			AuthRequired: false,
			Handler: *NewMiddlewares(
				main_gateways_ws_middlewares.NewCheckTokenMiddleware().ServeHTTP,
				main_gateways_ws_interceptors.NewRunAfterTestImpl().ServeHTTP,
			),
		},
		{
			URI:          API_V1_PREFIX + "/features/{key}/disable",
			Method:       http.MethodPost,
			Function:     *NewControllerParams(beans.FeatureControllerV1Bean.DisableFeatureByKey),
			AuthRequired: false,
			Handler: *NewMiddlewares(
				main_gateways_ws_middlewares.NewCheckTokenMiddleware().ServeHTTP,
				main_gateways_ws_interceptors.NewRunAfterTestImpl().ServeHTTP,
			),
		},
		{
			URI:          API_V1_PREFIX + "/rabbitmq/send-event",
			Method:       http.MethodPost,
			Function:     *NewControllerParams(beans.RabbitMqControllerV1Bean.CreateRabbitMqTransactionEvent),
			AuthRequired: false,
			Handler: *NewMiddlewares(
				main_gateways_ws_middlewares.NewCheckTokenMiddleware().ServeHTTP,
				main_gateways_ws_interceptors.NewRunAfterTestImpl().ServeHTTP,
			),
		},
		{
			URI:          API_V1_PREFIX + "/transactions",
			Method:       http.MethodPost,
			Function:     *NewControllerParams(beans.TransactionControllerV1Bean.CreateTransaction),
			AuthRequired: false,
			Handler: *NewMiddlewares(
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
		for _, v := range route.Handler.funcs {
			subRouter.Use(v)
		}
		subRouter.HandleFunc("", route.Function.controllerFunc).Methods(route.Method)
	}

	sh := http.StripPrefix("/swagger-ui/", http.FileServer(http.Dir("./main/configs/doc/dist/")))
	r.PathPrefix("/swagger-ui/").Handler(sh)
	return r
}
