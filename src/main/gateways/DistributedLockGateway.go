package main_gateways

import (
	"baseapplicationgo/main/domains/lock"
	"context"
	"time"
)

type DistributedLockGateway interface {
	Get(ctx context.Context, key string, ttl time.Duration) lock.SingleLock
}
