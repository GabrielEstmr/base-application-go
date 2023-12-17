package main_gateways_rabbitmq_listeners

import (
	main_configs_error "baseapplicationgo/main/configs/error"
	main_configs_rabbitmq "baseapplicationgo/main/configs/rabbitmq"
	main_configs_rabbitmq_paramaters "baseapplicationgo/main/configs/rabbitmq/paramaters"
	main_configs_yml "baseapplicationgo/main/configs/yml"
	"log"
)

const _RABBITMQ_URI_YML_IDX = "RabbitMQ.URI"

func ListenTest() {

	rabbitMqURI := main_configs_yml.GetYmlValueByName(_RABBITMQ_URI_YML_IDX)
	log.Println("========> rabbitMqURI LISTENER", rabbitMqURI)

	conn := main_configs_rabbitmq.GetConnection()
	defer main_configs_rabbitmq.CloseRabbitMqConnection(conn)

	ch := main_configs_rabbitmq.GetChannel(conn)
	defer main_configs_rabbitmq.CloseRabbitMqChannel(ch)

	consumerParams := main_configs_rabbitmq_paramaters.AmqpConsumerParameters[main_configs_rabbitmq_paramaters.QUEUE_BASE_APP_GO_AMQP_TEST]
	msgs, err := ch.Consume(
		consumerParams.GetQueueName(),
		consumerParams.GetConsumerTag(),
		consumerParams.GetAutoAck(),
		consumerParams.GetExclusive(),
		consumerParams.GetNoLocal(),
		consumerParams.GetNoWait(),
		consumerParams.GetArgs(),
	)
	main_configs_error.FailOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf(" [MESSAGE RECEIVED FROM RABBITMQ] %s", d.Body)
			//d.Headers
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}
