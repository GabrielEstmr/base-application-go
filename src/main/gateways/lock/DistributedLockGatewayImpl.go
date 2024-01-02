package main_gateways_lock

import (
	main_configs_distributedlock "baseapplicationgo/main/configs/distributedlock"
	main_configs_logs "baseapplicationgo/main/configs/log"
	main_domains "baseapplicationgo/main/domains"
	main_gateways "baseapplicationgo/main/gateways"
	main_gateways_lock_resources "baseapplicationgo/main/gateways/lock/resources"
	main_gateways_spans "baseapplicationgo/main/gateways/spans"
	"context"
	"fmt"
	"github.com/go-redsync/redsync/v4"
	"log/slog"
	"time"
)

const _MSG_DISTRIBUTED_LOCK_GATEWAY = "Initiating lock. Key: %s, Ttl: %s"

type DistributedLockGatewayImpl struct {
	lockConfig  *redsync.Redsync
	spanGateway main_gateways.SpanGateway
	apLog       *slog.Logger
}

func NewDistributedLockGatewayImpl() *DistributedLockGatewayImpl {
	return &DistributedLockGatewayImpl{
		lockConfig:  main_configs_distributedlock.GetLockClientBean(),
		spanGateway: main_gateways_spans.NewSpanGatewayImpl(),
		apLog:       main_configs_logs.GetLogConfigBean(),
	}
}

func (this *DistributedLockGatewayImpl) Get(ctx context.Context, key string, ttl time.Duration) main_domains.SingleLock {
	span := this.spanGateway.Get(ctx, "DistributedLockGatewayImpl-Get")
	defer span.End()
	this.apLog.Debug(fmt.Sprintf(_MSG_DISTRIBUTED_LOCK_GATEWAY, key, ttl))

	lock := this.lockConfig.NewMutex(key, redsync.WithExpiry(ttl))
	return main_gateways_lock_resources.NewSingleLockResource(lock).ToDomain()
}
