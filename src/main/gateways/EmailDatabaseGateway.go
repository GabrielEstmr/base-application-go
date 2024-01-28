package main_gateways

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	"context"
)

type EmailDatabaseGateway interface {
	Save(ctx context.Context, email main_domains.Email) (main_domains.Email, main_domains_exceptions.ApplicationException)
	FindById(ctx context.Context, id string) (main_domains.Email, main_domains_exceptions.ApplicationException)
	FindByEventId(ctx context.Context, eventId string) (main_domains.Email, main_domains_exceptions.ApplicationException)
	Update(ctx context.Context, email main_domains.Email) (main_domains.Email, main_domains_exceptions.ApplicationException)
	FindByFilter(ctx context.Context, filter main_domains.FindEmailFilter, pageable main_domains.Pageable) (
		main_domains.Page, main_domains_exceptions.ApplicationException)
}
