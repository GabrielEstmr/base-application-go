package main_gateways_rabbitmq

import "baseapplicationgo/main/gateways/rabbitmq/listeners"

func SubscribeListeners() {
	go main_gateways_rabbitmq_listeners.ListenTest()
}
