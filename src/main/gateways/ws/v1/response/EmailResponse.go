package main_gateways_ws_v1_response

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_enums "baseapplicationgo/main/domains/enums"
	"time"
)

type EmailResponse struct {
	Id               string                         `json:"id,omitempty"`
	EmailParams      EmailParamsResponse            `json:"emailParams,omitempty"`
	Status           main_domains_enums.EmailStatus `json:"status,omitempty"`
	ErrorMsg         string                         `json:"errorMsg,omitempty"`
	CreatedDate      time.Time                      `json:"createdDate,omitempty"`
	LastModifiedDate time.Time                      `json:"lastModifiedDate,omitempty"`
}

func NewEmailResponse(
	email main_domains.Email,
) *EmailResponse {
	return &EmailResponse{
		Id:               email.GetId(),
		EmailParams:      *NewEmailParamsResponse(email.GetEmailParams()),
		Status:           email.GetStatus(),
		ErrorMsg:         email.GetErrorMsg(),
		CreatedDate:      email.GetCreatedDate(),
		LastModifiedDate: email.GetLastModifiedDate(),
	}
}
