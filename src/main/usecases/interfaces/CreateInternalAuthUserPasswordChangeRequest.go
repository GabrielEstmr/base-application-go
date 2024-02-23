package main_usecases_interfaces

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	"context"
)

type CreateInternalAuthUserPasswordChangeRequest interface {
	Execute(
		ctx context.Context,
		userId string,
		databaseOptions main_domains.DatabaseOptions,
	) main_domains_exceptions.ApplicationException
}
