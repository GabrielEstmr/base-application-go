package main_configs_ff_lib_mongorepo_documents

import (
	main_configs_ff_lib_resources "baseapplicationgo/main/configs/ff/lib/resources"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FeaturesDataDocument struct {
	Key          primitive.ObjectID `json:"key,omitempty" bson:"key,omitempty"`
	Group        string             `json:"group,omitempty" bson:"group,omitempty"`
	Description  string             `json:"description,omitempty" bson:"description,omitempty"`
	DefaultValue bool               `json:"defaultValue,omitempty" bson:"defaultValue,omitempty"`
}

func NewFeaturesDataDocument(featuresData main_configs_ff_lib_resources.FeaturesData) FeaturesDataDocument {
	return FeaturesDataDocument{
		Group:        featuresData.GetGroup(),
		Description:  featuresData.GetDescription(),
		DefaultValue: featuresData.GetDefaultValue(),
	}
}

func (this *FeaturesDataDocument) IsEmpty() bool {
	return *this == FeaturesDataDocument{}
}

func (this *FeaturesDataDocument) ToDomain() main_configs_ff_lib_resources.FeaturesData {
	if this.IsEmpty() {
		return main_configs_ff_lib_resources.FeaturesData{}
	}
	return *main_configs_ff_lib_resources.NewFeaturesData(
		this.Key.Hex(),
		this.Group,
		this.Description,
		this.DefaultValue,
	)
}
