package main_gateways_mongodb

import (
	mainDomains "baseapplicationgo/main/domains"
	gatewaysMongodbDocuments "baseapplicationgo/main/gateways/mongodb/documents"
	gatewaysMongodbRepositories "baseapplicationgo/main/gateways/mongodb/repositories"
)

type UserDatabaseGatewayImpl struct {
	userRepository *gatewaysMongodbRepositories.UserRepository
}

func NewUserDatabaseGatewayImpl(userRepository *gatewaysMongodbRepositories.UserRepository) *UserDatabaseGatewayImpl {
	return &UserDatabaseGatewayImpl{userRepository}
}

func (this *UserDatabaseGatewayImpl) Save(user mainDomains.User) (string, error) {
	userDocument := gatewaysMongodbDocuments.NewUserDocument(user)
	id, err := this.userRepository.Save(userDocument)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (thisGateway *UserDatabaseGatewayImpl) FindById(id string) (mainDomains.User, error) {
	userDocument, err := thisGateway.userRepository.FindById(id)
	return userDocument.ToDomain(), err
}
