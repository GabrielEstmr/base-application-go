package main_configs_rabbitmq

import (
	main_configs_error "baseapplicationgo/main/configs/error"
	main_configs_rabbitmq_paramaters "baseapplicationgo/main/configs/rabbitmq/paramaters"
	main_configs_yml "baseapplicationgo/main/configs/yml"
	"fmt"
	"github.com/rabbitmq/amqp091-go"
	"log"
)

const _MSG_TRYING_TO_CONNECT_TO_RABBITMQ = "Trying to connect to rabbitmq in the url: %s"
const _MSG_RABBITMQ_DECLARE_EXCHANGE_FAILURE = "Failed to declare an exchange"
const _MSG_RABBITMQ_CONNECT_FAILURE = "Failed to connect to RabbitMQ"
const _MSG_RABBITMQ_OPEN_CHANNEL_FAILURE = "Failed to open a channel"
const _MSG_RABBITMQ_CLOSE_CONNECT_FAILURE = "Failed to close RabbitMQ connection"
const _MSG_RABBITMQ_CLOSE_CHANNEL_FAILURE = "Failed to close RabbitMQ channel"
const _MSG_RABBITMQ_DECLARE_QUEUE_FAILURE = "Failed to declare a queue"

const _RABBITMQ_URI_YML_IDX = "RabbitMQ.URI"

func CloseRabbitMqConnection(conn *amqp091.Connection) {
	err := conn.Close()
	main_configs_error.FailOnError(err, _MSG_RABBITMQ_CLOSE_CONNECT_FAILURE)
}

func CloseRabbitMqChannel(ch *amqp091.Channel) {
	err := ch.Close()
	main_configs_error.FailOnError(err, _MSG_RABBITMQ_CLOSE_CHANNEL_FAILURE)
}

func GetConnection() *amqp091.Connection {
	rabbitMqURI := main_configs_yml.GetYmlValueByName(_RABBITMQ_URI_YML_IDX)
	log.Println(fmt.Sprintf(_MSG_TRYING_TO_CONNECT_TO_RABBITMQ, rabbitMqURI))

	conn, err := amqp091.Dial(rabbitMqURI)
	main_configs_error.FailOnError(err, _MSG_RABBITMQ_CONNECT_FAILURE)
	return conn
}

func GetChannel(conn *amqp091.Connection) *amqp091.Channel {
	ch, err := conn.Channel()
	main_configs_error.FailOnError(err, _MSG_RABBITMQ_OPEN_CHANNEL_FAILURE)
	return ch
}

func SetAmqpConfig() {
	rabbitMqURI := main_configs_yml.GetYmlValueByName(_RABBITMQ_URI_YML_IDX)
	log.Println(fmt.Sprintf(_MSG_TRYING_TO_CONNECT_TO_RABBITMQ, rabbitMqURI))

	conn := GetConnection()
	defer CloseRabbitMqConnection(conn)

	ch := GetChannel(conn)
	defer CloseRabbitMqChannel(ch)

	declareQueues(ch)
	declareExchanges(ch)
	declareBindingQueues(ch)
}

func declareQueues(ch *amqp091.Channel) {
	rabbitMQQueuesParameters := main_configs_rabbitmq_paramaters.AmqpQueueParameters
	for _, value := range rabbitMQQueuesParameters {
		_, err := ch.QueueDeclare(
			value.GetQueueName(),        // name
			value.GetDurable(),          // durable
			value.GetDeleteWhenUnused(), // delete when unused
			value.GetExclusive(),        // exclusive
			value.GetNowait(),           // no-wait
			nil,                         // arguments
		)
		main_configs_error.FailOnError(err, _MSG_RABBITMQ_DECLARE_QUEUE_FAILURE)
	}
}

func declareExchanges(ch *amqp091.Channel) {
	exchanges := main_configs_rabbitmq_paramaters.AmqpExchangeParameters
	for _, value := range exchanges {
		err := ch.ExchangeDeclare(
			value.GetExchangeName(), // name
			value.GetKind(),         // type
			value.GetDurable(),      // durable
			value.GetAutoDelete(),   // auto-deleted
			value.GetInternal(),     // internal
			value.GetNowait(),       // no-wait
			value.GetArgs(),         // arguments
		)
		main_configs_error.FailOnError(err, _MSG_RABBITMQ_DECLARE_EXCHANGE_FAILURE)
	}
}

func declareBindingQueues(ch *amqp091.Channel) {
	bindingParameters := main_configs_rabbitmq_paramaters.AmqpBindingQueue
	for _, value := range bindingParameters {
		err := ch.QueueBind(
			value.GetQueueName(),  // queue name
			value.GetRoutingKey(), // routing key
			value.GetExchange(),   // exchange
			value.GetNowait(),
			nil)
		main_configs_error.FailOnError(err, _MSG_RABBITMQ_DECLARE_EXCHANGE_FAILURE)
	}
}
