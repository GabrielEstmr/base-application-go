package main_usecases

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_gateways "baseapplicationgo/main/gateways"
	main_utils_messages "baseapplicationgo/main/utils/messages"
	"context"
	"fmt"
)

type RefreshSession struct {
	_MSG_TOKEN_MUST_NOT_BE_EMPTY    string
	_MSG_ARCH_ISSUE                 string
	_MSG_USER_NOT_FOUND             string
	_MSG_INVALID_CREDENTIALS        string
	_MSG_CONFLICT_CREATE_USER_ISSUE string
	authProvider                    main_gateways.AuthProviderGateway
	logsMonitoringGateway           main_gateways.LogsMonitoringGateway
	spanGateway                     main_gateways.SpanGateway
	messageUtils                    main_utils_messages.ApplicationMessages
}

func NewRefreshSession(
	authProvider main_gateways.AuthProviderGateway,
	logsMonitoringGateway main_gateways.LogsMonitoringGateway,
	spanGateway main_gateways.SpanGateway,
	messageBeans main_utils_messages.ApplicationMessages,
) *RefreshSession {
	return &RefreshSession{
		_MSG_TOKEN_MUST_NOT_BE_EMPTY:    "providers.create.session.by.identity.provider.token.empty",
		_MSG_ARCH_ISSUE:                 "exceptions.architecture.application.issue",
		_MSG_USER_NOT_FOUND:             "providers.find.user.user.not.found",
		_MSG_INVALID_CREDENTIALS:        "providers.create.session.invalid.credentials",
		_MSG_CONFLICT_CREATE_USER_ISSUE: "providers.create.user.lock.issue",
		authProvider:                    authProvider,
		logsMonitoringGateway:           logsMonitoringGateway,
		spanGateway:                     spanGateway,
		messageUtils:                    messageBeans,
	}
}

func (this *RefreshSession) Execute(ctx context.Context, refreshToken string,
) (main_domains.SessionCredentials, main_domains_exceptions.ApplicationException) {

	span := this.spanGateway.Get(ctx, "CreateRefreshToken-Execute")
	defer span.End()
	this.logsMonitoringGateway.INFO(span, "CreateRefreshToken")

	sessionCredentials, errS := this.authProvider.RefreshSession(span.GetCtx(), refreshToken)
	if errS != nil {
		this.logsMonitoringGateway.ERROR(span, fmt.Sprintf(errS.Error()))
		return *new(main_domains.SessionCredentials), errS
	}
	return sessionCredentials, nil
}
