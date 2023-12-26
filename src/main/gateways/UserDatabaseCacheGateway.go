package main_gateways

import (
	main_domains "baseapplicationgo/main/domains"
	"context"
)

type UserDatabaseCacheGateway interface {
	Save(ctx context.Context, user main_domains.User) (main_domains.User, error)
	FindById(ctx context.Context, id string) (main_domains.User, error)
	FindByDocumentNumber(ctx context.Context, documentNumber string) (main_domains.User, error)
}
