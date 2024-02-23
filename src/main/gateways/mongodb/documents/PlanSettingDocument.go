package main_gateways_mongodb_documents

import (
	main_domains "baseapplicationgo/main/domains"
	main_domains_enums "baseapplicationgo/main/domains/enums"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"reflect"
)

type PlanSettingDocument struct {
	Id                primitive.ObjectID   `json:"_id,omitempty" bson:"_id,omitempty"`
	PlanType          string               `json:"planType,omitempty" bson:"planType,omitempty"`
	Metadata          PlanMetadataDocument `json:"metadata,omitempty" bson:"metadata,omitempty"`
	Status            string               `json:"status,omitempty" bson:"status,omitempty"`
	CreationUserEmail string               `json:"creationUserEmail,omitempty" bson:"creationUserEmail,omitempty"`
	StartDate         primitive.DateTime   `json:"startDate,omitempty" bson:"startDate"`
	EndDate           primitive.DateTime   `json:"endDate,omitempty" bson:"endDate"`
	CreatedDate       primitive.DateTime   `json:"createdDate,omitempty" bson:"createdDate"`
	LastModifiedDate  primitive.DateTime   `json:"lastModifiedDate,omitempty" bson:"lastModifiedDate"`
}

func NewPlanSettingDocument(planSetting main_domains.PlanSetting) PlanSettingDocument {
	oId, _ := primitive.ObjectIDFromHex(planSetting.GetId())
	return PlanSettingDocument{
		Id:                oId,
		PlanType:          planSetting.GetPlanType().Name(),
		Metadata:          *NewPlanMetadataDocument(planSetting.GetMetadata()),
		CreationUserEmail: planSetting.GetCreationUserEmail(),
		StartDate:         primitive.NewDateTimeFromTime(planSetting.GetStartDate()),
		EndDate:           primitive.NewDateTimeFromTime(planSetting.GetEndDate()),
		CreatedDate:       primitive.NewDateTimeFromTime(planSetting.GetCreatedDate()),
		LastModifiedDate:  primitive.NewDateTimeFromTime(planSetting.GetLastModifiedDate()),
	}
}

func (this PlanSettingDocument) IsEmpty() bool {
	document := *new(PlanSettingDocument)
	return reflect.DeepEqual(this, document)
}

func (this PlanSettingDocument) ToDomain() main_domains.PlanSetting {
	if this.IsEmpty() {
		return *new(main_domains.PlanSetting)
	}
	return *main_domains.NewPlanSettingAllArgs(
		this.Id.Hex(),
		new(main_domains_enums.PlanType).FromValue(this.PlanType),
		this.Metadata.ToDomain(),
		new(main_domains_enums.PlanSettingStatus).FromValue(this.Status),
		this.CreationUserEmail,
		this.StartDate.Time(),
		this.EndDate.Time(),
		this.CreatedDate.Time(),
		this.LastModifiedDate.Time(),
	)
}
