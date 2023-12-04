package main_gateways_mongodb

import (
	main_domains "baseapplicationgo/main/domains"
	main_gateways_mongodb_documents "baseapplicationgo/main/gateways/mongodb/documents"
	main_gateways_mongodb_repositories "baseapplicationgo/main/gateways/mongodb/repositories"
	"time"
)

type UserDatabaseGatewayImpl struct {
	userRepository main_gateways_mongodb_repositories.UserRepository
}

func NewUserDatabaseGatewayImpl(userRepository main_gateways_mongodb_repositories.UserRepository) *UserDatabaseGatewayImpl {
	return &UserDatabaseGatewayImpl{userRepository}
}

func (this *UserDatabaseGatewayImpl) Save(user main_domains.User) (main_domains.User, error) {
	now := time.Now()
	user.CreatedDate = now
	user.LastModifiedDate = now
	userDocument := main_gateways_mongodb_documents.NewUserDocument(user)
	persistedUserDocument, err := this.userRepository.Save(userDocument)
	if err != nil {
		return main_domains.User{}, err
	}
	return persistedUserDocument.ToDomain(), nil
}

func (this *UserDatabaseGatewayImpl) FindById(id string) (main_domains.User, error) {
	userDocument, err := this.userRepository.FindById(id)
	return userDocument.ToDomain(), err
}

func (this *UserDatabaseGatewayImpl) FindByDocumentNumber(documentNumber string) (main_domains.User, error) {
	userDocument, err := this.userRepository.FindByDocumentNumber(documentNumber)
	return userDocument.ToDomain(), err
}

func (this *UserDatabaseGatewayImpl) FindByFilter(filter main_domains.FindUserFilter, pageable main_domains.Pageable) (main_domains.Page, error) {
	page, err := this.userRepository.FindByFilter(filter, pageable)
	return *page, err
}
