package main_configs_ff_lib_mongo_documents

import (
	main_configs_ff_lib_resources "baseapplicationgo/main/configs/ff/lib/resources"
	"reflect"
)

type FeaturesDataDocument struct {
	Key          string `json:"_id,omitempty" bson:"_id,omitempty"`
	Group        string `json:"group,omitempty" bson:"group,omitempty"`
	Description  string `json:"description,omitempty" bson:"description,omitempty"`
	DefaultValue bool   `json:"defaultValue" bson:"defaultValue"`
}

func NewFeaturesDataDocument(featuresData main_configs_ff_lib_resources.FeaturesData) FeaturesDataDocument {
	return FeaturesDataDocument{
		Key:          featuresData.GetKey(),
		Group:        featuresData.GetGroup(),
		Description:  featuresData.GetDescription(),
		DefaultValue: featuresData.GetDefaultValue(),
	}
}

func (this FeaturesDataDocument) IsEmpty() bool {
	document := *new(FeaturesDataDocument)
	return reflect.DeepEqual(this, document)
}

func (this FeaturesDataDocument) ToDomain() main_configs_ff_lib_resources.FeaturesData {
	if this.IsEmpty() {
		return *new(main_configs_ff_lib_resources.FeaturesData)
	}
	return *main_configs_ff_lib_resources.NewFeaturesData(
		this.Key,
		this.Group,
		this.Description,
		this.DefaultValue,
	)
}

func (this FeaturesDataDocument) IsEnabled() bool {
	return this.DefaultValue == true
}

func (this FeaturesDataDocument) IsDisabled() bool {
	return this.DefaultValue == false
}
