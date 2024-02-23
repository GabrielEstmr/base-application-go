package main_usecases_beans_factories

import (
	main_gateways "baseapplicationgo/main/gateways"
	main_gateways_lock "baseapplicationgo/main/gateways/lock"
	main_gateways_logs "baseapplicationgo/main/gateways/logs"
	main_gateways_mongodb "baseapplicationgo/main/gateways/mongodb"
	main_gateways_mongodb_repositories "baseapplicationgo/main/gateways/mongodb/repositories"
	main_gateways_redis "baseapplicationgo/main/gateways/redis"
	main_gateways_redis_repositories "baseapplicationgo/main/gateways/redis/repositories"
	main_gateways_spans "baseapplicationgo/main/gateways/spans"
	main_usecases_lockers "baseapplicationgo/main/usecases/lockers"
	main_utils_messages "baseapplicationgo/main/utils/messages"
)

type AtomicLockedCreateInternalAuthUserPasswordChangeRequestBean struct {
}

func NewAtomicLockedCreateInternalAuthUserPasswordChangeRequestBean() *AtomicLockedCreateInternalAuthUserPasswordChangeRequestBean {
	return &AtomicLockedCreateInternalAuthUserPasswordChangeRequestBean{}
}

func (this *AtomicLockedCreateInternalAuthUserPasswordChangeRequestBean) Get() *main_usecases_lockers.AtomicLockedCreateInternalAuthUserPasswordChangeRequest {

	var distributedLockGateway main_gateways.DistributedLockGateway = main_gateways_lock.NewDistributedLockGatewayImpl()
	var logsMonitoringGateway main_gateways.LogsMonitoringGateway = main_gateways_logs.NewLogsMonitoringGatewayImpl()
	var spanGatewayImpl main_gateways.SpanGateway = main_gateways_spans.NewSpanGatewayImpl()

	userRepository := main_gateways_mongodb_repositories.NewUserRepository()
	var userDatabaseGateway main_gateways.UserDatabaseGateway = main_gateways_mongodb.NewUserDatabaseGatewayImpl(*userRepository)
	userRedisRepository := main_gateways_redis_repositories.NewUserRedisRepository()
	var userDatabaseCacheGateway main_gateways.UserDatabaseCacheGateway = main_gateways_redis.NewUserDatabaseCacheGatewayImpl(*userRedisRepository)
	var cachedUserDatabaseGateway main_gateways.UserDatabaseGateway = main_gateways_mongodb.NewCachedUserDatabaseGatewayImpl(userDatabaseGateway, userDatabaseCacheGateway)

	return main_usecases_lockers.NewAtomicLockedCreateInternalAuthUserPasswordChangeRequest(
		NewCreateInternalAuthUserPasswordChangeRequestBean().Get(),
		distributedLockGateway,
		cachedUserDatabaseGateway,
		logsMonitoringGateway,
		spanGatewayImpl,
		*main_utils_messages.NewApplicationMessages())
}
