package main_gateways_ws_v1_request

import (
	main_domains "baseapplicationgo/main/domains"
	main_utils "baseapplicationgo/main/utils"
	"net/http"
	"time"
)

type FindUserFilterRequest struct {
	Name                  []string    `json:"name"`
	DocumentNumber        []string    `json:"documentNumber"`
	Birthday              []time.Time `json:"birthday"`
	StartCreatedDate      []time.Time `json:"startCreatedDate"`
	EndCreatedDate        []time.Time `json:"endCreatedDate"`
	StartLastModifiedDate []time.Time `json:"startLastModifiedDate"`
	EndLastModifiedDate   []time.Time `json:"endLastModifiedDate"`
}

func (this *FindUserFilterRequest) FindUserFilterToObject(w http.ResponseWriter, r *http.Request) (*FindUserFilterRequest, error) {
	filter := main_utils.NewQueryParams(this)
	object, err := main_utils.QueryParamsToObject(filter, w, r)
	obj := object.GetObj()
	return obj.(*FindUserFilterRequest), err
}

func (this *FindUserFilterRequest) ToDomain() main_domains.FindUserFilter {
	return *main_domains.NewFindUserFilter(
		this.Name,
		this.DocumentNumber,
		this.Birthday,
		this.StartCreatedDate,
		this.EndCreatedDate,
		this.StartLastModifiedDate,
		this.EndLastModifiedDate)
}
