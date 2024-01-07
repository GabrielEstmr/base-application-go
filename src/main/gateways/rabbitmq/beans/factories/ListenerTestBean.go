package main_gateways_rabbitmq_beans_factories

import (
	main_gateways "baseapplicationgo/main/gateways"
	main_gateways_rabbitmq_listeners "baseapplicationgo/main/gateways/rabbitmq/listeners"
	main_gateways_spans "baseapplicationgo/main/gateways/spans"
	main_usecases_beans "baseapplicationgo/main/usecases/beans"
)

type ListenerTestBean struct {
}

func NewListenerTestBean() *ListenerTestBean {
	return &ListenerTestBean{}
}

func (this *ListenerTestBean) Get() *main_gateways_rabbitmq_listeners.ListenerTest {
	var spanGatewayImpl main_gateways.SpanGateway = main_gateways_spans.NewSpanGatewayImpl()
	usecaseBeans := main_usecases_beans.GetUsecaseBeans()
	return main_gateways_rabbitmq_listeners.NewListenerTest(*usecaseBeans.PersistTransaction, spanGatewayImpl)
}
