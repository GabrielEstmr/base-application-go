package main_configs_ff_lib_redis_documents

import (
	main_configs_ff_lib_resources "baseapplicationgo/main/configs/ff/lib/resources"
	"reflect"
)

type FeaturesDataRedisDocument struct {
	Key          string `json:"_id,omitempty"`
	Group        string `json:"group,omitempty"`
	Description  string `json:"description,omitempty"`
	DefaultValue bool   `json:"defaultValue"`
}

func NewFeaturesDataRedisDocument(featuresData main_configs_ff_lib_resources.FeaturesData) FeaturesDataRedisDocument {
	return FeaturesDataRedisDocument{
		Key:          featuresData.GetKey(),
		Group:        featuresData.GetGroup(),
		Description:  featuresData.GetDescription(),
		DefaultValue: featuresData.GetDefaultValue(),
	}
}

func (this FeaturesDataRedisDocument) IsEmpty() bool {
	document := *new(FeaturesDataRedisDocument)
	return reflect.DeepEqual(this, document)
}

func (this FeaturesDataRedisDocument) ToDomain() main_configs_ff_lib_resources.FeaturesData {
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
