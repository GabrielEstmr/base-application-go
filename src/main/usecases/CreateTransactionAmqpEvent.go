package main_usecases

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_gateways "baseapplicationgo/main/gateways"
	main_utils_messages "baseapplicationgo/main/utils/messages"
	"context"
	"fmt"
)

type CreateTransactionAmqpEvent struct {
	transactionEventProducerGateway main_gateways.TransactionEventProducerGateway
	logsMonitoringGateway           main_gateways.LogsMonitoringGateway
	messageUtils                    main_utils_messages.ApplicationMessages
	spanGateway                     main_gateways.SpanGateway
}

func NewCreateTransactionAmqpEventAllArgs(
	transactionEventProducerGateway main_gateways.TransactionEventProducerGateway,
	logsMonitoringGateway main_gateways.LogsMonitoringGateway,
	messageBeans main_utils_messages.ApplicationMessages,
	spanGateway main_gateways.SpanGateway,
) *CreateTransactionAmqpEvent {
	return &CreateTransactionAmqpEvent{
		transactionEventProducerGateway: transactionEventProducerGateway,
		logsMonitoringGateway:           logsMonitoringGateway,
		messageUtils:                    messageBeans,
		spanGateway:                     spanGateway,
	}
}

func NewCreateTransactionAmqpEvent(
	transactionEventProducerGateway main_gateways.TransactionEventProducerGateway,
	logsMonitoringGateway main_gateways.LogsMonitoringGateway,
	spanGateway main_gateways.SpanGateway,
) *CreateTransactionAmqpEvent {
	return &CreateTransactionAmqpEvent{
		transactionEventProducerGateway,
		logsMonitoringGateway,
		*main_utils_messages.NewApplicationMessages(),
		spanGateway,
	}
}

func (this *CreateTransactionAmqpEvent) Execute(
	ctx context.Context, transaction main_domains.Transaction) main_domains_exceptions.ApplicationException {
	span := this.spanGateway.Get(ctx, "CreateTransactionAmqpEvent-Execute")
	defer span.End()

	this.logsMonitoringGateway.INFO(span,
		fmt.Sprintf("Creating new transaction event with accountId: %s", transaction.GetAccountId()))

	return this.transactionEventProducerGateway.Send(span.GetCtx(), transaction)
}
