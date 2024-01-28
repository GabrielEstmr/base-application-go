package main_usecases_interfaces

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	"context"
)

type CreateEmailBodySendAndPersistAsSent interface {
	Execute(ctx context.Context, email main_domains.Email) (main_domains.Email, main_domains_exceptions.ApplicationException)
}
