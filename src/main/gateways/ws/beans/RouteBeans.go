package main_gateways_ws_beans

import (
	main_gateways_ws_commons "baseapplicationgo/main/gateways/ws/commons"
	main_gateways_ws_middlewares "baseapplicationgo/main/gateways/ws/middlewares"
	"fmt"
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

func (this *Route) GetRoutKey() string {
	return fmt.Sprintf("%s - %s", this.Method, this.URI)
}

var onceR sync.Once

var routes *map[string]Route = nil
var routesBean map[string]Route = nil

func GetRoutesBean() *map[string]Route {
	onceR.Do(func() {
		if routesBean == nil {
			routesBean = getFunctionBeans()
			routes = &routesBean
		}
	})
	return routes
}

func getFunctionBeans() map[string]Route {
	beans := GetControllerBeans()

	var RoutesConfig = []Route{
		{
			URI:    API_V1_PREFIX + "/users",
			Method: http.MethodPost,
			ControllerParams: *main_gateways_ws_commons.NewControllerParams(
				beans.UserControllerV1Bean.CreateUser),
			AuthRequired: false,
			Handler: *main_gateways_ws_commons.NewMiddlewares(
				main_gateways_ws_middlewares.NewAcceptJsonAcceptHeaderMiddleware().ServeHTTP,
				main_gateways_ws_middlewares.NewAcceptJsonContentTypeHeaderMiddleware().ServeHTTP,
			),
		},
		{
			URI:    API_V1_PREFIX + "/users/{id}",
			Method: http.MethodGet,
			ControllerParams: *main_gateways_ws_commons.NewControllerParams(
				beans.UserControllerV1Bean.FindUserById),
			AuthRequired: false,
			Handler: *main_gateways_ws_commons.NewMiddlewares(
				main_gateways_ws_middlewares.NewAcceptJsonAcceptHeaderMiddleware().ServeHTTP),
		},
		{
			URI:    API_V1_PREFIX + "/users",
			Method: http.MethodGet,
			ControllerParams: *main_gateways_ws_commons.NewControllerParams(
				beans.UserControllerV1Bean.FindUser),
			AuthRequired: false,
			Handler: *main_gateways_ws_commons.NewMiddlewares(
				main_gateways_ws_middlewares.NewAcceptJsonAcceptHeaderMiddleware().ServeHTTP),
		},
		{
			URI:    API_V1_PREFIX + "/features/{key}/enable",
			Method: http.MethodPost,
			ControllerParams: *main_gateways_ws_commons.NewControllerParams(
				beans.FeatureControllerV1Bean.EnableFeatureByKey),
			AuthRequired: false,
			Handler: *main_gateways_ws_commons.NewMiddlewares(
				main_gateways_ws_middlewares.NewAcceptJsonAcceptHeaderMiddleware().ServeHTTP),
		},
		{
			URI:    API_V1_PREFIX + "/features/{key}/disable",
			Method: http.MethodPost,
			ControllerParams: *main_gateways_ws_commons.NewControllerParams(
				beans.FeatureControllerV1Bean.DisableFeatureByKey),
			AuthRequired: false,
			Handler: *main_gateways_ws_commons.NewMiddlewares(
				main_gateways_ws_middlewares.NewAcceptJsonAcceptHeaderMiddleware().ServeHTTP),
		},
		{
			URI:    API_V1_PREFIX + "/rabbitmq/send-event",
			Method: http.MethodPost,
			ControllerParams: *main_gateways_ws_commons.NewControllerParams(
				beans.RabbitMqControllerV1Bean.CreateRabbitMqTransactionEvent),
			AuthRequired: false,
			Handler: *main_gateways_ws_commons.NewMiddlewares(
				main_gateways_ws_middlewares.NewAcceptJsonAcceptHeaderMiddleware().ServeHTTP,
				main_gateways_ws_middlewares.NewAcceptJsonContentTypeHeaderMiddleware().ServeHTTP,
			),
		},
		{
			URI:    API_V1_PREFIX + "/transactions",
			Method: http.MethodPost,
			ControllerParams: *main_gateways_ws_commons.NewControllerParams(
				beans.TransactionControllerV1Bean.CreateTransaction),
			AuthRequired: false,
			Handler: *main_gateways_ws_commons.NewMiddlewares(
				main_gateways_ws_middlewares.NewAcceptJsonAcceptHeaderMiddleware().ServeHTTP,
				main_gateways_ws_middlewares.NewAcceptJsonContentTypeHeaderMiddleware().ServeHTTP,
			),
		},
	}

	RoutesConfigMap := make(map[string]Route)
	for _, route := range RoutesConfig {
		RoutesConfigMap[route.GetRoutKey()] = route
	}
	return RoutesConfigMap
}
