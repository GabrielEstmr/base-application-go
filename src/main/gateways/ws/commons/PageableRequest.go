package main_gateways_ws_commons

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_utils "baseapplicationgo/main/utils"
	"net/http"
	"strconv"
)

type PageableRequest struct {
	Page []string `json:"page,omitempty" validate:"required,min=0"`
	Size []string `json:"size,omitempty" validate:"required,min=1,max=200"`
	Sort []string `json:"sort,omitempty"`
}

func NewPageableRequest(
	page []string,
	size []string,
	sort []string) *PageableRequest {
	return &PageableRequest{
		Page: page,
		Size: size,
		Sort: sort}
}

func NewPageableRequestNoArgs() *PageableRequest {
	return &PageableRequest{
		Page: []string{"0"},
		Size: []string{"20"},
		Sort: nil}
}

func (this *PageableRequest) QueryParamsToObject(
	w http.ResponseWriter,
	r *http.Request) (*PageableRequest, main_domains_exceptions.ApplicationException) {
	filter := main_utils.NewQueryParams(this)
	object, err := main_utils.QueryParamsToObject(filter, w, r)
	if err != nil {
		return new(PageableRequest), err
	}
	obj := object.GetObj()
	return obj.(*PageableRequest), err
}

func (this *PageableRequest) ToDomain() (main_domains.Pageable, error) {
	page, err := strconv.ParseInt(this.Page[0], 10, 64)
	if err != nil {
		return main_domains.Pageable{}, err
	}
	size, err := strconv.ParseInt(this.Size[0], 10, 64)
	if err != nil {
		return main_domains.Pageable{}, err
	}
	return *main_domains.NewPageable(
		page,
		size,
		main_utils.NewPageableUtils().BuildPageableSortToMap(this.Sort)), nil
}

func (this *PageableRequest) Validate() main_domains_exceptions.ApplicationException {
	structValidatorMessages := main_utils.NewStructValidatorMessages(this)
	if len(structValidatorMessages.GetMessages()) != 0 {
		return main_domains_exceptions.NewBadRequestException(structValidatorMessages.GetMessages())
	}
	return nil
}
