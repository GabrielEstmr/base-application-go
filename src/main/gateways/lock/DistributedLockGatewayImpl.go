package main_gateways_lock

import (
	main_configs_distributedlock "baseapplicationgo/main/configs/distributedlock"
	main_configs_yml "baseapplicationgo/main/configs/yml"
	main_domains_enums "baseapplicationgo/main/domains/enums"
	"baseapplicationgo/main/domains/lock"
	main_gateways "baseapplicationgo/main/gateways"
	main_gateways_lock_resources "baseapplicationgo/main/gateways/lock/resources"
	main_gateways_logs "baseapplicationgo/main/gateways/logs"
	main_gateways_spans "baseapplicationgo/main/gateways/spans"
	"context"
	"github.com/go-redsync/redsync/v4"
	"time"
)

const _SCOPE_SEPARATOR = " scope: "
const _KEY_SEPARATOR = " key: "

type DistributedLockGatewayImpl struct {
	appName               string
	lockConfig            *redsync.Redsync
	spanGateway           main_gateways.SpanGateway
	logsMonitoringGateway main_gateways.LogsMonitoringGateway
}

func NewDistributedLockGatewayImpl() *DistributedLockGatewayImpl {
	return &DistributedLockGatewayImpl{
		appName:               main_configs_yml.GetYmlValueByName(main_configs_yml.ApmServerName),
		lockConfig:            main_configs_distributedlock.GetLockClientBean(),
		spanGateway:           main_gateways_spans.NewSpanGatewayImpl(),
		logsMonitoringGateway: main_gateways_logs.NewLogsMonitoringGatewayImpl(),
	}
}

func (this *DistributedLockGatewayImpl) Get(ctx context.Context, key string, ttl time.Duration) *lock.SingleLock {
	span := this.spanGateway.Get(ctx, "DistributedLockGatewayImpl-Get")
	defer span.End()
	this.logsMonitoringGateway.DEBUG(span, "Finding User by id")

	redisLock := this.lockConfig.NewMutex(this.appName+key, redsync.WithExpiry(ttl))
	return main_gateways_lock_resources.NewSingleLockResource(redisLock).ToDomain()
}

func (this *DistributedLockGatewayImpl) GetWithScope(
	ctx context.Context,
	scope main_domains_enums.LockScope,
	key string,
	ttl time.Duration,
) *lock.SingleLock {
	span := this.spanGateway.Get(ctx, "DistributedLockGatewayImpl-Get")
	defer span.End()
	this.logsMonitoringGateway.DEBUG(span, "Finding User by id")

	redisLock := this.lockConfig.NewMutex(this.appName+_SCOPE_SEPARATOR+scope.Name()+_KEY_SEPARATOR+key, redsync.WithExpiry(ttl))
	return main_gateways_lock_resources.NewSingleLockResource(redisLock).ToDomain()
}

func (this *DistributedLockGatewayImpl) UnlockAndLogIfError(ctx context.Context, lock lock.SingleLock) {
	span := this.spanGateway.Get(ctx, "DistributedLockGatewayImpl-UnlockAndLogIfError")
	defer span.End()
	this.logsMonitoringGateway.DEBUG(span, "UnlockAndLogIfError")

	_, errUnlock := lock.Unlock()
	if errUnlock != nil {
		this.logsMonitoringGateway.ERROR(span, "Error during unlock")
	}
}
