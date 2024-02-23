package main_domains_features

import (
	main_configs_ff_lib_resources "baseapplicationgo/main/configs/ff/lib/resources"
	"reflect"
)

type FeaturesData struct {
	key          string
	group        string
	description  string
	defaultValue bool
}

func NewFeaturesDataAllArgs(
	key string,
	group string,
	description string,
	defaultValue bool) *FeaturesData {
	return &FeaturesData{
		key:          key,
		group:        group,
		description:  description,
		defaultValue: defaultValue}
}

func NewFeaturesData(
	feature main_configs_ff_lib_resources.FeaturesData) *FeaturesData {
	return &FeaturesData{
		key:          feature.GetKey(),
		group:        feature.GetGroup(),
		description:  feature.GetDescription(),
		defaultValue: feature.GetDefaultValue()}
}

func (this FeaturesData) ToGatewayResource() *main_configs_ff_lib_resources.FeaturesData {
	return main_configs_ff_lib_resources.NewFeaturesData(
		this.key,
		this.group,
		this.description,
		this.defaultValue,
	)
}

func (this FeaturesData) IsEmpty() bool {
	document := *new(FeaturesData)
	return reflect.DeepEqual(this, document)
}

func (this FeaturesData) GetKey() string {
	return this.key
}

func (this FeaturesData) GetGroup() string {
	return this.group
}

func (this FeaturesData) GetDescription() string {
	return this.description
}

func (this FeaturesData) GetDefaultValue() bool {
	return this.defaultValue
}

func (this FeaturesData) SetDefaultValue(defaultValue bool) {
	this.defaultValue = defaultValue
}

func (this FeaturesData) IsEnabled() bool {
	return this.defaultValue == true
}

func (this FeaturesData) IsDisabled() bool {
	return this.defaultValue == false
}
