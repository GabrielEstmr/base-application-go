package main_gateways_mongodb

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_enums "baseapplicationgo/main/domains/enums"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_gateways "baseapplicationgo/main/gateways"
	main_gateways_logs "baseapplicationgo/main/gateways/logs"
	main_gateways_mongodb_documents "baseapplicationgo/main/gateways/mongodb/documents"
	main_gateways_mongodb_repositories "baseapplicationgo/main/gateways/mongodb/repositories"
	main_gateways_spans "baseapplicationgo/main/gateways/spans"
	"context"
)

type PlanSettingDatabaseGatewayImpl struct {
	planSettingRepository main_gateways_mongodb_repositories.PlanSettingRepository
	logsMonitoringGateway main_gateways.LogsMonitoringGateway
	spanGateway           main_gateways.SpanGateway
}

func NewPlanSettingDatabaseGatewayImpl(
	planSettingRepository main_gateways_mongodb_repositories.PlanSettingRepository,
) *PlanSettingDatabaseGatewayImpl {
	return &PlanSettingDatabaseGatewayImpl{
		planSettingRepository: planSettingRepository,
		logsMonitoringGateway: main_gateways_logs.NewLogsMonitoringGatewayImpl(),
		spanGateway:           main_gateways_spans.NewSpanGatewayImpl(),
	}
}

func NewPlanSettingDatabaseGatewayImplAllArgs(
	planSettingRepository main_gateways_mongodb_repositories.PlanSettingRepository,
	logsMonitoringGateway main_gateways.LogsMonitoringGateway,
	spanGateway main_gateways.SpanGateway,
) *PlanSettingDatabaseGatewayImpl {
	return &PlanSettingDatabaseGatewayImpl{
		planSettingRepository: planSettingRepository,
		logsMonitoringGateway: logsMonitoringGateway,
		spanGateway:           spanGateway,
	}
}

func (this PlanSettingDatabaseGatewayImpl) Save(
	ctx context.Context,
	planSetting main_domains.PlanSetting,
) (main_domains.PlanSetting, main_domains_exceptions.ApplicationException) {
	span := this.spanGateway.Get(ctx, "PlanSettingDatabaseGatewayImpl-Save")
	defer span.End()

	planSettingDoc := main_gateways_mongodb_documents.NewPlanSettingDocument(planSetting)
	persistedPlanSettingDoc, err := this.planSettingRepository.Save(span.GetCtx(), planSettingDoc)
	if err != nil {
		this.logsMonitoringGateway.ERROR(span, err.Error())
		return *new(main_domains.PlanSetting), main_domains_exceptions.
			NewInternalServerErrorExceptionSglMsg(err.Error())
	}

	return persistedPlanSettingDoc.ToDomain(), nil
}

func (this PlanSettingDatabaseGatewayImpl) FindById(
	ctx context.Context,
	id string,
) (main_domains.PlanSetting, main_domains_exceptions.ApplicationException) {
	span := this.spanGateway.Get(ctx, "PlanSettingDatabaseGatewayImpl-FindById")
	defer span.End()

	planSettingDoc, err := this.planSettingRepository.FindById(span.GetCtx(), id)
	if err != nil {
		this.logsMonitoringGateway.ERROR(span, err.Error())
		return *new(main_domains.PlanSetting), main_domains_exceptions.
			NewInternalServerErrorExceptionSglMsg(err.Error())
	}

	return planSettingDoc.ToDomain(), nil
}

func (this PlanSettingDatabaseGatewayImpl) Update(
	ctx context.Context,
	planSetting main_domains.PlanSetting,
) (main_domains.PlanSetting, main_domains_exceptions.ApplicationException) {
	span := this.spanGateway.Get(ctx, "PlanSettingDatabaseGatewayImpl-Update")
	defer span.End()

	planSettingDoc, err := this.planSettingRepository.Update(
		span.GetCtx(), main_gateways_mongodb_documents.NewPlanSettingDocument(planSetting))
	if err != nil {
		this.logsMonitoringGateway.ERROR(span, err.Error())
		return *new(main_domains.PlanSetting), main_domains_exceptions.
			NewInternalServerErrorExceptionSglMsg(err.Error())
	}

	return planSettingDoc.ToDomain(), nil
}

func (this PlanSettingDatabaseGatewayImpl) FindByPlanTypeAndHasEndDate(
	ctx context.Context,
	planType main_domains_enums.PlanType, hasEndDate bool,
) (
	main_domains.PlanSetting, main_domains_exceptions.ApplicationException) {

	span := this.spanGateway.Get(ctx, "PlanSettingDatabaseGatewayImpl-FindById")
	defer span.End()

	planSettingDoc, err := this.planSettingRepository.FindByPlanTypeAndHasEndDate(span.GetCtx(), planType, hasEndDate)
	if err != nil {
		this.logsMonitoringGateway.ERROR(span, err.Error())
		return *new(main_domains.PlanSetting), main_domains_exceptions.
			NewInternalServerErrorExceptionSglMsg(err.Error())
	}

	return planSettingDoc.ToDomain(), nil
}

func (this PlanSettingDatabaseGatewayImpl) FindByFilter(
	ctx context.Context,
	filter main_domains.FindPlanSettingFilter,
	pageable main_domains.Pageable,
) (
	main_domains.Page,
	main_domains_exceptions.ApplicationException,
) {
	span := this.spanGateway.Get(ctx, "PlanSettingDatabaseGatewayImpl-FindByFilter")
	defer span.End()

	page, err := this.planSettingRepository.FindByFilter(span.GetCtx(), filter, pageable)
	if err != nil {
		this.logsMonitoringGateway.ERROR(span, err.Error())
		return *new(main_domains.Page), main_domains_exceptions.
			NewInternalServerErrorExceptionSglMsg(err.Error())
	}
	contentDoc := page.GetContent()
	var content []any
	for _, value := range contentDoc {
		emailDoc := value.(main_gateways_mongodb_documents.PlanSettingDocument)
		content = append(content, emailDoc.ToDomain())
	}
	return *main_domains.NewPageFromContentAndPage(content, page), nil
}
