package main_gateways_ws_beans_factories

import (
	gatewaysWsV1 "baseapplicationgo/main/gateways/ws/v1"
	main_usecases_beans "baseapplicationgo/main/usecases/beans"
)

func UserControllerBeanFactory() gatewaysWsV1.UserController {
	usecaseBeans := main_usecases_beans.GetUsecaseBeans()
	userController := gatewaysWsV1.NewUserController(usecaseBeans.CreateNewUser)
	return userController
}
