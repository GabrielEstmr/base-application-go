package main_usecases

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_gateways "baseapplicationgo/main/gateways"
	main_gateways_spans "baseapplicationgo/main/gateways/spans"
	main_utils_messages "baseapplicationgo/main/utils/messages"
	"context"
	"fmt"
	"time"
)

const _MSG_CREATE_NEW_TRANSACTION_ARCH_ISSUE = "exceptions.architecture.application.issue"
const _MSG_CREATE_NEW_TRANSACTION_LOCK_ISSUE = "providers.general.lock.issue"

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

	lock := this.lockGateway.Get(span.GetCtx(), "Key", 60*time.Second)

	errLock := lock.Lock()

	if errLock != nil {
		return main_domains.Transaction{},
			main_domains_exceptions.NewConflictExceptionSglMsg(this.messageUtils.GetDefaultLocale(
				_MSG_CREATE_NEW_TRANSACTION_LOCK_ISSUE))
	}

	// For check lock purposes only
	time.Sleep(1 * time.Second)

	persistedTransaction, err := this.transactionDatabaseGateway.Save(span.GetCtx(), transaction)
	if err != nil {
		return main_domains.Transaction{},
			main_domains_exceptions.NewInternalServerErrorExceptionSglMsg(
				this.messageUtils.GetDefaultLocale(
					_MSG_CREATE_NEW_TRANSACTION_ARCH_ISSUE))
	}

	unlocked, errUnlock := lock.Unlock()
	if errUnlock != nil {
		return main_domains.Transaction{},
			main_domains_exceptions.NewInternalServerErrorExceptionSglMsg(
				this.messageUtils.GetDefaultLocale(
					_MSG_CREATE_NEW_TRANSACTION_ARCH_ISSUE))
	}

	fmt.Println("UNLOCKED =============> ", unlocked)
	return persistedTransaction, nil
}
