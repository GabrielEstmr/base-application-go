package main_gateways_rabbitmq

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_gateways "baseapplicationgo/main/gateways"
	main_gateways_logs_resources "baseapplicationgo/main/gateways/logs/resources"
	main_gateways_rabbitmq_producers "baseapplicationgo/main/gateways/rabbitmq/producers"
	main_utils_messages "baseapplicationgo/main/utils/messages"
	"context"
	"fmt"
)

const _MSG_CREATE_EVENT_TRANSACTION_DOC_ARCH_ISSUE = "exceptions.architecture.application.issue"

type TransactionEventProducerGatewayImpl struct {
	transactionProducer   main_gateways_rabbitmq_producers.RabbiMQTransactionProducer
	logsMonitoringGateway main_gateways.LogsMonitoringGateway
	messageUtils          main_utils_messages.ApplicationMessages
}

func NewTransactionEventProducerGatewayImpl(
	transactionProducer main_gateways_rabbitmq_producers.RabbiMQTransactionProducer,
	logsMonitoringGateway main_gateways.LogsMonitoringGateway,
) *TransactionEventProducerGatewayImpl {
	return &TransactionEventProducerGatewayImpl{
		transactionProducer:   transactionProducer,
		logsMonitoringGateway: logsMonitoringGateway,
		messageUtils:          *main_utils_messages.NewApplicationMessages(),
	}
}

func (this *TransactionEventProducerGatewayImpl) Send(
	ctx context.Context, transaction main_domains.Transaction) main_domains_exceptions.ApplicationException {

	span := *main_gateways_logs_resources.NewSpanLogInfo(ctx)
	defer span.End()

	this.logsMonitoringGateway.DEBUG(span,
		fmt.Sprintf("Creating new transaction event with accountId: %s", transaction.GetAccountId()))

	err := this.transactionProducer.Produce(ctx, transaction)
	if err != nil {
		return main_domains_exceptions.
			NewInternalServerErrorExceptionSglMsg(this.messageUtils.
				GetDefaultLocale(_MSG_CREATE_EVENT_TRANSACTION_DOC_ARCH_ISSUE))
	}
	return nil
}
