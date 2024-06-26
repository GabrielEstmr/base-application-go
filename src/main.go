package main

import (
	main_configs "baseapplicationgo/main/configs"
	main_configs_apm "baseapplicationgo/main/configs/apm"
	main_configs_error "baseapplicationgo/main/configs/error"
	mainConfigsRouterHttp "baseapplicationgo/main/configs/router"
	main_configs_yml "baseapplicationgo/main/configs/yml"
	main_gateways_rabbitmq "baseapplicationgo/main/gateways/rabbitmq/subscribers"
	main_gateways_ws "baseapplicationgo/main/gateways/ws"
	main_gateways_ws_beans "baseapplicationgo/main/gateways/ws/beans"
	"context"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

const MSG_APPLICATION_FAILED = "Application has failed to start"
const MSG_LISTENER = "Listener on port: %s"

const IDX_APPLICATION_PORT = "Application.Port"
const IDX_TRACER_APM_SERVER_NAME_YML = "Apm.server.name"

func init() {
	main_configs.InitConfigBeans()
	go main_gateways_rabbitmq.SubscribeListeners()
}

func main() {
	ctx := context.Background()
	main_configs_apm.InitiateApmConfig(&ctx)
	defer main_configs.TerminateConfigBeans(&ctx)
	applicationPort := main_configs_yml.GetYmlValueByName(IDX_APPLICATION_PORT)
	routes := main_gateways_ws_beans.GetRoutesBean()
	router := main_gateways_ws.ConfigRoutes(mainConfigsRouterHttp.GetRouterBean(), *routes)
	router.Handle("/metrics", promhttp.Handler())
	log.Printf(MSG_LISTENER, applicationPort)

	err2 := http.ListenAndServe(":"+applicationPort, router)
	if err2 != nil {
		main_configs_error.FailOnError(err2, MSG_APPLICATION_FAILED)
	}
}
