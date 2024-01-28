package main_usecases_factories

import (
	main_domains_enums "baseapplicationgo/main/domains/enums"
	main_gateways "baseapplicationgo/main/gateways"
	main_usecases_interfaces "baseapplicationgo/main/usecases/interfaces"
	"context"
	"fmt"
)

type CreateEmailBodyFactory struct {
	createWelcomeEmailTemplateBody main_usecases_interfaces.CreateEmailBody
	logsMonitoringGateway          main_gateways.LogsMonitoringGateway
	spanGateway                    main_gateways.SpanGateway
}

func NewCreateEmailBodyFactoryAllArgs(
	createWelcomeEmailTemplateBody main_usecases_interfaces.CreateEmailBody,
	logsMonitoringGateway main_gateways.LogsMonitoringGateway,
	spanGateway main_gateways.SpanGateway,
) *CreateEmailBodyFactory {
	return &CreateEmailBodyFactory{
		createWelcomeEmailTemplateBody: createWelcomeEmailTemplateBody,
		logsMonitoringGateway:          logsMonitoringGateway,
		spanGateway:                    spanGateway,
	}
}

func (this *CreateEmailBodyFactory) Get(ctx context.Context,
	emailType main_domains_enums.EmailTemplateType) main_usecases_interfaces.CreateEmailBody {

	span := this.spanGateway.Get(ctx, "CreateEmailBodyFactory-Get")
	defer span.End()
	this.logsMonitoringGateway.DEBUG(span,
		fmt.Sprintf("CreateEmailBodyFactory. emailType: %s",
			main_domains_enums.GetEmailTemplateTypeDescription(emailType)))

	if emailType == main_domains_enums.EMAIL_TYPE_WELCOME_EMAIL {
		return this.createWelcomeEmailTemplateBody
	}
	return this.createWelcomeEmailTemplateBody
}
