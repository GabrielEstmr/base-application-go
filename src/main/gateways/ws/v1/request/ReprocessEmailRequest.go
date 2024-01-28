package main_gateways_ws_v1_request

import (
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_utils "baseapplicationgo/main/utils"
	"encoding/json"
	"io"
)

type ReprocessEmailRequest struct {
	Ids []string `json:"ids" validate:"required,min=1,max=15"`
}

func (this *ReprocessEmailRequest) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(this)
}

func (this *ReprocessEmailRequest) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(this)
}

func (this *ReprocessEmailRequest) Validate() main_domains_exceptions.ApplicationException {
	structValidatorMessages := main_utils.NewStructValidatorMessages(this)
	if len(structValidatorMessages.GetMessages()) != 0 {
		return main_domains_exceptions.NewBadRequestException(structValidatorMessages.GetMessages())
	}
	return nil
}
