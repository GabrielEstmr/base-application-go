package main_gateways_rabbitmq_beans_factories

import (
	main_gateways_rabbitmq_listeners "baseapplicationgo/main/gateways/rabbitmq/listeners"
	main_usecases_beans "baseapplicationgo/main/usecases/beans"
)

type ListenerTestBean struct {
}

func NewListenerTestBean() *ListenerTestBean {
	return &ListenerTestBean{}
}

func (this *ListenerTestBean) Get() *main_gateways_rabbitmq_listeners.ListenerTest {
	usecaseBeans := main_usecases_beans.GetUsecaseBeans()
	test := main_gateways_rabbitmq_listeners.NewListenerTest(*usecaseBeans.PersistTransaction)
	return test
}
