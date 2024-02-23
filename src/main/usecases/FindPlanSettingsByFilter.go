package main_usecases

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_gateways "baseapplicationgo/main/gateways"
	main_utils_messages "baseapplicationgo/main/utils/messages"
	"context"
	"fmt"
)

const _MSG_FIND_PLAN_SETTING_BY_FILTER_ARCH_ISSUE = "exceptions.architecture.application.issue"

type FindPlanSettingsByFilter struct {
	planSettingDatabaseGateway main_gateways.PlanSettingDatabaseGateway
	messageUtils               main_utils_messages.ApplicationMessages
	spanGateway                main_gateways.SpanGateway
	logsMonitoringGateway      main_gateways.LogsMonitoringGateway
}

func NewFindPlanSettingsByFilter(
	planSettingDatabaseGateway main_gateways.PlanSettingDatabaseGateway,
	messageUtils main_utils_messages.ApplicationMessages,
	spanGateway main_gateways.SpanGateway,
	logsMonitoringGateway main_gateways.LogsMonitoringGateway,
) *FindPlanSettingsByFilter {
	return &FindPlanSettingsByFilter{
		planSettingDatabaseGateway: planSettingDatabaseGateway,
		messageUtils:               messageUtils,
		spanGateway:                spanGateway,
		logsMonitoringGateway:      logsMonitoringGateway,
	}
}

func (this *FindPlanSettingsByFilter) Execute(
	ctx context.Context,
	filter main_domains.FindPlanSettingFilter,
	pageable main_domains.Pageable) (main_domains.Page, main_domains_exceptions.ApplicationException) {
	span := this.spanGateway.Get(ctx, "FindPlanSettingsByFilter-Execute")
	defer span.End()
	this.logsMonitoringGateway.INFO(span, fmt.Sprintf("FindPlanSettingsByFilter-Execute"))

	page, err := this.planSettingDatabaseGateway.FindByFilter(span.GetCtx(), filter, pageable)
	if err != nil {
		return *new(main_domains.Page),
			main_domains_exceptions.NewInternalServerErrorExceptionSglMsg(
				this.messageUtils.GetDefaultLocale(
					_MSG_FIND_PLAN_SETTING_BY_FILTER_ARCH_ISSUE))
	}
	return page, nil
}
