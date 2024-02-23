package main_usecases_interfaces

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	"context"
)

type BuildTokenClaim interface {
	Execute(
		ctx context.Context,
		token string) (
		main_domains.TokenClaims,
		main_domains_exceptions.ApplicationException)
}
