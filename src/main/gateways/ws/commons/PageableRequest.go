package main_gateways_ws_commons

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_utils "baseapplicationgo/main/utils"
	"encoding/json"
	"io"
)

type PageableRequest struct {
	Page int    `json:"page" validate:"required,min=0"`
	Size int    `json:"size" validate:"required,min=1,max=200"`
	Sort string `json:"sort" validate:"required"`
}

func (this *PageableRequest) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(this)
}

func (this *PageableRequest) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(this)
}

func (this *PageableRequest) ToDomain() main_domains.Pageable {
	return *main_domains.NewPageable(
		this.Page,
		this.Size,
		this.Sort)
}

func (this *PageableRequest) Validate() main_domains_exceptions.ApplicationException {
	structValidatorMessages := main_utils.NewStructValidatorMessages(this)
	if len(structValidatorMessages.GetMessages()) != 0 {
		return main_domains_exceptions.NewBadRequestException(structValidatorMessages.GetMessages())
	}
	return nil
}
