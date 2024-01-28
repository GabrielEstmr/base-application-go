package main_usecases_factories_interfaces

import (
	main_domains_enums "baseapplicationgo/main/domains/enums"
	main_gateways "baseapplicationgo/main/gateways"
	"context"
)

type SendEmailGatewayFactory interface {
	Get(ctx context.Context,
		emailType main_domains_enums.EmailTemplateType) main_gateways.EmailGateway
}
