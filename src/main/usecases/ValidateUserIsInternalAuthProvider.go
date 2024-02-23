package main_usecases

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_gateways "baseapplicationgo/main/gateways"
	main_utils_messages "baseapplicationgo/main/utils/messages"
	"context"
	"fmt"
)

const _MSG_PROVIDER_CONFLICT_USER_NOT_INTERNAL_PROVIDER = "providers.enable.user.not.internal.provider"

type ValidateUserIsInternalAuthProvider struct {
	_MSG_ARCH_ISSUE       string
	userDatabaseGateway   main_gateways.UserDatabaseGateway
	logsMonitoringGateway main_gateways.LogsMonitoringGateway
	spanGateway           main_gateways.SpanGateway
	messageUtils          main_utils_messages.ApplicationMessages
}

func NewValidateUserIsInternalAuthProvider(
	userDatabaseGateway main_gateways.UserDatabaseGateway,
	logsMonitoringGateway main_gateways.LogsMonitoringGateway,
	spanGateway main_gateways.SpanGateway,
	messageBeans main_utils_messages.ApplicationMessages) *ValidateUserIsInternalAuthProvider {
	return &ValidateUserIsInternalAuthProvider{
		_MSG_ARCH_ISSUE:       "exceptions.architecture.application.issue",
		userDatabaseGateway:   userDatabaseGateway,
		logsMonitoringGateway: logsMonitoringGateway,
		spanGateway:           spanGateway,
		messageUtils:          messageBeans,
	}
}

func (this *ValidateUserIsInternalAuthProvider) Execute(ctx context.Context, userId string, databaseOptions main_domains.DatabaseOptions,
) (main_domains.User, main_domains_exceptions.ApplicationException) {

	span := this.spanGateway.Get(ctx, "ValidateUserIsInternalAuthProvider-Execute")
	defer span.End()
	this.logsMonitoringGateway.DEBUG(span,
		fmt.Sprintf("Verifing user is internal provider. id: %s", userId))

	user, errF := this.userDatabaseGateway.FindById(span.GetCtx(), userId, databaseOptions)
	if errF != nil {
		return *new(main_domains.User), main_domains_exceptions.
			NewInternalServerErrorExceptionSglMsg(this.messageUtils.
				GetDefaultLocale(this._MSG_ARCH_ISSUE))
	}
	if !user.IsInternalAuthProvider() {
		return user, main_domains_exceptions.NewConflictExceptionSglMsg(this.messageUtils.GetDefaultLocale(
			_MSG_PROVIDER_CONFLICT_USER_NOT_INTERNAL_PROVIDER))
	}

	return user, nil
}
