package main_usecases

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_gateways "baseapplicationgo/main/gateways"
	main_usecases_interfaces "baseapplicationgo/main/usecases/interfaces"
	main_utils_messages "baseapplicationgo/main/utils/messages"
	"context"
)

type ValidateHasUserByProperties struct {
	validateUserByPropertyImpls []main_usecases_interfaces.ValidateUserByProperty
	logsMonitoringGateway       main_gateways.LogsMonitoringGateway
	spanGateway                 main_gateways.SpanGateway
	messageUtils                main_utils_messages.ApplicationMessages
}

func NewValidateHasUserByProperties(
	validateUserByPropertyImpls []main_usecases_interfaces.ValidateUserByProperty,
	logsMonitoringGateway main_gateways.LogsMonitoringGateway,
	spanGateway main_gateways.SpanGateway,
	messageUtils main_utils_messages.ApplicationMessages) *ValidateHasUserByProperties {
	return &ValidateHasUserByProperties{
		validateUserByPropertyImpls: validateUserByPropertyImpls,
		logsMonitoringGateway:       logsMonitoringGateway,
		spanGateway:                 spanGateway,
		messageUtils:                messageUtils}
}

func (this *ValidateHasUserByProperties) Execute(
	ctx context.Context, user main_domains.User, databaseOptions main_domains.DatabaseOptions) main_domains_exceptions.ApplicationException {

	span := this.spanGateway.Get(ctx, "ValidateHasUserByProperties-Execute")
	defer span.End()
	this.logsMonitoringGateway.DEBUG(span, "Validating user properties")

	for _, v := range this.validateUserByPropertyImpls {
		err := v.Execute(span.GetCtx(), user, databaseOptions)
		if err != nil {
			return err
		}
	}

	return nil
}
