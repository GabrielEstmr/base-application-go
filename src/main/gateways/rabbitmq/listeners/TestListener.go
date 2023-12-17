package main_gateways_rabbitmq_listeners

import (
	main_configs_error "baseapplicationgo/main/configs/error"
	main_configs_rabbitmq "baseapplicationgo/main/configs/rabbitmq"
	main_configs_rabbitmq_paramaters "baseapplicationgo/main/configs/rabbitmq/paramaters"
	main_configs_yml "baseapplicationgo/main/configs/yml"
	"github.com/rabbitmq/amqp091-go"
	"log"
)

const _MSG_RABBITMQ_CONNECT_FAILURE = "Failed to connect to RabbitMQ"
const _MSG_RABBITMQ_OPEN_CHANNEL_FAILURE = "Failed to open a channel"
const _RABBITMQ_URI_YML_IDX = "RabbitMQ.URI"

func ListenTest() {

	rabbitMqURI := main_configs_yml.GetYmlValueByName(_RABBITMQ_URI_YML_IDX)
	log.Println("========> rabbitMqURI LISTENER", rabbitMqURI)

	conn, err := amqp091.Dial(rabbitMqURI)
	main_configs_error.FailOnError(err, _MSG_RABBITMQ_CONNECT_FAILURE)
	defer main_configs_rabbitmq.CloseRabbitMqConnection(conn)

	ch, err := conn.Channel()
	main_configs_error.FailOnError(err, _MSG_RABBITMQ_OPEN_CHANNEL_FAILURE)
	defer main_configs_rabbitmq.CloseRabbitMqChannel(ch)

	//log.Println("AQUI ===================>", main_configurations_rabbitmq.AmqpQueues["mp-indicator-aqmpq-test"])

	msgs, err := ch.Consume(
		main_configs_rabbitmq_paramaters.QUEUE_BASE_APP_GO_AMQP_TEST, // queue
		"",    // consumer
		true,  // auto ack
		false, // exclusive
		false, // no local
		false, // no wait
		nil,   // args
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
