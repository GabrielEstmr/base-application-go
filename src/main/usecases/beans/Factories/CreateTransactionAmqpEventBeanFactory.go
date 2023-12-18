package main_usecases_beans_factories

import (
	main_configs_apm_logs_impl "baseapplicationgo/main/configs/apm/logs/impl"
	main_gateways "baseapplicationgo/main/gateways"
	main_gateways_logs "baseapplicationgo/main/gateways/logs"
	main_gateways_rabbitmq "baseapplicationgo/main/gateways/rabbitmq"
	main_gateways_rabbitmq_producers "baseapplicationgo/main/gateways/rabbitmq/producers"
	main_usecases "baseapplicationgo/main/usecases"
)

type CreateTransactionAmqpEventBean struct {
}

func NewCreateTransactionAmqpEventBean() *CreateTransactionAmqpEventBean {
	return &CreateTransactionAmqpEventBean{}
}

func (this *CreateTransactionAmqpEventBean) Get() *main_usecases.CreateTransactionAmqpEvent {

	rabbitProducer := *main_gateways_rabbitmq_producers.NewRabbiMQTransactionProducer()
	var logsMonitoringGateway main_gateways.LogsMonitoringGateway = main_gateways_logs.NewLogsMonitoringGatewayImpl(
		main_configs_apm_logs_impl.NewLogsGatewayImpl())

	var producerGateway main_gateways.TransactionEventProducerGateway = main_gateways_rabbitmq.NewTransactionEventProducerGatewayImpl(rabbitProducer, logsMonitoringGateway)

	return main_usecases.NewCreateTransactionAmqpEvent(producerGateway, logsMonitoringGateway)
}
