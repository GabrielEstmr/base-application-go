package main_usecases

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_enums "baseapplicationgo/main/domains/enums"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_gateways "baseapplicationgo/main/gateways"
	main_usecases_interfaces "baseapplicationgo/main/usecases/interfaces"
	main_utils_messages "baseapplicationgo/main/utils/messages"
	"context"
	"fmt"
)

type EnableInternalProviderUser struct {
	_MSG_KEY_ARCH_ISSUE                  string
	_MSG_EMAIL_VERIFICATION_LOCK_ISSUE   string
	_MSG_USER_ALREADY_ENABLED            string
	validateUserIsInternalAuthProvider   main_usecases_interfaces.ValidateUserAuthProviderOrigin
	findValidUserEmailVerification       main_usecases_interfaces.FindValidUserEmailVerification
	userEmailVerificationDatabaseGateway main_gateways.UserEmailVerificationDatabaseGateway
	setUserToEnabledAndEmailVerified     main_usecases_interfaces.SetUserToEnabledAndEmailVerified
	lockGateway                          main_gateways.DistributedLockGateway
	logsMonitoringGateway                main_gateways.LogsMonitoringGateway
	spanGateway                          main_gateways.SpanGateway
	messageUtils                         main_utils_messages.ApplicationMessages
}

func NewEnableInternalProviderUser(
	validateUserIsInternalAuthProvider main_usecases_interfaces.ValidateUserAuthProviderOrigin,
	findValidUserEmailVerification main_usecases_interfaces.FindValidUserEmailVerification,
	userEmailVerificationDatabaseGateway main_gateways.UserEmailVerificationDatabaseGateway,
	setUserToEnabledAndEmailVerified main_usecases_interfaces.SetUserToEnabledAndEmailVerified,
	lockGateway main_gateways.DistributedLockGateway,
	logsMonitoringGateway main_gateways.LogsMonitoringGateway,
	spanGateway main_gateways.SpanGateway,
	messageBeans main_utils_messages.ApplicationMessages) *EnableInternalProviderUser {
	return &EnableInternalProviderUser{
		_MSG_KEY_ARCH_ISSUE:                  "exceptions.architecture.application.issue",
		_MSG_EMAIL_VERIFICATION_LOCK_ISSUE:   "providers.modify.user.verification.email.error.lock.issue",
		_MSG_USER_ALREADY_ENABLED:            "providers.enable.user.already.enable",
		validateUserIsInternalAuthProvider:   validateUserIsInternalAuthProvider,
		findValidUserEmailVerification:       findValidUserEmailVerification,
		userEmailVerificationDatabaseGateway: userEmailVerificationDatabaseGateway,
		setUserToEnabledAndEmailVerified:     setUserToEnabledAndEmailVerified,
		lockGateway:                          lockGateway,
		logsMonitoringGateway:                logsMonitoringGateway,
		spanGateway:                          spanGateway,
		messageUtils:                         messageBeans}
}

func (this *EnableInternalProviderUser) Execute(
	ctx context.Context,
	userId string,
	verificationCode string,
	databaseOptions main_domains.DatabaseOptions,
) (
	main_domains.User, main_domains_exceptions.ApplicationException) {
	span := this.spanGateway.Get(ctx, "VerifyUser-Execute")
	defer span.End()
	this.logsMonitoringGateway.DEBUG(span,
		fmt.Sprintf("Verifing user. code: %s", verificationCode))

	userVerification, errValidate := this.validateUserIsInternalAuthProvider.Execute(span.GetCtx(), userId, databaseOptions)
	if errValidate != nil {
		return *new(main_domains.User), errValidate
	}

	if userVerification.IsStatusEnabled() {
		return *new(main_domains.User),
			main_domains_exceptions.NewConflictExceptionSglMsg(this.messageUtils.GetDefaultLocale(
				this._MSG_USER_ALREADY_ENABLED))
	}

	userEmailVerification, errUV := this.findValidUserEmailVerification.Execute(
		span.GetCtx(),
		userId,
		main_domains_enums.EMAIL_VERIFICATION_SCOPE_ENABLE_USER,
		verificationCode,
		databaseOptions)
	if errUV != nil {
		return *new(main_domains.User), errUV
	}

	updatedUser, errUP := this.setUserToEnabledAndEmailVerified.Execute(span.GetCtx(), userId, databaseOptions)
	if errUP != nil {
		return updatedUser, errUP
	}

	_, errUpdate := this.userEmailVerificationDatabaseGateway.Update(
		span.GetCtx(),
		userEmailVerification.CloneAsUsed(),
		databaseOptions)
	if errUpdate != nil {
		return *new(main_domains.User), main_domains_exceptions.
			NewInternalServerErrorExceptionSglMsg(this.messageUtils.
				GetDefaultLocale(this._MSG_KEY_ARCH_ISSUE))
	}

	return updatedUser, nil
}
