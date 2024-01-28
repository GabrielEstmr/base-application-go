package main_gateways_rabbitmq_beans

import (
	main_gateways_rabbitmq_beans_factories "baseapplicationgo/main/gateways/rabbitmq/beans/factories"
	main_gateways_rabbitmq_listeners "baseapplicationgo/main/gateways/rabbitmq/listeners"
	"sync"
)

var once sync.Once

var rabbitMqListenerBeans *RabbitMqListenerBeans

type RabbitMqListenerBeans struct {
	ListenerTestBean            *main_gateways_rabbitmq_listeners.ListenerTest
	CreateEmailListener         *main_gateways_rabbitmq_listeners.CreateEmailListener
	ReprocessEmailListener      *main_gateways_rabbitmq_listeners.ReprocessEmailListener
	CreateEmailFallbackListener *main_gateways_rabbitmq_listeners.CreateEmailFallbackListener
}

func GetRabbitMqListenerBeans() *RabbitMqListenerBeans {
	once.Do(func() {
		if rabbitMqListenerBeans == nil {
			rabbitMqListenerBeans = subscriptRabbitMqListenerBeans()
		}
	})
	return rabbitMqListenerBeans
}

func subscriptRabbitMqListenerBeans() *RabbitMqListenerBeans {
	return &RabbitMqListenerBeans{
		ListenerTestBean:            main_gateways_rabbitmq_beans_factories.NewListenerTestBean().Get(),
		CreateEmailListener:         main_gateways_rabbitmq_beans_factories.NewCreateEmailListener().Get(),
		ReprocessEmailListener:      main_gateways_rabbitmq_beans_factories.NewReprocessEmailListener().Get(),
		CreateEmailFallbackListener: main_gateways_rabbitmq_beans_factories.NewCreateEmailFallbackListenerBean().Get(),
	}
}
