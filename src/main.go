package main

import (
	main_configs "baseapplicationgo/main/configs"
	main_configs_apm "baseapplicationgo/main/configs/apm"
	main_configs_error "baseapplicationgo/main/configs/error"
	mainConfigsRouterHttp "baseapplicationgo/main/configs/router"
	main_configs_yml "baseapplicationgo/main/configs/yml"
	main_gateways_rabbitmq "baseapplicationgo/main/gateways/rabbitmq/subscribers"
	mainGatewaysWs "baseapplicationgo/main/gateways/ws"
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
	routes := mainGatewaysWs.GetRoutesBean()
	router := mainGatewaysWs.ConfigRoutes(mainConfigsRouterHttp.GetRouterBean(), *routes)
	//router.Use(main_gateways_ws_middlewares.NewLogMiddleware().Handler)
	//router.Use(main_gateways_ws_middlewares.NewAcceptLanguageMiddleware().Handler)
	//router.Use(main_gateways_ws_middlewares.NewCheckTokenMiddleware().Handler)
	//router.Use(otelmux.Middleware(main_configs_yml.GetYmlValueByName(IDX_TRACER_APM_SERVER_NAME_YML)))
	router.Handle("/metrics", promhttp.Handler())
	log.Printf(MSG_LISTENER, applicationPort)

	err2 := http.ListenAndServe(":"+applicationPort, runsafter(runsbefore(router)))
	if err2 != nil {
		main_configs_error.FailOnError(err2, MSG_APPLICATION_FAILED)
	}
}

func runsafter(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
		log.Println("runsafter")
		//w.Write([]byte("run after, "))
	}
	return http.HandlerFunc(fn)
}

func runsbefore(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Println("runsbefore")
		//w.Write([]byte("run before, "))
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
