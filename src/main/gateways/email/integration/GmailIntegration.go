package main_gateways_email_integration

import (
	main_configs_email "baseapplicationgo/main/configs/email"
	main_configs_email_resources "baseapplicationgo/main/configs/email/resources"
	main_gateways "baseapplicationgo/main/gateways"
	main_gateways_logs "baseapplicationgo/main/gateways/logs"
	main_gateways_spans "baseapplicationgo/main/gateways/spans"
	"context"
	"fmt"
	"net/smtp"
)

type GmailIntegration struct {
	gmailClientProps      main_configs_email_resources.EmailClientProps
	logsMonitoringGateway main_gateways.LogsMonitoringGateway
	spanGateway           main_gateways.SpanGateway
}

func NewGmailIntegration() *GmailIntegration {
	configBean := *main_configs_email.GetEmailConfigsBean()
	return &GmailIntegration{
		gmailClientProps:      configBean[main_configs_email_resources.GMAIL],
		spanGateway:           main_gateways_spans.NewSpanGatewayImpl(),
		logsMonitoringGateway: main_gateways_logs.NewLogsMonitoringGatewayImpl(),
	}
}

func (this *GmailIntegration) SendMail(ctx context.Context, to []string, body []byte) error {
	span := this.spanGateway.Get(ctx, "GmailIntegration-SendMail")
	defer span.End()

	this.logsMonitoringGateway.DEBUG(span,
		fmt.Sprintf("Creating new email. To: %s", to))

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	auth := smtp.PlainAuth("", this.gmailClientProps.GetEmail(), this.gmailClientProps.GetPassword(), smtpHost)
	return smtp.SendMail(smtpHost+":"+smtpPort, auth, this.gmailClientProps.GetEmail(), to, body)
}
