package main_gateways_rabbitmq

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_gateways "baseapplicationgo/main/gateways"
	main_gateways_logs "baseapplicationgo/main/gateways/logs"
	main_gateways_rabbitmq_producers "baseapplicationgo/main/gateways/rabbitmq/producers"
	main_gateways_rabbitmq_resources "baseapplicationgo/main/gateways/rabbitmq/resources"
	main_gateways_spans "baseapplicationgo/main/gateways/spans"
	main_utils_messages "baseapplicationgo/main/utils/messages"
	"context"
	"fmt"
)

const _MSG_CREATE_EVENT_EMAIL_DOC_ARCH_ISSUE = "exceptions.architecture.application.issue"

type EmailInternalProviderGateway struct {
	emailEventProducer    main_gateways_rabbitmq_producers.EmailEventProducer
	logsMonitoringGateway main_gateways.LogsMonitoringGateway
	messageUtils          main_utils_messages.ApplicationMessages
	spanGateway           main_gateways.SpanGateway
}

func NewEmailInternalProviderGateway(
	emailEventProducer main_gateways_rabbitmq_producers.EmailEventProducer,
) *EmailInternalProviderGateway {
	return &EmailInternalProviderGateway{
		emailEventProducer:    emailEventProducer,
		logsMonitoringGateway: main_gateways_logs.NewLogsMonitoringGatewayImpl(),
		messageUtils:          *main_utils_messages.NewApplicationMessages(),
		spanGateway:           main_gateways_spans.NewSpanGatewayImpl(),
	}
}

func (this *EmailInternalProviderGateway) SendMail(
	ctx context.Context,
	emailParams main_domains.EmailParams) main_domains_exceptions.ApplicationException {

	span := this.spanGateway.Get(ctx, "EmailInternalProviderGateway-Send")
	defer span.End()

	this.logsMonitoringGateway.DEBUG(span,
		fmt.Sprintf("Sending email to: %s", emailParams.GetTo()))

	err := this.emailEventProducer.Produce(
		span.GetCtx(),
		*main_gateways_rabbitmq_resources.NewEmailParamsResource(emailParams))
	if err != nil {
		this.logsMonitoringGateway.ERROR(span, err.Error())
		return main_domains_exceptions.
			NewInternalServerErrorExceptionSglMsg(this.messageUtils.
				GetDefaultLocale(_MSG_CREATE_EVENT_EMAIL_DOC_ARCH_ISSUE))
	}
	return nil
}
