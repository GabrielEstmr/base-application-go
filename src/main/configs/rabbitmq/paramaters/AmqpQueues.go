package main_configs_rabbitmq_paramaters

import main_configs_rabbitmq_paramaters "baseapplicationgo/main/configs/rabbitmq/resources"

const (
	QUEUE_BASE_APP_GO_AMQP_TEST = "QueueBaseAppAmqpTest"
)

var AmqpQueueParameters = []main_configs_rabbitmq_paramaters.AmqpQueueProperties{
	*main_configs_rabbitmq_paramaters.NewAmqpQueueParameters(
		QUEUE_BASE_APP_GO_AMQP_TEST,
		true,
		false,
		false,
		false,
		nil,
	),
}
