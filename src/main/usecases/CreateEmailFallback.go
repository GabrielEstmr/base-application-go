package main_usecases

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_apm "baseapplicationgo/main/domains/apm"
	main_domains_enums "baseapplicationgo/main/domains/enums"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_gateways "baseapplicationgo/main/gateways"
	main_usecases_interfaces "baseapplicationgo/main/usecases/interfaces"
	main_utils_messages "baseapplicationgo/main/utils/messages"
	"context"
	"fmt"
	"time"
)

const _MSG_CREATE_EMAIL_FALLBACK_ARCH_ISSUE = "exceptions.architecture.application.issue"

type CreateEmailFallback struct {
	emailDatabaseGateway                main_gateways.EmailDatabaseGateway
	createEmailBodySendAndPersistAsSent main_usecases_interfaces.CreateEmailBodySendAndPersistAsSent
	lockGateway                         main_gateways.DistributedLockGateway
	logsMonitoringGateway               main_gateways.LogsMonitoringGateway
	spanGateway                         main_gateways.SpanGateway
	messageUtils                        main_utils_messages.ApplicationMessages
}

func NewCreateEmailFallbackAllArgs(
	emailDatabaseGateway main_gateways.EmailDatabaseGateway,
	createEmailBodySendAndPersistAsSent main_usecases_interfaces.CreateEmailBodySendAndPersistAsSent,
	lockGateway main_gateways.DistributedLockGateway,
	logsMonitoringGateway main_gateways.LogsMonitoringGateway,
	spanGateway main_gateways.SpanGateway,
	messageUtils main_utils_messages.ApplicationMessages,
) *CreateEmailFallback {
	return &CreateEmailFallback{
		emailDatabaseGateway:                emailDatabaseGateway,
		createEmailBodySendAndPersistAsSent: createEmailBodySendAndPersistAsSent,
		lockGateway:                         lockGateway,
		logsMonitoringGateway:               logsMonitoringGateway,
		spanGateway:                         spanGateway,
		messageUtils:                        messageUtils,
	}
}

func (this *CreateEmailFallback) Execute(
	ctx context.Context,
	msgId string,
	emailParams main_domains.EmailParams,
) (
	main_domains.Email,
	main_domains_exceptions.ApplicationException,
) {

	span := this.spanGateway.Get(ctx, "CreateEmailFallback-Execute")
	defer span.End()

	this.logsMonitoringGateway.DEBUG(span,
		fmt.Sprintf("Creating new email by fallback process. eventId: %s", msgId))

	emailDB, errF := this.emailDatabaseGateway.FindByEventId(span.GetCtx(), msgId)
	if errF != nil {
		return this.logAndReturnArchError(span, errF)
	}

	var reprocessEmail main_domains.Email
	if emailDB.IsEmpty() {
		email := *main_domains.NewEmail(msgId, emailParams, main_domains_enums.EMAIL_STATUS_STARTED)
		persistedEmail, errSave := this.emailDatabaseGateway.Save(span.GetCtx(), email)
		if errSave != nil {
			return this.logAndReturnArchError(span, errSave)
		}
		reprocessEmail = persistedEmail
	} else {
		reprocessEmail = emailDB
	}

	if reprocessEmail.IsAbleToReprocess() {
		lock := this.lockGateway.Get(span.GetCtx(), reprocessEmail.GetEventId(), 90*time.Second)
		errLock := lock.Lock()
		if errLock == nil {
			persistedSentEmail, errUpdateSent := this.createEmailBodySendAndPersistAsSent.Execute(span.GetCtx(), reprocessEmail)
			if errUpdateSent != nil {
				return reprocessEmail, main_domains_exceptions.
					NewInternalServerErrorExceptionSglMsg(this.messageUtils.
						GetDefaultLocale(_MSG_CREATE_EMAIL_FALLBACK_ARCH_ISSUE))
			}
			return persistedSentEmail, nil
		} else {
			return reprocessEmail, main_domains_exceptions.NewConflictExceptionSglMsg(this.messageUtils.GetDefaultLocale(
				_MSG_CREATE_NEW_TRANSACTION_LOCK_ISSUE))
		}
	}

	return reprocessEmail, nil
}

func (this *CreateEmailFallback) logAndReturnArchError(
	span main_domains_apm.SpanLogInfo,
	err error,
) (main_domains.Email, main_domains_exceptions.ApplicationException) {
	this.logsMonitoringGateway.ERROR(span, err.Error())
	return main_domains.Email{}, main_domains_exceptions.
		NewInternalServerErrorExceptionSglMsg(this.messageUtils.
			GetDefaultLocale(_MSG_CREATE_EMAIL_FALLBACK_ARCH_ISSUE))
}
