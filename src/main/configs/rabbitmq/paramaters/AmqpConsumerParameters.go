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
	QUEUE_RM_NOTIFICATION_CREATE_EMAIL: *main_configs_rabbitmq_paramaters.NewAmqpConsumerProperties(
		QUEUE_RM_NOTIFICATION_CREATE_EMAIL,
		"",
		false,
		false,
		false,
		false,
		nil,
	),
	QUEUE_RM_NOTIFICATION_CREATE_EMAIL_DLQ: *main_configs_rabbitmq_paramaters.NewAmqpConsumerProperties(
		QUEUE_RM_NOTIFICATION_CREATE_EMAIL_DLQ,
		"",
		false,
		false,
		false,
		false,
		nil,
	),
	QUEUE_RM_NOTIFICATION_REPROCESS_EMAIL: *main_configs_rabbitmq_paramaters.NewAmqpConsumerProperties(
		QUEUE_RM_NOTIFICATION_REPROCESS_EMAIL,
		"",
		false,
		false,
		false,
		false,
		nil,
	),
}
