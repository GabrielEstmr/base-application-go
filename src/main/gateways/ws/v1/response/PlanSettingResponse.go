package main_gateways_ws_v1_response

import (
	main_domains "baseapplicationgo/main/domains"
	"time"
)

type PlanSettingResponse struct {
	Id                string               `json:"id,omitempty"`
	PlanType          string               `json:"planType,omitempty"`
	Metadata          PlanMetadataResponse `json:"metadata,omitempty"`
	Status            string               `json:"status,omitempty"`
	CreationUserEmail string               `json:"creationUserEmail,omitempty"`
	StartDate         time.Time            `json:"startDate,omitempty"`
	EndDate           time.Time            `json:"endDate,omitempty"`
	CreatedDate       time.Time            `json:"createdDate,omitempty"`
	LastModifiedDate  time.Time            `json:"lastModifiedDate,omitempty"`
}

func NewPlanSettingResponse(
	planSetting main_domains.PlanSetting,
) *PlanSettingResponse {
	return &PlanSettingResponse{
		Id:                planSetting.GetId(),
		PlanType:          planSetting.GetPlanType().Name(),
		Metadata:          *NewPlanMetadataResponse(planSetting.GetMetadata()),
		Status:            planSetting.GetStatus().Name(),
		CreationUserEmail: planSetting.GetCreationUserEmail(),
		StartDate:         planSetting.GetStartDate(),
		EndDate:           planSetting.GetEndDate(),
		CreatedDate:       planSetting.GetCreatedDate(),
		LastModifiedDate:  planSetting.GetLastModifiedDate(),
	}
}
