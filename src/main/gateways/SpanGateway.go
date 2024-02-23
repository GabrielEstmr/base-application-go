package main_gateways

import (
	main_domains_apm "baseapplicationgo/main/domains/apm"
	"context"
)

type SpanGateway interface {
	Get(ctx context.Context, scope string) main_domains_apm.SpanLogInfo
}
