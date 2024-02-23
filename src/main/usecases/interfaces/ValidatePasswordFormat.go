package main_usecases_interfaces

import (
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	"context"
)

type ValidatePasswordFormat interface {
	Execute(ctx context.Context, password string) main_domains_exceptions.ApplicationException
}
