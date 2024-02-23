package main_usecases

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_apm "baseapplicationgo/main/domains/apm"
	main_domains_enums "baseapplicationgo/main/domains/enums"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	"baseapplicationgo/main/domains/lock"
	main_gateways "baseapplicationgo/main/gateways"
	main_usecases_interfaces "baseapplicationgo/main/usecases/interfaces"
	main_utils_messages "baseapplicationgo/main/utils/messages"
	"context"
	"fmt"
)

type CreateAndSendUserVerificationEmail struct {
	userEmailVerificationDatabaseGateway main_gateways.UserEmailVerificationDatabaseGateway
	emailInternalProviderGateway         main_gateways.EmailInternalProviderGateway
	createUserEmailVerification          main_usecases_interfaces.CreateUserVerificationEmail
	lockGateway                          main_gateways.DistributedLockGateway
	logsMonitoringGateway                main_gateways.LogsMonitoringGateway
	spanGateway                          main_gateways.SpanGateway
	messageUtils                         main_utils_messages.ApplicationMessages
}

func NewCreateAndSendUserVerificationEmail(
	userEmailVerificationDatabaseGateway main_gateways.UserEmailVerificationDatabaseGateway,
	emailInternalProviderGateway main_gateways.EmailInternalProviderGateway,
	createUserEmailVerification main_usecases_interfaces.CreateUserVerificationEmail,
	lockGateway main_gateways.DistributedLockGateway,
	logsMonitoringGateway main_gateways.LogsMonitoringGateway,
	spanGateway main_gateways.SpanGateway,
	messageUtils main_utils_messages.ApplicationMessages,
) *CreateAndSendUserVerificationEmail {
	return &CreateAndSendUserVerificationEmail{
		userEmailVerificationDatabaseGateway: userEmailVerificationDatabaseGateway,
		emailInternalProviderGateway:         emailInternalProviderGateway,
		createUserEmailVerification:          createUserEmailVerification,
		lockGateway:                          lockGateway,
		logsMonitoringGateway:                logsMonitoringGateway,
		spanGateway:                          spanGateway,
		messageUtils:                         messageUtils,
	}
}

func (this *CreateAndSendUserVerificationEmail) Execute(
	ctx context.Context,
	user main_domains.User,
	scope main_domains_enums.UserEmailVerificationScope,
	databaseOptions main_domains.DatabaseOptions,
) (main_domains.UserEmailVerification, main_domains_exceptions.ApplicationException,
) {

	span := this.spanGateway.Get(ctx, "CreateUserVerificationEmail-Execute")
	defer span.End()
	this.logsMonitoringGateway.DEBUG(span,
		fmt.Sprintf("Creating verification email for user: %s", user.GetEmail()))

	verificationEmail := this.createUserEmailVerification.Execute(span.GetCtx(), user, scope)

	persistedVerificationEmail, errS := this.userEmailVerificationDatabaseGateway.Save(span.GetCtx(), verificationEmail, databaseOptions)
	if errS != nil {
		return *new(main_domains.UserEmailVerification), errS
	}

	errSend := this.emailInternalProviderGateway.SendMail(span.GetCtx(), persistedVerificationEmail.GetEmailParams())
	if errSend != nil {
		return persistedVerificationEmail, errSend
	}
	sentVerificationEmail := persistedVerificationEmail.CloneAsSent()

	sentPersistedVerificationEmail, errU := this.userEmailVerificationDatabaseGateway.Update(span.GetCtx(), sentVerificationEmail, databaseOptions)
	if errU != nil {
		return sentVerificationEmail, errU
	}

	return sentPersistedVerificationEmail, nil
}

func (this *CreateAndSendUserVerificationEmail) unlockAndLogIfError(span main_domains_apm.SpanLogInfo, lock *lock.SingleLock) {
	_, errUnlock := lock.Unlock()
	if errUnlock != nil {
		this.logsMonitoringGateway.ERROR(span, "CreateEmail-execute: error during unlock")
	}
}
