package main_gateways_ws_beans_factories

import (
	main_gateways_ws_v1 "baseapplicationgo/main/gateways/ws/v1"
	main_usecases_beans "baseapplicationgo/main/usecases/beans"
)

type EmailControllerBean struct {
}

func NewEmailControllerBean() *EmailControllerBean {
	return &EmailControllerBean{}
}

func (this *EmailControllerBean) Get() *main_gateways_ws_v1.EmailController {
	usecaseBeans := main_usecases_beans.GetUsecaseBeans()
	userController := *main_gateways_ws_v1.NewEmailController(
		usecaseBeans.SendEmailEventsToReprocess,
		usecaseBeans.FindEmailsByFilter,
	)
	return &userController
}
