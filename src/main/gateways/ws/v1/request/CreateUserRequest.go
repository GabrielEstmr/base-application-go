package main_gateways_ws_v1_request

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_utils "baseapplicationgo/main/utils"
	"encoding/json"
	"io"
	"time"
)

type CreateUserRequest struct {
	Name           string    `json:"name" validate:"required,min=4,max=15"`
	DocumentNumber string    `json:"documentNumber" validate:"required"`
	Age            int       `json:"age" validate:"required"`
	Birthday       time.Time `json:"birthday" validate:"required"`
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
	return main_domains.User{
		Name:           this.Name,
		DocumentNumber: this.DocumentNumber,
		Birthday:       this.Birthday,
	}
}

func (this *CreateUserRequest) Validate() main_domains_exceptions.ApplicationException {
	structValidatorMessages := main_utils.NewStructValidatorMessages(this)
	if len(structValidatorMessages.GetMessages()) != 0 {
		return main_domains_exceptions.NewBadRequestException(structValidatorMessages.GetMessages())
	}
	return nil
}
