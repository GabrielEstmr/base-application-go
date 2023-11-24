package main_gateways_ws_beans_factories

import (
	gatewaysWsV1 "baseapplicationgo/main/gateways/ws/v1"
	main_usecases_beans "baseapplicationgo/main/usecases/beans/Factories"
)

func UserControllerBeanFactory() gatewaysWsV1.UserController {
	createNewUse := main_usecases_beans.CreateNewUserBean()
	userController := gatewaysWsV1.NewUserController(createNewUse)
	return userController
}
