package main_gateways_rabbitmq_subscribers

import (
	main_gateways_rabbitmq_beans "baseapplicationgo/main/gateways/rabbitmq/beans"
)

func SubscribeListeners() {
	go main_gateways_rabbitmq_beans.GetRabbitMqListenerBeans().ListenerTestBean.Listen()
	go main_gateways_rabbitmq_beans.GetRabbitMqListenerBeans().CreateEmailListener.Listen()
	go main_gateways_rabbitmq_beans.GetRabbitMqListenerBeans().ReprocessEmailListener.Listen()
	go main_gateways_rabbitmq_beans.GetRabbitMqListenerBeans().CreateEmailFallbackListener.Listen()
}
