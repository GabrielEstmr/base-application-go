package main_usecases

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_enums "baseapplicationgo/main/domains/enums"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_gateways "baseapplicationgo/main/gateways"
	main_usecases_interfaces "baseapplicationgo/main/usecases/interfaces"
	main_utils_messages "baseapplicationgo/main/utils/messages"
	"context"
	"fmt"
)

type CreateInternalAuthUserPasswordChangeRequest struct {
	_MSG_ARCH_ISSUE                    string
	validateUserIsInternalAuthProvider main_usecases_interfaces.ValidateUserAuthProviderOrigin
	userDatabaseGateway                main_gateways.UserDatabaseGateway
	createUserVerificationEmail        main_usecases_interfaces.CreateAndSendUserVerificationEmail
	logsMonitoringGateway              main_gateways.LogsMonitoringGateway
	spanGateway                        main_gateways.SpanGateway
	messageUtils                       main_utils_messages.ApplicationMessages
}

func NewCreateInternalAuthUserPasswordChangeRequest(
	userDatabaseGateway main_gateways.UserDatabaseGateway,
	createUserVerificationEmail main_usecases_interfaces.CreateAndSendUserVerificationEmail,
	logsMonitoringGateway main_gateways.LogsMonitoringGateway,
	spanGateway main_gateways.SpanGateway,
	messageUtils main_utils_messages.ApplicationMessages,
) *CreateInternalAuthUserPasswordChangeRequest {
	return &CreateInternalAuthUserPasswordChangeRequest{
		_MSG_ARCH_ISSUE:             "exceptions.architecture.application.issue",
		createUserVerificationEmail: createUserVerificationEmail,
		userDatabaseGateway:         userDatabaseGateway,
		logsMonitoringGateway:       logsMonitoringGateway,
		spanGateway:                 spanGateway,
		messageUtils:                messageUtils,
	}
}

func (this *CreateInternalAuthUserPasswordChangeRequest) Execute(
	ctx context.Context,
	userId string,
	databaseOptions main_domains.DatabaseOptions,
) main_domains_exceptions.ApplicationException {

	span := this.spanGateway.Get(ctx, "CreateInternalAuthUserPasswordChangeRequest-Execute")
	defer span.End()
	this.logsMonitoringGateway.DEBUG(span,
		fmt.Sprintf("Creating password change request for user id: %s", userId))

	persistedUser, errValidate := this.validateUserIsInternalAuthProvider.Execute(span.GetCtx(), userId, databaseOptions)
	if errValidate != nil {
		return errValidate
	}

	_, errS := this.createUserVerificationEmail.Execute(
		span.GetCtx(),
		persistedUser,
		main_domains_enums.EMAIL_VERIFICATION_SCOPE_CHANGE_PASSWORD, databaseOptions)

	if errS != nil {
		this.logsMonitoringGateway.ERROR(span, errS.Error())
		return main_domains_exceptions.NewInternalServerErrorExceptionSglMsg(
			this.messageUtils.GetDefaultLocale(
				this._MSG_ARCH_ISSUE))
	}

	return nil
}
