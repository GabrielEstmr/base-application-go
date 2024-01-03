package main_gateways_rabbitmq_listeners

import (
	main_configs_error "baseapplicationgo/main/configs/error"
	main_configs_rabbitmq "baseapplicationgo/main/configs/rabbitmq"
	main_configs_rabbitmq_paramaters "baseapplicationgo/main/configs/rabbitmq/paramaters"
	main_gateways "baseapplicationgo/main/gateways"
	main_gateways_logs "baseapplicationgo/main/gateways/logs"
	main_gateways_rabbitmq_resources "baseapplicationgo/main/gateways/rabbitmq/resources"
	main_usecases "baseapplicationgo/main/usecases"
	"context"
	"encoding/json"
	"fmt"
	"log"
)

const _MSG_RABBITMQ_TEST_LISTENER_INSTANTIATION = "Rabbitmq test listener instantiation. Queue: %s"

type ListenerTest struct {
	persistTransaction    main_usecases.PersistTransaction
	logsMonitoringGateway main_gateways.LogsMonitoringGateway
	spanGateway           main_gateways.SpanGateway
}

func NewListenerTest(
	persistTransaction main_usecases.PersistTransaction,
	spanGateway main_gateways.SpanGateway) *ListenerTest {
	return &ListenerTest{
		persistTransaction:    persistTransaction,
		logsMonitoringGateway: main_gateways_logs.NewLogsMonitoringGatewayImpl(),
		spanGateway:           spanGateway}
}

func (this *ListenerTest) Listen() {

	conn := main_configs_rabbitmq.GetConnection()
	defer main_configs_rabbitmq.CloseRabbitMqConnection(conn)

	ch := main_configs_rabbitmq.GetChannel(conn)
	defer main_configs_rabbitmq.CloseRabbitMqChannel(ch)

	consumerParams := main_configs_rabbitmq_paramaters.
		AmqpConsumerParameters[main_configs_rabbitmq_paramaters.QUEUE_BASE_APP_GO_AMQP_TEST]
	log.Println(
		fmt.Sprintf(_MSG_RABBITMQ_TEST_LISTENER_INSTANTIATION, consumerParams.GetQueueName()))

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

			span := this.spanGateway.Get(ctx, "ListenerTest-Listen")

			this.logsMonitoringGateway.INFO(span, fmt.Sprintf(
				"Message has been received. Queue: %s MessageId: %s",
				consumerParams.GetQueueName(),
				d.MessageId,
			))

			var event main_gateways_rabbitmq_resources.Event
			if err = json.Unmarshal(d.Body, &event); err != nil {
				log.Fatal(err)
				return
			}

			msgMap := event.Message.(map[string]interface{})
			msg := main_gateways_rabbitmq_resources.NewTransactionResourceFromProps(msgMap)
			_, errT := this.persistTransaction.Execute(span.GetCtx(), msg.ToDomain())
			if errT != nil {
				log.Fatal(errT)
			}
			span.End()
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}
