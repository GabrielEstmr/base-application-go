package main_gateways

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	"context"
)

type TransactionEventProducerGateway interface {
	Send(ctx context.Context, transaction main_domains.Transaction) main_domains_exceptions.ApplicationException
}
