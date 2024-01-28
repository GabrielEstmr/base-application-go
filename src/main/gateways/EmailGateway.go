package main_gateways

import (
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	"context"
)

type EmailGateway interface {
	SendMail(ctx context.Context, to []string, body []byte) main_domains_exceptions.ApplicationException
}
