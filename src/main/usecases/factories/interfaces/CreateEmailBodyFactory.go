package main_usecases_factories_interfaces

import (
	main_domains_enums "baseapplicationgo/main/domains/enums"
	main_usecases_interfaces "baseapplicationgo/main/usecases/interfaces"
	"context"
)

type CreateEmailBodyFactory interface {
	Get(ctx context.Context,
		emailType main_domains_enums.EmailTemplateType) main_usecases_interfaces.CreateEmailBody
}
