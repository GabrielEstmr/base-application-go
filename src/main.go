package main

import (
	main_configs "baseapplicationgo/main/configs"
	main_error "baseapplicationgo/main/configs/error"
	mainConfigsRouterHttp "baseapplicationgo/main/configs/router"
	main_configs_yml "baseapplicationgo/main/configs/yml"
	mainGatewaysWs "baseapplicationgo/main/gateways/ws"
	"log"
	"net/http"
)

const MSG_APPLICATION_FAILED = "Application has failed to start"
const MSG_LISTENER = "Listener on port: %s"
const IDX_APPLICATION_PORT = "Application.Port"

func init() {
	main_configs.InitConfigBeans()
}

func main() {
	defer main_configs.TerminateConfigBeans()

	applicationPort := main_configs_yml.GetYmlValueByName(IDX_APPLICATION_PORT)
	routes := mainGatewaysWs.GetRoutesBean()
	router := mainGatewaysWs.ConfigRoutes(mainConfigsRouterHttp.GetRouterBean(), *routes)

	err := http.ListenAndServe(":"+applicationPort, router)
	if err != nil {
		main_error.FailOnError(err, MSG_APPLICATION_FAILED)
	}
	log.Printf(MSG_LISTENER, applicationPort)
}
