package main_gateways

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	"context"
)

type UserDatabaseCacheGateway interface {
	Save(ctx context.Context, user main_domains.User) (main_domains.User, main_domains_exceptions.ApplicationException)
	Update(ctx context.Context, user main_domains.User) (main_domains.User, main_domains_exceptions.ApplicationException)
	FindById(ctx context.Context, id string) (main_domains.User, main_domains_exceptions.ApplicationException)
	FindByDocumentId(ctx context.Context, documentId string) (main_domains.User, main_domains_exceptions.ApplicationException)
	FindByUserName(ctx context.Context, userName string) (main_domains.User, main_domains_exceptions.ApplicationException)
	FindByEmail(ctx context.Context, email string) (main_domains.User, main_domains_exceptions.ApplicationException)
}
