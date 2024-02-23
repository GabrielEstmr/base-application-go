package main_gateways_ws_v1_request

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_enums "baseapplicationgo/main/domains/enums"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_utils "baseapplicationgo/main/utils"
	"net/http"
	"time"
)

type FindPlanSettingFilterRequest struct {
	Ids                   []string                      `json:"id,omitempty"`
	PlanTypes             []main_domains_enums.PlanType `json:"planType,omitempty"`
	CreationUserEmails    []string                      `json:"creationUserEmail,omitempty"`
	StartStartDate        []time.Time                   `json:"startStartDate,omitempty"`
	EndStartDate          []time.Time                   `json:"endStartDate,omitempty"`
	StartEndDate          []time.Time                   `json:"startEndDate,omitempty"`
	EndEndDate            []time.Time                   `json:"endEndDate,omitempty"`
	StartCreatedDate      []time.Time                   `json:"startCreatedDate,omitempty"`
	EndCreatedDate        []time.Time                   `json:"endCreatedDate,omitempty"`
	StartLastModifiedDate []time.Time                   `json:"startLastModifiedDate,omitempty"`
	EndLastModifiedDate   []time.Time                   `json:"endLastModifiedDate,omitempty"`
	HasEndDate            []bool                        `json:"hasEndDate,omitempty"`
}

func (this *FindPlanSettingFilterRequest) QueryParamsToObject(
	w http.ResponseWriter,
	r *http.Request) (*FindPlanSettingFilterRequest, main_domains_exceptions.ApplicationException) {
	filter := main_utils.NewQueryParams(this)
	object, err := main_utils.QueryParamsToObject(filter, w, r)
	if err != nil {
		return new(FindPlanSettingFilterRequest), err
	}
	obj := object.GetObj()
	return obj.(*FindPlanSettingFilterRequest), err
}

func (this *FindPlanSettingFilterRequest) ToDomain() main_domains.FindPlanSettingFilter {
	return main_domains.FindPlanSettingFilter{
		Ids:                   this.Ids,
		PlanTypes:             this.PlanTypes,
		CreationUserEmails:    this.CreationUserEmails,
		StartStartDate:        this.getMinStartStartDate(),
		EndStartDate:          this.getMaxEndStartDate(),
		StartEndDate:          this.getMinStartEndDate(),
		EndEndDate:            this.getMaxEndEndDate(),
		StartCreatedDate:      this.getMinStartCreatedDate(),
		EndCreatedDate:        this.getMaxEndCreatedDate(),
		StartLastModifiedDate: this.getMinStartLastModifiedDate(),
		EndLastModifiedDate:   this.getMaxEndLastModifiedDate(),
		HasEndDate:            this.HasEndDate,
	}
}

func (this *FindPlanSettingFilterRequest) getMinStartStartDate() time.Time {
	if this.StartCreatedDate == nil {
		return time.Time{}
	}
	return main_utils.NewDateUtils().GetMinValue(this.StartStartDate)
}

func (this *FindPlanSettingFilterRequest) getMaxEndStartDate() time.Time {
	if this.EndCreatedDate == nil {
		return time.Time{}
	}
	return main_utils.NewDateUtils().GetMaxValue(this.EndStartDate)
}

func (this *FindPlanSettingFilterRequest) getMinStartEndDate() time.Time {
	if this.StartLastModifiedDate == nil {
		return time.Time{}
	}
	return main_utils.NewDateUtils().GetMinValue(this.StartEndDate)
}

func (this *FindPlanSettingFilterRequest) getMaxEndEndDate() time.Time {
	if this.EndLastModifiedDate == nil {
		return time.Time{}
	}
	return main_utils.NewDateUtils().GetMaxValue(this.EndEndDate)
}

func (this *FindPlanSettingFilterRequest) getMinStartCreatedDate() time.Time {
	if this.StartCreatedDate == nil {
		return time.Time{}
	}
	return main_utils.NewDateUtils().GetMinValue(this.StartCreatedDate)
}

func (this *FindPlanSettingFilterRequest) getMaxEndCreatedDate() time.Time {
	if this.EndCreatedDate == nil {
		return time.Time{}
	}
	return main_utils.NewDateUtils().GetMaxValue(this.EndCreatedDate)
}

func (this *FindPlanSettingFilterRequest) getMinStartLastModifiedDate() time.Time {
	if this.StartLastModifiedDate == nil {
		return time.Time{}
	}
	return main_utils.NewDateUtils().GetMinValue(this.StartLastModifiedDate)
}

func (this *FindPlanSettingFilterRequest) getMaxEndLastModifiedDate() time.Time {
	if this.EndLastModifiedDate == nil {
		return time.Time{}
	}
	return main_utils.NewDateUtils().GetMaxValue(this.EndLastModifiedDate)
}
