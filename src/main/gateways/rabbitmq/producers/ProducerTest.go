package main_gateways_rabbitmq_producers

import (
	main_configs_error "baseapplicationgo/main/configs/error"
	main_configs_rabbitmq "baseapplicationgo/main/configs/rabbitmq"
	main_configs_yml "baseapplicationgo/main/configs/yml"
	"context"
	"github.com/rabbitmq/amqp091-go"
	"log"
	"os"
	"strings"
)

const _MSG_RABBITMQ_CONNECT_FAILURE = "Failed to connect to RabbitMQ"
const _MSG_RABBITMQ_OPEN_CHANNEL_FAILURE = "Failed to open a channel"
const _MSG_RABBITMQ_PUBLISH_MESSAGE_FAILURE = "Failed to publish a message"

const _RABBITMQ_URI_YML_IDX = "RabbitMQ.URI"

func Produce(ctx *context.Context) {

	rabbitMqURI := main_configs_yml.GetYmlValueByName(_RABBITMQ_URI_YML_IDX)

	conn, err := amqp091.Dial(rabbitMqURI)
	main_configs_error.FailOnError(err, _MSG_RABBITMQ_CONNECT_FAILURE)
	defer main_configs_rabbitmq.CloseRabbitMqConnection(conn)

	ch, err := conn.Channel()
	main_configs_error.FailOnError(err, _MSG_RABBITMQ_OPEN_CHANNEL_FAILURE)
	defer main_configs_rabbitmq.CloseRabbitMqChannel(ch)

	body := bodyFrom(os.Args)
	err = ch.PublishWithContext(*ctx,
		"logs_topic",     // exchange
		"AmqpRoutingKey", // routing key
		false,            // mandatory
		false,            // immediate
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	main_configs_error.FailOnError(err, _MSG_RABBITMQ_PUBLISH_MESSAGE_FAILURE)

	log.Printf(" [x] Sent %s", body)
}

func bodyFrom(args []string) string {
	var s string
	if (len(args) < 3) || os.Args[2] == "" {
		s = "OIEEEE"
	} else {
		s = strings.Join(args[2:], " ")
	}
	return s
}
