package main_usecases_interfaces

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	"context"
)

type CreateEmailBody interface {
	Execute(ctx context.Context, emailParams main_domains.EmailParams) ([]byte, main_domains_exceptions.ApplicationException)
}
