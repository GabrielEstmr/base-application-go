package main_gateways_ws_beans_factories

import (
	gatewaysWsV1 "baseapplicationgo/main/gateways/ws/v1"
	main_usecases_beans "baseapplicationgo/main/usecases/beans"
)

func UserControllerBeanFunc() *gatewaysWsV1.UserController {
	useCase := main_usecases_beans.CreateNewUserBeanFactory()
	userController := gatewaysWsV1.NewUserController(useCase)
	return userController
}
