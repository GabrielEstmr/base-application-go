package main_usecases_interfaces

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	"context"
)

type CreateExternalAuthProviderUserSession interface {
	Execute(
		ctx context.Context,
		args main_domains.ExternalProviderSessionArgs,
		tokenClaims main_domains.TokenClaims,
		databaseOptions main_domains.DatabaseOptions,
	) main_domains_exceptions.ApplicationException
}
