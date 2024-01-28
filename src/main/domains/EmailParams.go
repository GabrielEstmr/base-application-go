package main_domains

import (
	main_domains_enums "baseapplicationgo/main/domains/enums"
	"reflect"
)

type EmailParams struct {
	emailTemplateType main_domains_enums.EmailTemplateType
	appOwner          string
	requestUserId     string
	to                []string
	subject           string
	bodyParams        map[string]string
}

func NewEmailParams(
	emailTemplateType main_domains_enums.EmailTemplateType,
	appOwner string,
	requestUserId string,
	to []string,
	subject string,
	bodyParams map[string]string,
) *EmailParams {
	return &EmailParams{
		appOwner:          appOwner,
		emailTemplateType: emailTemplateType,
		requestUserId:     requestUserId,
		to:                to,
		subject:           subject,
		bodyParams:        bodyParams,
	}
}

func (this EmailParams) GetEmailTemplateType() main_domains_enums.EmailTemplateType {
	return this.emailTemplateType
}

func (this EmailParams) GetAppOwner() string {
	return this.appOwner
}

func (this EmailParams) GetRequestUserId() string {
	return this.requestUserId
}

func (this EmailParams) GetTo() []string {
	return this.to
}

func (this EmailParams) GetSubject() string {
	return this.subject
}

func (this EmailParams) GetBodyParams() map[string]string {
	return this.bodyParams
}

func (this EmailParams) IsEmpty() bool {
	document := *new(EmailParams)
	return reflect.DeepEqual(this, document)
}
