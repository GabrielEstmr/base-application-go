package main_usecases_beans_factories

import (
	main_gateways "baseapplicationgo/main/gateways"
	main_gateways_features "baseapplicationgo/main/gateways/features"
	main_gateways_mongodb "baseapplicationgo/main/gateways/mongodb"
	main_gateways_mongodb_repositories "baseapplicationgo/main/gateways/mongodb/repositories"
	main_gateways_redis "baseapplicationgo/main/gateways/redis"
	main_gateways_redis_repositories "baseapplicationgo/main/gateways/redis/repositories"
	main_usecases "baseapplicationgo/main/usecases"
)

type FindUserByIdBean struct {
}

func NewFindUserByIdBean() *FindUserByIdBean {
	return &FindUserByIdBean{}
}

func (this *FindUserByIdBean) Get() *main_usecases.FindUserById {

	userRepository := main_gateways_mongodb_repositories.NewUserRepository()
	var userDatabaseGateway main_gateways.UserDatabaseGateway = main_gateways_mongodb.NewUserDatabaseGatewayImpl(*userRepository)

	userRedisRepository := main_gateways_redis_repositories.NewUserRedisRepository()
	var userDatabaseCacheGateway main_gateways.UserDatabaseCacheGateway = main_gateways_redis.NewUserDatabaseCacheGatewayImpl(*userRedisRepository)

	var cachedUserDatabaseGateway main_gateways.UserDatabaseGateway = main_gateways_mongodb.NewCachedUserDatabaseGatewayImpl(userDatabaseGateway, userDatabaseCacheGateway)

	var featuresGateway main_gateways.FeaturesGateway = main_gateways_features.NewFeaturesGatewayImpl()

	return main_usecases.NewFindUserById(cachedUserDatabaseGateway, featuresGateway)
}
