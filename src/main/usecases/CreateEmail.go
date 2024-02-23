package main_usecases

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_apm "baseapplicationgo/main/domains/apm"
	main_domains_enums "baseapplicationgo/main/domains/enums"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	"baseapplicationgo/main/domains/lock"
	main_gateways "baseapplicationgo/main/gateways"
	main_usecases_interfaces "baseapplicationgo/main/usecases/interfaces"
	main_utils "baseapplicationgo/main/utils"
	main_utils_messages "baseapplicationgo/main/utils/messages"
	"context"
	"errors"
	"fmt"
	"time"
)

type CreateEmail struct {
	_MSG_ARCH_ISSUE                     string
	_MSG_LOCK_ISSUE                     string
	emailDatabaseGateway                main_gateways.EmailDatabaseGateway
	createEmailBodySendAndPersistAsSent main_usecases_interfaces.CreateEmailBodySendAndPersistAsSent
	lockGateway                         main_gateways.DistributedLockGateway
	logsMonitoringGateway               main_gateways.LogsMonitoringGateway
	spanGateway                         main_gateways.SpanGateway
	messageUtils                        main_utils_messages.ApplicationMessages
}

func NewCreateEmailAllArgs(
	emailDatabaseGateway main_gateways.EmailDatabaseGateway,
	createEmailBodySendAndPersistAsSent main_usecases_interfaces.CreateEmailBodySendAndPersistAsSent,
	lockGateway main_gateways.DistributedLockGateway,
	logsMonitoringGateway main_gateways.LogsMonitoringGateway,
	spanGateway main_gateways.SpanGateway,
	messageUtils main_utils_messages.ApplicationMessages,
) *CreateEmail {
	return &CreateEmail{
		_MSG_ARCH_ISSUE:                     "exceptions.architecture.application.issue",
		_MSG_LOCK_ISSUE:                     "providers.general.lock.issue",
		emailDatabaseGateway:                emailDatabaseGateway,
		createEmailBodySendAndPersistAsSent: createEmailBodySendAndPersistAsSent,
		lockGateway:                         lockGateway,
		logsMonitoringGateway:               logsMonitoringGateway,
		spanGateway:                         spanGateway,
		messageUtils:                        messageUtils,
	}
}

func (this *CreateEmail) Execute(
	ctx context.Context,
	msgId string,
	emailParams main_domains.EmailParams,
) (
	main_domains.Email,
	main_domains_exceptions.ApplicationException,
) {

	span := this.spanGateway.Get(ctx, "CreateEmail-Execute")
	defer span.End()
	this.logsMonitoringGateway.DEBUG(span,
		fmt.Sprintf("Creating new email. eventId: %s", msgId))

	if main_utils.NewStringUtils().IsEmpty(msgId) {
		return this.logAndReturnArchError(span, errors.New("empty msgId"))
	}

	singleLock := this.lockGateway.GetWithScope(span.GetCtx(), "CreateEmail-Execute", msgId, 90*time.Second)
	errLock := singleLock.Lock()

	if errLock == nil {
		defer this.unlockAndLogIfError(span, singleLock)

		email := *main_domains.NewEmail(msgId, emailParams, main_domains_enums.EMAIL_STATUS_STARTED)
		persistedEmail, errSave := this.emailDatabaseGateway.Save(span.GetCtx(), email)
		if errSave != nil {
			return this.logAndReturnArchError(span, errSave)
		}

		persistedSentEmail, errUpdateSent := this.createEmailBodySendAndPersistAsSent.Execute(span.GetCtx(), persistedEmail)
		if errUpdateSent != nil {
			return this.logAndReturnArchError(span, errUpdateSent)
		}

		return persistedSentEmail, nil
	} else {
		return *new(main_domains.Email),
			main_domains_exceptions.NewConflictExceptionSglMsg(this.messageUtils.GetDefaultLocale(
				this._MSG_LOCK_ISSUE))
	}
}

func (this *CreateEmail) logAndReturnArchError(
	span main_domains_apm.SpanLogInfo,
	err error,
) (main_domains.Email, main_domains_exceptions.ApplicationException) {
	this.logsMonitoringGateway.ERROR(span, err.Error())
	return main_domains.Email{}, main_domains_exceptions.
		NewInternalServerErrorExceptionSglMsg(this.messageUtils.
			GetDefaultLocale(this._MSG_ARCH_ISSUE))
}

func (this *CreateEmail) unlockAndLogIfError(span main_domains_apm.SpanLogInfo, lock *lock.SingleLock) {
	_, errUnlock := lock.Unlock()
	if errUnlock != nil {
		this.logsMonitoringGateway.ERROR(span, "CreateEmail-execute: error during unlock")
	}
}
