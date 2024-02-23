package main_usecases

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_gateways "baseapplicationgo/main/gateways"
	main_gateways_logs "baseapplicationgo/main/gateways/logs"
	main_gateways_spans "baseapplicationgo/main/gateways/spans"
	main_utils_messages "baseapplicationgo/main/utils/messages"
	"context"
	"fmt"
)

const _MSG_FIND_USER_BY_ID_DOC_NOT_FOUND = "find.user.user.not.found"
const _MSG_FIND_USER_BY_ID_ARCH_ISSUE = "exceptions.architecture.application.issue"

type FindUserById struct {
	userDatabaseGateway   main_gateways.UserDatabaseGateway
	messageUtils          main_utils_messages.ApplicationMessages
	featuresGateway       main_gateways.FeaturesGateway
	spanGateway           main_gateways.SpanGateway
	logsMonitoringGateway main_gateways.LogsMonitoringGateway
}

func NewFindUserByIdAllArgs(
	userDatabaseGateway main_gateways.UserDatabaseGateway,
	messageUtils main_utils_messages.ApplicationMessages,
	featuresGateway main_gateways.FeaturesGateway,
	spanGateway main_gateways.SpanGateway,
	logsMonitoringGateway main_gateways.LogsMonitoringGateway,
) *FindUserById {
	return &FindUserById{
		userDatabaseGateway:   userDatabaseGateway,
		messageUtils:          messageUtils,
		featuresGateway:       featuresGateway,
		spanGateway:           spanGateway,
		logsMonitoringGateway: logsMonitoringGateway,
	}
}

func NewFindUserById(
	userDatabaseGateway main_gateways.UserDatabaseGateway,
	featuresGateway main_gateways.FeaturesGateway,
) *FindUserById {
	return &FindUserById{
		userDatabaseGateway,
		*main_utils_messages.NewApplicationMessages(),
		featuresGateway,
		main_gateways_spans.NewSpanGatewayImpl(),
		main_gateways_logs.NewLogsMonitoringGatewayImpl(),
	}
}

func (this *FindUserById) Execute(ctx context.Context, id string) (main_domains.User, main_domains_exceptions.ApplicationException) {
	span := this.spanGateway.Get(ctx, "FindUserById-Execute")
	defer span.End()
	this.logsMonitoringGateway.INFO(span, fmt.Sprintf("FindUserById-Execute. id: %s", id))

	user, err := this.userDatabaseGateway.FindById(span.GetCtx(), id, nil)
	if err != nil {
		return main_domains.User{},
			main_domains_exceptions.NewInternalServerErrorExceptionSglMsg(_MSG_FIND_USER_BY_ID_ARCH_ISSUE)
	}
	if user.IsEmpty() {
		return main_domains.User{}, main_domains_exceptions.NewResourceNotFoundExceptionSglMsg(
			this.messageUtils.GetDefaultLocale(_MSG_FIND_USER_BY_ID_DOC_NOT_FOUND))
	}

	return user, nil
}
