package main_usecases

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_gateways "baseapplicationgo/main/gateways"
	main_gateways_logs_resources "baseapplicationgo/main/gateways/logs/resources"
	main_utils_messages "baseapplicationgo/main/utils/messages"
	"context"
	"fmt"
)

const _MSG_CREATE_NEW_DOC_DOC_ALREADY_EXISTS = "providers.create.user.user.with.given.document.already.exists"
const _MSG_CREATE_NEW_DOC_ARCH_ISSUE = "exceptions.architecture.application.issue"

type CreateNewUser struct {
	userDatabaseGateway   main_gateways.UserDatabaseGateway
	logsMonitoringGateway main_gateways.LogsMonitoringGateway
	messageUtils          main_utils_messages.ApplicationMessages
}

func NewCreateNewUserAllArgs(
	userDatabaseGateway main_gateways.UserDatabaseGateway,
	logLoki main_gateways.LogsMonitoringGateway,
	messageBeans main_utils_messages.ApplicationMessages) *CreateNewUser {
	return &CreateNewUser{
		userDatabaseGateway:   userDatabaseGateway,
		logsMonitoringGateway: logLoki,
		messageUtils:          messageBeans}
}

func NewCreateNewUser(
	userDatabaseGateway main_gateways.UserDatabaseGateway,
	logLoki main_gateways.LogsMonitoringGateway,
) *CreateNewUser {
	return &CreateNewUser{
		userDatabaseGateway,
		logLoki,
		*main_utils_messages.NewApplicationMessages(),
	}
}

func (this *CreateNewUser) Execute(ctx context.Context, user main_domains.User) (
	main_domains.User, main_domains_exceptions.ApplicationException) {

	span := *main_gateways_logs_resources.NewSpanLogInfo(ctx)
	defer span.End()

	this.logsMonitoringGateway.INFO(span,
		fmt.Sprintf("Creating new User with documentNumber: %s", user.DocumentNumber))
	userAlreadyPersisted, err := this.userDatabaseGateway.FindByDocumentNumber(user.DocumentNumber)
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

	persistedUser, err := this.userDatabaseGateway.Save(user)
	if err != nil {
		return main_domains.User{}, main_domains_exceptions.
			NewInternalServerErrorExceptionSglMsg(_MSG_CREATE_NEW_DOC_ARCH_ISSUE)
	}
	return persistedUser, nil
}
