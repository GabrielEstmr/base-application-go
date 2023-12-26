package main_usecases

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_gateways "baseapplicationgo/main/gateways"
	main_utils_messages "baseapplicationgo/main/utils/messages"
	"context"
	"fmt"
)

const _MSG_CREATE_NEW_DOC_DOC_ALREADY_EXISTS = "providers.create.user.user.with.given.document.already.exists"
const _MSG_CREATE_NEW_DOC_ARCH_ISSUE = "exceptions.architecture.application.issue"

type CreateNewUser struct {
	userDatabaseGateway   main_gateways.UserDatabaseGateway
	logsMonitoringGateway main_gateways.LogsMonitoringGateway
	spanGateway           main_gateways.SpanGateway
	messageUtils          main_utils_messages.ApplicationMessages
}

func NewCreateNewUserAllArgs(
	userDatabaseGateway main_gateways.UserDatabaseGateway,
	logsMonitoringGateway main_gateways.LogsMonitoringGateway,
	spanGateway main_gateways.SpanGateway,
	messageBeans main_utils_messages.ApplicationMessages) *CreateNewUser {
	return &CreateNewUser{
		userDatabaseGateway:   userDatabaseGateway,
		logsMonitoringGateway: logsMonitoringGateway,
		spanGateway:           spanGateway,
		messageUtils:          messageBeans}
}

func NewCreateNewUser(
	userDatabaseGateway main_gateways.UserDatabaseGateway,
	logsMonitoringGateway main_gateways.LogsMonitoringGateway,
	spanGateway main_gateways.SpanGateway,
) *CreateNewUser {
	return &CreateNewUser{
		userDatabaseGateway,
		logsMonitoringGateway,
		spanGateway,
		*main_utils_messages.NewApplicationMessages(),
	}
}

func (this *CreateNewUser) Execute(ctx context.Context, user main_domains.User) (
	main_domains.User, main_domains_exceptions.ApplicationException) {

	span := this.spanGateway.Get(ctx, "CreateNewUser-Execute")
	defer span.End()

	this.logsMonitoringGateway.INFO(span,
		fmt.Sprintf("Creating new User with documentNumber: %s", user.DocumentNumber))
	userAlreadyPersisted, err := this.userDatabaseGateway.FindByDocumentNumber(ctx, user.DocumentNumber)
	if err != nil {
		return main_domains.User{}, main_domains_exceptions.
			NewInternalServerErrorExceptionSglMsg(this.messageUtils.
				GetDefaultLocale(_MSG_CREATE_NEW_DOC_ARCH_ISSUE))
	}
	if !userAlreadyPersisted.IsEmpty() {
		return main_domains.User{}, main_domains_exceptions.NewConflictExceptionSglMsg(
			this.messageUtils.
				GetDefaultLocale(_MSG_CREATE_NEW_DOC_DOC_ALREADY_EXISTS))
	}

	persistedUser, err := this.userDatabaseGateway.Save(ctx, user)
	if err != nil {
		return main_domains.User{}, main_domains_exceptions.
			NewInternalServerErrorExceptionSglMsg(this.messageUtils.
				GetDefaultLocale(_MSG_CREATE_NEW_DOC_ARCH_ISSUE))
	}
	return persistedUser, nil
}
