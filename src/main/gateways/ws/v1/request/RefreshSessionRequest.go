package main_gateways_ws_v1_request

import (
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_utils "baseapplicationgo/main/utils"
	"net/http"
)

type RefreshSessionRequest struct {
	RefreshToken []string `json:"refresh_token,omitempty" validate:"required,min=1,max=1"`
}

func (this *RefreshSessionRequest) QueryParamsToObject(
	w http.ResponseWriter,
	r *http.Request,
) (*RefreshSessionRequest, main_domains_exceptions.ApplicationException) {
	filter := main_utils.NewQueryParams(this)
	object, err := main_utils.QueryParamsToObject(filter, w, r)
	if err != nil {
		return new(RefreshSessionRequest), err
	}
	obj := object.GetObj()
	return obj.(*RefreshSessionRequest), err
}

func (this *RefreshSessionRequest) GetFirstRefreshToken() string {
	if this.RefreshToken == nil {
		return main_utils.STRING_UTILS_EMPTY_STRING
	}
	if len(this.RefreshToken) == 0 {
		return main_utils.STRING_UTILS_EMPTY_STRING
	}
	return this.RefreshToken[0]
}

func (this *RefreshSessionRequest) Validate() main_domains_exceptions.ApplicationException {
	structValidatorMessages := main_utils.NewStructValidatorMessages(this)
	if len(structValidatorMessages.GetMessages()) != 0 {
		return main_domains_exceptions.NewBadRequestException(structValidatorMessages.GetMessages())
	}
	return nil
}
