package main_gateways_ws_v1_request

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_enums "baseapplicationgo/main/domains/enums"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_utils "baseapplicationgo/main/utils"
	"encoding/json"
	"io"
	"time"
)

type CreateUserRequest struct {
	DocumentId    string                `json:"documentId" validate:"required,min=4,max=15"`
	UserName      string                `json:"userName" validate:"required,min=8,max=15"`
	Password      string                `json:"password" validate:"required,min=8,max=20"`
	FirstName     string                `json:"firstName" validate:"required,min=4,max=15"`
	LastName      string                `json:"lastName" validate:"required,min=4,max=15"`
	Email         string                `json:"email" validate:"required,min=4,max=50"`
	Birthday      time.Time             `json:"birthday" validate:"required"`
	PhoneContacts []PhoneContactRequest `json:"phoneContacts" validate:"required,min=1,max=5"`
}

func (this *CreateUserRequest) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(this)
}

func (this *CreateUserRequest) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(this)
}

func (this *CreateUserRequest) ToDomain() main_domains.User {
	phones := make([]main_domains.PhoneContact, 0)
	for _, v := range this.PhoneContacts {
		phones = append(phones, v.ToDomain())
	}
	return *main_domains.NewUserAsCreated(
		this.DocumentId,
		this.UserName,
		this.Password,
		this.FirstName,
		this.LastName,
		this.Email,
		this.Birthday,
		phones,
		main_domains_enums.AUTH_PROVIDER_KEYCLOAK,
	)
}

func (this *CreateUserRequest) Validate() main_domains_exceptions.ApplicationException {
	structValidatorMessages := main_utils.NewStructValidatorMessages(this)
	if len(structValidatorMessages.GetMessages()) != 0 {
		return main_domains_exceptions.NewBadRequestException(structValidatorMessages.GetMessages())
	}

	for _, v := range this.PhoneContacts {
		err := v.Validate()
		if err != nil {
			return err
		}
	}

	return nil
}
