package main_gateways_rabbitmq_producers

import (
	main_configs_rabbitmq_paramaters "baseapplicationgo/main/configs/rabbitmq/paramaters"
	main_configs_yml "baseapplicationgo/main/configs/yml"
	main_gateways "baseapplicationgo/main/gateways"
	main_gateways_rabbitmq_resources "baseapplicationgo/main/gateways/rabbitmq/resources"
	main_gateways_spans "baseapplicationgo/main/gateways/spans"
	"context"
	"encoding/json"
	"github.com/rabbitmq/amqp091-go"
	"log"
)

const _REPROCESS_EMAIL_RABBITMQ_URI_YML_IDX = "RabbitMQ.URI"

type ReprocessEmailEventProducer struct {
	spanGateway main_gateways.SpanGateway
}

func NewReprocessEmailEventProducer() *ReprocessEmailEventProducer {
	return &ReprocessEmailEventProducer{spanGateway: main_gateways_spans.NewSpanGatewayImpl()}
}

func (this *ReprocessEmailEventProducer) Produce(
	ctx context.Context, emailId string) error {

	span := this.spanGateway.Get(ctx, "ReprocessEmailEventProducer - Produce")
	defer span.End()

	rabbitMqURI := main_configs_yml.GetYmlValueByName(_REPROCESS_EMAIL_RABBITMQ_URI_YML_IDX)
	conn, errConn := amqp091.Dial(rabbitMqURI)
	if errConn != nil {
		return errConn
	}

	ch, errCh := conn.Channel()
	if errCh != nil {
		return errCh
	}

	event := main_gateways_rabbitmq_resources.NewEvent(emailId)
	marshal, errJson := json.Marshal(event)
	if errJson != nil {
		return errJson
	}

	err := ch.PublishWithContext(span.GetCtx(),
		main_configs_rabbitmq_paramaters.EXCHANGE_RM_NOTIFICATION_REPROCESS_EMAIL,
		main_configs_rabbitmq_paramaters.ROUTING_RM_NOTIFICATION_REPROCESS_EMAIL,
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
