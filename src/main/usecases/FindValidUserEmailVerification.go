package main_usecases

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_enums "baseapplicationgo/main/domains/enums"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_gateways "baseapplicationgo/main/gateways"
	main_utils_messages "baseapplicationgo/main/utils/messages"
	"context"
	"fmt"
)

type FindValidUserEmailVerification struct {
	_MSG_USER_NOT_FOUND                  string
	userEmailVerificationDatabaseGateway main_gateways.UserEmailVerificationDatabaseGateway
	logsMonitoringGateway                main_gateways.LogsMonitoringGateway
	spanGateway                          main_gateways.SpanGateway
	messageUtils                         main_utils_messages.ApplicationMessages
}

func NewFindValidUserEmailVerification(
	userEmailVerificationDatabaseGateway main_gateways.UserEmailVerificationDatabaseGateway,
	logsMonitoringGateway main_gateways.LogsMonitoringGateway,
	spanGateway main_gateways.SpanGateway,
	messageBeans main_utils_messages.ApplicationMessages) *FindValidUserEmailVerification {
	return &FindValidUserEmailVerification{
		_MSG_USER_NOT_FOUND:                  "providers.find.user.user.not.found",
		userEmailVerificationDatabaseGateway: userEmailVerificationDatabaseGateway,
		logsMonitoringGateway:                logsMonitoringGateway,
		spanGateway:                          spanGateway,
		messageUtils:                         messageBeans}
}

func (this *FindValidUserEmailVerification) Execute(
	ctx context.Context,
	userId string,
	scope main_domains_enums.UserEmailVerificationScope,
	verificationCode string,
	dbOpt main_domains.DatabaseOptions,
) (main_domains.UserEmailVerification, main_domains_exceptions.ApplicationException) {

	span := this.spanGateway.Get(ctx, "FindValidUserEmailVerification-Execute")
	defer span.End()
	this.logsMonitoringGateway.DEBUG(span,
		fmt.Sprintf("FindValidUserEmailVerification user. code: %s, scope: %s", verificationCode, scope))

	filter := main_domains.FindUserEmailVerificationFilter{
		Scopes:            []string{scope.Name()},
		UserIds:           []string{userId},
		VerificationCodes: []string{verificationCode},
		Statuses:          []main_domains_enums.EmailVerificationStatus{main_domains_enums.EMAIL_VERIFICATION_SENT},
	}

	pageable := *main_domains.NewPageableNoSort(0, 1)

	page, errF := this.userEmailVerificationDatabaseGateway.FindByFilter(span.GetCtx(), filter, pageable, dbOpt)
	if errF != nil {
		return *new(main_domains.UserEmailVerification), errF
	}

	if len(page.GetContent()) == 0 {
		return *new(main_domains.UserEmailVerification), main_domains_exceptions.
			NewResourceNotFoundExceptionSglMsg(this.messageUtils.
				GetDefaultLocale(this._MSG_USER_NOT_FOUND))
	}

	userEmailVerification := page.GetContent()[0].(main_domains.UserEmailVerification)
	if userEmailVerification.IsEmpty() {
		return *new(main_domains.UserEmailVerification), main_domains_exceptions.
			NewResourceNotFoundExceptionSglMsg(this.messageUtils.
				GetDefaultLocale(this._MSG_USER_NOT_FOUND))
	}

	return userEmailVerification, nil
}
