package main

import (
	main_configs "baseapplicationgo/main/configs"
	main_configs_apm "baseapplicationgo/main/configs/apm"
	main_configs_error "baseapplicationgo/main/configs/error"
	mainConfigsRouterHttp "baseapplicationgo/main/configs/router"
	main_configs_yml "baseapplicationgo/main/configs/yml"
	main_gateways_rabbitmq_listeners "baseapplicationgo/main/gateways/rabbitmq/listeners"
	mainGatewaysWs "baseapplicationgo/main/gateways/ws"
	"context"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"
	"log"
	"net/http"
)

const MSG_APPLICATION_FAILED = "Application has failed to start"
const MSG_LISTENER = "Listener on port: %s"

const IDX_APPLICATION_PORT = "Application.Port"
const IDX_TRACER_APM_SERVER_NAME_YML = "Apm.server.name"

func init() {
	main_configs.InitConfigBeans()
	go main_gateways_rabbitmq_listeners.SubscribeListeners()
}

func main() {
	ctx := context.Background()
	main_configs_apm.InitiateApmConfig(&ctx)
	defer main_configs.TerminateConfigBeans(&ctx)

	applicationPort := main_configs_yml.GetYmlValueByName(IDX_APPLICATION_PORT)
	routes := mainGatewaysWs.GetRoutesBean()
	router := mainGatewaysWs.ConfigRoutes(mainConfigsRouterHttp.GetRouterBean(), *routes)
	router.Use(otelmux.Middleware(main_configs_yml.GetYmlValueByName(IDX_TRACER_APM_SERVER_NAME_YML)))
	router.Handle("/metrics", promhttp.Handler())
	log.Printf(MSG_LISTENER, applicationPort)

	err2 := http.ListenAndServe(":"+applicationPort, router)
	if err2 != nil {
		main_configs_error.FailOnError(err2, MSG_APPLICATION_FAILED)
	}
}
