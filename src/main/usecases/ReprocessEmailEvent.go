package main_usecases

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_apm "baseapplicationgo/main/domains/apm"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	"baseapplicationgo/main/domains/lock"
	main_gateways "baseapplicationgo/main/gateways"
	main_usecases_interfaces "baseapplicationgo/main/usecases/interfaces"
	main_utils_messages "baseapplicationgo/main/utils/messages"
	"context"
	"fmt"
	"time"
)

type ReprocessEmailEvent struct {
	emailDatabaseGateway                main_gateways.EmailDatabaseGateway
	createEmailBodySendAndPersistAsSent main_usecases_interfaces.CreateEmailBodySendAndPersistAsSent
	lockGateway                         main_gateways.DistributedLockGateway
	logsMonitoringGateway               main_gateways.LogsMonitoringGateway
	spanGateway                         main_gateways.SpanGateway
	messageUtils                        main_utils_messages.ApplicationMessages
}

func NewReprocessEmailEventAllArgs(
	emailDatabaseGateway main_gateways.EmailDatabaseGateway,
	createEmailBodySendAndPersistAsSent main_usecases_interfaces.CreateEmailBodySendAndPersistAsSent,
	lockGateway main_gateways.DistributedLockGateway,
	logsMonitoringGateway main_gateways.LogsMonitoringGateway,
	spanGateway main_gateways.SpanGateway,
	messageUtils main_utils_messages.ApplicationMessages,
) *ReprocessEmailEvent {
	return &ReprocessEmailEvent{
		emailDatabaseGateway:                emailDatabaseGateway,
		createEmailBodySendAndPersistAsSent: createEmailBodySendAndPersistAsSent,
		lockGateway:                         lockGateway,
		logsMonitoringGateway:               logsMonitoringGateway,
		spanGateway:                         spanGateway,
		messageUtils:                        messageUtils,
	}
}

func (this *ReprocessEmailEvent) Execute(ctx context.Context, id string,
) (main_domains.Email, main_domains_exceptions.ApplicationException) {

	span := this.spanGateway.Get(ctx, "ReprocessEmail-Execute")
	defer span.End()
	this.logsMonitoringGateway.DEBUG(span,
		fmt.Sprintf("Reprocessing emails. Id: %s", id))

	email, err := this.emailDatabaseGateway.FindById(span.GetCtx(), id)
	if err != nil {
		return *new(main_domains.Email),
			*main_domains_exceptions.NewResourceNotFoundExceptionSglMsg(
				this.messageUtils.GetDefaultLocale(
					_MSG_FIND_USER_BY_FILTER_ARCH_ISSUE))
	}

	if email.IsEmpty() {
		this.logsMonitoringGateway.DEBUG(span,
			fmt.Sprintf("Email not found. Id: %s", id))
		return *new(main_domains.Email), nil
	}

	if email.IsAbleToReprocess() {
		lock := this.lockGateway.Get(span.GetCtx(), email.GetEventId(), 90*time.Second)
		errLock := lock.Lock()
		if errLock == nil {
			updatedEmail, errS := this.createEmailBodySendAndPersistAsSent.Execute(span.GetCtx(), email)
			if errS != nil {
				return *new(main_domains.Email), *main_domains_exceptions.NewInternalServerErrorExceptionSglMsg(
					this.messageUtils.GetDefaultLocale(
						_MSG_FIND_USER_BY_FILTER_ARCH_ISSUE))
			}
			this.unlockAndLogIfError(span, lock)
			return updatedEmail, nil
		} else {
			return email, main_domains_exceptions.NewConflictExceptionSglMsg(this.messageUtils.GetDefaultLocale(
				_MSG_CREATE_NEW_TRANSACTION_LOCK_ISSUE))
		}
	}

	return *new(main_domains.Email), nil
}

func (this *ReprocessEmailEvent) unlockAndLogIfError(span main_domains_apm.SpanLogInfo, lock *lock.SingleLock) {
	_, errUnlock := lock.Unlock()
	if errUnlock != nil {
		this.logsMonitoringGateway.ERROR(span, "CreateEmail-execute: error during unlock")
	}
}
