package main_gateways_mongodb

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_gateways "baseapplicationgo/main/gateways"
	main_gateways_logs "baseapplicationgo/main/gateways/logs"
	main_gateways_mongodb_documents "baseapplicationgo/main/gateways/mongodb/documents"
	main_gateways_mongodb_repositories "baseapplicationgo/main/gateways/mongodb/repositories"
	main_gateways_spans "baseapplicationgo/main/gateways/spans"
	"context"
)

type EmailDatabaseGatewayImpl struct {
	emailRepository       main_gateways_mongodb_repositories.EmailRepository
	logsMonitoringGateway main_gateways.LogsMonitoringGateway
	spanGateway           main_gateways.SpanGateway
}

func NewEmailDatabaseGatewayImpl(
	emailRepository main_gateways_mongodb_repositories.EmailRepository,
) *EmailDatabaseGatewayImpl {
	return &EmailDatabaseGatewayImpl{
		emailRepository:       emailRepository,
		logsMonitoringGateway: main_gateways_logs.NewLogsMonitoringGatewayImpl(),
		spanGateway:           main_gateways_spans.NewSpanGatewayImpl(),
	}
}

func NewEmailDatabaseGatewayImplAllArgs(
	emailRepository main_gateways_mongodb_repositories.EmailRepository,
	logsMonitoringGateway main_gateways.LogsMonitoringGateway,
	spanGateway main_gateways.SpanGateway,
) *EmailDatabaseGatewayImpl {
	return &EmailDatabaseGatewayImpl{
		emailRepository:       emailRepository,
		logsMonitoringGateway: logsMonitoringGateway,
		spanGateway:           spanGateway,
	}
}

func (this EmailDatabaseGatewayImpl) Save(
	ctx context.Context,
	email main_domains.Email,
) (main_domains.Email, main_domains_exceptions.ApplicationException) {
	span := this.spanGateway.Get(ctx, "EmailDatabaseGatewayImpl-Save")
	defer span.End()

	emailDoc := main_gateways_mongodb_documents.NewEmailDocument(email)
	persistedEmailDoc, err := this.emailRepository.Save(span.GetCtx(), emailDoc)
	if err != nil {
		this.logsMonitoringGateway.ERROR(span, err.Error())
		return *new(main_domains.Email), main_domains_exceptions.
			NewInternalServerErrorExceptionSglMsg(err.Error())
	}

	return persistedEmailDoc.ToDomain(), nil
}

func (this EmailDatabaseGatewayImpl) FindById(
	ctx context.Context,
	id string,
) (main_domains.Email, main_domains_exceptions.ApplicationException) {
	span := this.spanGateway.Get(ctx, "EmailDatabaseGatewayImpl-FindById")
	defer span.End()

	emailDoc, err := this.emailRepository.FindById(span.GetCtx(), id)
	if err != nil {
		this.logsMonitoringGateway.ERROR(span, err.Error())
		return *new(main_domains.Email), main_domains_exceptions.
			NewInternalServerErrorExceptionSglMsg(err.Error())
	}

	return emailDoc.ToDomain(), nil
}

func (this EmailDatabaseGatewayImpl) FindByEventId(
	ctx context.Context,
	eventId string,
) (main_domains.Email, main_domains_exceptions.ApplicationException) {
	span := this.spanGateway.Get(ctx, "EmailDatabaseGatewayImpl-FindById")
	defer span.End()

	emailDoc, err := this.emailRepository.FindByEventId(span.GetCtx(), eventId)
	if err != nil {
		this.logsMonitoringGateway.ERROR(span, err.Error())
		return *new(main_domains.Email), main_domains_exceptions.
			NewInternalServerErrorExceptionSglMsg(err.Error())
	}

	return emailDoc.ToDomain(), nil
}

func (this EmailDatabaseGatewayImpl) Update(
	ctx context.Context,
	email main_domains.Email,
) (main_domains.Email, main_domains_exceptions.ApplicationException) {
	span := this.spanGateway.Get(ctx, "EmailDatabaseGatewayImpl-Update")
	defer span.End()

	emailDoc, err := this.emailRepository.Update(span.GetCtx(), main_gateways_mongodb_documents.NewEmailDocument(email))
	if err != nil {
		this.logsMonitoringGateway.ERROR(span, err.Error())
		return *new(main_domains.Email), main_domains_exceptions.
			NewInternalServerErrorExceptionSglMsg(err.Error())
	}

	return emailDoc.ToDomain(), nil
}

func (this EmailDatabaseGatewayImpl) FindByFilter(
	ctx context.Context,
	filter main_domains.FindEmailFilter,
	pageable main_domains.Pageable,
) (
	main_domains.Page,
	main_domains_exceptions.ApplicationException,
) {
	span := this.spanGateway.Get(ctx, "EmailDatabaseGatewayImpl-FindByFilter")
	defer span.End()

	page, err := this.emailRepository.FindByFilter(span.GetCtx(), filter, pageable)
	if err != nil {
		this.logsMonitoringGateway.ERROR(span, err.Error())
		return *new(main_domains.Page), main_domains_exceptions.
			NewInternalServerErrorExceptionSglMsg(err.Error())
	}
	contentDoc := page.GetContent()
	var content []any
	for _, value := range contentDoc {
		emailDoc := value.(main_gateways_mongodb_documents.EmailDocument)
		content = append(content, emailDoc.ToDomain())
	}
	return *main_domains.NewPageFromContentAndPage(content, page), nil
}
