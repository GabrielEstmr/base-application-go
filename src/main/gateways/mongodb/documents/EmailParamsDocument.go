package main_gateways_mongodb_documents

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_enums "baseapplicationgo/main/domains/enums"
	"reflect"
)

type EmailParamsDocument struct {
	EmailTemplateType string            `json:"emailTemplateType,omitempty" bson:"emailTemplateType,omitempty"`
	AppOwner          string            `json:"appOwner,omitempty" bson:"appOwner,omitempty"`
	RequestUserId     string            `json:"requestUserId,omitempty" bson:"requestUserId,omitempty"`
	To                []string          `json:"to,omitempty" bson:"to,omitempty"`
	Subject           string            `json:"subject,omitempty" bson:"subject,omitempty"`
	BodyParams        map[string]string `json:"bodyParams,omitempty" bson:"bodyParams,omitempty"`
}

func NewEmailParamsDocument(
	emailParams main_domains.EmailParams,
) *EmailParamsDocument {
	return &EmailParamsDocument{
		AppOwner:          emailParams.GetAppOwner(),
		EmailTemplateType: main_domains_enums.GetEmailTemplateTypeDescription(emailParams.GetEmailTemplateType()),
		RequestUserId:     emailParams.GetRequestUserId(),
		To:                emailParams.GetTo(),
		Subject:           emailParams.GetSubject(),
		BodyParams:        emailParams.GetBodyParams(),
	}
}

func (this *EmailParamsDocument) GetEmailTemplateType() string {
	return this.EmailTemplateType
}

func (this *EmailParamsDocument) GetAppOwner() string {
	return this.AppOwner
}

func (this *EmailParamsDocument) GetRequestUserId() string {
	return this.RequestUserId
}

func (this *EmailParamsDocument) GetTo() []string {
	return this.To
}

func (this *EmailParamsDocument) GetSubject() string {
	return this.Subject
}

func (this *EmailParamsDocument) GetBodyParams() map[string]string {
	return this.BodyParams
}

func (this *EmailParamsDocument) IsEmpty() bool {
	document := new(EmailParamsDocument)
	return reflect.DeepEqual(this, document)
}

func (this *EmailParamsDocument) ToDomain() main_domains.EmailParams {
	if this.IsEmpty() {
		return *new(main_domains.EmailParams)
	}
	return *main_domains.NewEmailParams(
		main_domains_enums.GetEmailTemplateTypeFromDescription(this.EmailTemplateType),
		this.AppOwner,
		this.RequestUserId,
		this.To,
		this.Subject,
		this.BodyParams,
	)
}
