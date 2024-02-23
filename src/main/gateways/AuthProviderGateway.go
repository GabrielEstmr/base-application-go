package main_gateways

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	"context"
)

type AuthProviderGateway interface {
	CreateUser(
		ctx context.Context,
		user main_domains.User) (main_domains.User, main_domains_exceptions.ApplicationException)

	CreateOauthExchangeSession(
		ctx context.Context, args main_domains.ExternalProviderSessionArgs,
	) (main_domains.SessionCredentials, main_domains_exceptions.ApplicationException)

	CreateSession(
		ctx context.Context, username string, password string,
	) (main_domains.SessionCredentials, main_domains_exceptions.ApplicationException)

	RefreshSession(
		ctx context.Context, refreshToken string,
	) (main_domains.SessionCredentials, main_domains_exceptions.ApplicationException)

	EndSession(
		ctx context.Context, refreshToken string,
	) main_domains_exceptions.ApplicationException

	ChangeUserStatusAndEmailVerification(
		ctx context.Context,
		user main_domains.User,
		enabled bool,
	) (main_domains.User, main_domains_exceptions.ApplicationException)

	ChangeUsersPassword(
		ctx context.Context,
		user main_domains.User,
	) (main_domains.User, main_domains_exceptions.ApplicationException)

	GetUsers(
		ctx context.Context,
		email string,
	) ([]main_domains.AuthProviderUser, main_domains_exceptions.ApplicationException)
}
