package main_configs_rabbitmq_paramaters

import main_configs_rabbitmq_paramaters "baseapplicationgo/main/configs/rabbitmq/resources"

var AmqpBindingQueue = []main_configs_rabbitmq_paramaters.AmqpBindingQueueProperties{
	*main_configs_rabbitmq_paramaters.NewAmqpBindingQueueProperties(
		QUEUE_BASE_APP_GO_AMQP_TEST,
		ROUTING_KEY_BASE_APP_GO_AMQP_TEST,
		EXCHANGE_BASE_APP_GO_AMQP_TEST,
		false,
		nil,
	),
	*main_configs_rabbitmq_paramaters.NewAmqpBindingQueueProperties(
		QUEUE_BASE_APP_GO_AMQP_DLQ_TEST,
		ROUTING_KEY_BASE_APP_GO_AMQP_DLQ_TEST,
		EXCHANGE_BASE_APP_GO_AMQP_TEST,
		false,
		nil,
	),
	*main_configs_rabbitmq_paramaters.NewAmqpBindingQueueProperties(
		QUEUE_RM_NOTIFICATION_CREATE_EMAIL,
		ROUTING_RM_NOTIFICATION_CREATE_EMAIL,
		EXCHANGE_RM_NOTIFICATION_CREATE_EMAIL,
		false,
		nil,
	),
	*main_configs_rabbitmq_paramaters.NewAmqpBindingQueueProperties(
		QUEUE_RM_NOTIFICATION_CREATE_EMAIL_DLQ,
		ROUTING_RM_NOTIFICATION_CREATE_EMAIL_DLQ,
		EXCHANGE_RM_NOTIFICATION_CREATE_EMAIL,
		false,
		nil,
	),
	*main_configs_rabbitmq_paramaters.NewAmqpBindingQueueProperties(
		QUEUE_RM_NOTIFICATION_REPROCESS_EMAIL,
		ROUTING_RM_NOTIFICATION_REPROCESS_EMAIL,
		EXCHANGE_RM_NOTIFICATION_REPROCESS_EMAIL,
		false,
		nil,
	),
}
