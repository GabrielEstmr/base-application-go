package main_gateways_ws_v1_request

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_enums "baseapplicationgo/main/domains/enums"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_utils "baseapplicationgo/main/utils"
	"time"
)

type CreatePlanSetting struct {
	PlanType          string             `json:"planType" validate:"required"`
	Metadata          CreatePlanMetadata `json:"metadata" validate:"required"`
	CreationUserEmail string             `json:"creationUserEmail" validate:"required"`
	StartDate         time.Time          `json:"startDate" validate:"required"`
	EndDate           time.Time          `json:"endDate"`
}

func (this CreatePlanSetting) ToDomain() main_domains.PlanSetting {
	return *main_domains.NewPlanSetting(
		new(main_domains_enums.PlanType).FromValue(this.PlanType),
		this.Metadata.ToDomain(),
		main_domains_enums.PLAN_SETTING_ENABLED,
		this.CreationUserEmail,
		this.StartDate,
		this.EndDate,
	)
}

func (this CreatePlanSetting) Validate() main_domains_exceptions.ApplicationException {
	structValidatorMessages := main_utils.NewStructValidatorMessages(this)
	if len(structValidatorMessages.GetMessages()) != 0 {
		return main_domains_exceptions.NewBadRequestException(structValidatorMessages.GetMessages())
	}
	return nil
}
