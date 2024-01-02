package main_gateways

import (
	main_domains "baseapplicationgo/main/domains"
	"context"
	"time"
)

type DistributedLockGateway interface {
	Get(ctx context.Context, key string, ttl time.Duration) main_domains.SingleLock
}
