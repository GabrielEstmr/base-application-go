package main_gateways

import (
	main_domains "baseapplicationgo/main/domains"
	"context"
)

type TransactionDatabaseGateway interface {
	Save(ctx context.Context, transaction main_domains.Transaction) (main_domains.Transaction, error)
}
