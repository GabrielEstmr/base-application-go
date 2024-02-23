package main_usecases_interfaces

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_enums "baseapplicationgo/main/domains/enums"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	"context"
)

type CreateAndSendUserVerificationEmail interface {
	Execute(ctx context.Context, user main_domains.User,
		scope main_domains_enums.UserEmailVerificationScope,
		databaseOptions main_domains.DatabaseOptions,
	) (main_domains.UserEmailVerification, main_domains_exceptions.ApplicationException)
}
