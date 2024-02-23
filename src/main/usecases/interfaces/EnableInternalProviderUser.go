package main_usecases_interfaces

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	"context"
)

type EnableInternalProviderUser interface {
	Execute(
		ctx context.Context,
		userId string,
		verificationCode string,
		databaseOptions main_domains.DatabaseOptions) (main_domains.User, main_domains_exceptions.ApplicationException)
}
