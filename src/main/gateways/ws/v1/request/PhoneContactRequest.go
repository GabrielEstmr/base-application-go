package main_gateways_ws_v1_request

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_enums "baseapplicationgo/main/domains/enums"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_utils "baseapplicationgo/main/utils"
)

type PhoneContactRequest struct {
	PhoneCountry string `json:"phoneCountry" validate:"required,min=2,max=2"`
	PhoneType    string `json:"phoneType" validate:"required,min=4,max=15"`
	PhoneNumber  string `json:"phoneNumber" validate:"required,min=8,max=20"`
}

func (this *PhoneContactRequest) ToDomain() main_domains.PhoneContact {
	return *main_domains.NewPhoneContact(
		new(main_domains_enums.Country).FromValue(this.PhoneCountry),
		new(main_domains_enums.PhoneType).FromValue(this.PhoneType),
		this.PhoneNumber,
	)
}

func (this *PhoneContactRequest) Validate() main_domains_exceptions.ApplicationException {
	structValidatorMessages := main_utils.NewStructValidatorMessages(this)
	if len(structValidatorMessages.GetMessages()) != 0 {
		return main_domains_exceptions.NewBadRequestException(structValidatorMessages.GetMessages())
	}
	return nil
}
