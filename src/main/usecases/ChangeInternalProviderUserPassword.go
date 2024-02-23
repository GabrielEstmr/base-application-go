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

type ChangeInternalProviderUserPassword struct {
	_MSG_ARCH_ISSUE                      string
	validatePasswordFormat               main_usecases_interfaces.ValidatePasswordFormat
	validateUserIsInternalAuthProvider   main_usecases_interfaces.ValidateUserAuthProviderOrigin
	findValidUserEmailVerification       main_usecases_interfaces.FindValidUserEmailVerification
	userEmailVerificationDatabaseGateway main_gateways.UserEmailVerificationDatabaseGateway
	authProviderGateway                  main_gateways.AuthProviderGateway
	userDatabaseGateway                  main_gateways.UserDatabaseGateway
	logsMonitoringGateway                main_gateways.LogsMonitoringGateway
	spanGateway                          main_gateways.SpanGateway
	messageUtils                         main_utils_messages.ApplicationMessages
}

func NewChangeInternalProviderUserPassword(
	validatePasswordFormat main_usecases_interfaces.ValidatePasswordFormat,
	validateUserIsInternalAuthProvider main_usecases_interfaces.ValidateUserAuthProviderOrigin,
	findValidUserEmailVerification main_usecases_interfaces.FindValidUserEmailVerification,
	userEmailVerificationDatabaseGateway main_gateways.UserEmailVerificationDatabaseGateway,
	authProviderGateway main_gateways.AuthProviderGateway,
	userDatabaseGateway main_gateways.UserDatabaseGateway,
	logsMonitoringGateway main_gateways.LogsMonitoringGateway,
	spanGateway main_gateways.SpanGateway,
	messageUtils main_utils_messages.ApplicationMessages,
) *ChangeInternalProviderUserPassword {
	return &ChangeInternalProviderUserPassword{
		_MSG_ARCH_ISSUE:                      "exceptions.architecture.application.issue",
		validatePasswordFormat:               validatePasswordFormat,
		validateUserIsInternalAuthProvider:   validateUserIsInternalAuthProvider,
		findValidUserEmailVerification:       findValidUserEmailVerification,
		userEmailVerificationDatabaseGateway: userEmailVerificationDatabaseGateway,
		authProviderGateway:                  authProviderGateway,
		userDatabaseGateway:                  userDatabaseGateway,
		logsMonitoringGateway:                logsMonitoringGateway,
		spanGateway:                          spanGateway,
		messageUtils:                         messageUtils,
	}
}

func (this ChangeInternalProviderUserPassword) Execute(
	ctx context.Context,
	userId string,
	password string,
	verificationCode string,
	dbOpts main_domains.DatabaseOptions,
) (main_domains.User, main_domains_exceptions.ApplicationException) {

	span := this.spanGateway.Get(ctx, "ChangeInternalProviderUserPassword-Execute")
	defer span.End()
	this.logsMonitoringGateway.INFO(span, fmt.Sprintf("ChangeInternalProviderUserPassword User by id: %s", userId))

	errPassValidation := this.validatePasswordFormat.Execute(span.GetCtx(), password)
	if errPassValidation != nil {
		return *new(main_domains.User), errPassValidation
	}

	persistedUser, errValidate := this.validateUserIsInternalAuthProvider.Execute(span.GetCtx(), userId, dbOpts)
	if errValidate != nil {
		return *new(main_domains.User), errValidate
	}

	userEmailVerification, errUV := this.findValidUserEmailVerification.Execute(
		span.GetCtx(),
		userId,
		main_domains_enums.EMAIL_VERIFICATION_SCOPE_CHANGE_PASSWORD,
		verificationCode,
		dbOpts)
	if errUV != nil {
		return *new(main_domains.User), errUV
	}

	upPassUser := persistedUser.CloneWithNewPassword(password)
	upPassIntegratedUser, errUP := this.authProviderGateway.ChangeUsersPassword(span.GetCtx(), upPassUser)
	if errUP != nil {
		return *new(main_domains.User), errUP
	}

	upPassSavedUser, errS := this.userDatabaseGateway.Update(span.GetCtx(), upPassIntegratedUser, dbOpts)
	if errS != nil {
		this.logsMonitoringGateway.ERROR(span, errS.Error())
		return *new(main_domains.User), main_domains_exceptions.NewInternalServerErrorExceptionSglMsg(
			this.messageUtils.GetDefaultLocale(
				this._MSG_ARCH_ISSUE))
	}

	_, errUpdate := this.userEmailVerificationDatabaseGateway.Update(span.GetCtx(), userEmailVerification.CloneAsUsed(), dbOpts)
	if errUpdate != nil {
		return *new(main_domains.User), main_domains_exceptions.
			NewInternalServerErrorExceptionSglMsg(this.messageUtils.
				GetDefaultLocale(this._MSG_ARCH_ISSUE))
	}

	return upPassSavedUser, nil
}
