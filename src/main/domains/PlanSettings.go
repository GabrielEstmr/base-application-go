package main_domains

import (
	main_domains_enums "baseapplicationgo/main/domains/enums"
	main_utils "baseapplicationgo/main/utils"
	"reflect"
	"time"
)

type PlanSetting struct {
	id                string
	planType          main_domains_enums.PlanType
	metadata          PlanMetadata
	status            main_domains_enums.PlanSettingStatus
	creationUserEmail string
	startDate         time.Time
	endDate           time.Time
	createdDate       time.Time
	lastModifiedDate  time.Time
}

func NewPlanSetting(
	planType main_domains_enums.PlanType,
	metadata PlanMetadata,
	status main_domains_enums.PlanSettingStatus,
	creationUserEmail string,
	startDate time.Time,
	endDate time.Time,
) *PlanSetting {
	return &PlanSetting{
		planType:          planType,
		metadata:          metadata,
		status:            status,
		creationUserEmail: creationUserEmail,
		startDate:         startDate,
		endDate:           endDate,
	}
}

func NewPlanSettingAllArgs(
	id string,
	planType main_domains_enums.PlanType,
	metadata PlanMetadata,
	status main_domains_enums.PlanSettingStatus,
	creationUserEmail string,
	startDate time.Time,
	endDate time.Time,
	createdDate time.Time,
	lastModifiedDate time.Time,
) *PlanSetting {
	return &PlanSetting{
		id:                id,
		planType:          planType,
		metadata:          metadata,
		status:            status,
		creationUserEmail: creationUserEmail,
		startDate:         startDate,
		endDate:           endDate,
		createdDate:       createdDate,
		lastModifiedDate:  lastModifiedDate,
	}
}

func (this PlanSetting) GetId() string {
	return this.id
}

func (this PlanSetting) GetPlanType() main_domains_enums.PlanType {
	return this.planType
}

func (this PlanSetting) GetMetadata() PlanMetadata {
	return this.metadata
}

func (this PlanSetting) GetCreationUserEmail() string {
	return this.creationUserEmail
}

func (this PlanSetting) GetStartDate() time.Time {
	return this.startDate
}

func (this PlanSetting) GetEndDate() time.Time {
	return this.endDate
}

func (this PlanSetting) GetCreatedDate() time.Time {
	return this.createdDate
}

func (this PlanSetting) GetLastModifiedDate() time.Time {
	return this.lastModifiedDate
}

func (this PlanSetting) GetStatus() main_domains_enums.PlanSettingStatus {
	return this.status
}

func (this PlanSetting) SetStartDate(startDate time.Time) {
	this.startDate = startDate
}

func (this PlanSetting) CloneWithStartDateAtStartOfTheDay() PlanSetting {
	return PlanSetting{
		id:                this.id,
		planType:          this.planType,
		metadata:          this.metadata,
		status:            this.status,
		creationUserEmail: this.creationUserEmail,
		startDate:         new(main_utils.DateUtils).GetDateUTCAtStartOfTheDay(this.startDate),
		endDate:           this.endDate,
		createdDate:       this.createdDate,
		lastModifiedDate:  this.lastModifiedDate,
	}

}

func (this PlanSetting) IsEmpty() bool {
	document := *new(PlanSetting)
	return reflect.DeepEqual(this, document)
}

func (this PlanSetting) CloneAsDisabled(endDate time.Time) PlanSetting {
	return PlanSetting{
		id:                this.id,
		planType:          this.planType,
		metadata:          this.metadata,
		status:            main_domains_enums.PLAN_SETTING_DISABLED,
		creationUserEmail: this.creationUserEmail,
		startDate:         this.startDate,
		endDate:           endDate,
		createdDate:       this.createdDate,
		lastModifiedDate:  this.lastModifiedDate,
	}
}
