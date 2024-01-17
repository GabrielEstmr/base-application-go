package main_configs_rabbitmq_paramaters

import main_configs_rabbitmq_paramaters "baseapplicationgo/main/configs/rabbitmq/resources"

const (
	EXCHANGE_BASE_APP_GO_AMQP_TEST = "app-base-golang-exchange-test"
)

const (
	TOPIC_KIND = "topic"
)

var AmqpExchangeParameters = []main_configs_rabbitmq_paramaters.AmqpExchangeProperties{
	*main_configs_rabbitmq_paramaters.NewAmqpExchangeParameters(
		EXCHANGE_BASE_APP_GO_AMQP_TEST,
		TOPIC_KIND,
		true,
		false,
		false,
		false,
		nil,
	),
}
