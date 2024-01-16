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

const _MSG_FIND_USER_BY_FILTER_ARCH_ISSUE = "exceptions.architecture.application.issue"

type FindUsersByFilter struct {
	userDatabaseGateway   main_gateways.UserDatabaseGateway
	messageUtils          main_utils_messages.ApplicationMessages
	spanGateway           main_gateways.SpanGateway
	logsMonitoringGateway main_gateways.LogsMonitoringGateway
}

func NewFindUsersByFilter(
	userDatabaseGateway main_gateways.UserDatabaseGateway,
) *FindUsersByFilter {
	return &FindUsersByFilter{
		userDatabaseGateway,
		*main_utils_messages.NewApplicationMessages(),
		main_gateways_spans.NewSpanGatewayImpl(),
		main_gateways_logs.NewLogsMonitoringGatewayImpl(),
	}
}

func (this *FindUsersByFilter) Execute(
	ctx context.Context,
	filter main_domains.FindUserFilter,
	pageable main_domains.Pageable) (main_domains.Page, main_domains_exceptions.ApplicationException) {
	span := this.spanGateway.Get(ctx, "FindUsersByFilter-Execute")
	defer span.End()
	this.logsMonitoringGateway.DEBUG(span, fmt.Sprintf("FindUsersByFilter-Execute"))

	page, err := this.userDatabaseGateway.FindByFilter(span.GetCtx(), filter, pageable)
	if err != nil {
		return main_domains.Page{},
			main_domains_exceptions.NewInternalServerErrorExceptionSglMsg(
				this.messageUtils.GetDefaultLocale(
					_MSG_FIND_USER_BY_FILTER_ARCH_ISSUE))
	}
	return page, nil
}
