package main_usecases_interfaces

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_enums "baseapplicationgo/main/domains/enums"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	"context"
)

type FindValidUserEmailVerification interface {
	Execute(
		ctx context.Context,
		userId string,
		scope main_domains_enums.UserEmailVerificationScope,
		verificationCode string,
		dbOpt main_domains.DatabaseOptions,
	) (main_domains.UserEmailVerification, main_domains_exceptions.ApplicationException)
}
