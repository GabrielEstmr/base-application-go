package main_usecases_beans_factories

import (
	main_gateways "baseapplicationgo/main/gateways"
	main_gateways_logs "baseapplicationgo/main/gateways/logs"
	main_gateways_mongodb "baseapplicationgo/main/gateways/mongodb"
	main_gateways_mongodb_repositories "baseapplicationgo/main/gateways/mongodb/repositories"
	main_gateways_redis "baseapplicationgo/main/gateways/redis"
	main_gateways_redis_repositories "baseapplicationgo/main/gateways/redis/repositories"
	main_gateways_spans "baseapplicationgo/main/gateways/spans"
	main_usecases "baseapplicationgo/main/usecases"
)

type CreateNewUserBean struct {
}

func NewCreateNewUserBean() *CreateNewUserBean {
	return &CreateNewUserBean{}
}

func (this *CreateNewUserBean) Get() *main_usecases.CreateNewUser {

	userRepository := main_gateways_mongodb_repositories.NewUserRepository()
	var userDatabaseGateway main_gateways.UserDatabaseGateway = main_gateways_mongodb.NewUserDatabaseGatewayImpl(*userRepository)

	userRedisRepository := main_gateways_redis_repositories.NewUserRedisRepository()
	var userDatabaseCacheGateway main_gateways.UserDatabaseCacheGateway = main_gateways_redis.NewUserDatabaseCacheGatewayImpl(*userRedisRepository)

	var cachedUserDatabaseGateway main_gateways.UserDatabaseGateway = main_gateways_mongodb.NewCachedUserDatabaseGatewayImpl(userDatabaseGateway, userDatabaseCacheGateway)

	var logsMonitoringGateway main_gateways.LogsMonitoringGateway = main_gateways_logs.NewLogsMonitoringGatewayImpl()

	var spanGatewayImpl main_gateways.SpanGateway = main_gateways_spans.NewSpanGatewayImpl()

	return main_usecases.NewCreateNewUser(cachedUserDatabaseGateway, logsMonitoringGateway, spanGatewayImpl)
}
