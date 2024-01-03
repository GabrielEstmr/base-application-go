package main_usecases

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_gateways "baseapplicationgo/main/gateways"
	main_gateways_spans "baseapplicationgo/main/gateways/spans"
	main_utils_messages "baseapplicationgo/main/utils/messages"
	"context"
	"time"
)

const _MSG_CREATE_NEW_TRANSACTION_ARCH_ISSUE = "exceptions.architecture.application.issue"

type PersistTransaction struct {
	transactionDatabaseGateway main_gateways.TransactionDatabaseGateway
	logsMonitoringGateway      main_gateways.LogsMonitoringGateway
	lockGateway                main_gateways.DistributedLockGateway
	spanGateway                main_gateways.SpanGateway
	messageUtils               main_utils_messages.ApplicationMessages
}

func NewPersistTransactionAllArgs(
	transactionDatabaseGateway main_gateways.TransactionDatabaseGateway,
	logsMonitoringGateway main_gateways.LogsMonitoringGateway,
	lockGateway main_gateways.DistributedLockGateway,
	spanGateway main_gateways.SpanGateway,
	messageUtils main_utils_messages.ApplicationMessages) *PersistTransaction {
	return &PersistTransaction{
		transactionDatabaseGateway: transactionDatabaseGateway,
		logsMonitoringGateway:      logsMonitoringGateway,
		lockGateway:                lockGateway,
		spanGateway:                spanGateway,
		messageUtils:               messageUtils}
}

func NewPersistTransaction(
	transactionDatabaseGateway main_gateways.TransactionDatabaseGateway,
	logsMonitoringGateway main_gateways.LogsMonitoringGateway,
	lockGateway main_gateways.DistributedLockGateway,
) *PersistTransaction {
	return &PersistTransaction{
		transactionDatabaseGateway: transactionDatabaseGateway,
		logsMonitoringGateway:      logsMonitoringGateway,
		lockGateway:                lockGateway,
		spanGateway:                main_gateways_spans.NewSpanGatewayImpl(),
		messageUtils:               *main_utils_messages.NewApplicationMessages()}
}

func (this *PersistTransaction) Execute(
	ctx context.Context,
	transaction main_domains.Transaction) (main_domains.Transaction, main_domains_exceptions.ApplicationException) {
	span := this.spanGateway.Get(ctx, "PersistTransaction-Execute")
	defer span.End()

	lock := this.lockGateway.Get(span.GetCtx(), "Key", 8*time.Second)

	errLock := lock.Lock()
	if errLock != nil {
		return main_domains.Transaction{},
			main_domains_exceptions.NewInternalServerErrorExceptionSglMsg(
				this.messageUtils.GetDefaultLocale(
					_MSG_CREATE_NEW_TRANSACTION_ARCH_ISSUE))
	}

	persistedTransaction, err := this.transactionDatabaseGateway.Save(span.GetCtx(), transaction)
	if err != nil {
		return main_domains.Transaction{},
			main_domains_exceptions.NewInternalServerErrorExceptionSglMsg(
				this.messageUtils.GetDefaultLocale(
					_MSG_CREATE_NEW_TRANSACTION_ARCH_ISSUE))
	}

	_, errUnlock := lock.Unlock()
	if errUnlock != nil {
		return main_domains.Transaction{},
			main_domains_exceptions.NewInternalServerErrorExceptionSglMsg(
				this.messageUtils.GetDefaultLocale(
					_MSG_CREATE_NEW_TRANSACTION_ARCH_ISSUE))
	}

	return persistedTransaction, nil
}
