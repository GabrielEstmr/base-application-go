package main_gateways_ws_v1_request

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_utils "baseapplicationgo/main/utils"
	"net/http"
	"time"
)

type FindUserFilterRequest struct {
	Name                  []string    `json:"name,omitempty"`
	DocumentNumber        []string    `json:"documentNumber,omitempty"`
	Birthday              []time.Time `json:"birthday,omitempty"`
	StartCreatedDate      []time.Time `json:"startCreatedDate,omitempty"`
	EndCreatedDate        []time.Time `json:"endCreatedDate,omitempty"`
	StartLastModifiedDate []time.Time `json:"startLastModifiedDate,omitempty"`
	EndLastModifiedDate   []time.Time `json:"endLastModifiedDate,omitempty"`
}

func (this *FindUserFilterRequest) QueryParamsToObject(
	w http.ResponseWriter,
	r *http.Request) (*FindUserFilterRequest, main_domains_exceptions.ApplicationException) {
	filter := main_utils.NewQueryParams(this)
	object, err := main_utils.QueryParamsToObject(filter, w, r)
	if err != nil {
		return new(FindUserFilterRequest), err
	}
	obj := object.GetObj()
	return obj.(*FindUserFilterRequest), err
}

func (this *FindUserFilterRequest) ToDomain() main_domains.FindUserFilter {
	return *main_domains.NewFindUserFilter(
		this.Name,
		this.DocumentNumber,
		this.Birthday,
		this.getMinStartCreatedDate(),
		this.getMaxEndCreatedDate(),
		this.getMinStartLastModifiedDate(),
		this.getMaxEndLastModifiedDate(),
	)
}

func (this *FindUserFilterRequest) getMinStartCreatedDate() time.Time {
	if this.StartCreatedDate == nil {
		return time.Time{}
	}
	return main_utils.NewDateUtils().GetMinValue(this.StartCreatedDate)
}

func (this *FindUserFilterRequest) getMaxEndCreatedDate() time.Time {
	if this.EndCreatedDate == nil {
		return time.Time{}
	}
	return main_utils.NewDateUtils().GetMaxValue(this.EndCreatedDate)
}

func (this *FindUserFilterRequest) getMinStartLastModifiedDate() time.Time {
	if this.StartLastModifiedDate == nil {
		return time.Time{}
	}
	return main_utils.NewDateUtils().GetMinValue(this.StartLastModifiedDate)
}

func (this *FindUserFilterRequest) getMaxEndLastModifiedDate() time.Time {
	if this.EndLastModifiedDate == nil {
		return time.Time{}
	}
	return main_utils.NewDateUtils().GetMaxValue(this.EndLastModifiedDate)
}
