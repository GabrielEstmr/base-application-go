package main_gateways_ws_v1_request

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_utils "baseapplicationgo/main/utils"
	"time"
)

type EnableExternalUserRequest struct {
	DocumentId    string                `json:"documentId" validate:"required,min=4,max=15"`
	UserName      string                `json:"userName" validate:"required,min=8,max=15"`
	Birthday      time.Time             `json:"birthday"`
	PhoneContacts []PhoneContactRequest `json:"phoneContacts" validate:"required,min=1,max=5"`
}

func (this EnableExternalUserRequest) Validate() main_domains_exceptions.ApplicationException {
	structValidatorMessages := main_utils.NewStructValidatorMessages(this)
	if len(structValidatorMessages.GetMessages()) != 0 {
		return main_domains_exceptions.NewBadRequestException(structValidatorMessages.GetMessages())
	}
	return nil
}

func (this EnableExternalUserRequest) ToDomain() main_domains.EnableExternalUserArgs {
	phones := make([]main_domains.PhoneContact, 0)
	for _, v := range this.PhoneContacts {
		phones = append(phones, v.ToDomain())
	}
	return *main_domains.NewEnableExternalUserArgs(
		this.DocumentId,
		this.UserName,
		this.Birthday,
		phones,
	)
}
