package main

import (
	mainConfigsRouterHttp "baseapplicationgo/main/configs/router"
	mainConfigsYml "baseapplicationgo/main/configs/yml"
	mainGatewaysWs "baseapplicationgo/main/gateways/ws"
	main_utils "baseapplicationgo/main/utils"
	"log"
	"net/http"
)

const MSG_INITIALIZING_APPLICATION = "Initializing application."
const MSG_APPLICATION_FAILED = "Application has failed to start"
const MSG_LISTENER = "Listener on port: %s"

const IDX_APPLICATION_PORT = "Application.Port"
const IDX_TRACING_SERVER_NAME = "Tracing.server.name"

func init() {
	log.Println(MSG_INITIALIZING_APPLICATION)
	mainConfigsYml.GetYmlConfigBean()

}

func main() {

	applicationPort := mainConfigsYml.GetYmlValueByName(IDX_APPLICATION_PORT)
	//tracingServerName := mainConfigsYml.GetYmlValueByName(IDX_TRACING_SERVER_NAME)
	routes := mainGatewaysWs.GetRoutesBean()

	router := mainGatewaysWs.ConfigRoutes(mainConfigsRouterHttp.GetRouterBean(), *routes)
	//router.Use(otelmux.Middleware(tracingServerName))

	err := http.ListenAndServe(":"+applicationPort, router)
	if err != nil {
		main_utils.FailOnError(err, MSG_APPLICATION_FAILED)
	}
	log.Printf(MSG_LISTENER, applicationPort)
}
