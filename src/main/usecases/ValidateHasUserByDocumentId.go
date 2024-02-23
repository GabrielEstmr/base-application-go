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

type ValidateHasUserByDocumentId struct {
	userDatabaseGateway   main_gateways.UserDatabaseGateway
	logsMonitoringGateway main_gateways.LogsMonitoringGateway
	spanGateway           main_gateways.SpanGateway
	messageUtils          main_utils_messages.ApplicationMessages
}

func NewValidateHasUserByDocumentId(
	userDatabaseGateway main_gateways.UserDatabaseGateway,
	logsMonitoringGateway main_gateways.LogsMonitoringGateway,
	spanGateway main_gateways.SpanGateway,
	messageUtils main_utils_messages.ApplicationMessages,
) *ValidateHasUserByDocumentId {
	return &ValidateHasUserByDocumentId{
		userDatabaseGateway:   userDatabaseGateway,
		logsMonitoringGateway: logsMonitoringGateway,
		spanGateway:           spanGateway,
		messageUtils:          messageUtils,
	}
}

func (this *ValidateHasUserByDocumentId) Execute(
	ctx context.Context, user main_domains.User, databaseOptions main_domains.DatabaseOptions) main_domains_exceptions.ApplicationException {

	span := this.spanGateway.Get(ctx, "ValidateHasUserByUserName-Execute")
	defer span.End()
	this.logsMonitoringGateway.DEBUG(span,
		fmt.Sprintf("Validating user with documentId: %s", user.GetDocumentId()))

	userAlreadyPersisted, err := this.userDatabaseGateway.FindByDocumentId(ctx, user.GetDocumentId(), databaseOptions)
	if err != nil {
		return main_domains_exceptions.
			NewInternalServerErrorExceptionSglMsg(err.Error())
	}
	if !userAlreadyPersisted.IsEmpty() {
		return main_domains_exceptions.NewConflictExceptionSglMsg(
			this.messageUtils.
				GetDefaultLocale(_MSG_CREATE_NEW_DOC_DOC_ALREADY_EXISTS))
	}
	return nil
}
