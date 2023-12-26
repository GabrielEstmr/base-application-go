package main_gateways_mongodb

import (
	main_domains "baseapplicationgo/main/domains"
	main_gateways_mongodb_documents "baseapplicationgo/main/gateways/mongodb/documents"
	main_gateways_mongodb_repositories "baseapplicationgo/main/gateways/mongodb/repositories"
	"context"
	"time"
)

type UserDatabaseGatewayImpl struct {
	userRepository main_gateways_mongodb_repositories.UserRepository
}

func NewUserDatabaseGatewayImpl(userRepository main_gateways_mongodb_repositories.UserRepository) *UserDatabaseGatewayImpl {
	return &UserDatabaseGatewayImpl{userRepository}
}

func (this *UserDatabaseGatewayImpl) Save(ctx context.Context, user main_domains.User) (main_domains.User, error) {
	now := time.Now()
	user.CreatedDate = now
	user.LastModifiedDate = now
	userDocument := main_gateways_mongodb_documents.NewUserDocument(user)
	persistedUserDocument, err := this.userRepository.Save(ctx, userDocument)
	if err != nil {
		return main_domains.User{}, err
	}
	return persistedUserDocument.ToDomain(), nil
}

func (this *UserDatabaseGatewayImpl) FindById(ctx context.Context, id string) (main_domains.User, error) {
	userDocument, err := this.userRepository.FindById(ctx, id)
	return userDocument.ToDomain(), err
}

func (this *UserDatabaseGatewayImpl) FindByDocumentNumber(ctx context.Context, documentNumber string) (main_domains.User, error) {
	userDocument, err := this.userRepository.FindByDocumentNumber(ctx, documentNumber)
	return userDocument.ToDomain(), err
}

func (this *UserDatabaseGatewayImpl) FindByFilter(ctx context.Context, filter main_domains.FindUserFilter, pageable main_domains.Pageable) (main_domains.Page, error) {
	page, err := this.userRepository.FindByFilter(ctx, filter, pageable)
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
