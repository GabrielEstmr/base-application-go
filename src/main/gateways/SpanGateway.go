package main_gateways

import (
	main_domains "baseapplicationgo/main/domains"
	"context"
)

type SpanGateway interface {
	Get(ctx context.Context, scope string) main_domains.SpanLogInfo
}
