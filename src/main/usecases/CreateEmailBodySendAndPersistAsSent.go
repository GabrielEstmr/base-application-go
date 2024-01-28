package main_usecases

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_apm "baseapplicationgo/main/domains/apm"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_gateways "baseapplicationgo/main/gateways"
	main_usecases_factories_interfaces "baseapplicationgo/main/usecases/factories/interfaces"
	"context"
	"errors"
	"fmt"
)

type CreateEmailBodySendAndPersistAsSent struct {
	emailDatabaseGateway    main_gateways.EmailDatabaseGateway
	sendEmailGatewayFactory main_usecases_factories_interfaces.SendEmailGatewayFactory
	createEmailBodyFactory  main_usecases_factories_interfaces.CreateEmailBodyFactory
	logsMonitoringGateway   main_gateways.LogsMonitoringGateway
	spanGateway             main_gateways.SpanGateway
}

func NewCreateEmailBodySendAndPersistAsSentAllArgs(
	emailDatabaseGateway main_gateways.EmailDatabaseGateway,
	sendEmailGatewayFactory main_usecases_factories_interfaces.SendEmailGatewayFactory,
	createEmailBodyFactory main_usecases_factories_interfaces.CreateEmailBodyFactory,
	logsMonitoringGateway main_gateways.LogsMonitoringGateway,
	spanGateway main_gateways.SpanGateway,
) *CreateEmailBodySendAndPersistAsSent {
	return &CreateEmailBodySendAndPersistAsSent{
		emailDatabaseGateway:    emailDatabaseGateway,
		sendEmailGatewayFactory: sendEmailGatewayFactory,
		createEmailBodyFactory:  createEmailBodyFactory,
		logsMonitoringGateway:   logsMonitoringGateway,
		spanGateway:             spanGateway,
	}
}

func (this *CreateEmailBodySendAndPersistAsSent) Execute(
	ctx context.Context,
	email main_domains.Email,
) (main_domains.Email, main_domains_exceptions.ApplicationException) {

	span := this.spanGateway.Get(ctx, "CreateEmailBodySendAndPersistAsSent-Execute")
	defer span.End()
	this.logsMonitoringGateway.INFO(span,
		fmt.Sprintf("Creating email body and setting as sent. eventId: %s", email.GetEventId()))

	if email.GetEmailParams().IsEmpty() {
		errEmpty := errors.New("empty email params")
		updatedEmail := this.updateEmailAsErrorAndLogIfError(span, email, errEmpty)
		return updatedEmail, this.logAndReturnArchError(span, errEmpty)
	}

	body, errB := this.createEmailBodyFactory.Get(span.GetCtx(),
		email.GetEmailParams().GetEmailTemplateType()).Execute(
		span.GetCtx(), email.GetEmailParams())
	if errB != nil {
		updatedEmail := this.updateEmailAsErrorAndLogIfError(span, email, errB)
		return updatedEmail, this.logAndReturnArchError(span, errB)
	}

	errSend := this.sendEmailGatewayFactory.Get(span.GetCtx(),
		email.GetEmailParams().GetEmailTemplateType()).SendMail(span.GetCtx(), email.GetEmailParams().GetTo(), body)
	if errSend != nil {
		updatedEmail := this.saveEmailAsIntegrationErrorAndLogIfError(span, email, errSend)
		return updatedEmail, this.logAndReturnArchError(span, errSend)
	}

	// Here: transactional
	persistedSentEmail, errUpdateSent := this.emailDatabaseGateway.Update(span.GetCtx(), email.CloneAsSent())
	if errUpdateSent != nil {
		return email, this.logAndReturnArchError(span, errUpdateSent)
	}
	return persistedSentEmail, nil
}

func (this *CreateEmailBodySendAndPersistAsSent) logAndReturnArchError(
	span main_domains_apm.SpanLogInfo,
	err error,
) main_domains_exceptions.ApplicationException {
	this.logsMonitoringGateway.ERROR(span, err.Error())
	return main_domains_exceptions.
		NewInternalServerErrorExceptionSglMsg(err.Error())
}

func (this *CreateEmailBodySendAndPersistAsSent) updateEmailAsErrorAndLogIfError(
	span main_domains_apm.SpanLogInfo,
	email main_domains.Email, err error) main_domains.Email {
	updatedEmail, errUpdateError := this.emailDatabaseGateway.Update(span.GetCtx(), email.CloneAsError(err.Error()))
	if errUpdateError != nil {
		this.logsMonitoringGateway.ERROR(span, errUpdateError.Error())
		return email
	}
	return updatedEmail
}

func (this *CreateEmailBodySendAndPersistAsSent) saveEmailAsIntegrationErrorAndLogIfError(
	span main_domains_apm.SpanLogInfo,
	email main_domains.Email, err error) main_domains.Email {
	updatedEmail, errUpdateError := this.emailDatabaseGateway.Update(span.GetCtx(), email.CloneAsIntegrationError(err.Error()))
	if errUpdateError != nil {
		this.logsMonitoringGateway.ERROR(span, errUpdateError.Error())
		return email
	}
	return updatedEmail
}
