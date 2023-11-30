package main_gateways_ws_beans_factories

import (
	main_gateways_ws_v1 "baseapplicationgo/main/gateways/ws/v1"
	main_usecases_beans "baseapplicationgo/main/usecases/beans"
)

func UserControllerBeanFactory() *main_gateways_ws_v1.UserController {
	usecaseBeans := main_usecases_beans.GetUsecaseBeans()
	userController := *main_gateways_ws_v1.NewUserController(usecaseBeans.CreateNewUser, usecaseBeans.FindUserById)
	return &userController
}
