package main_gateways_rabbitmq_listeners

import (
	"log"
	main_configurations_rabbitmq "mpindicator/main/configurations/rabbitmq"
	main_configurations_yml "mpindicator/main/configurations/yml"
	main_utils "mpindicator/main/utils"

	"github.com/rabbitmq/amqp091-go"
)

const MSG_RABBITMQ_REGISTER_CONSUMER_FAILURE = "Failed to register a consumer"
const MSG_RABBITMQ_OPEN_CHANNEL_FAILURE = "Failed to open a channel"
const MSG_RABBITMQ_CONNECT_FAILURE = "Failed to connect to RabbitMQ"
const RABBITMQ_URI = "RabbitMQ.URI"

func Listen() {

	rabbitMqURI := main_configurations_yml.GetBeanPropertyByName(RABBITMQ_URI)
	log.Println("========> rabbitMqURI LISTENER", rabbitMqURI)

	conn, err := amqp091.Dial(rabbitMqURI)
	main_utils.FailOnError(err, MSG_RABBITMQ_CONNECT_FAILURE)
	defer conn.Close()

	ch, err := conn.Channel()
	main_utils.FailOnError(err, MSG_RABBITMQ_CONNECT_FAILURE)
	defer ch.Close()

	log.Println("AQUI ===================>", main_configurations_rabbitmq.AmqpQueues["mp-indicator-aqmpq-test"])

	msgs, err := ch.Consume(
		"mp-indicator-aqmpq-test", // queue
		"",                        // consumer
		true,                      // auto ack
		false,                     // exclusive
		false,                     // no local
		false,                     // no wait
		nil,                       // args
	)
	main_utils.FailOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf(" [x] %s", d.Body)
			//d.Headers
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}
