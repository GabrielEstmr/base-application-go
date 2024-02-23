package main_gateways_ws_v1_request

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_enums "baseapplicationgo/main/domains/enums"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_utils "baseapplicationgo/main/utils"
	"net/http"
	"time"
)

type FindUserFilterRequest struct {
	Ids                   []string    `json:"id,omitempty"`
	AccountIds            []string    `json:"accountId,omitempty"`
	AuthProviderIds       []string    `json:"authProviderId,omitempty"`
	DocumentIds           []string    `json:"documentId,omitempty"`
	UserNames             []string    `json:"userName,omitempty"`
	FirstNames            []string    `json:"firstName,omitempty"`
	LastNames             []string    `json:"lastName,omitempty"`
	Emails                []string    `json:"email,omitempty"`
	EmailsVerified        []string    `json:"emailVerified,omitempty"`
	Statuses              []string    `json:"status,omitempty"`
	Roles                 []string    `json:"roles,omitempty"`
	ProviderTypes         []string    `json:"providerType,omitempty"`
	StartBirthdayDate     []time.Time `json:"startBirthdayDate,omitempty"`
	EndBirthdayDate       []time.Time `json:"endBirthdayDate,omitempty"`
	StartCreatedDate      []time.Time `json:"startCreatedDate,omitempty"`
	EndCreatedDate        []time.Time `json:"endCreatedDate,omitempty"`
	StartLastModifiedDate []time.Time `json:"startLastModifiedDate,omitempty"`
	EndLastModifiedDate   []time.Time `json:"endLastModifiedDate,omitempty"`
}

// understand why pointer here
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
	statuses := make([]main_domains_enums.UserStatus, 0)
	userStatus := *new(main_domains_enums.UserStatus)
	for _, v := range this.Statuses {
		statuses = append(statuses, userStatus.FromValue(v))
	}

	providerTypes := make([]main_domains_enums.AuthProviderType, 0)
	providerType := *new(main_domains_enums.AuthProviderType)
	for _, v := range this.ProviderTypes {
		providerTypes = append(providerTypes, providerType.FromValue(v))
	}

	return main_domains.FindUserFilter{
		Ids:             this.Ids,
		AccountIds:      this.AccountIds,
		AuthProviderIds: this.AuthProviderIds,
		DocumentIds:     this.DocumentIds,
		UserNames:       this.UserNames,
		FirstNames:      this.FirstNames,
		LastNames:       this.LastNames,
		Emails:          this.Emails,
		//EmailsVerified:        this.EmailsVerified,
		Statuses:              statuses,
		Roles:                 this.Roles,
		ProviderTypes:         providerTypes,
		StartBirthdayDate:     this.getMinStartBirthdayDate(),
		EndBirthdayDate:       this.getMaxEndBirthdayDate(),
		StartCreatedDate:      this.getMinStartCreatedDate(),
		EndCreatedDate:        this.getMaxEndCreatedDate(),
		StartLastModifiedDate: this.getMinStartLastModifiedDate(),
		EndLastModifiedDate:   this.getMaxEndLastModifiedDate(),
	}
}

func (this *FindUserFilterRequest) getMinStartBirthdayDate() time.Time {
	if this.StartCreatedDate == nil {
		return time.Time{}
	}
	return main_utils.NewDateUtils().GetMinValue(this.StartBirthdayDate)
}

func (this *FindUserFilterRequest) getMaxEndBirthdayDate() time.Time {
	if this.EndCreatedDate == nil {
		return time.Time{}
	}
	return main_utils.NewDateUtils().GetMaxValue(this.EndBirthdayDate)
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
