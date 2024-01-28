package main_gateways

import (
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	"context"
)

type ReprocessEmailEventProducerGateway interface {
	Send(ctx context.Context, id string) main_domains_exceptions.ApplicationException
}
