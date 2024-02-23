package main_usecases

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_gateways "baseapplicationgo/main/gateways"
	main_utils "baseapplicationgo/main/utils"
	main_utils_messages "baseapplicationgo/main/utils/messages"
	"errors"
	"fmt"
	"golang.org/x/net/context"
	"time"
)

type CreateNewPlanSetting struct {
	_MSG_ARCH_ISSUE            string
	planSettingDatabaseGateway main_gateways.PlanSettingDatabaseGateway
	dateUtils                  main_utils.DateUtils
	logsMonitoringGateway      main_gateways.LogsMonitoringGateway
	spanGateway                main_gateways.SpanGateway
	messageUtils               main_utils_messages.ApplicationMessages
}

func NewCreateNewPlanSettingAllArgs(
	planSettingDatabaseGateway main_gateways.PlanSettingDatabaseGateway,
	dateUtils main_utils.DateUtils,
	logsMonitoringGateway main_gateways.LogsMonitoringGateway,
	spanGateway main_gateways.SpanGateway,
	messageBeans main_utils_messages.ApplicationMessages) *CreateNewPlanSetting {
	return &CreateNewPlanSetting{
		planSettingDatabaseGateway: planSettingDatabaseGateway,
		dateUtils:                  dateUtils,
		logsMonitoringGateway:      logsMonitoringGateway,
		spanGateway:                spanGateway,
		messageUtils:               messageBeans}
}

func NewCreateNewPlanSetting(
	planSettingDatabaseGateway main_gateways.PlanSettingDatabaseGateway,
	dateUtils main_utils.DateUtils,
	logsMonitoringGateway main_gateways.LogsMonitoringGateway,
	spanGateway main_gateways.SpanGateway,
) *CreateNewPlanSetting {
	return &CreateNewPlanSetting{
		_MSG_ARCH_ISSUE:            "exceptions.architecture.application.issue",
		planSettingDatabaseGateway: planSettingDatabaseGateway,
		dateUtils:                  dateUtils,
		logsMonitoringGateway:      logsMonitoringGateway,
		spanGateway:                spanGateway,
		messageUtils:               *main_utils_messages.NewApplicationMessages(),
	}
}

func (this *CreateNewPlanSetting) getDisableDate(planSettingToCreate main_domains.PlanSetting) (time.Time, main_domains_exceptions.ApplicationException) {
	startDate := planSettingToCreate.GetStartDate()
	var disableDate time.Time
	if startDate.IsZero() {
		disableDate = this.dateUtils.GetDateUTCAtEndOfTheDay(time.Now())
	} else {
		disableDate = this.dateUtils.GetDateUTCAtEndOfTheDay(this.dateUtils.AddDays(startDate, -1))
	}
	return disableDate, nil
}

func (this *CreateNewPlanSetting) getReferenceDates(
	planSettingToCreate main_domains.PlanSetting) (time.Time, time.Time) {
	planSettingStartDate := planSettingToCreate.GetStartDate()
	var startDate time.Time
	if planSettingStartDate.IsZero() {
		startDate = this.dateUtils.GetDateUTCAtStartOfTheDay(this.dateUtils.AddDays(time.Now(), 1))
	} else {
		startDate = this.dateUtils.GetDateUTCAtStartOfTheDay(startDate)
	}
	disableDate := this.dateUtils.GetDateUTCAtEndOfTheDay(this.dateUtils.AddDays(startDate, -1))
	return startDate, disableDate
}

func (this *CreateNewPlanSetting) validateStartDate(persistedPlanSetting main_domains.PlanSetting, startDate time.Time) error {
	if startDate.Compare(persistedPlanSetting.GetStartDate()) == -1 {
		return errors.New(this.messageUtils.GetDefaultLocale(this._MSG_ARCH_ISSUE))
	}
	return nil
}

func (this *CreateNewPlanSetting) Execute(ctx context.Context, planSetting main_domains.PlanSetting) (
	main_domains.PlanSetting, main_domains_exceptions.ApplicationException) {

	span := this.spanGateway.Get(ctx, "CreateNewPlanSetting-Execute")
	defer span.End()
	this.logsMonitoringGateway.INFO(span,
		fmt.Sprintf("Creating new PlanSetting for planType: %s", planSetting.GetPlanType()))

	persistedPlanSetting, err := this.planSettingDatabaseGateway.FindByPlanTypeAndHasEndDate(
		span.GetCtx(), planSetting.GetPlanType(), false)
	if err != nil {
		return *new(main_domains.PlanSetting), main_domains_exceptions.
			NewInternalServerErrorExceptionSglMsg(this.messageUtils.
				GetDefaultLocale(this._MSG_ARCH_ISSUE))
	}

	startDate, disableDate := this.getReferenceDates(planSetting)

	errV := this.validateStartDate(persistedPlanSetting, startDate)
	if errV != nil {
		return *new(main_domains.PlanSetting), main_domains_exceptions.
			NewConflictExceptionSglMsg(this.messageUtils.
				GetDefaultLocale(this._MSG_ARCH_ISSUE))
	}

	if !persistedPlanSetting.IsEmpty() {
		planSettingToDisable := persistedPlanSetting.CloneAsDisabled(disableDate)
		_, errDisable := this.planSettingDatabaseGateway.Update(span.GetCtx(), planSettingToDisable)
		if errDisable != nil {
			return *new(main_domains.PlanSetting), main_domains_exceptions.
				NewInternalServerErrorExceptionSglMsg(this.messageUtils.
					GetDefaultLocale(this._MSG_ARCH_ISSUE))
		}
	}

	// create New One
	savedPlanSetting, errS := this.planSettingDatabaseGateway.Save(span.GetCtx(), planSetting.CloneWithStartDateAtStartOfTheDay())
	if errS != nil {
		return *new(main_domains.PlanSetting), main_domains_exceptions.
			NewInternalServerErrorExceptionSglMsg(this.messageUtils.
				GetDefaultLocale(this._MSG_ARCH_ISSUE))
	}
	return savedPlanSetting, nil

}
