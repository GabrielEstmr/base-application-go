package main_gateways_email

import (
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_gateways "baseapplicationgo/main/gateways"
	main_gateways_email_integration "baseapplicationgo/main/gateways/email/integration"
	main_gateways_logs "baseapplicationgo/main/gateways/logs"
	main_gateways_spans "baseapplicationgo/main/gateways/spans"
	"context"
	"fmt"
)

type GmailEmailGatewayImpl struct {
	gmailIntegration      main_gateways_email_integration.GmailIntegration
	logsMonitoringGateway main_gateways.LogsMonitoringGateway
	spanGateway           main_gateways.SpanGateway
}

func NewGmailEmailGatewayImpl() *GmailEmailGatewayImpl {
	return &GmailEmailGatewayImpl{
		gmailIntegration:      *main_gateways_email_integration.NewGmailIntegration(),
		spanGateway:           main_gateways_spans.NewSpanGatewayImpl(),
		logsMonitoringGateway: main_gateways_logs.NewLogsMonitoringGatewayImpl(),
	}
}

func (this *GmailEmailGatewayImpl) SendMail(ctx context.Context, to []string, body []byte) main_domains_exceptions.ApplicationException {
	span := this.spanGateway.Get(ctx, "GmailEmailGatewayImpl-SendMail")
	defer span.End()

	this.logsMonitoringGateway.DEBUG(span,
		fmt.Sprintf("Creating new email. To: %s", to))

	err := this.gmailIntegration.SendMail(span.GetCtx(), to, body)
	if err != nil {
		this.logsMonitoringGateway.ERROR(span, err.Error())
		return main_domains_exceptions.
			NewInternalServerErrorExceptionSglMsg(err.Error())
	}
	return nil
}
