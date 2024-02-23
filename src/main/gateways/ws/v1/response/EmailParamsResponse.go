package main_gateways_ws_v1_response

import (
	main_domains "baseapplicationgo/main/domains"
)

type EmailParamsResponse struct {
	EmailTemplateType string            `json:"emailTemplateType,omitempty"`
	AppOwner          string            `json:"appOwner,omitempty"`
	RequestUserId     string            `json:"requestUserId,omitempty"`
	To                []string          `json:"to,omitempty"`
	Subject           string            `json:"subject,omitempty"`
	BodyParams        map[string]string `json:"bodyParams,omitempty"`
}

func NewEmailParamsResponse(
	emailParams main_domains.EmailParams,
) *EmailParamsResponse {
	return &EmailParamsResponse{
		AppOwner:          emailParams.GetAppOwner(),
		EmailTemplateType: emailParams.GetEmailTemplateType().Name(),
		RequestUserId:     emailParams.GetRequestUserId(),
		To:                emailParams.GetTo(),
		Subject:           emailParams.GetSubject(),
		BodyParams:        emailParams.GetBodyParams(),
	}
}
