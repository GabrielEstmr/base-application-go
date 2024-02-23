package main_gateways_rabbitmq_resources

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_enums "baseapplicationgo/main/domains/enums"
	"encoding/json"
)

type EmailParamsResource struct {
	EmailTemplateType string            `json:"emailTemplateType"`
	AppOwner          string            `json:"appOwner"`
	RequestUserId     string            `json:"requestUserId"`
	To                []string          `json:"to"`
	Subject           string            `json:"subject"`
	BodyParams        map[string]string `json:"bodyParams"`
}

func NewEmailParamsResource(
	emailParams main_domains.EmailParams,
) *EmailParamsResource {
	return &EmailParamsResource{
		EmailTemplateType: emailParams.GetEmailTemplateType().Name(),
		AppOwner:          emailParams.GetAppOwner(),
		RequestUserId:     emailParams.GetRequestUserId(),
		To:                emailParams.GetTo(),
		Subject:           emailParams.GetSubject(),
		BodyParams:        emailParams.GetBodyParams(),
	}
}

func NewEmailParamsResourceAllArgs(
	emailTemplateType string,
	appOwner string,
	requestUserId string,
	to []string,
	subject string,
	bodyParams map[string]string,
) *EmailParamsResource {
	return &EmailParamsResource{
		EmailTemplateType: emailTemplateType,
		AppOwner:          appOwner,
		RequestUserId:     requestUserId,
		To:                to,
		Subject:           subject,
		BodyParams:        bodyParams,
	}
}

func (this *EmailParamsResource) FromJSON(data []byte) (EmailParamsResource, error) {
	var eventMessage EmailParamsResource
	if errU := json.Unmarshal(data, &eventMessage); errU != nil {
		return *new(EmailParamsResource), errU
	}
	return eventMessage, nil
}

func (this *EmailParamsResource) ToDomain() main_domains.EmailParams {
	return *main_domains.NewEmailParams(
		new(main_domains_enums.EmailTemplateType).FromValue(this.EmailTemplateType),
		this.AppOwner,
		this.RequestUserId,
		this.To,
		this.Subject,
		this.BodyParams,
	)
}
