package main_gateways

import (
	main_domains_enums "baseapplicationgo/main/domains/enums"
	"baseapplicationgo/main/domains/lock"
	"context"
	"time"
)

type DistributedLockGateway interface {
	Get(ctx context.Context, key string, ttl time.Duration) *lock.SingleLock
	GetWithScope(ctx context.Context, scope main_domains_enums.LockScope, key string, ttl time.Duration) *lock.SingleLock
	UnlockAndLogIfError(ctx context.Context, lock lock.SingleLock)
}
