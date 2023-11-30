package main_gateways_mongodb

import (
	main_domains "baseapplicationgo/main/domains"
	main_gateways_mongodb_documents "baseapplicationgo/main/gateways/mongodb/documents"
	main_gateways_mongodb_repositories "baseapplicationgo/main/gateways/mongodb/repositories"
)

type UserDatabaseGatewayImpl struct {
	userRepository main_gateways_mongodb_repositories.UserRepository
}

func NewUserDatabaseGatewayImpl(userRepository main_gateways_mongodb_repositories.UserRepository) *UserDatabaseGatewayImpl {
	return &UserDatabaseGatewayImpl{userRepository}
}

func (this *UserDatabaseGatewayImpl) Save(user main_domains.User) (string, error) {
	userDocument := main_gateways_mongodb_documents.NewUserDocument(user)
	id, err := this.userRepository.Save(userDocument)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (this *UserDatabaseGatewayImpl) FindById(id string) (main_domains.User, error) {
	userDocument, err := this.userRepository.FindById(id)
	return userDocument.ToDomain(), err
}

func (this *UserDatabaseGatewayImpl) FindByDocumentNumber(documentNumber string) (main_domains.User, error) {
	userDocument, err := this.userRepository.FindByDocumentNumber(documentNumber)
	return userDocument.ToDomain(), err
}
