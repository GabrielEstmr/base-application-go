package main_gateways_ws_v1_response

import main_domains "baseapplicationgo/main/domains"

type PhoneContactResponse struct {
	PhoneCountry string `json:"phoneCountry"`
	PhoneType    string `json:"phoneType"`
	PhoneNumber  string `json:"phoneNumber"`
}

func NewPhoneContactResponse(
	phoneContact main_domains.PhoneContact,
) *PhoneContactResponse {
	return &PhoneContactResponse{
		PhoneCountry: phoneContact.GetPhoneCountry().Name(),
		PhoneType:    phoneContact.GetPhoneType().Name(),
		PhoneNumber:  phoneContact.GetPhoneNumber(),
	}
}
