package main_gateways_ws_v1_request

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_exceptions "baseapplicationgo/main/domains/exceptions"
	main_utils "baseapplicationgo/main/utils"
)

type CreatePlanMetadata struct {
	Value             float64 `json:"value" validate:"required"`
	DurationDays      int16   `json:"durationDays" validate:"required"`
	NumberOfUsers     int16   `json:"numberOfUsers" validate:"required"`
	NumberOfProjects  int16   `json:"numberOfProjects" validate:"required"`
	NumberOfCompanies int16   `json:"numberOfCompanies" validate:"required"`
}

func (this CreatePlanMetadata) ToDomain() main_domains.PlanMetadata {
	return *main_domains.NewPlanMetadata(
		this.Value,
		this.DurationDays,
		this.NumberOfUsers,
		this.NumberOfProjects,
		this.NumberOfCompanies,
	)
}

func (this CreatePlanMetadata) Validate() main_domains_exceptions.ApplicationException {
	structValidatorMessages := main_utils.NewStructValidatorMessages(this)
	if len(structValidatorMessages.GetMessages()) != 0 {
		return main_domains_exceptions.NewBadRequestException(structValidatorMessages.GetMessages())
	}
	return nil
}
