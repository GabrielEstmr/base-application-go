package main_usecases_factories

import (
	main_domains_enums "baseapplicationgo/main/domains/enums"
	main_gateways "baseapplicationgo/main/gateways"
	"context"
	"fmt"
)

type SendEmailGatewayFactory struct {
	gmailEmailGatewayImpl main_gateways.EmailGateway
	logsMonitoringGateway main_gateways.LogsMonitoringGateway
	spanGateway           main_gateways.SpanGateway
}

func NewSendEmailGatewayFactoryAllArgs(
	gmailEmailGatewayImpl main_gateways.EmailGateway,
	logsMonitoringGateway main_gateways.LogsMonitoringGateway,
	spanGateway main_gateways.SpanGateway,
) *SendEmailGatewayFactory {
	return &SendEmailGatewayFactory{
		gmailEmailGatewayImpl: gmailEmailGatewayImpl,
		logsMonitoringGateway: logsMonitoringGateway,
		spanGateway:           spanGateway,
	}
}

func (this *SendEmailGatewayFactory) Get(ctx context.Context,
	emailType main_domains_enums.EmailTemplateType) main_gateways.EmailGateway {

	span := this.spanGateway.Get(ctx, "CreateEmail-Execute")
	defer span.End()
	this.logsMonitoringGateway.DEBUG(span, fmt.Sprintf("EmailGatewayFactory. emailType: %s", emailType.Name()))

	if emailType == main_domains_enums.EMAIL_TYPE_WELCOME_EMAIL {
		return this.gmailEmailGatewayImpl
	}
	return this.gmailEmailGatewayImpl
}
