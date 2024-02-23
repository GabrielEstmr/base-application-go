package main_usecases_interfaces

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	"context"
)

type SetUserToEnabledAndEmailVerified interface {
	Execute(ctx context.Context, id string, dbOpt main_domains.DatabaseOptions) (
		main_domains.User, main_domains_exceptions.ApplicationException)
}
