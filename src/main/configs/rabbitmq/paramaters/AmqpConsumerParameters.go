package main_configs_rabbitmq_paramaters

import main_configs_rabbitmq_paramaters "baseapplicationgo/main/configs/rabbitmq/resources"

var AmqpConsumerParameters = map[string]main_configs_rabbitmq_paramaters.AmqpConsumerProperties{
	QUEUE_BASE_APP_GO_AMQP_TEST: *main_configs_rabbitmq_paramaters.NewAmqpConsumerProperties(
		QUEUE_BASE_APP_GO_AMQP_TEST,
		"",
		false,
		false,
		false,
		false,
		nil,
	),
}
