package main_gateways_redis

import (
	main_domains "baseapplicationgo/main/domains"
	main_gateways_redis_documents "baseapplicationgo/main/gateways/redis/documents"
	main_gateways_redis_repositories "baseapplicationgo/main/gateways/redis/repositories"
)

type UserDatabaseCacheGatewayImpl struct {
	userRedisRepository main_gateways_redis_repositories.UserRedisRepository
}

func NewUserDatabaseCacheGatewayImpl(
	redisUserRepo main_gateways_redis_repositories.UserRedisRepository) *UserDatabaseCacheGatewayImpl {
	return &UserDatabaseCacheGatewayImpl{redisUserRepo}
}

//Save(user main_domains.User) (string, error)
//FindById(id string) (main_domains.User, error)

func (this *UserDatabaseCacheGatewayImpl) Save(user main_domains.User) (main_domains.User, error) {
	userDocument := main_gateways_redis_documents.NewUserRedisDocument(user)
	userDocument, err := this.userRedisRepository.Save(userDocument)
	if err != nil {
		return userDocument.ToDomain(), err
	}
	return userDocument.ToDomain(), nil
}

func (this *UserDatabaseCacheGatewayImpl) FindById(id string) (main_domains.User, error) {
	userDocument, err := this.userRedisRepository.FindById(id)
	return userDocument.ToDomain(), err
}

func (this *UserDatabaseCacheGatewayImpl) FindByDocumentNumber(documentNumber string) (main_domains.User, error) {
	userDocument, err := this.userRedisRepository.FindByDocumentNumber(documentNumber)
	return userDocument.ToDomain(), err
}
