package main_usecases

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_gateways "baseapplicationgo/main/gateways"
	main_gateways_logs_resources "baseapplicationgo/main/gateways/logs/resources"
	main_utils_messages "baseapplicationgo/main/utils/messages"
	"context"
	"fmt"
)

type CreateTransactionAmqpEvent struct {
	transactionEventProducerGateway main_gateways.TransactionEventProducerGateway
	logsMonitoringGateway           main_gateways.LogsMonitoringGateway
	messageUtils                    main_utils_messages.ApplicationMessages
}

func NewCreateTransactionAmqpEventAllArgs(
	transactionEventProducerGateway main_gateways.TransactionEventProducerGateway,
	logsMonitoringGateway main_gateways.LogsMonitoringGateway,
	messageBeans main_utils_messages.ApplicationMessages) *CreateTransactionAmqpEvent {
	return &CreateTransactionAmqpEvent{
		transactionEventProducerGateway: transactionEventProducerGateway,
		logsMonitoringGateway:           logsMonitoringGateway,
		messageUtils:                    messageBeans}
}

func NewCreateTransactionAmqpEvent(
	transactionEventProducerGateway main_gateways.TransactionEventProducerGateway,
	logsMonitoringGateway main_gateways.LogsMonitoringGateway,
) *CreateTransactionAmqpEvent {
	return &CreateTransactionAmqpEvent{
		transactionEventProducerGateway,
		logsMonitoringGateway,
		*main_utils_messages.NewApplicationMessages(),
	}
}

func (this *CreateTransactionAmqpEvent) Execute(
	ctx context.Context, transaction main_domains.Transaction) main_domains_exceptions.ApplicationException {

	span := *main_gateways_logs_resources.NewSpanLogInfo(ctx)
	defer span.End()

	this.logsMonitoringGateway.DEBUG(span,
		fmt.Sprintf("Creating new transaction event with accountId: %s", transaction.GetAccountId()))

	return this.transactionEventProducerGateway.Send(ctx, transaction)
}
