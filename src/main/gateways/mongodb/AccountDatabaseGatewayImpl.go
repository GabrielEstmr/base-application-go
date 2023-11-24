package main_gateways_mongodb

import (
	mainDomains "baseapplicationgo/main/domains"
	gatewaysMongodbDocuments "baseapplicationgo/main/gateways/mongodb/documents"
	gatewaysMongodbRepositories "baseapplicationgo/main/gateways/mongodb/repositories"
)

type AccountDatabaseGatewayImpl struct {
	accountRepository *gatewaysMongodbRepositories.AccountRepository
}

func NewAccountDatabaseGatewayImpl(indicatorRepository *gatewaysMongodbRepositories.AccountRepository) *AccountDatabaseGatewayImpl {
	return &AccountDatabaseGatewayImpl{indicatorRepository}
}

func (thisGateway *AccountDatabaseGatewayImpl) Save(indicator mainDomains.Account) (string, error) {
	indicatorDocument := gatewaysMongodbDocuments.NewAccountDocument(indicator)
	id, err := thisGateway.accountRepository.Save(indicatorDocument)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (thisGateway *AccountDatabaseGatewayImpl) FindById(id string) (mainDomains.Account, error) {
	accountDocument, err := thisGateway.accountRepository.FindById(id)
	return accountDocument.ToDomain(), err
}
