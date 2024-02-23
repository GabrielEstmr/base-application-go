package main_usecases

import (
	main_gateways "baseapplicationgo/main/gateways"
	main_utils_messages "baseapplicationgo/main/utils/messages"
	"context"
)

type SendEmailEventsToReprocess struct {
	reprocessEmailEventProducerGateway main_gateways.ReprocessEmailEventProducerGateway
	logsMonitoringGateway              main_gateways.LogsMonitoringGateway
	spanGateway                        main_gateways.SpanGateway
	messageUtils                       main_utils_messages.ApplicationMessages
}

func NewSendEmailEventsToReprocessAllArgs(
	reprocessEmailEventProducerGateway main_gateways.ReprocessEmailEventProducerGateway,
	logsMonitoringGateway main_gateways.LogsMonitoringGateway,
	spanGateway main_gateways.SpanGateway,
	messageUtils main_utils_messages.ApplicationMessages,
) *SendEmailEventsToReprocess {
	return &SendEmailEventsToReprocess{
		reprocessEmailEventProducerGateway: reprocessEmailEventProducerGateway,
		logsMonitoringGateway:              logsMonitoringGateway,
		spanGateway:                        spanGateway,
		messageUtils:                       messageUtils,
	}
}

func (this *SendEmailEventsToReprocess) Execute(ctx context.Context, ids []string,
) {
	span := this.spanGateway.Get(ctx, "ReprocessEmails-Execute")
	defer span.End()
	this.logsMonitoringGateway.DEBUG(span, "Reprocessing emails")

	for _, v := range ids {
		go this.reprocessEmailEventProducerGateway.Send(span.GetCtx(), v)
	}
}
