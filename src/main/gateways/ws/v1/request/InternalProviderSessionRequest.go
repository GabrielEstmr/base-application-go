package main_gateways_ws_v1_request

import (
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_utils "baseapplicationgo/main/utils"
	"net/http"
)

type InternalProviderSessionRequest struct {
	Username []string `json:"username,omitempty" validate:"required,min=1,max=1"`
	Password []string `json:"password,omitempty" validate:"required,min=1,max=1"`
}

func (this *InternalProviderSessionRequest) QueryParamsToObject(
	w http.ResponseWriter,
	r *http.Request,
) (*InternalProviderSessionRequest, main_domains_exceptions.ApplicationException) {
	filter := main_utils.NewQueryParams(this)
	object, err := main_utils.QueryParamsToObject(filter, w, r)
	if err != nil {
		return new(InternalProviderSessionRequest), err
	}
	obj := object.GetObj()
	return obj.(*InternalProviderSessionRequest), err
}

func (this *InternalProviderSessionRequest) GetFirstUsername() string {
	if this.Username == nil {
		return main_utils.STRING_UTILS_EMPTY_STRING
	}
	if len(this.Username) == 0 {
		return main_utils.STRING_UTILS_EMPTY_STRING
	}
	return this.Username[0]
}

func (this *InternalProviderSessionRequest) GetFirstPassword() string {
	if this.Password == nil {
		return main_utils.STRING_UTILS_EMPTY_STRING
	}
	if len(this.Password) == 0 {
		return main_utils.STRING_UTILS_EMPTY_STRING
	}
	return this.Password[0]
}

func (this *InternalProviderSessionRequest) Validate() main_domains_exceptions.ApplicationException {
	structValidatorMessages := main_utils.NewStructValidatorMessages(this)
	if len(structValidatorMessages.GetMessages()) != 0 {
		return main_domains_exceptions.NewBadRequestException(structValidatorMessages.GetMessages())
	}
	return nil
}
