package main_gateways_ws_v1_request

import (
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_utils "baseapplicationgo/main/utils"
	"encoding/json"
	"io"
)

type EnableInternalUserRequest struct {
	VerificationCode string `json:"verificationCode" validate:"required,min=4,max=15"`
	Email            string `json:"email" validate:"required,min=4,max=50"`
}

func (this EnableInternalUserRequest) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(this)
}

func (this EnableInternalUserRequest) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(this)
}

func (this EnableInternalUserRequest) Validate() main_domains_exceptions.ApplicationException {
	structValidatorMessages := main_utils.NewStructValidatorMessages(this)
	if len(structValidatorMessages.GetMessages()) != 0 {
		return main_domains_exceptions.NewBadRequestException(structValidatorMessages.GetMessages())
	}
	return nil
}
