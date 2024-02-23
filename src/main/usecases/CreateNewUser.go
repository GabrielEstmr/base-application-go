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

type CreateNewUser struct {
	_MSG_ARCH_ISSUE                      string
	createOrRetrieveUserFromAuthProvider main_usecases_interfaces.CreateOrRetrieveUserFromAuthProvider
	validatePasswordFormat               main_usecases_interfaces.ValidatePasswordFormat
	validateHasUserByProperties          main_usecases_interfaces.ValidateUserByProperty
	userDatabaseGateway                  main_gateways.UserDatabaseGateway
	createUserVerificationEmail          main_usecases_interfaces.CreateAndSendUserVerificationEmail
	logsMonitoringGateway                main_gateways.LogsMonitoringGateway
	spanGateway                          main_gateways.SpanGateway
	messageUtils                         main_utils_messages.ApplicationMessages
}

func NewCreateNewUser(
	createOrRetrieveUserFromAuthProvider main_usecases_interfaces.CreateOrRetrieveUserFromAuthProvider,
	validatePasswordFormat main_usecases_interfaces.ValidatePasswordFormat,
	validateHasUserByProperties main_usecases_interfaces.ValidateUserByProperty,
	userDatabaseGateway main_gateways.UserDatabaseGateway,
	createUserVerificationEmail main_usecases_interfaces.CreateAndSendUserVerificationEmail,
	logsMonitoringGateway main_gateways.LogsMonitoringGateway,
	spanGateway main_gateways.SpanGateway,
	messageBeans main_utils_messages.ApplicationMessages) *CreateNewUser {
	return &CreateNewUser{
		_MSG_ARCH_ISSUE:                      "exceptions.architecture.application.issue",
		createOrRetrieveUserFromAuthProvider: createOrRetrieveUserFromAuthProvider,
		validatePasswordFormat:               validatePasswordFormat,
		validateHasUserByProperties:          validateHasUserByProperties,
		userDatabaseGateway:                  userDatabaseGateway,
		createUserVerificationEmail:          createUserVerificationEmail,
		logsMonitoringGateway:                logsMonitoringGateway,
		spanGateway:                          spanGateway,
		messageUtils:                         messageBeans}
}

func (this *CreateNewUser) Execute(
	ctx context.Context,
	user main_domains.User,
	databaseOptions main_domains.DatabaseOptions) (
	main_domains.User, main_domains_exceptions.ApplicationException) {

	span := this.spanGateway.Get(ctx, "CreateNewUser-Execute")
	defer span.End()
	this.logsMonitoringGateway.INFO(span,
		fmt.Sprintf("Creating new User with email: %s", user.GetEmail()))

	errPassValidation := this.validatePasswordFormat.Execute(span.GetCtx(), user.GetPassword())
	if errPassValidation != nil {
		return *new(main_domains.User), errPassValidation
	}

	errV := this.validateHasUserByProperties.Execute(span.GetCtx(), user, databaseOptions)
	if errV != nil {
		return *new(main_domains.User), errV
	}

	updatedUser, errAP := this.createOrRetrieveUserFromAuthProvider.Execute(span.GetCtx(), user)
	if errAP != nil {
		return *new(main_domains.User), errAP
	}

	persistedUser, err := this.userDatabaseGateway.Save(span.GetCtx(), updatedUser, databaseOptions)
	if err != nil {
		return updatedUser, main_domains_exceptions.
			NewInternalServerErrorExceptionSglMsg(this.messageUtils.
				GetDefaultLocale(this._MSG_ARCH_ISSUE))
	}

	_, errE := this.createUserVerificationEmail.Execute(
		span.GetCtx(),
		persistedUser,
		main_domains_enums.EMAIL_VERIFICATION_SCOPE_ENABLE_USER,
		databaseOptions)
	if errE != nil {
		return persistedUser, errE
	}

	return persistedUser, nil

}
