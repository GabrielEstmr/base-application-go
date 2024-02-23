package main_gateways_ws_v1_request

import (
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_utils "baseapplicationgo/main/utils"
)

type ChangeUserPasswordRequest struct {
	Password         string `json:"password" validate:"required,min=8,max=20"`
	VerificationCode string `json:"verificationCode" validate:"required,min=8,max=20"`
}

func (this *ChangeUserPasswordRequest) Validate() main_domains_exceptions.ApplicationException {
	structValidatorMessages := main_utils.NewStructValidatorMessages(this)
	if len(structValidatorMessages.GetMessages()) != 0 {
		return main_domains_exceptions.NewBadRequestException(structValidatorMessages.GetMessages())
	}
	return nil
}
