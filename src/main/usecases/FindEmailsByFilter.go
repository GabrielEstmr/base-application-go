package main_usecases

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_gateways "baseapplicationgo/main/gateways"
	main_utils_messages "baseapplicationgo/main/utils/messages"
	"context"
	"fmt"
)

const _MSG_FIND_EMAILS_BY_FILTER_ARCH_ISSUE = "exceptions.architecture.application.issue"

type FindEmailsByFilter struct {
	emailDatabaseGateway  main_gateways.EmailDatabaseGateway
	messageUtils          main_utils_messages.ApplicationMessages
	spanGateway           main_gateways.SpanGateway
	logsMonitoringGateway main_gateways.LogsMonitoringGateway
}

func NewFindEmailsByFilter(
	emailDatabaseGateway main_gateways.EmailDatabaseGateway,
	messageUtils main_utils_messages.ApplicationMessages,
	spanGateway main_gateways.SpanGateway,
	logsMonitoringGateway main_gateways.LogsMonitoringGateway,
) *FindEmailsByFilter {
	return &FindEmailsByFilter{
		emailDatabaseGateway:  emailDatabaseGateway,
		messageUtils:          messageUtils,
		spanGateway:           spanGateway,
		logsMonitoringGateway: logsMonitoringGateway,
	}
}

func (this *FindEmailsByFilter) Execute(
	ctx context.Context,
	filter main_domains.FindEmailFilter,
	pageable main_domains.Pageable) (main_domains.Page, main_domains_exceptions.ApplicationException) {
	span := this.spanGateway.Get(ctx, "FindEmailsByFilter-Execute")
	defer span.End()
	this.logsMonitoringGateway.DEBUG(span, fmt.Sprintf("FindEmailsByFilter-Execute"))

	page, err := this.emailDatabaseGateway.FindByFilter(span.GetCtx(), filter, pageable)
	if err != nil {
		return *new(main_domains.Page),
			main_domains_exceptions.NewInternalServerErrorExceptionSglMsg(
				this.messageUtils.GetDefaultLocale(
					_MSG_FIND_EMAILS_BY_FILTER_ARCH_ISSUE))
	}
	return page, nil
}
