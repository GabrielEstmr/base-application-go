package main_gateways_mongodb_documents

import (
	main_domains "baseapplicationgo/main/domains"
	"reflect"
)

type PlanMetadataDocument struct {
	Value             float64 `json:"value,omitempty" bson:"value,omitempty"`
	DurationDays      int16   `json:"durationDays,omitempty" bson:"durationDays,omitempty"`
	NumberOfUsers     int16   `json:"numberOfUsers,omitempty" bson:"numberOfUsers,omitempty"`
	NumberOfProjects  int16   `json:"numberOfProjects,omitempty" bson:"numberOfProjects,omitempty"`
	NumberOfCompanies int16   `json:"numberOfCompanies,omitempty" bson:"numberOfCompanies,omitempty"`
}

func NewPlanMetadataDocument(
	metadata main_domains.PlanMetadata,
) *PlanMetadataDocument {
	return &PlanMetadataDocument{
		Value:             metadata.GetValue(),
		DurationDays:      metadata.GetDurationDays(),
		NumberOfUsers:     metadata.GetNumberOfUsers(),
		NumberOfProjects:  metadata.GetNumberOfProjects(),
		NumberOfCompanies: metadata.GetNumberOfCompanies(),
	}
}

func (this PlanMetadataDocument) IsEmpty() bool {
	document := *new(PlanMetadataDocument)
	return reflect.DeepEqual(this, document)
}

func (this PlanMetadataDocument) ToDomain() main_domains.PlanMetadata {
	if this.IsEmpty() {
		return *new(main_domains.PlanMetadata)
	}
	return *main_domains.NewPlanMetadata(
		this.Value,
		this.DurationDays,
		this.NumberOfUsers,
		this.NumberOfProjects,
		this.NumberOfCompanies,
	)
}
