package main_gateways

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	"context"
)

type UserDatabaseGateway interface {
	Save(ctx context.Context, user main_domains.User, options main_domains.DatabaseOptions) (main_domains.User, main_domains_exceptions.ApplicationException)
	Update(ctx context.Context, user main_domains.User, options main_domains.DatabaseOptions) (main_domains.User, main_domains_exceptions.ApplicationException)
	FindById(ctx context.Context, id string, options main_domains.DatabaseOptions) (main_domains.User, main_domains_exceptions.ApplicationException)
	FindByDocumentId(ctx context.Context, documentId string, options main_domains.DatabaseOptions) (main_domains.User, main_domains_exceptions.ApplicationException)
	FindByUserName(ctx context.Context, userName string, options main_domains.DatabaseOptions) (main_domains.User, main_domains_exceptions.ApplicationException)
	FindByEmail(ctx context.Context, email string, options main_domains.DatabaseOptions) (main_domains.User, main_domains_exceptions.ApplicationException)
	FindByFilter(ctx context.Context, filter main_domains.FindUserFilter, pageable main_domains.Pageable, options main_domains.DatabaseOptions) (main_domains.Page, main_domains_exceptions.ApplicationException)
}
