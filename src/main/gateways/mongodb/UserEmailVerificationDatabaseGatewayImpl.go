package main_gateways_mongodb

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_gateways "baseapplicationgo/main/gateways"
	main_gateways_mongodb_documents "baseapplicationgo/main/gateways/mongodb/documents"
	main_gateways_mongodb_repositories "baseapplicationgo/main/gateways/mongodb/repositories"
	main_gateways_spans "baseapplicationgo/main/gateways/spans"
	"context"
)

type UserEmailVerificationDatabaseGatewayImpl struct {
	userEmailVerificationRepository main_gateways_mongodb_repositories.UserEmailVerificationRepository
	spanGateway                     main_gateways.SpanGateway
}

func NewUserEmailVerificationDatabaseGatewayImpl(
	userEmailVerificationRepository main_gateways_mongodb_repositories.UserEmailVerificationRepository,
) *UserEmailVerificationDatabaseGatewayImpl {
	return &UserEmailVerificationDatabaseGatewayImpl{
		userEmailVerificationRepository: userEmailVerificationRepository,
		spanGateway:                     main_gateways_spans.NewSpanGatewayImpl(),
	}
}

func (this *UserEmailVerificationDatabaseGatewayImpl) Save(
	ctx context.Context,
	userEmailVerification main_domains.UserEmailVerification,
	databaseOptions main_domains.DatabaseOptions,
) (main_domains.UserEmailVerification, main_domains_exceptions.ApplicationException) {
	span := this.spanGateway.Get(ctx, "UserEmailVerificationDatabaseGatewayImpl-Save")
	defer span.End()

	userEmailVerificationDoc := main_gateways_mongodb_documents.NewUserEmailVerificationDocument(userEmailVerification)
	persistedUserEmailVerificationDoc, err := this.userEmailVerificationRepository.Save(span.GetCtx(), userEmailVerificationDoc, databaseOptions)
	if err != nil {
		return *new(main_domains.UserEmailVerification), main_domains_exceptions.NewInternalServerErrorExceptionSglMsg(err.Error())
	}
	return persistedUserEmailVerificationDoc.ToDomain(), nil
}

func (this *UserEmailVerificationDatabaseGatewayImpl) Update(
	ctx context.Context,
	userEmailVerification main_domains.UserEmailVerification,
	databaseOptions main_domains.DatabaseOptions,
) (main_domains.UserEmailVerification, main_domains_exceptions.ApplicationException) {
	span := this.spanGateway.Get(ctx, "UserEmailVerificationDatabaseGatewayImpl-Save")
	defer span.End()

	userEmailVerificationDoc := main_gateways_mongodb_documents.NewUserEmailVerificationDocument(userEmailVerification)
	persistedUserEmailVerificationDoc, err := this.userEmailVerificationRepository.Update(span.GetCtx(), userEmailVerificationDoc, databaseOptions)
	if err != nil {
		return *new(main_domains.UserEmailVerification), main_domains_exceptions.NewInternalServerErrorExceptionSglMsg(err.Error())
	}
	return persistedUserEmailVerificationDoc.ToDomain(), nil
}

func (this *UserEmailVerificationDatabaseGatewayImpl) FindById(
	ctx context.Context,
	id string,
	databaseOptions main_domains.DatabaseOptions,
) (main_domains.UserEmailVerification, main_domains_exceptions.ApplicationException) {
	span := this.spanGateway.Get(ctx, "UserEmailVerificationDatabaseGatewayImpl-FindById")
	defer span.End()

	userEmailVerificationDoc, err := this.userEmailVerificationRepository.FindById(span.GetCtx(), id, databaseOptions)
	if err != nil {
		return *new(main_domains.UserEmailVerification), main_domains_exceptions.NewInternalServerErrorExceptionSglMsg(err.Error())
	}
	return userEmailVerificationDoc.ToDomain(), nil
}

func (this *UserEmailVerificationDatabaseGatewayImpl) FindByFilter(
	ctx context.Context,
	filter main_domains.FindUserEmailVerificationFilter,
	pageable main_domains.Pageable,
	databaseOptions main_domains.DatabaseOptions,
) (
	main_domains.Page, main_domains_exceptions.ApplicationException,
) {
	span := this.spanGateway.Get(ctx, "UserEmailVerificationDatabaseGatewayImpl-FindByFilter")
	defer span.End()

	page, err := this.userEmailVerificationRepository.FindByFilter(span.GetCtx(), filter, pageable, databaseOptions)
	if err != nil {
		return *new(main_domains.Page), main_domains_exceptions.NewInternalServerErrorExceptionSglMsg(err.Error())
	}
	contentDoc := page.GetContent()
	var content []any
	for _, value := range contentDoc {
		userDoc := value.(main_gateways_mongodb_documents.UserEmailVerificationDocument)
		content = append(content, userDoc.ToDomain())
	}
	return *main_domains.NewPageFromContentAndPage(content, page), nil
}
