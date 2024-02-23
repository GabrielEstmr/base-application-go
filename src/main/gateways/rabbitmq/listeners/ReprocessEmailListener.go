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
	"encoding/json"
	"fmt"
	"log"
)

const _MSG_RABBITMQ_REPROCESS_EMAIL_LISTENER_INSTANTIATION = "Rabbitmq middlewares listener instantiation. Queue: %s"

type ReprocessEmailListener struct {
	reprocessEmailEvent   main_usecases.ReprocessEmailEvent
	logsMonitoringGateway main_gateways.LogsMonitoringGateway
	spanGateway           main_gateways.SpanGateway
}

func NewReprocessEmailListener(
	reprocessEmailEvent main_usecases.ReprocessEmailEvent,
) *ReprocessEmailListener {
	return &ReprocessEmailListener{
		reprocessEmailEvent:   reprocessEmailEvent,
		logsMonitoringGateway: main_gateways_logs.NewLogsMonitoringGatewayImpl(),
		spanGateway:           main_gateways_spans.NewSpanGatewayImpl(),
	}
}

func NewReprocessEmailListenerAllArgs(
	createEmail main_usecases.CreateEmail,
	logsMonitoringGateway main_gateways.LogsMonitoringGateway,
	spanGateway main_gateways.SpanGateway,
) *CreateEmailListener {
	return &CreateEmailListener{
		createEmail:           createEmail,
		logsMonitoringGateway: logsMonitoringGateway,
		spanGateway:           spanGateway,
	}
}

func (this *ReprocessEmailListener) Listen() {

	conn := main_configs_rabbitmq.GetConnection()
	defer main_configs_rabbitmq.CloseRabbitMqConnection(conn)

	ch := main_configs_rabbitmq.GetChannel(conn)
	defer main_configs_rabbitmq.CloseRabbitMqChannel(ch)

	consumerParams := main_configs_rabbitmq_paramaters.
		AmqpConsumerParameters[main_configs_rabbitmq_paramaters.QUEUE_RM_NOTIFICATION_REPROCESS_EMAIL]
	log.Println(
		fmt.Sprintf(_MSG_RABBITMQ_REPROCESS_EMAIL_LISTENER_INSTANTIATION, consumerParams.GetQueueName()))

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

			span := this.spanGateway.Get(ctx, "ReprocessEmailListener-Listen")

			messageQueue := fmt.Sprintf(
				"Message has been received. Queue: %s MessageId: %s",
				consumerParams.GetQueueName(),
				d.MessageId,
			)

			this.logsMonitoringGateway.DEBUG(span, messageQueue)

			var event main_gateways_rabbitmq_resources.Event
			msgB, errM := event.GetMessageJSON(d.Body)
			if errM != nil {
				this.logsMonitoringGateway.ERROR(span, messageQueue, errM)
				_ = d.Ack(false)
				return
			}

			var eventMessage string
			if errU := json.Unmarshal(msgB, &eventMessage); errU != nil {
				this.logsMonitoringGateway.ERROR(span, messageQueue, errU)
				_ = d.Ack(false)
				return
			}

			_, errE := this.reprocessEmailEvent.Execute(span.GetCtx(), eventMessage)
			if errE != nil {
				this.logsMonitoringGateway.ERROR(span, messageQueue, errE)
				_ = d.Ack(true)
				return
			}
			_ = d.Ack(true)
			span.End()
		}
	}()

	<-forever
}
