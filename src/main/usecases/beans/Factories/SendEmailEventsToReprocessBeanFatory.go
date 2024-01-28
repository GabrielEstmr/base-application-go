package main_usecases_beans_factories

import (
	main_gateways "baseapplicationgo/main/gateways"
	main_gateways_logs "baseapplicationgo/main/gateways/logs"
	main_gateways_rabbitmq "baseapplicationgo/main/gateways/rabbitmq"
	main_gateways_rabbitmq_producers "baseapplicationgo/main/gateways/rabbitmq/producers"
	main_gateways_spans "baseapplicationgo/main/gateways/spans"
	main_usecases "baseapplicationgo/main/usecases"
	main_utils_messages "baseapplicationgo/main/utils/messages"
)

type SendEmailEventsToReprocessBean struct {
}

func NewSendEmailEventsToReprocessBean() *SendEmailEventsToReprocessBean {
	return &SendEmailEventsToReprocessBean{}
}

func (this *SendEmailEventsToReprocessBean) Get() *main_usecases.SendEmailEventsToReprocess {

	producer := *main_gateways_rabbitmq_producers.NewReprocessEmailEventProducer()

	var logsMonitoringGateway main_gateways.ReprocessEmailEventProducerGateway = main_gateways_rabbitmq.NewReprocessEmailEventProducerGatewayImpl(
		producer)

	return main_usecases.NewSendEmailEventsToReprocessAllArgs(
		logsMonitoringGateway,
		main_gateways_logs.NewLogsMonitoringGatewayImpl(),
		main_gateways_spans.NewSpanGatewayImpl(),
		*main_utils_messages.NewApplicationMessages(),
	)
}
