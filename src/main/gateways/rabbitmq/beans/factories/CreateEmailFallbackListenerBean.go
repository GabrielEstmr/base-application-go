package main_gateways_rabbitmq_beans_factories

import (
	main_gateways_rabbitmq_listeners "baseapplicationgo/main/gateways/rabbitmq/listeners"
	main_usecases_beans "baseapplicationgo/main/usecases/beans"
)

type CreateEmailFallbackListenerBean struct {
}

func NewCreateEmailFallbackListenerBean() *CreateEmailFallbackListenerBean {
	return &CreateEmailFallbackListenerBean{}
}

func (this *CreateEmailFallbackListenerBean) Get() *main_gateways_rabbitmq_listeners.CreateEmailFallbackListener {
	usecaseBeans := main_usecases_beans.GetUsecaseBeans()
	return main_gateways_rabbitmq_listeners.NewCreateEmailFallbackListener(*usecaseBeans.CreateEmailFallback)
}
