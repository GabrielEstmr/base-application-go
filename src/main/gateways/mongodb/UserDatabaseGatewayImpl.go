package main_gateways_mongodb

import (
	main_domains "baseapplicationgo/main/domains"
	main_gateways "baseapplicationgo/main/gateways"
	main_gateways_mongodb_documents "baseapplicationgo/main/gateways/mongodb/documents"
	main_gateways_mongodb_repositories "baseapplicationgo/main/gateways/mongodb/repositories"
	main_gateways_spans "baseapplicationgo/main/gateways/spans"
	"context"
	"time"
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

func (this *UserDatabaseGatewayImpl) Save(ctx context.Context, user main_domains.User) (main_domains.User, error) {
	span := this.spanGateway.Get(ctx, "UserDatabaseGatewayImpl-Save")
	defer span.End()

	now := time.Now()
	user.CreatedDate = now
	user.LastModifiedDate = now
	userDocument := main_gateways_mongodb_documents.NewUserDocument(user)
	persistedUserDocument, err := this.userRepository.Save(span.GetCtx(), userDocument)
	if err != nil {
		return main_domains.User{}, err
	}
	return persistedUserDocument.ToDomain(), nil
}

func (this *UserDatabaseGatewayImpl) FindById(ctx context.Context, id string) (main_domains.User, error) {
	span := this.spanGateway.Get(ctx, "UserDatabaseGatewayImpl-FindById")
	defer span.End()

	userDocument, err := this.userRepository.FindById(span.GetCtx(), id)
	return userDocument.ToDomain(), err
}

func (this *UserDatabaseGatewayImpl) FindByDocumentNumber(ctx context.Context, documentNumber string) (main_domains.User, error) {
	span := this.spanGateway.Get(ctx, "UserDatabaseGatewayImpl-FindByDocumentNumber")
	defer span.End()

	userDocument, err := this.userRepository.FindByDocumentNumber(span.GetCtx(), documentNumber)
	return userDocument.ToDomain(), err
}

func (this *UserDatabaseGatewayImpl) FindByFilter(ctx context.Context, filter main_domains.FindUserFilter, pageable main_domains.Pageable) (main_domains.Page, error) {
	span := this.spanGateway.Get(ctx, "UserDatabaseGatewayImpl-FindByFilter")
	defer span.End()

	page, err := this.userRepository.FindByFilter(span.GetCtx(), filter, pageable)
	if err != nil {
		return main_domains.Page{}, err
	}
	contentDoc := page.GetContent()
	var content []main_domains.Content
	for _, value := range contentDoc {
		userDoc := value.GetObj().(main_gateways_mongodb_documents.UserDocument)
		content = append(content, *main_domains.NewContent(userDoc.ToDomain()))
	}
	return *main_domains.NewPageFromContentAndPage(content, *page), err
}
