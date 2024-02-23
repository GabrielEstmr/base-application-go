package main_usecases

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_gateways "baseapplicationgo/main/gateways"
	main_utils_messages "baseapplicationgo/main/utils/messages"
	"context"
	"fmt"
	"net/http"
)

type CreateOrRetrieveUserFromAuthProvider struct {
	_MSG_ARCH_ISSUE       string
	authProviderGateway   main_gateways.AuthProviderGateway
	logsMonitoringGateway main_gateways.LogsMonitoringGateway
	spanGateway           main_gateways.SpanGateway
	messageUtils          main_utils_messages.ApplicationMessages
}

func NewCreateOrRetrieveUserFromAuthProvider(
	authProviderGateway main_gateways.AuthProviderGateway,
	logsMonitoringGateway main_gateways.LogsMonitoringGateway,
	spanGateway main_gateways.SpanGateway,
	messageBeans main_utils_messages.ApplicationMessages) *CreateOrRetrieveUserFromAuthProvider {
	return &CreateOrRetrieveUserFromAuthProvider{
		_MSG_ARCH_ISSUE:       "exceptions.architecture.application.issue",
		authProviderGateway:   authProviderGateway,
		logsMonitoringGateway: logsMonitoringGateway,
		spanGateway:           spanGateway,
		messageUtils:          messageBeans}
}

func (this *CreateOrRetrieveUserFromAuthProvider) Execute(
	ctx context.Context,
	user main_domains.User) (
	main_domains.User, main_domains_exceptions.ApplicationException) {

	span := this.spanGateway.Get(ctx, "CreateOrRetrieveUserFromAuthProvider-Execute")
	defer span.End()
	this.logsMonitoringGateway.DEBUG(span,
		fmt.Sprintf("Creating new into auth provider: %s", user.GetEmail()))

	updatedUser, errAP := this.authProviderGateway.CreateUser(span.GetCtx(), user)
	if errAP != nil && errAP.GetCode() != http.StatusConflict {
		return *new(main_domains.User), errAP
	}

	if errAP != nil && errAP.GetCode() == http.StatusConflict {
		users, errGU := this.authProviderGateway.GetUsers(
			span.GetCtx(),
			user.GetEmail(),
		)
		if errGU != nil {
			return *new(main_domains.User), errGU
		}

		if len(users) > 0 {
			return user.CloneWithNewAuthProviderId(users[0].GetId()), nil
		} else {
			this.logsMonitoringGateway.ERROR(span, "Error to get user from internal authProvider")
			return *new(main_domains.User), main_domains_exceptions.
				NewConflictExceptionSglMsg(this.messageUtils.
					GetDefaultLocale(this._MSG_ARCH_ISSUE))
		}
	}
	return updatedUser, nil
}
