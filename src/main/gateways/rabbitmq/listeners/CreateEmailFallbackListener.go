package main_gateways_rabbitmq_listeners

import (
	main_configs_error "baseapplicationgo/main/configs/error"
	main_configs_rabbitmq "baseapplicationgo/main/configs/rabbitmq"
	main_configs_rabbitmq_paramaters "baseapplicationgo/main/configs/rabbitmq/paramaters"
	main_gateways "baseapplicationgo/main/gateways"
	main_gateways_logs "baseapplicationgo/main/gateways/logs"
	main_gateways_rabbitmq_resources "baseapplicationgo/main/gateways/rabbitmq/resources"
	main_gateways_spans "baseapplicationgo/main/gateways/spans"
	main_usecases "baseapplicationgo/main/usecases"
	"context"
	"fmt"
	"log"
)

const _MSG_RABBITMQ_CREATE_EMAIL_FALLBACK_LISTENER_INSTANTIATION = "Rabbitmq middlewares listener instantiation. Queue: %s"

type CreateEmailFallbackListener struct {
	createEmailFallback   main_usecases.CreateEmailFallback
	logsMonitoringGateway main_gateways.LogsMonitoringGateway
	spanGateway           main_gateways.SpanGateway
}

func NewCreateEmailFallbackListener(
	createEmailFallback main_usecases.CreateEmailFallback,
) *CreateEmailFallbackListener {
	return &CreateEmailFallbackListener{
		createEmailFallback:   createEmailFallback,
		logsMonitoringGateway: main_gateways_logs.NewLogsMonitoringGatewayImpl(),
		spanGateway:           main_gateways_spans.NewSpanGatewayImpl(),
	}
}

func NewCreateEmailFallbackListenerAllArgs(
	createEmailFallback main_usecases.CreateEmailFallback,
	logsMonitoringGateway main_gateways.LogsMonitoringGateway,
	spanGateway main_gateways.SpanGateway,
) *CreateEmailFallbackListener {
	return &CreateEmailFallbackListener{
		createEmailFallback:   createEmailFallback,
		logsMonitoringGateway: logsMonitoringGateway,
		spanGateway:           spanGateway,
	}
}

func (this *CreateEmailFallbackListener) Listen() {

	conn := main_configs_rabbitmq.GetConnection()
	defer main_configs_rabbitmq.CloseRabbitMqConnection(conn)

	ch := main_configs_rabbitmq.GetChannel(conn)
	defer main_configs_rabbitmq.CloseRabbitMqChannel(ch)

	consumerParams := main_configs_rabbitmq_paramaters.
		AmqpConsumerParameters[main_configs_rabbitmq_paramaters.QUEUE_RM_NOTIFICATION_CREATE_EMAIL_DLQ]
	log.Println(
		fmt.Sprintf(
			_MSG_RABBITMQ_CREATE_EMAIL_FALLBACK_LISTENER_INSTANTIATION,
			consumerParams.GetQueueName()))

	ctx := context.TODO()
	msgs, err := ch.ConsumeWithContext(
		ctx,
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

			span := this.spanGateway.Get(ctx, "CreateEmailListener-Listen")

			messageQueue := fmt.Sprintf(
				"Message has been received. Queue: %s MessageId: %s",
				consumerParams.GetQueueName(),
				d.MessageId,
			)

			this.logsMonitoringGateway.INFO(span, messageQueue)

			var event main_gateways_rabbitmq_resources.EventUUID
			event, errEv := event.FromJSON(d.Body)
			if errEv != nil {
				this.logsMonitoringGateway.ERROR(span, messageQueue, errEv)
				_ = d.Nack(false, false)
				return
			}

			msgB, errM := event.GetMessageJSON(d.Body)
			if errM != nil {
				this.logsMonitoringGateway.ERROR(span, messageQueue, errM)
				_ = d.Nack(false, false)
				return
			}

			var eventMessage main_gateways_rabbitmq_resources.EmailParamsResource
			eventMessage, errJ := eventMessage.FromJSON(msgB)
			if errJ != nil {
				this.logsMonitoringGateway.ERROR(span, messageQueue, err)
				_ = d.Nack(false, false)
				return
			}

			_, errE := this.createEmailFallback.Execute(span.GetCtx(), event.EventId, eventMessage.ToDomain())

			if errE != nil {
				this.logsMonitoringGateway.ERROR(span, messageQueue, errE)
				_ = d.Nack(false, false)
			}
			_ = d.Ack(true)
			span.End()
		}
	}()

	<-forever
}
