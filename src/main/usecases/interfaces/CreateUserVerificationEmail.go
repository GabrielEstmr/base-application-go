package main_usecases_interfaces

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_enums "baseapplicationgo/main/domains/enums"
	"context"
)

type CreateUserVerificationEmail interface {
	Execute(
		ctx context.Context,
		user main_domains.User,
		scope main_domains_enums.UserEmailVerificationScope,
	) main_domains.UserEmailVerification
}
