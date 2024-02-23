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

type EmailEventProducer struct {
	spanGateway main_gateways.SpanGateway
}

func NewEmailEventProducer() *EmailEventProducer {
	return &EmailEventProducer{
		spanGateway: main_gateways_spans.NewSpanGatewayImpl(),
	}
}

func (this *EmailEventProducer) Produce(
	ctx context.Context, emailParamsResource main_gateways_rabbitmq_resources.EmailParamsResource) error {

	span := this.spanGateway.Get(ctx, "EmailEventProducer - Produce")
	defer span.End()

	rabbitMqURI := main_configs_yml.GetYmlValueByName(main_configs_yml.RabbitMQURI)
	conn, errConn := amqp091.Dial(rabbitMqURI)
	if errConn != nil {
		return errConn
	}

	ch, errCh := conn.Channel()
	if errCh != nil {
		return errCh
	}

	event := main_gateways_rabbitmq_resources.NewEvent(emailParamsResource)
	marshal, errJson := json.Marshal(event)
	if errJson != nil {
		return errJson
	}

	err := ch.PublishWithContext(span.GetCtx(),
		main_configs_rabbitmq_paramaters.EXCHANGE_RM_NOTIFICATION_CREATE_EMAIL,
		main_configs_rabbitmq_paramaters.ROUTING_RM_NOTIFICATION_CREATE_EMAIL,
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
