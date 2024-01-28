package main_usecases

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_apm "baseapplicationgo/main/domains/apm"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_gateways "baseapplicationgo/main/gateways"
	main_utils_messages "baseapplicationgo/main/utils/messages"
	"bytes"
	"context"
	"fmt"
	"text/template"
)

const _MSG_CREATE_EMAIL_BODY_WELCOME_ARCH_ISSUE_KEY = "exceptions.architecture.application.issue"

type CreateEmailBodyForWelcomeEmail struct {
	templatePath          string
	logsMonitoringGateway main_gateways.LogsMonitoringGateway
	spanGateway           main_gateways.SpanGateway
	messageUtils          main_utils_messages.ApplicationMessages
}

func NewCreateEmailBodyForWelcomeEmail(
	logsMonitoringGateway main_gateways.LogsMonitoringGateway,
	spanGateway main_gateways.SpanGateway,
	messageUtils main_utils_messages.ApplicationMessages,
) *CreateEmailBodyForWelcomeEmail {
	return &CreateEmailBodyForWelcomeEmail{
		templatePath:          "./zresources/emailtemplates/welcome-email-template.html",
		logsMonitoringGateway: logsMonitoringGateway,
		spanGateway:           spanGateway,
		messageUtils:          messageUtils,
	}
}

func NewCreateEmailBodyForWelcomeEmailAllArgs(
	templatePath string,
	logsMonitoringGateway main_gateways.LogsMonitoringGateway,
	spanGateway main_gateways.SpanGateway,
	messageUtils main_utils_messages.ApplicationMessages,
) *CreateEmailBodyForWelcomeEmail {
	return &CreateEmailBodyForWelcomeEmail{
		templatePath:          templatePath,
		logsMonitoringGateway: logsMonitoringGateway,
		spanGateway:           spanGateway,
		messageUtils:          messageUtils,
	}
}

func (this *CreateEmailBodyForWelcomeEmail) Execute(
	ctx context.Context,
	emailParams main_domains.EmailParams,
) ([]byte, main_domains_exceptions.ApplicationException) {
	span := this.spanGateway.Get(ctx, "CreateEmail-Execute")
	defer span.End()

	t, errP := template.ParseFiles(this.templatePath)
	if errP != nil {
		return this.logAndReturnArchError(span, errP)
	}

	var body bytes.Buffer
	subject := "Subject:" + emailParams.GetSubject() + "\n%s\n\n"
	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf(subject, mimeHeaders)))

	emailVariables := struct {
		Name    string
		Message string
	}{
		Name:    emailParams.GetBodyParams()["Name"],
		Message: emailParams.GetBodyParams()["Message"],
	}

	err := t.Execute(&body, emailVariables)
	if err != nil {
		return this.logAndReturnArchError(span, err)
	}

	return body.Bytes(), nil

}

func (this *CreateEmailBodyForWelcomeEmail) logAndReturnArchError(span main_domains_apm.SpanLogInfo, err error) (
	[]byte, main_domains_exceptions.ApplicationException) {
	this.logsMonitoringGateway.ERROR(span, err.Error())
	return nil, main_domains_exceptions.
		NewInternalServerErrorExceptionSglMsg(this.messageUtils.
			GetDefaultLocale(_MSG_CREATE_EMAIL_BODY_WELCOME_ARCH_ISSUE_KEY))
}
