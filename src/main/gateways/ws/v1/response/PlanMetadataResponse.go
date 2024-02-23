package main_gateways_ws_v1_response

import (
	main_domains "baseapplicationgo/main/domains"
)

type PlanMetadataResponse struct {
	Value             float64 `json:"value,omitempty"`
	DurationDays      int16   `json:"durationDays,omitempty"`
	NumberOfUsers     int16   `json:"numberOfUsers,omitempty"`
	NumberOfProjects  int16   `json:"numberOfProjects,omitempty"`
	NumberOfCompanies int16   `json:"numberOfCompanies,omitempty"`
}

func NewPlanMetadataResponse(
	metadata main_domains.PlanMetadata,
) *PlanMetadataResponse {
	return &PlanMetadataResponse{
		Value:             metadata.GetValue(),
		DurationDays:      metadata.GetDurationDays(),
		NumberOfUsers:     metadata.GetNumberOfUsers(),
		NumberOfProjects:  metadata.GetNumberOfProjects(),
		NumberOfCompanies: metadata.GetNumberOfCompanies(),
	}
}
