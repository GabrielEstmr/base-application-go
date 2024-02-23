package main_usecases

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_gateways "baseapplicationgo/main/gateways"
	main_utils_messages "baseapplicationgo/main/utils/messages"
	"context"
	"fmt"
)

type CreateSessionByIdentityProvider struct {
	_MSG_TOKEN_MUST_NOT_BE_EMPTY    string
	_MSG_ARCH_ISSUE                 string
	_MSG_CONFLICT_CREATE_USER_ISSUE string
	userDatabaseGateway             main_gateways.UserDatabaseGateway
	lockGateway                     main_gateways.DistributedLockGateway
	logsMonitoringGateway           main_gateways.LogsMonitoringGateway
	spanGateway                     main_gateways.SpanGateway
	messageUtils                    main_utils_messages.ApplicationMessages
}

func NewCreateSessionByIdentityProvider(
	userDatabaseGateway main_gateways.UserDatabaseGateway,
	lockGateway main_gateways.DistributedLockGateway,
	logsMonitoringGateway main_gateways.LogsMonitoringGateway,
	spanGateway main_gateways.SpanGateway,
	messageBeans main_utils_messages.ApplicationMessages,
) *CreateSessionByIdentityProvider {
	return &CreateSessionByIdentityProvider{
		_MSG_TOKEN_MUST_NOT_BE_EMPTY:    "providers.create.session.by.identity.provider.token.empty",
		_MSG_ARCH_ISSUE:                 "exceptions.architecture.application.issue",
		_MSG_CONFLICT_CREATE_USER_ISSUE: "providers.create.user.lock.issue",
		userDatabaseGateway:             userDatabaseGateway,
		lockGateway:                     lockGateway,
		logsMonitoringGateway:           logsMonitoringGateway,
		spanGateway:                     spanGateway,
		messageUtils:                    messageBeans,
	}
}

func (this *CreateSessionByIdentityProvider) Execute(
	ctx context.Context,
	args main_domains.ExternalProviderSessionArgs,
	tokenClaims main_domains.TokenClaims,
	databaseOptions main_domains.DatabaseOptions,
) main_domains_exceptions.ApplicationException {

	span := this.spanGateway.Get(ctx, "CreateSessionByIdentityProvider-Execute")
	defer span.End()
	this.logsMonitoringGateway.DEBUG(span,
		fmt.Sprintf("CreateSessionByIdentityProvider. provider: %s", args.GetProvider()))

	user, errF := this.userDatabaseGateway.FindByEmail(span.GetCtx(), tokenClaims.Email, databaseOptions)
	if errF != nil {
		this.logsMonitoringGateway.ERROR(span, fmt.Sprintf(errF.Error()))
		return main_domains_exceptions.NewInternalServerErrorExceptionSglMsg(
			this.messageUtils.GetDefaultLocale(
				this._MSG_ARCH_ISSUE))
	}

	if user.IsEmpty() {
		_, errSave := this.userDatabaseGateway.Save(
			span.GetCtx(),
			*main_domains.NewUserAsCreatedFromProvider(tokenClaims, args.GetProvider()), databaseOptions)
		if errSave != nil {
			this.logsMonitoringGateway.ERROR(span, fmt.Sprintf(errSave.Error()))
			return main_domains_exceptions.NewInternalServerErrorExceptionSglMsg(
				this.messageUtils.GetDefaultLocale(
					this._MSG_ARCH_ISSUE))
		}
	}
	return nil
}
