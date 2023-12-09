package main

import (
	main_configs "baseapplicationgo/main/configs"
	main_configs_apm_metrics "baseapplicationgo/main/configs/apm/metrics"
	main_configs_apm_tracer "baseapplicationgo/main/configs/apm/tracer"
	main_configs_error "baseapplicationgo/main/configs/error"
	main_configs_profile "baseapplicationgo/main/configs/profile"
	mainConfigsRouterHttp "baseapplicationgo/main/configs/router"
	main_configs_yml "baseapplicationgo/main/configs/yml"
	mainGatewaysWs "baseapplicationgo/main/gateways/ws"
	"context"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"
	"go.opentelemetry.io/otel"
	"log"
	"net/http"
)

const MSG_APPLICATION_FAILED = "Application has failed to start"
const MSG_LISTENER = "Listener on port: %s"

const IDX_APPLICATION_PORT = "Application.Port"
const IDX_TRACER_APM_SERVER_NAME_YML = "Apm.server.name"

func init() {
	main_configs.InitConfigBeans()
}

func main() {
	ctx := context.Background()
	defer main_configs.TerminateConfigBeans(&ctx)

	profile := main_configs_profile.GetProfileBean().GetLowerCaseName()
	if profile != "local" {
		tracerProvider := main_configs_apm_tracer.GetTracerProviderBean(&ctx)
		meterProvider := main_configs_apm_metrics.GetMetricProviderBean(&ctx)

		otel.SetTracerProvider(tracerProvider)
		otel.SetMeterProvider(meterProvider)
	}

	applicationPort := main_configs_yml.GetYmlValueByName(IDX_APPLICATION_PORT)
	routes := mainGatewaysWs.GetRoutesBean()
	router := mainGatewaysWs.ConfigRoutes(mainConfigsRouterHttp.GetRouterBean(), *routes)
	router.Use(otelmux.Middleware(main_configs_yml.GetYmlValueByName(IDX_TRACER_APM_SERVER_NAME_YML)))

	err2 := http.ListenAndServe(":"+applicationPort, router)
	if err2 != nil {
		main_configs_error.FailOnError(err2, MSG_APPLICATION_FAILED)
	}
	log.Printf(MSG_LISTENER, applicationPort)
}
