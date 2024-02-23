package main_usecases_interfaces

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	"context"
)

type ChangeInternalProviderUserPassword interface {
	Execute(
		ctx context.Context,
		userId string,
		password string,
		verificationCode string,
		dbOpts main_domains.DatabaseOptions,
	) (main_domains.User, main_domains_exceptions.ApplicationException)
}
