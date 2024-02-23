package main_usecases

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_gateways "baseapplicationgo/main/gateways"
	main_utils_messages "baseapplicationgo/main/utils/messages"
	"context"
	"fmt"
)

const _MSG_VERIFY_USER_USER_NOT_FOUND = "providers.find.user.user.not.found"

type SetUserToEnabledAndEmailVerified struct {
	authProviderGateway   main_gateways.AuthProviderGateway
	userDatabaseGateway   main_gateways.UserDatabaseGateway
	logsMonitoringGateway main_gateways.LogsMonitoringGateway
	spanGateway           main_gateways.SpanGateway
	messageUtils          main_utils_messages.ApplicationMessages
}

func NewSetUserToEnabledAndEmailVerified(
	authProviderGateway main_gateways.AuthProviderGateway,
	userDatabaseGateway main_gateways.UserDatabaseGateway,
	logsMonitoringGateway main_gateways.LogsMonitoringGateway,
	spanGateway main_gateways.SpanGateway,
	messageBeans main_utils_messages.ApplicationMessages) *SetUserToEnabledAndEmailVerified {
	return &SetUserToEnabledAndEmailVerified{
		authProviderGateway:   authProviderGateway,
		userDatabaseGateway:   userDatabaseGateway,
		logsMonitoringGateway: logsMonitoringGateway,
		spanGateway:           spanGateway,
		messageUtils:          messageBeans}
}

func (this *SetUserToEnabledAndEmailVerified) Execute(ctx context.Context, id string, dbOpt main_domains.DatabaseOptions) (
	main_domains.User, main_domains_exceptions.ApplicationException) {

	span := this.spanGateway.Get(ctx, "SetUserToEnabledAndEmailVerified-Execute")
	defer span.End()
	this.logsMonitoringGateway.DEBUG(span,
		fmt.Sprintf("SetUserToEnabledAndEmailVerified user with id: %s", id))

	persistedUser, errF := this.userDatabaseGateway.FindById(span.GetCtx(), id, dbOpt)
	if errF != nil {
		return *new(main_domains.User), errF
	}

	if persistedUser.IsEmpty() {
		return *new(main_domains.User),
			main_domains_exceptions.NewResourceNotFoundExceptionSglMsg(_MSG_VERIFY_USER_USER_NOT_FOUND)
	}

	verifiedUser := persistedUser.CloneEnabledAsVerified()

	verifiedAndUpdatedUser, errT := this.authProviderGateway.ChangeUserStatusAndEmailVerification(span.GetCtx(), verifiedUser, true)
	if errT != nil {
		return persistedUser, errT
	}

	updatedUser, errUP := this.userDatabaseGateway.Update(span.GetCtx(), verifiedAndUpdatedUser, dbOpt)
	if errUP != nil {
		return verifiedAndUpdatedUser, errUP
	}

	return updatedUser, nil
}
