package main_usecases

import (
	main_configs_apm_logs_impl "baseapplicationgo/main/configs/apm/logs/impl"
	main_configs_messages "baseapplicationgo/main/configs/messages"
	main_domains "baseapplicationgo/main/domains"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_gateways "baseapplicationgo/main/gateways"
	"context"
	"fmt"
	"go.opentelemetry.io/otel/trace"
)

const _MSG_CREATE_NEW_DOC_DOC_ALREADY_EXISTS = "providers.create.user.user.with.given.document.already.exists"
const _MSG_CREATE_NEW_DOC_ARCH_ISSUE = "exceptions.architecture.application.issue"

type CreateNewUser struct {
	userDatabaseGateway main_gateways.UserDatabaseGateway
	logLoki             main_configs_apm_logs_impl.LogsGateway
}

func NewCreateNewUserAllArgs(
	userDatabaseGateway main_gateways.UserDatabaseGateway,
	logLoki main_configs_apm_logs_impl.LogsGateway) *CreateNewUser {
	return &CreateNewUser{
		userDatabaseGateway: userDatabaseGateway,
		logLoki:             logLoki,
	}
}

func NewCreateNewUser(
	userDatabaseGateway main_gateways.UserDatabaseGateway,
) *CreateNewUser {
	return &CreateNewUser{
		userDatabaseGateway,
		main_configs_apm_logs_impl.NewLogsGatewayImpl(),
	}
}

func (this *CreateNewUser) Execute(ctx context.Context, user main_domains.User) (main_domains.User, main_domains_exceptions.ApplicationException) {

	span := trace.SpanFromContext(ctx)
	defer span.End()

	this.logLoki.INFO(span, fmt.Sprintf("Creating new User with documentNumber: %s", user.DocumentNumber))
	userAlreadyPersisted, err := this.userDatabaseGateway.FindByDocumentNumber(user.DocumentNumber)
	if err != nil {
		return main_domains.User{}, main_domains_exceptions.NewInternalServerErrorExceptionSglMsg(_MSG_CREATE_NEW_DOC_ARCH_ISSUE)
	}
	if !userAlreadyPersisted.IsEmpty() {
		return main_domains.User{}, main_domains_exceptions.NewConflictExceptionSglMsg(
			main_configs_messages.GetMessagesConfigBean().GetDefaultLocale(_MSG_CREATE_NEW_DOC_DOC_ALREADY_EXISTS))
	}

	persistedUser, err := this.userDatabaseGateway.Save(user)
	if err != nil {
		return main_domains.User{}, main_domains_exceptions.NewInternalServerErrorExceptionSglMsg(_MSG_CREATE_NEW_DOC_ARCH_ISSUE)
	}
	return persistedUser, nil
}
