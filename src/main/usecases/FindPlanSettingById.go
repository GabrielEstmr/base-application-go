package main_usecases

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_gateways "baseapplicationgo/main/gateways"
	main_utils_messages "baseapplicationgo/main/utils/messages"
	"context"
	"fmt"
)

const _MSG_FIND_PLAN_SETTING_BY_ID_DOC_NOT_FOUND = "find.user.user.not.found"
const _MSG_FIND_PLAN_SETTING_BY_ID_ARCH_ISSUE = "exceptions.architecture.application.issue"

type FindPlanSettingById struct {
	planSettingDatabaseGateway main_gateways.PlanSettingDatabaseGateway
	messageUtils               main_utils_messages.ApplicationMessages
	featuresGateway            main_gateways.FeaturesGateway
	spanGateway                main_gateways.SpanGateway
	logsMonitoringGateway      main_gateways.LogsMonitoringGateway
}

func NewFindPlanSettingById(
	planSettingDatabaseGateway main_gateways.PlanSettingDatabaseGateway,
	messageUtils main_utils_messages.ApplicationMessages,
	featuresGateway main_gateways.FeaturesGateway,
	spanGateway main_gateways.SpanGateway,
	logsMonitoringGateway main_gateways.LogsMonitoringGateway,
) *FindPlanSettingById {
	return &FindPlanSettingById{
		planSettingDatabaseGateway: planSettingDatabaseGateway,
		messageUtils:               messageUtils,
		featuresGateway:            featuresGateway,
		spanGateway:                spanGateway,
		logsMonitoringGateway:      logsMonitoringGateway,
	}
}

func (this *FindPlanSettingById) Execute(ctx context.Context, id string,
) (main_domains.PlanSetting, main_domains_exceptions.ApplicationException) {
	span := this.spanGateway.Get(ctx, "FindPlanSettingById-Execute")
	defer span.End()
	this.logsMonitoringGateway.INFO(span, fmt.Sprintf("FindPlanSettingById-Execute. id: %s", id))

	planSetting, err := this.planSettingDatabaseGateway.FindById(span.GetCtx(), id)
	if err != nil {
		return *new(main_domains.PlanSetting),
			main_domains_exceptions.NewInternalServerErrorExceptionSglMsg(_MSG_FIND_PLAN_SETTING_BY_ID_ARCH_ISSUE)
	}
	if planSetting.IsEmpty() {
		return *new(main_domains.PlanSetting), main_domains_exceptions.NewResourceNotFoundExceptionSglMsg(
			this.messageUtils.GetDefaultLocale(_MSG_FIND_PLAN_SETTING_BY_ID_DOC_NOT_FOUND))
	}

	return planSetting, nil
}
