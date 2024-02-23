package main_domains

import (
	main_domains_enums "baseapplicationgo/main/domains/enums"
	"time"
)

type FindPlanSettingFilter struct {
	Ids                   []string
	PlanTypes             []main_domains_enums.PlanType
	CreationUserEmails    []string
	StartStartDate        time.Time
	EndStartDate          time.Time
	StartEndDate          time.Time
	EndEndDate            time.Time
	StartCreatedDate      time.Time
	EndCreatedDate        time.Time
	StartLastModifiedDate time.Time
	EndLastModifiedDate   time.Time
	HasEndDate            []bool
}
