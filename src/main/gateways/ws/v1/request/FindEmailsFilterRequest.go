package main_gateways_ws_v1_request

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_enums "baseapplicationgo/main/domains/enums"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_utils "baseapplicationgo/main/utils"
	"net/http"
	"time"
)

type FindEmailsFilterRequest struct {
	Ids                   []string    `json:"id,omitempty"`
	Statuses              []string    `json:"status,omitempty"`
	EmailTemplateTypes    []string    `json:"emailTemplateType,omitempty"`
	AppOwners             []string    `json:"appOwner,omitempty"`
	StartCreatedDate      []time.Time `json:"startCreatedDate,omitempty"`
	EndCreatedDate        []time.Time `json:"endCreatedDate,omitempty"`
	StartLastModifiedDate []time.Time `json:"startLastModifiedDate,omitempty"`
	EndLastModifiedDate   []time.Time `json:"endLastModifiedDate,omitempty"`
}

func (this *FindEmailsFilterRequest) QueryParamsToObject(
	w http.ResponseWriter,
	r *http.Request,
) (
	*FindEmailsFilterRequest,
	main_domains_exceptions.ApplicationException,
) {
	filter := main_utils.NewQueryParams(this)
	object, err := main_utils.QueryParamsToObject(filter, w, r)
	if err != nil {
		return new(FindEmailsFilterRequest), err
	}
	obj := object.GetObj()
	return obj.(*FindEmailsFilterRequest), err
}

func (this *FindEmailsFilterRequest) ToDomain() main_domains.FindEmailFilter {
	return *main_domains.NewFindEmailFilter(
		this.Ids,
		main_domains_enums.GetEmailStatusesFromDescriptions(this.Statuses),
		this.EmailTemplateTypes,
		this.AppOwners,
		this.getMinStartCreatedDate(),
		this.getMaxEndCreatedDate(),
		this.getMinStartLastModifiedDate(),
		this.getMaxEndLastModifiedDate(),
	)
}

func (this *FindEmailsFilterRequest) getMinStartCreatedDate() time.Time {
	if this.StartCreatedDate == nil {
		return time.Time{}
	}
	return main_utils.NewDateUtils().GetMinValue(this.StartCreatedDate)
}

func (this *FindEmailsFilterRequest) getMaxEndCreatedDate() time.Time {
	if this.EndCreatedDate == nil {
		return time.Time{}
	}
	return main_utils.NewDateUtils().GetMaxValue(this.EndCreatedDate)
}

func (this *FindEmailsFilterRequest) getMinStartLastModifiedDate() time.Time {
	if this.StartLastModifiedDate == nil {
		return time.Time{}
	}
	return main_utils.NewDateUtils().GetMinValue(this.StartLastModifiedDate)
}

func (this *FindEmailsFilterRequest) getMaxEndLastModifiedDate() time.Time {
	if this.EndLastModifiedDate == nil {
		return time.Time{}
	}
	return main_utils.NewDateUtils().GetMaxValue(this.EndLastModifiedDate)
}
