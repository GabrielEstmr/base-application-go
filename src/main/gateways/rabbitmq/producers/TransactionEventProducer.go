package main_gateways_rabbitmq_producers

import (
	main_configs_rabbitmq_paramaters "baseapplicationgo/main/configs/rabbitmq/paramaters"
	main_configs_yml "baseapplicationgo/main/configs/yml"
	main_domains "baseapplicationgo/main/domains"
	main_gateways "baseapplicationgo/main/gateways"
	main_gateways_rabbitmq_resources "baseapplicationgo/main/gateways/rabbitmq/resources"
	"context"
	"encoding/json"
	"github.com/rabbitmq/amqp091-go"
	"log"
)

const _RABBITMQ_URI_YML_IDX = "RabbitMQ.URI"

type RabbiMQTransactionProducer struct {
	spanGateway main_gateways.SpanGateway
}

func NewRabbiMQTransactionProducer(spanGateway main_gateways.SpanGateway,
) *RabbiMQTransactionProducer {
	return &RabbiMQTransactionProducer{spanGateway}
}

func (this *RabbiMQTransactionProducer) Produce(
	ctx context.Context, transaction main_domains.Transaction) error {

	span := this.spanGateway.Get(ctx, "RabbiMQTransactionProducer - Produce")
	defer span.End()

	rabbitMqURI := main_configs_yml.GetYmlValueByName(_RABBITMQ_URI_YML_IDX)
	conn, errConn := amqp091.Dial(rabbitMqURI)
	if errConn != nil {
		return errConn
	}

	ch, errCh := conn.Channel()
	if errCh != nil {
		return errCh
	}

	event := main_gateways_rabbitmq_resources.NewEvent(*main_gateways_rabbitmq_resources.NewTransactionResource(transaction))
	marshal, errJson := json.Marshal(event)
	if errJson != nil {
		return errJson
	}

	err := ch.PublishWithContext(span.GetCtx(),
		main_configs_rabbitmq_paramaters.EXCHANGE_BASE_APP_GO_AMQP_TEST,    // exchange
		main_configs_rabbitmq_paramaters.ROUTING_KEY_BASE_APP_GO_AMQP_TEST, // routing key
		false,
		false,
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        marshal,
		})
	if err != nil {
		return err
	}

	log.Printf("Sent %s", marshal)
	return nil
}
