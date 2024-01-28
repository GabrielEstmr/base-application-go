package main_gateways_rabbitmq_beans_factories

import (
	main_gateways_rabbitmq_listeners "baseapplicationgo/main/gateways/rabbitmq/listeners"
	main_usecases_beans "baseapplicationgo/main/usecases/beans"
)

type CreateEmailListener struct {
}

func NewCreateEmailListener() *CreateEmailListener {
	return &CreateEmailListener{}
}

func (this *CreateEmailListener) Get() *main_gateways_rabbitmq_listeners.CreateEmailListener {
	usecaseBeans := main_usecases_beans.GetUsecaseBeans()
	return main_gateways_rabbitmq_listeners.NewCreateEmailListener(*usecaseBeans.CreateEmail)
}
