package main_configs_rabbitmq_paramaters

import main_configs_rabbitmq_paramaters "baseapplicationgo/main/configs/rabbitmq/resources"

const (
	QUEUE_BASE_APP_GO_AMQP_TEST     = "app-base-golang-queue-test"
	QUEUE_BASE_APP_GO_AMQP_DLQ_TEST = "app-base-golang-queue-dlq-test"
)

var AmqpQueueParameters = []main_configs_rabbitmq_paramaters.AmqpQueueProperties{
	*main_configs_rabbitmq_paramaters.NewAmqpQueueParameters(
		QUEUE_BASE_APP_GO_AMQP_TEST,
		true,
		false,
		false,
		false,
		map[string]any{
			"x-message-ttl":             10000,
			"x-dead-letter-exchange":    EXCHANGE_BASE_APP_GO_AMQP_TEST,
			"x-dead-letter-routing-key": ROUTING_KEY_BASE_APP_GO_AMQP_DLQ_TEST,
		},
	),
	*main_configs_rabbitmq_paramaters.NewAmqpQueueParameters(
		QUEUE_BASE_APP_GO_AMQP_DLQ_TEST,
		true,
		false,
		false,
		false,
		map[string]any{
			"x-dead-letter-exchange": EXCHANGE_BASE_APP_GO_AMQP_TEST,
		},
	),
}
