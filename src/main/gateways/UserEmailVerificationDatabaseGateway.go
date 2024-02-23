package main_gateways

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	"context"
)

type UserEmailVerificationDatabaseGateway interface {
	Save(ctx context.Context, user main_domains.UserEmailVerification, databaseOptions main_domains.DatabaseOptions) (main_domains.UserEmailVerification, main_domains_exceptions.ApplicationException)
	Update(ctx context.Context, user main_domains.UserEmailVerification, databaseOptions main_domains.DatabaseOptions) (main_domains.UserEmailVerification, main_domains_exceptions.ApplicationException)
	FindById(ctx context.Context, id string, databaseOptions main_domains.DatabaseOptions) (main_domains.UserEmailVerification, main_domains_exceptions.ApplicationException)
	FindByFilter(ctx context.Context, filter main_domains.FindUserEmailVerificationFilter, pageable main_domains.Pageable, databaseOptions main_domains.DatabaseOptions) (main_domains.Page, main_domains_exceptions.ApplicationException)
}
