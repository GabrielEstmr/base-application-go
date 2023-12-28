package main_gateways_rabbitmq_listeners

import (
	main_configs_apm_logs_impl "baseapplicationgo/main/configs/apm/logs/impl"
	main_configs_error "baseapplicationgo/main/configs/error"
	main_configs_rabbitmq "baseapplicationgo/main/configs/rabbitmq"
	main_configs_rabbitmq_paramaters "baseapplicationgo/main/configs/rabbitmq/paramaters"
	main_gateways "baseapplicationgo/main/gateways"
	main_gateways_logs "baseapplicationgo/main/gateways/logs"
	main_gateways_mongodb "baseapplicationgo/main/gateways/mongodb"
	main_gateways_mongodb_repositories "baseapplicationgo/main/gateways/mongodb/repositories"
	main_gateways_rabbitmq_resources "baseapplicationgo/main/gateways/rabbitmq/resources"
	main_usecases "baseapplicationgo/main/usecases"
	"context"
	"encoding/json"
	"fmt"
	"log"
)

const _MSG_RABBITMQ_TEST_LISTENER_INSTANTIATION = "Rabbitmq test listener instantiation. Queue: %s"

func ListenTest() {

	conn := main_configs_rabbitmq.GetConnection()
	defer main_configs_rabbitmq.CloseRabbitMqConnection(conn)

	ch := main_configs_rabbitmq.GetChannel(conn)
	defer main_configs_rabbitmq.CloseRabbitMqChannel(ch)

	consumerParams := main_configs_rabbitmq_paramaters.
		AmqpConsumerParameters[main_configs_rabbitmq_paramaters.QUEUE_BASE_APP_GO_AMQP_TEST]
	log.Println(
		fmt.Sprintf(_MSG_RABBITMQ_TEST_LISTENER_INSTANTIATION, consumerParams.GetQueueName()))

	msgs, err := ch.Consume(
		consumerParams.GetQueueName(),
		consumerParams.GetConsumerTag(),
		consumerParams.GetAutoAck(),
		consumerParams.GetExclusive(),
		consumerParams.GetNoLocal(),
		consumerParams.GetNoWait(),
		consumerParams.GetArgs(),
	)
	main_configs_error.FailOnError(err, "Failed to register a consumer")

	transactionRepository := main_gateways_mongodb_repositories.NewTransactionRepository()
	var transactionDatabaseGateway main_gateways.TransactionDatabaseGateway = main_gateways_mongodb.NewTransactionDatabaseGatewayImpl(*transactionRepository)

	var logsMonitoringGateway main_gateways.LogsMonitoringGateway = main_gateways_logs.NewLogsMonitoringGatewayImpl(
		main_configs_apm_logs_impl.NewLogsGatewayImpl())

	//var spanGatewayImpl main_gateways.SpanGateway = main_gateways_spans.NewSpanGatewayImpl()

	usecase := main_usecases.NewPersistTransaction(transactionDatabaseGateway, logsMonitoringGateway)

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf(" [MESSAGE RECEIVED FROM RABBITMQ] %s", d.Body)
			//d.Headers
			var event main_gateways_rabbitmq_resources.Event
			if err = json.Unmarshal(d.Body, &event); err != nil {
				log.Fatal(err)
				return
			}

			var msg main_gateways_rabbitmq_resources.TransactionResource
			msgMap := event.Message.(map[string]interface{})
			msg.AccountId = msgMap["accountId"].(string)
			msg.OperationTypeId = msgMap["operationTypeId"].(string)
			msg.Amount = msgMap["amount"].(float64)

			ctx := context.Background()
			_, errT := usecase.Execute(ctx, msg.ToDomain())

			if errT != nil {
				log.Fatal(errT)
			}

		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}
