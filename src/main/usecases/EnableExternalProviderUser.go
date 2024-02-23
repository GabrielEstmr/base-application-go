package main_usecases

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_gateways "baseapplicationgo/main/gateways"
	main_usecases_interfaces "baseapplicationgo/main/usecases/interfaces"
	main_utils_messages "baseapplicationgo/main/utils/messages"
	"context"
	"fmt"
)

//AtomicLockedEnableExternalProviderUser

type EnableExternalProviderUser struct {
	_MSG_ENABLE_USER_CODE_NOT_FOUND    string
	_MSG_ENABLE_EMAIL_LOCK_ISSUE       string
	authProvider                       main_gateways.AuthProviderGateway
	userDatabaseGateway                main_gateways.UserDatabaseGateway
	validateUserIsExternalAuthProvider main_usecases_interfaces.ValidateUserAuthProviderOrigin
	logsMonitoringGateway              main_gateways.LogsMonitoringGateway
	spanGateway                        main_gateways.SpanGateway
	messageUtils                       main_utils_messages.ApplicationMessages
}

func NewEnableExternalProviderUser(
	authProvider main_gateways.AuthProviderGateway,
	userDatabaseGateway main_gateways.UserDatabaseGateway,
	validateUserIsExternalAuthProvider main_usecases_interfaces.ValidateUserAuthProviderOrigin,
	logsMonitoringGateway main_gateways.LogsMonitoringGateway,
	spanGateway main_gateways.SpanGateway,
	messageBeans main_utils_messages.ApplicationMessages,
) *EnableExternalProviderUser {
	return &EnableExternalProviderUser{
		_MSG_ENABLE_USER_CODE_NOT_FOUND:    "exceptions.architecture.application.issue",
		_MSG_ENABLE_EMAIL_LOCK_ISSUE:       "providers.modify.user.verification.email.error.lock.issue",
		authProvider:                       authProvider,
		userDatabaseGateway:                userDatabaseGateway,
		validateUserIsExternalAuthProvider: validateUserIsExternalAuthProvider,
		logsMonitoringGateway:              logsMonitoringGateway,
		spanGateway:                        spanGateway,
		messageUtils:                       messageBeans,
	}
}

func (this *EnableExternalProviderUser) Execute(
	ctx context.Context,
	currentUser main_domains.User,
	args main_domains.EnableExternalUserArgs,
	dbOpts main_domains.DatabaseOptions,
) (
	main_domains.User,
	main_domains_exceptions.ApplicationException,
) {

	span := this.spanGateway.Get(ctx, "EnableExternalProviderUser-Execute")
	defer span.End()
	this.logsMonitoringGateway.DEBUG(span,
		fmt.Sprintf("Verifing user. id: %s", currentUser.GetId()))

	_, errValidate := this.validateUserIsExternalAuthProvider.Execute(span.GetCtx(), currentUser.GetId(), dbOpts)
	if errValidate != nil {
		return *new(main_domains.User), errValidate
	}

	enabledUser := currentUser.CloneEnabledAsVerifiedAndAttData(args)
	persistedUpdatedUser, errU := this.userDatabaseGateway.Update(span.GetCtx(), enabledUser, dbOpts)
	if errU != nil {
		return *new(main_domains.User), errU
	}

	authProviderUpdatedUser, errAU := this.authProvider.ChangeUserStatusAndEmailVerification(
		span.GetCtx(),
		persistedUpdatedUser,
		true)
	if errAU != nil {
		return *new(main_domains.User), errAU
	}

	return authProviderUpdatedUser, nil
}
