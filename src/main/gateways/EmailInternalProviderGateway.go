package main_gateways

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	"golang.org/x/net/context"
)

type EmailInternalProviderGateway interface {
	SendMail(ctx context.Context, emailParams main_domains.EmailParams) main_domains_exceptions.ApplicationException
}
