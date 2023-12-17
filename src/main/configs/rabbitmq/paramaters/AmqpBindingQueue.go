package main_configs_rabbitmq_paramaters

import main_configs_rabbitmq_paramaters "baseapplicationgo/main/configs/rabbitmq/resources"

var AmqpBindingQueue = []main_configs_rabbitmq_paramaters.AmqpBindingQueueProperties{
	*main_configs_rabbitmq_paramaters.NewAmqpBindingQueueProperties(
		QUEUE_BASE_APP_GO_AMQP_TEST,
		ROUTING_KEY_BASE_APP_GO_AMQP_TEST,
		EXCHANGE_BASE_APP_GO_AMQP_TEST,
		false,
		false,
		false,
		false,
	),
}
