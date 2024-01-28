package main_configs_rabbitmq_paramaters

import main_configs_rabbitmq_paramaters "baseapplicationgo/main/configs/rabbitmq/resources"

const (
	EXCHANGE_BASE_APP_GO_AMQP_TEST           = "app-base-golang-exchange-test"
	EXCHANGE_RM_NOTIFICATION_CREATE_EMAIL    = "app-rm-notification-exchange-create-email"
	EXCHANGE_RM_NOTIFICATION_REPROCESS_EMAIL = "app-rm-notification-exchange-reprocess-email"
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
	*main_configs_rabbitmq_paramaters.NewAmqpExchangeParameters(
		EXCHANGE_RM_NOTIFICATION_CREATE_EMAIL,
		TOPIC_KIND,
		true,
		false,
		false,
		false,
		nil,
	),
	*main_configs_rabbitmq_paramaters.NewAmqpExchangeParameters(
		EXCHANGE_RM_NOTIFICATION_REPROCESS_EMAIL,
		TOPIC_KIND,
		true,
		false,
		false,
		false,
		nil,
	),
}
