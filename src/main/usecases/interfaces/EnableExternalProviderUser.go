package main_usecases_interfaces

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	"context"
)

type EnableExternalProviderUser interface {
	Execute(
		ctx context.Context,
		currentUser main_domains.User,
		args main_domains.EnableExternalUserArgs,
		dbOpts main_domains.DatabaseOptions,
	) (
		main_domains.User,
		main_domains_exceptions.ApplicationException,
	)
}
