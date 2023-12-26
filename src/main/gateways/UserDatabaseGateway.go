package main_gateways

import (
	main_domains "baseapplicationgo/main/domains"
	"context"
)

type UserDatabaseGateway interface {
	Save(ctx context.Context, user main_domains.User) (main_domains.User, error)
	FindById(ctx context.Context, id string) (main_domains.User, error)
	FindByDocumentNumber(ctx context.Context, documentNumber string) (main_domains.User, error)
	FindByFilter(ctx context.Context, filter main_domains.FindUserFilter, pageable main_domains.Pageable) (main_domains.Page, error)
}
