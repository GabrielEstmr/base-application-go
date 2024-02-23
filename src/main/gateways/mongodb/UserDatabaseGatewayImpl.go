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

type UserDatabaseGatewayImpl struct {
	userRepository main_gateways_mongodb_repositories.UserRepository
	spanGateway    main_gateways.SpanGateway
}

func NewUserDatabaseGatewayImpl(userRepository main_gateways_mongodb_repositories.UserRepository) *UserDatabaseGatewayImpl {
	return &UserDatabaseGatewayImpl{userRepository,
		main_gateways_spans.NewSpanGatewayImpl(),
	}
}

func (this *UserDatabaseGatewayImpl) Save(
	ctx context.Context,
	user main_domains.User,
	options main_domains.DatabaseOptions,
) (main_domains.User, main_domains_exceptions.ApplicationException) {
	span := this.spanGateway.Get(ctx, "UserDatabaseGatewayImpl-Save")
	defer span.End()

	userDocument := main_gateways_mongodb_documents.NewUserDocument(user)
	persistedUserDocument, err := this.userRepository.Save(span.GetCtx(), userDocument, options)
	if err != nil {
		return *new(main_domains.User), main_domains_exceptions.NewInternalServerErrorExceptionSglMsg(err.Error())
	}
	return persistedUserDocument.ToDomain(), nil
}

func (this *UserDatabaseGatewayImpl) Update(
	ctx context.Context,
	user main_domains.User,
	options main_domains.DatabaseOptions,
) (main_domains.User, main_domains_exceptions.ApplicationException) {
	span := this.spanGateway.Get(ctx, "UserDatabaseGatewayImpl-Save")
	defer span.End()

	userDocument := main_gateways_mongodb_documents.NewUserDocument(user)
	persistedUserDocument, err := this.userRepository.Update(span.GetCtx(), userDocument, options)
	if err != nil {
		return *new(main_domains.User), main_domains_exceptions.NewInternalServerErrorExceptionSglMsg(err.Error())
	}
	return persistedUserDocument.ToDomain(), nil
}

func (this *UserDatabaseGatewayImpl) FindById(
	ctx context.Context,
	id string,
	options main_domains.DatabaseOptions,
) (main_domains.User, main_domains_exceptions.ApplicationException) {
	span := this.spanGateway.Get(ctx, "UserDatabaseGatewayImpl-FindById")
	defer span.End()

	userDocument, err := this.userRepository.FindById(span.GetCtx(), id, options)
	if err != nil {
		return *new(main_domains.User), main_domains_exceptions.NewInternalServerErrorExceptionSglMsg(err.Error())
	}
	return userDocument.ToDomain(), nil
}

func (this *UserDatabaseGatewayImpl) FindByDocumentId(
	ctx context.Context,
	documentId string,
	options main_domains.DatabaseOptions,
) (main_domains.User, main_domains_exceptions.ApplicationException) {
	span := this.spanGateway.Get(ctx, "UserDatabaseGatewayImpl-FindByDocumentId")
	defer span.End()

	userDocument, err := this.userRepository.FindByDocumentId(span.GetCtx(), documentId, options)
	if err != nil {
		return *new(main_domains.User), main_domains_exceptions.NewInternalServerErrorExceptionSglMsg(err.Error())
	}
	return userDocument.ToDomain(), nil
}

func (this *UserDatabaseGatewayImpl) FindByUserName(
	ctx context.Context,
	userName string,
	options main_domains.DatabaseOptions,
) (main_domains.User, main_domains_exceptions.ApplicationException) {
	span := this.spanGateway.Get(ctx, "UserDatabaseGatewayImpl-FindByUserName")
	defer span.End()

	userDocument, err := this.userRepository.FindByUserName(span.GetCtx(), userName, options)
	if err != nil {
		return *new(main_domains.User), main_domains_exceptions.NewInternalServerErrorExceptionSglMsg(err.Error())
	}
	return userDocument.ToDomain(), nil
}

func (this *UserDatabaseGatewayImpl) FindByEmail(
	ctx context.Context,
	email string,
	options main_domains.DatabaseOptions,
) (main_domains.User, main_domains_exceptions.ApplicationException) {
	span := this.spanGateway.Get(ctx, "UserDatabaseGatewayImpl-FindByEmail")
	defer span.End()

	userDocument, err := this.userRepository.FindByEmail(span.GetCtx(), email, options)
	if err != nil {
		return *new(main_domains.User), main_domains_exceptions.NewInternalServerErrorExceptionSglMsg(err.Error())
	}
	return userDocument.ToDomain(), nil
}

func (this *UserDatabaseGatewayImpl) FindByFilter(
	ctx context.Context,
	filter main_domains.FindUserFilter,
	pageable main_domains.Pageable,
	options main_domains.DatabaseOptions,
) (main_domains.Page, main_domains_exceptions.ApplicationException) {
	span := this.spanGateway.Get(ctx, "UserDatabaseGatewayImpl-FindByFilter")
	defer span.End()

	page, err := this.userRepository.FindByFilter(span.GetCtx(), filter, pageable, options)
	if err != nil {
		return main_domains.Page{}, main_domains_exceptions.NewInternalServerErrorExceptionSglMsg(err.Error())
	}
	contentDoc := page.GetContent()
	var content []any
	for _, value := range contentDoc {
		userDoc := value.(main_gateways_mongodb_documents.UserDocument)
		content = append(content, userDoc.ToDomain())
	}
	return *main_domains.NewPageFromContentAndPage(content, page), nil
}
