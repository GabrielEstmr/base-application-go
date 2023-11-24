package main_gateways_ws_beans_factories

import gatewaysWsV1 "baseapplicationgo/main/gateways/ws/v1"

func accountControllerBeanFunc() *gatewaysWsV1.AccountController {
	//useCase := main_usecases_beans.GetBeans().CreateUserBean
	userController := gatewaysWsV1.NewIndicatorController()
	return userController
}
